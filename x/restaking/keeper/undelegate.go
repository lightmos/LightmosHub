package keeper

import (
	"cosmossdk.io/math"
	"encoding/binary"
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v7/modules/core/24-host"
	restakingtypes "lightmos/x/restaking/types"
	"time"
)

var (
	ShareUnbondingDelegationKey           = []byte{0x64}
	ShareUnbondingIDKey                   = []byte{0x65}
	ShareUnbondingDelegationByValIndexKey = []byte{0x68}
	ShareUnbondingQueueKey                = []byte{0x69}
)

// TransmitUndelegatePacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitUndelegatePacket(
	ctx sdk.Context,
	packetData restakingtypes.UndelegatePacketData,
	sourcePort,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
) (uint64, error) {
	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return 0, sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	packetBytes, err := packetData.GetBytes()
	if err != nil {
		return 0, sdkerrors.Wrapf(sdkerrors.ErrJSONMarshal, "cannot marshal the packet: %w", err)
	}

	return k.channelKeeper.SendPacket(ctx, channelCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, packetBytes)
}

// OnRecvUndelegatePacket processes packet reception
func (k Keeper) OnRecvUndelegatePacket(ctx sdk.Context, packet channeltypes.Packet, data restakingtypes.UndelegatePacketData) (packetAck restakingtypes.UndelegatePacketAck, err error) {
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	// TODO: packet reception logic
	accAddr, _ := sdk.AccAddressFromBech32(data.ValidatorAddress)
	valAddr := sdk.ValAddress(accAddr)
	bondDenom := k.stakingKeeper.BondDenom(ctx)
	if bondDenom != data.Amount.Denom {
		return packetAck, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid coin denomination: got %s, expected %s", data.Amount.Denom, bondDenom)
	}

	shares, err := k.stakingKeeper.ValidateUnbondAmount(
		ctx, accAddr, valAddr, data.Amount.Amount,
	)
	if err != nil {
		return packetAck, err
	}
	completionTime, err := k.stakingKeeper.Undelegate(ctx, accAddr, valAddr, shares)
	if err != nil {
		return packetAck, err
	}
	//save share validator
	ubds := k.SetShareUnbondingDelegationEntry(ctx, accAddr, valAddr, ctx.BlockHeight(), completionTime, shares.TruncateInt())
	k.InsertShareUBDQueue(ctx, ubds, completionTime)
	packetAck.Step = 1

	return packetAck, nil
}

// OnAcknowledgementUndelegatePacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementUndelegatePacket(ctx sdk.Context, packet channeltypes.Packet, data restakingtypes.UndelegatePacketData, ack channeltypes.Acknowledgement) error {
	log := k.Logger(ctx)
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		return errors.New(dispatchedAck.Error)
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck restakingtypes.UndelegatePacketAck

		if err := restakingtypes.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		// TODO: successful acknowledgement logic
		if packetAck.Step == 1 {
			log.Info("azh|OnAcknowledgementRetireSharePacket unbound")
		}
		return nil
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutUndelegatePacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutUndelegatePacket(ctx sdk.Context, packet channeltypes.Packet, data restakingtypes.UndelegatePacketData) error {

	// TODO: packet timeout logic

	return nil
}

func (k Keeper) CompleteShareUnbonding(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) (sdk.Coins, error) {
	ubd, found := k.GetShareUnbondingDelegation(ctx, delAddr, valAddr)
	if !found {
		return nil, types.ErrNoUnbondingDelegation
	}

	bondDenom := k.stakingKeeper.BondDenom(ctx)
	balances := sdk.NewCoins()
	ctxTime := ctx.BlockHeader().Time

	delegatorAddress, err := sdk.AccAddressFromBech32(ubd.DelegatorAddress)
	if err != nil {
		return nil, err
	}

	// loop through all the entries and complete unbonding mature entries
	for i := 0; i < len(ubd.Entries); i++ {
		entry := ubd.Entries[i]
		if entry.IsMature(ctxTime) {
			ubd.RemoveEntry(int64(i))
			i--

			// track undelegation only when remaining or truncated shares are non-zero
			if !entry.Balance.IsZero() {
				amt := sdk.NewCoin(bondDenom, entry.Balance)
				if err := k.bankKeeper.SendCoinsFromAccountToModule(
					ctx, delegatorAddress, restakingtypes.ModuleName, sdk.NewCoins(amt),
				); err != nil {
					return nil, err
				}

				balances = balances.Add(amt)
			}
		}
	}

	// set the unbonding delegation or remove it if there are no more entries
	if len(ubd.Entries) == 0 {
		k.RemoveShareUnbondingDelegation(ctx, ubd)
	} else {
		k.SetShareUnbondingDelegation(ctx, ubd)
	}

	return balances, nil
}

func (k Keeper) SetShareUnbondingDelegationEntry(
	ctx sdk.Context, delegatorAddr sdk.AccAddress, validatorAddr sdk.ValAddress,
	creationHeight int64, minTime time.Time, balance math.Int,
) types.UnbondingDelegation {
	ubd, found := k.GetShareUnbondingDelegation(ctx, delegatorAddr, validatorAddr)
	id := k.IncrementShareUnbondingID(ctx)
	if found {
		ubd.AddEntry(creationHeight, minTime, balance, id)
	} else {
		ubd = types.NewUnbondingDelegation(delegatorAddr, validatorAddr, creationHeight, minTime, balance, id)
	}

	k.SetShareUnbondingDelegation(ctx, ubd)

	return ubd
}

func (k Keeper) InsertShareUBDQueue(ctx sdk.Context, ubd types.UnbondingDelegation, completionTime time.Time) {
	dvPair := types.DVPair{DelegatorAddress: ubd.DelegatorAddress, ValidatorAddress: ubd.ValidatorAddress}

	timeSlice := k.GetShareUBDQueueTimeSlice(ctx, completionTime)
	if len(timeSlice) == 0 {
		k.SetUBDQueueTimeSlice(ctx, completionTime, []types.DVPair{dvPair})
	} else {
		timeSlice = append(timeSlice, dvPair)
		k.SetUBDQueueTimeSlice(ctx, completionTime, timeSlice)
	}
}

func (k Keeper) DequeueAllMatureUBDQueue(ctx sdk.Context, currTime time.Time) (matureUnbonds []types.DVPair) {
	store := ctx.KVStore(k.storeKey)

	// gets an iterator for all timeslices from time 0 until the current Blockheader time
	unbondingTimesliceIterator := k.ShareUBDQueueIterator(ctx, currTime)
	defer unbondingTimesliceIterator.Close()

	for ; unbondingTimesliceIterator.Valid(); unbondingTimesliceIterator.Next() {
		timeslice := types.DVPairs{}
		value := unbondingTimesliceIterator.Value()
		k.cdc.MustUnmarshal(value, &timeslice)

		matureUnbonds = append(matureUnbonds, timeslice.Pairs...)

		store.Delete(unbondingTimesliceIterator.Key())
	}

	return matureUnbonds
}

func (k Keeper) ShareUBDQueueIterator(ctx sdk.Context, endTime time.Time) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return store.Iterator(ShareUnbondingQueueKey,
		sdk.InclusiveEndBytes(GetShareUnbondingDelegationTimeKey(endTime)))
}

func (k Keeper) GetShareUnbondingDelegation(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) (ubd types.UnbondingDelegation, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := GetShareUBDKey(delAddr, valAddr)
	value := store.Get(key)

	if value == nil {
		return ubd, false
	}

	ubd = types.MustUnmarshalUBD(k.cdc, value)

	return ubd, true
}

func (k Keeper) SetShareUnbondingDelegation(ctx sdk.Context, ubd types.UnbondingDelegation) {
	delAddr := sdk.MustAccAddressFromBech32(ubd.DelegatorAddress)

	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalUBD(k.cdc, ubd)
	valAddr, err := sdk.ValAddressFromBech32(ubd.ValidatorAddress)
	if err != nil {
		panic(err)
	}
	key := GetShareUBDKey(delAddr, valAddr)
	store.Set(key, bz)
	store.Set(GetShareUBDByValIndexKey(delAddr, valAddr), []byte{}) // index, store empty bytes
}

func (k Keeper) RemoveShareUnbondingDelegation(ctx sdk.Context, ubd types.UnbondingDelegation) {
	delegatorAddress := sdk.MustAccAddressFromBech32(ubd.DelegatorAddress)

	store := ctx.KVStore(k.storeKey)
	addr, err := sdk.ValAddressFromBech32(ubd.ValidatorAddress)
	if err != nil {
		panic(err)
	}
	key := GetShareUBDKey(delegatorAddress, addr)
	store.Delete(key)
	store.Delete(types.GetUBDByValIndexKey(delegatorAddress, addr))
}

func (k Keeper) IncrementShareUnbondingID(ctx sdk.Context) (unbondingID uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(ShareUnbondingIDKey)
	if bz != nil {
		unbondingID = binary.BigEndian.Uint64(bz)
	}

	unbondingID++

	// Convert back into bytes for storage
	bz = make([]byte, 8)
	binary.BigEndian.PutUint64(bz, unbondingID)

	store.Set(types.UnbondingIDKey, bz)

	return unbondingID
}

func (k Keeper) GetShareUBDQueueTimeSlice(ctx sdk.Context, timestamp time.Time) (dvPairs []types.DVPair) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(GetShareUnbondingDelegationTimeKey(timestamp))
	if bz == nil {
		return []types.DVPair{}
	}

	pairs := types.DVPairs{}
	k.cdc.MustUnmarshal(bz, &pairs)

	return pairs.Pairs
}

func (k Keeper) SetUBDQueueTimeSlice(ctx sdk.Context, timestamp time.Time, keys []types.DVPair) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&types.DVPairs{Pairs: keys})
	store.Set(GetShareUnbondingDelegationTimeKey(timestamp), bz)
}

func GetShareUBDKey(delAddr sdk.AccAddress, valAddr sdk.ValAddress) []byte {
	return append(GetShareUBDsKey(delAddr.Bytes()), address.MustLengthPrefix(valAddr)...)
}

func GetShareUBDsKey(delAddr sdk.AccAddress) []byte {
	return append(ShareUnbondingDelegationKey, address.MustLengthPrefix(delAddr)...)
}

func GetShareUBDByValIndexKey(delAddr sdk.AccAddress, valAddr sdk.ValAddress) []byte {
	return append(GetShareUBDsByValIndexKey(valAddr), address.MustLengthPrefix(delAddr)...)
}

func GetShareUBDsByValIndexKey(valAddr sdk.ValAddress) []byte {
	return append(ShareUnbondingDelegationByValIndexKey, address.MustLengthPrefix(valAddr)...)
}

func GetShareUnbondingDelegationTimeKey(timestamp time.Time) []byte {
	bz := sdk.FormatTimeBytes(timestamp)
	return append(ShareUnbondingQueueKey, bz...)
}

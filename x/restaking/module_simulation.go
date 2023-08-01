package restaking

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"lightmos/testutil/sample"
	restakingsimulation "lightmos/x/restaking/simulation"
	"lightmos/x/restaking/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = restakingsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCancelSellOrder = "op_weight_msg_cancel_sell_order"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCancelSellOrder int = 100

	opWeightMsgCancelBuyOrder = "op_weight_msg_cancel_buy_order"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCancelBuyOrder int = 100

	opWeightMsgChangePairState = "op_weight_msg_change_pair_state"
	// TODO: Determine the simulation weight value
	defaultWeightMsgChangePairState int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	restakingGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&restakingGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCancelSellOrder int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCancelSellOrder, &weightMsgCancelSellOrder, nil,
		func(_ *rand.Rand) {
			weightMsgCancelSellOrder = defaultWeightMsgCancelSellOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancelSellOrder,
		restakingsimulation.SimulateMsgCancelSellOrder(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCancelBuyOrder int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCancelBuyOrder, &weightMsgCancelBuyOrder, nil,
		func(_ *rand.Rand) {
			weightMsgCancelBuyOrder = defaultWeightMsgCancelBuyOrder
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancelBuyOrder,
		restakingsimulation.SimulateMsgCancelBuyOrder(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgChangePairState int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgChangePairState, &weightMsgChangePairState, nil,
		func(_ *rand.Rand) {
			weightMsgChangePairState = defaultWeightMsgChangePairState
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgChangePairState,
		restakingsimulation.SimulateMsgChangePairState(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCancelSellOrder,
			defaultWeightMsgCancelSellOrder,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				restakingsimulation.SimulateMsgCancelSellOrder(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCancelBuyOrder,
			defaultWeightMsgCancelBuyOrder,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				restakingsimulation.SimulateMsgCancelBuyOrder(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgChangePairState,
			defaultWeightMsgChangePairState,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				restakingsimulation.SimulateMsgChangePairState(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}

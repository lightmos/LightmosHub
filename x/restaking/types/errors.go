package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/restaking module sentinel errors
var (
	ErrNoDelegation         = sdkerrors.Register(ModuleName, 19, "no delegation for (address, validator) tuple")
	ErrSample               = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidPacketTimeout = sdkerrors.Register(ModuleName, 1500, "invalid packet timeout")
	ErrInvalidVersion       = sdkerrors.Register(ModuleName, 1501, "invalid version")

	ErrEmptyValidatorPubKey          = sdkerrors.Register(ModuleName, 39, "empty validator public key")
	ErrCommissionNegative            = sdkerrors.Register(ModuleName, 9, "commission must be positive")
	ErrCommissionHuge                = sdkerrors.Register(ModuleName, 10, "commission cannot be more than 100%")
	ErrCommissionGTMaxRate           = sdkerrors.Register(ModuleName, 11, "commission cannot be more than the max rate")
	ErrCommissionUpdateTime          = sdkerrors.Register(ModuleName, 12, "commission cannot be changed more than once in 24h")
	ErrCommissionChangeRateNegative  = sdkerrors.Register(ModuleName, 13, "commission change rate must be positive")
	ErrCommissionChangeRateGTMaxRate = sdkerrors.Register(ModuleName, 14, "commission change rate cannot be more than the max rate")
	ErrCommissionGTMaxChangeRate     = sdkerrors.Register(ModuleName, 15, "commission cannot be changed more than max change rate")
	ErrSelfDelegationBelowMinimum    = sdkerrors.Register(ModuleName, 16, "validator's self delegation must be greater than their minimum self delegation")
	ErrRetireNotFound                = sdkerrors.Register(ModuleName, 17, "retire not found")
	ErrTokenTooLarge                 = sdkerrors.Register(ModuleName, 18, "token is too large")
	ErrValidatorNotFound             = sdkerrors.Register(ModuleName, 30, "validator not found")
	ErrExceedBondedAmount            = sdkerrors.Register(ModuleName, 31, "exceed bonded amount")
	ErrExceedBalance                 = sdkerrors.Register(ModuleName, 32, "exceed balance")
)

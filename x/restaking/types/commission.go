package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Validate performs basic sanity validation checks of initial commission
// parameters. If validation fails, an SDK error is returned.
func (cr CommissionRates) Validate() error {
	switch {
	case cr.MaxRate.IsNegative():
		// max rate cannot be negative
		return ErrCommissionNegative

	case cr.MaxRate.GT(sdk.OneDec()):
		// max rate cannot be greater than 1
		return ErrCommissionHuge

	case cr.Rate.IsNegative():
		// rate cannot be negative
		return ErrCommissionNegative

	case cr.Rate.GT(cr.MaxRate):
		// rate cannot be greater than the max rate
		return ErrCommissionGTMaxRate

	case cr.MaxChangeRate.IsNegative():
		// change rate cannot be negative
		return ErrCommissionChangeRateNegative

	case cr.MaxChangeRate.GT(cr.MaxRate):
		// change rate cannot be greater than the max rate
		return ErrCommissionChangeRateGTMaxRate
	}

	return nil
}

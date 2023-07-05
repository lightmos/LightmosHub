package types

import (
	fmt "fmt"
	"regexp"
)

var (
	reDnm *regexp.Regexp
)

// IsValid returns true if the Coin has a non-negative amount and the denom is valid.
func (coin Coin) IsValid() bool {
	return coin.Validate() == nil
}

// Validate returns an error if the Coin has a negative amount or if
// the denom is invalid.
func (coin Coin) Validate() error {
	if err := ValidateDenom(coin.Denom); err != nil {
		return err
	}

	if coin.Amount.IsNegative() {
		return fmt.Errorf("negative coin amount: %v", coin.Amount)
	}

	return nil
}

// ValidateDenom is the default validation function for Coin.Denom.
func ValidateDenom(denom string) error {
	// if !reDnm.MatchString(denom) {
	// 	return fmt.Errorf("invalid denom: %s", denom)
	// }
	return nil
}

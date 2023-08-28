package types_test

import (
	"testing"

	"github.com/QuadrateOrg/core/x/oracle/types"

	"github.com/stretchr/testify/require"
)

func TestParseExchangeRateTuples(t *testing.T) {
	valid := "123.0uqube,123.123demo"
	_, err := types.ParseExchangeRateTuples(valid)
	require.NoError(t, err)

	duplicatedDenom := "100.0uqube,123.123demo,121233.123demo"
	_, err = types.ParseExchangeRateTuples(duplicatedDenom)
	require.Error(t, err)

	invalidCoins := "123.123"
	_, err = types.ParseExchangeRateTuples(invalidCoins)
	require.Error(t, err)

	invalidCoinsWithValid := "123.0uqube,123.1"
	_, err = types.ParseExchangeRateTuples(invalidCoinsWithValid)
	require.Error(t, err)

	abstainCoinsWithValid := "0.0uqube,123.1demo"
	_, err = types.ParseExchangeRateTuples(abstainCoinsWithValid)
	require.NoError(t, err)
}

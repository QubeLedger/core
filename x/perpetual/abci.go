package perpetual

import (
	"fmt"
	"time"

	"github.com/QuadrateOrg/core/x/perpetual/keeper"
	"github.com/QuadrateOrg/core/x/perpetual/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper) error {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)
	for _, vault := range k.GetAllVault(ctx) {
		price_index, err := k.OracleKeeper.GetExchangeRate(ctx, vault.OracleAssetId)
		if err != nil {
			continue
		}

		price_perpetual := (vault.X.ToDec()).Quo(vault.Y.ToDec())

		funding_rate := (price_perpetual.Sub(price_index)).QuoInt64(24)
		fmt.Printf("%v\n", funding_rate)

		take_funding_payment_sum := sdk.NewDec(0)
		sum_all_position := sdk.NewDec(0)

		if funding_rate.GT(sdk.NewDec(0)) {
			for _, long_position_id := range vault.LongPositionId {
				position, _ := k.GetPositionByPositionId(ctx, long_position_id)
				take_funding_payment_sum = take_funding_payment_sum.Add(
					(position.ReturnAmount).Mul(funding_rate),
				)
			}

			for _, short_position_id := range vault.ShortPositionId {
				position, _ := k.GetPositionByPositionId(ctx, short_position_id)
				sum_all_position = sum_all_position.Add(
					(position.ReturnAmount),
				)
			}

			for _, short_position_id := range vault.ShortPositionId {
				position, _ := k.GetPositionByPositionId(ctx, short_position_id)
				plus_return_valut := ((position.ReturnAmount).Quo(sum_all_position)).Mul(take_funding_payment_sum)
				position.ReturnAmount = position.ReturnAmount.Add(plus_return_valut)
				k.SetPosition(ctx, position)
			}

		} else {
			for _, short_position_id := range vault.ShortPositionId {
				position, _ := k.GetPositionByPositionId(ctx, short_position_id)
				take_funding_payment_sum = take_funding_payment_sum.Add(
					(position.ReturnAmount).Mul(funding_rate),
				)
			}
			take_funding_payment_sum = take_funding_payment_sum.MulInt64(-1)

			for _, long_position_id := range vault.LongPositionId {
				position, _ := k.GetPositionByPositionId(ctx, long_position_id)
				sum_all_position = sum_all_position.Add(
					(position.ReturnAmount),
				)
			}

			for _, long_position_id := range vault.LongPositionId {
				position, _ := k.GetPositionByPositionId(ctx, long_position_id)
				plus_return_valut := ((position.ReturnAmount).Quo(sum_all_position)).Mul(take_funding_payment_sum)
				position.ReturnAmount = position.ReturnAmount.Add(plus_return_valut)
				k.SetPosition(ctx, position)
			}
		}

	}
	return nil
}

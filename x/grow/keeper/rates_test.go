package keeper_test

import "fmt"

func (s *GrowKeeperTestSuite) TestGetRatesByUtilizationRate() {
	s.Setup()

	custom_params := s.app.GrowKeeper.GetParams(s.ctx)
	custom_params.UStaticVolatile = 60
	custom_params.MaxRateVolatile = 300
	custom_params.Slope = 7

	s.app.GrowKeeper.SetParams(s.ctx, custom_params)

	s.app.GrowKeeper.AppendAsset(s.ctx, s.GetNormalAsset(0))

	asset, _ := s.app.GrowKeeper.GetAssetByAssetId(s.ctx, s.GetNormalAsset(0).AssetId)
	utilization_rate := (float64(60) / float64(100))
	bir, sir, _ := s.app.GrowKeeper.GetRatesByUtilizationRate(s.ctx, utilization_rate, asset)

	fmt.Printf("borrow_interest_rate: %f\n", bir)
	fmt.Printf("supply_interest_rate: %f\n", sir)
}

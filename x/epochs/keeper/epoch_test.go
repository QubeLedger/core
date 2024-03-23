package keeper_test

import (
	"time"

	"github.com/QuadrateOrg/core/x/epochs/types"
)

func (s *KeeperTestSuite) TestAddEpochInfo() {
	defaultIdentifier := "default_add_epoch_info_id"
	defaultDuration := time.Hour
	startBlockHeight := int64(100)
	startBlockTime := time.Unix(1656907200, 0).UTC()
	tests := map[string]struct {
		addedEpochInfo types.EpochInfo
		expErr         bool
		expEpochInfo   types.EpochInfo
	}{
		"simple_add": {
			addedEpochInfo: types.EpochInfo{
				Identifier:              defaultIdentifier,
				StartTime:               time.Time{},
				Duration:                defaultDuration,
				CurrentEpoch:            0,
				CurrentEpochStartHeight: 0,
				CurrentEpochStartTime:   time.Time{},
				EpochCountingStarted:    false,
			},
			expErr: false,
			expEpochInfo: types.EpochInfo{
				Identifier:              defaultIdentifier,
				StartTime:               startBlockTime,
				Duration:                defaultDuration,
				CurrentEpoch:            0,
				CurrentEpochStartHeight: startBlockHeight,
				CurrentEpochStartTime:   time.Time{},
				EpochCountingStarted:    false,
			},
		},
		"zero_duration": {
			addedEpochInfo: types.EpochInfo{
				Identifier:              defaultIdentifier,
				StartTime:               time.Time{},
				Duration:                time.Duration(0),
				CurrentEpoch:            0,
				CurrentEpochStartHeight: 0,
				CurrentEpochStartTime:   time.Time{},
				EpochCountingStarted:    false,
			},
			expErr: true,
		},
	}

	for name, test := range tests {
		//nolint:scopelint,testfile
		s.Run(name, func() {
			s.Setup()
			s.Ctx = s.Ctx.WithBlockHeight(startBlockHeight).WithBlockTime(startBlockTime)
			err := s.app.EpochsKeeper.AddEpochInfo(s.Ctx, test.addedEpochInfo)
			if !test.expErr {
				s.Require().NoError(err)
				actualEpochInfo := s.app.EpochsKeeper.GetEpochInfo(s.Ctx, test.addedEpochInfo.Identifier)
				s.Require().Equal(test.expEpochInfo, actualEpochInfo)
			} else {
				s.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestDuplicateAddEpochInfo() {
	identifier := "duplicate_add_epoch_info"
	epochInfo := types.NewGenesisEpochInfo(identifier, time.Hour*24*30)
	err := suite.app.EpochsKeeper.AddEpochInfo(suite.Ctx, epochInfo)
	suite.Require().NoError(err)
	err = suite.app.EpochsKeeper.AddEpochInfo(suite.Ctx, epochInfo)
	suite.Require().Error(err)
}

func (suite *KeeperTestSuite) TestEpochLifeCycle() {
	suite.Setup()

	epochInfo := types.NewGenesisEpochInfo("monthly", time.Hour*24*30)
	err := suite.app.EpochsKeeper.AddEpochInfo(suite.Ctx, epochInfo)
	suite.Require().NoError(err)
	epochInfoSaved := suite.app.EpochsKeeper.GetEpochInfo(suite.Ctx, "monthly")
	// setup expected epoch info
	expectedEpochInfo := epochInfo
	expectedEpochInfo.StartTime = suite.Ctx.BlockTime()
	expectedEpochInfo.CurrentEpochStartHeight = suite.Ctx.BlockHeight()
	suite.Require().Equal(expectedEpochInfo, epochInfoSaved)

	allEpochs := suite.app.EpochsKeeper.AllEpochInfos(suite.Ctx)
	suite.Require().Len(allEpochs, 4)
	suite.Require().Equal(allEpochs[0].Identifier, "day") // alphabetical order
	suite.Require().Equal(allEpochs[1].Identifier, "hour")
	suite.Require().Equal(allEpochs[2].Identifier, "monthly")
	suite.Require().Equal(allEpochs[3].Identifier, "week")
}

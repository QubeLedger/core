package keeper_test

import (
	"github.com/QuadrateOrg/core/x/liquidstakeibc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	icatypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/types"
	connectiontypes "github.com/cosmos/ibc-go/v4/modules/core/03-connection/types"
	commitmenttypes "github.com/cosmos/ibc-go/v4/modules/core/23-commitment/types"
	ibcexported "github.com/cosmos/ibc-go/v4/modules/core/exported"
	solomachinetypes "github.com/cosmos/ibc-go/v4/modules/light-clients/06-solomachine/types"
)

func (suite *IntegrationTestSuite) TestGetSetParams() {
	tc := []struct {
		name     string
		params   types.Params
		expected types.Params
	}{
		{
			name: "normal params",
			params: types.Params{
				AdminAddress: "qube1gmywgwu442ttkqgjl2r9pygum9zrlcv5rmcz20",
				FeeAddress:   "qube12qydd0w5ff4sww54dxm0sreznxlex8wfq0wzkd",
			},
			expected: types.Params{
				AdminAddress: "qube1gmywgwu442ttkqgjl2r9pygum9zrlcv5rmcz20",
				FeeAddress:   "qube12qydd0w5ff4sww54dxm0sreznxlex8wfq0wzkd",
			},
		},
	}

	for _, t := range tc {
		suite.Run(t.name, func() {
			pstakeApp, ctx := suite.app, suite.ctx

			pstakeApp.LiquidStakeIBCKeeper.SetParams(ctx, t.params)
			params := pstakeApp.LiquidStakeIBCKeeper.GetParams(ctx)
			suite.Require().Equal(params, t.expected)
		})
	}
}

func (suite *IntegrationTestSuite) TestSendProtocolFee() {
	tc := []struct {
		name       string
		fee        sdk.Coins
		module     string
		feeAddress string
		success    bool
	}{
		{
			name:       "successful case",
			fee:        sdk.Coins{sdk.Coin{Denom: MintDenom, Amount: sdk.NewInt(100)}},
			module:     types.ModuleName,
			feeAddress: FeeAddress,
			success:    true,
		},
		{
			name:       "invalid fee address",
			fee:        sdk.Coins{sdk.Coin{Denom: MintDenom, Amount: sdk.NewInt(100)}},
			module:     types.ModuleName,
			feeAddress: "1234",
			success:    false,
		},
		{
			name:       "not enough tokens",
			fee:        sdk.Coins{sdk.Coin{Denom: MintDenom, Amount: sdk.NewInt(1000)}},
			module:     types.ModuleName,
			feeAddress: FeeAddress,
			success:    false,
		},
	}

	//apptypes.SetConfig()
	suite.SetupTest()

	app, ctx := suite.app, suite.ctx
	hc, found := app.LiquidStakeIBCKeeper.GetHostChain(ctx, suite.chainB.ChainID)
	suite.Require().Equal(found, true)

	for _, t := range tc {
		suite.Run(t.name, func() {
			if t.success {

				baseFee := sdk.NewInt64Coin(hc.MintDenom(), 100)
				err := suite.FundModuleAccount(ctx, app.BankKeeper, t.module, sdk.NewCoins(baseFee))
				suite.Require().NoError(err)

				suite.Require().NoError(
					app.LiquidStakeIBCKeeper.SendProtocolFee(
						ctx,
						t.fee,
						t.module,
						t.feeAddress,
					),
				)

				feeAddress, _ := sdk.AccAddressFromBech32(t.feeAddress)
				currBalance := app.BankKeeper.GetBalance(ctx, feeAddress, hc.MintDenom())
				suite.Require().Equal(baseFee, currBalance)
			} else {
				suite.Require().Error(
					app.LiquidStakeIBCKeeper.SendProtocolFee(
						ctx,
						t.fee,
						t.module,
						t.feeAddress,
					),
				)
			}
		})
	}
}

func (suite *IntegrationTestSuite) TestDelegateAccountPortOwner() {
	pstakeApp, ctx := suite.app, suite.ctx
	hc, found := pstakeApp.LiquidStakeIBCKeeper.GetHostChain(ctx, suite.chainB.ChainID)
	suite.Require().Equal(found, true)

	suite.Require().Equal(
		hc.DelegationAccount.Owner,
		hc.ChainId+"."+types.DelegateICAType,
	)
}

func (suite *IntegrationTestSuite) TestRewardsAccountPortOwner() {
	pstakeApp, ctx := suite.app, suite.ctx
	hc, found := pstakeApp.LiquidStakeIBCKeeper.GetHostChain(ctx, suite.chainB.ChainID)
	suite.Require().Equal(found, true)

	suite.Require().Equal(
		hc.RewardsAccount.Owner,
		hc.ChainId+"."+types.RewardsICAType,
	)
}

func (suite *IntegrationTestSuite) TestGetEpochNumber() {
	pstakeApp, ctx := suite.app, suite.ctx

	suite.Require().Equal(
		pstakeApp.LiquidStakeIBCKeeper.GetEpochNumber(ctx, types.DelegationEpoch),
		pstakeApp.EpochsKeeper.GetEpochInfo(ctx, types.DelegationEpoch).CurrentEpoch,
	)
}

func (suite *IntegrationTestSuite) TestGetClientState() {
	app, ctx := suite.app, suite.ctx

	// check client state
	state, err := app.LiquidStakeIBCKeeper.GetClientState(ctx, suite.transferPathAB.EndpointA.ConnectionID)
	suite.Require().NoError(err)
	suite.Require().Equal(ibcexported.Tendermint, state.ClientType())

	// check localhost client exists
	//state, err = app.LiquidStakeIBCKeeper.GetClientState(ctx, ibcexported.Localhost)
	//suite.Require().NoError(err)
	//suite.Require().Equal(ibcexported.Localhost, state.ClientType())

	// no connection found
	_, err = app.LiquidStakeIBCKeeper.GetClientState(ctx, "connection-2")
	suite.Require().Error(err)

	// set connection without an active client-id
	app.IBCKeeper.ConnectionKeeper.SetConnection(ctx, "connection-2", connectiontypes.ConnectionEnd{ClientId: "client-1"})
	_, err = app.LiquidStakeIBCKeeper.GetClientState(ctx, "connection-2")
	suite.Require().Error(err)
}

func (suite *IntegrationTestSuite) TestGetChainID() {
	pstakeApp, ctx := suite.app, suite.ctx

	chainID, err := pstakeApp.LiquidStakeIBCKeeper.GetChainID(ctx, suite.transferPathAB.EndpointA.ConnectionID)
	suite.Require().NoError(err)
	suite.Require().Equal(suite.chainB.ChainID, chainID)

	// random type of client not supported
	solomachinetypes.RegisterInterfaces(pstakeApp.InterfaceRegistry())
	pstakeApp.IBCKeeper.ClientKeeper.SetClientState(ctx, "client-1", &solomachinetypes.ClientState{ConsensusState: &solomachinetypes.ConsensusState{}})
	pstakeApp.IBCKeeper.ConnectionKeeper.SetConnection(ctx, "connection-2", connectiontypes.NewConnectionEnd(connectiontypes.OPEN, "client-1", connectiontypes.NewCounterparty("--", "--", commitmenttypes.NewMerklePrefix([]byte("New"))), nil, 1))
	_, err = pstakeApp.LiquidStakeIBCKeeper.GetChainID(ctx, "connection-2")
	suite.Require().Error(err)

	// connection not found
	_, err = pstakeApp.LiquidStakeIBCKeeper.GetChainID(ctx, "connection-3")
	suite.Require().Error(err)
}

func (suite *IntegrationTestSuite) TestGetPortID() {
	portID := suite.app.LiquidStakeIBCKeeper.GetPortID("owner")
	suite.Require().Equal(icatypes.PortPrefix+"owner", portID)
}

func (suite *IntegrationTestSuite) TestRegisterICAAccount() {
	pstakeApp, ctx := suite.app, suite.ctx
	err := pstakeApp.LiquidStakeIBCKeeper.RegisterICAAccount(ctx, suite.transferPathAC.EndpointA.ConnectionID, types.DefaultDelegateAccountPortOwner(suite.chainB.ChainID))
	suite.Require().NoError(err)
}

func (suite *IntegrationTestSuite) TestSetWithdrawAddress() {
	app, ctx := suite.app, suite.ctx
	hc, found := app.LiquidStakeIBCKeeper.GetHostChain(ctx, suite.chainB.ChainID)
	suite.Require().Equal(true, found)
	suite.Require().NotNil(hc)

	err := app.LiquidStakeIBCKeeper.SetWithdrawAddress(ctx, hc)
	suite.Require().NoError(err)

	hc2 := hc
	hc2.ConnectionId = "connection-3"
	err = app.LiquidStakeIBCKeeper.SetWithdrawAddress(ctx, hc2)
	suite.Require().Error(err)
}

func (suite *IntegrationTestSuite) TestIsICAChannelActive() {
	pstakeApp, ctx := suite.app, suite.ctx
	hc, found := pstakeApp.LiquidStakeIBCKeeper.GetHostChain(ctx, suite.chainB.ChainID)
	suite.Require().Equal(true, found)
	suite.Require().NotNil(hc)

	active := pstakeApp.LiquidStakeIBCKeeper.IsICAChannelActive(ctx, hc, pstakeApp.LiquidStakeIBCKeeper.GetPortID(hc.DelegationAccount.Owner))
	suite.Require().Equal(true, active)
}

func (suite *IntegrationTestSuite) TestSendICATransfer() {
	pstakeApp, ctx := suite.app, suite.ctx
	hc, found := pstakeApp.LiquidStakeIBCKeeper.GetHostChain(ctx, suite.chainB.ChainID)
	suite.Require().Equal(true, found)
	suite.Require().NotNil(hc)

	_, err := pstakeApp.LiquidStakeIBCKeeper.SendICATransfer(
		ctx,
		hc,
		sdk.NewInt64Coin(hc.HostDenom, 10),
		hc.DelegationAccount.Address,
		authtypes.NewModuleAddress(types.UndelegationModuleAccount).String(),
		hc.DelegationAccount.Owner,
	)
	suite.Require().NoError(err)

	_, err = pstakeApp.LiquidStakeIBCKeeper.SendICATransfer(
		ctx,
		hc,
		sdk.NewInt64Coin(hc.HostDenom, 10),
		hc.DelegationAccount.Address,
		authtypes.NewModuleAddress(types.UndelegationModuleAccount).String(),
		"invalid owner",
	)
	suite.Require().Error(err)

	hc2 := hc
	hc2.PortId = ""
	_, err = pstakeApp.LiquidStakeIBCKeeper.SendICATransfer(ctx, hc2, sdk.NewInt64Coin(hc.HostDenom, 10),
		hc.DelegationAccount.Address, authtypes.NewModuleAddress(types.UndelegationModuleAccount).String(),
		hc.DelegationAccount.Owner)
	suite.Require().Error(err)
}

func (suite *IntegrationTestSuite) TestUpdateCValue() {
	pstakeApp, ctx := suite.app, suite.ctx
	hc, found := pstakeApp.LiquidStakeIBCKeeper.GetHostChain(ctx, suite.chainB.ChainID)
	suite.Require().Equal(true, found)
	suite.Require().NotNil(hc)

	suite.Require().NotPanics(func() { pstakeApp.LiquidStakeIBCKeeper.UpdateCValue(ctx, hc) })

	{
		epoch := pstakeApp.EpochsKeeper.GetEpochInfo(suite.chainA.GetContext(), types.DelegationEpoch)
		suite.NotNil(epoch)
		err := pstakeApp.LiquidStakeIBCKeeper.BeforeEpochStart(suite.chainA.GetContext(), epoch.Identifier, epoch.CurrentEpoch)
		suite.Require().NoError(err)

		senderAcc := suite.chainA.SenderAccount
		// user liquidstakes
		msgLiquidStake := types.NewMsgLiquidStake(sdk.NewInt64Coin(hc.IBCDenom(), 1000000), senderAcc.GetAddress())
		result, err := suite.app.MsgServiceRouter().Handler(msgLiquidStake)(suite.chainA.GetContext(), msgLiquidStake)
		suite.NotNil(result)
		suite.NoError(err)
	}
	suite.Require().NotPanics(func() { pstakeApp.LiquidStakeIBCKeeper.UpdateCValue(ctx, hc) })

	// lower limits so that chain goes out of limits
	hc.Params.UpperCValueLimit = sdk.MustNewDecFromStr("0.5")
	pstakeApp.LiquidStakeIBCKeeper.SetHostChain(ctx, hc)
	suite.Require().NotPanics(func() { pstakeApp.LiquidStakeIBCKeeper.UpdateCValue(ctx, hc) })
	hc, _ = pstakeApp.LiquidStakeIBCKeeper.GetHostChain(ctx, suite.chainB.ChainID)
	suite.Require().Equal(false, hc.Active)
}

func (suite *IntegrationTestSuite) TestRecalculateCValueLimits() {
	pstakeApp, ctx := suite.app, suite.ctx
	hc, found := pstakeApp.LiquidStakeIBCKeeper.GetHostChain(ctx, suite.chainB.ChainID)
	suite.Require().Equal(true, found)
	suite.Require().NotNil(hc)
	suite.Require().Equal(sdk.NewDec(1), hc.CValue)

	pstakeApp.LiquidStakeIBCKeeper.RecalculateCValueLimits(ctx, hc, sdk.NewInt(0), sdk.NewInt(0))
	suite.Require().Equal("0.950000000000000000", hc.Params.LowerCValueLimit.String())
	suite.Require().Equal("1.050000000000000000", hc.Params.UpperCValueLimit.String())

	hc.Validators[0].DelegatedAmount = sdk.NewInt(1000)
	pstakeApp.LiquidStakeIBCKeeper.SetHostChainValidator(ctx, hc, hc.Validators[0])

	pstakeApp.LiquidStakeIBCKeeper.RecalculateCValueLimits(ctx, hc, sdk.NewInt(1000), sdk.NewInt(1000))

	expectedLower := sdk.NewDec(1000).Quo(sdk.NewDec(1000).
		Add(sdk.NewDec(1000).Mul(hc.AutoCompoundFactor).Mul(sdk.NewDec(types.CValueDynamicLowerDiff))))

	suite.Require().Equal(expectedLower, hc.Params.LowerCValueLimit)

	expectedUpper := hc.CValue.Add(hc.CValue.Sub(hc.Params.LowerCValueLimit).Mul(sdk.NewDec(types.CValueDynamicUpperDiff)))

	suite.Require().Equal(expectedUpper, hc.Params.UpperCValueLimit)
}

func (s *IntegrationTestSuite) FundAccount(ctx sdk.Context, addr sdk.AccAddress, amounts sdk.Coins) error {
	if err := s.app.BankKeeper.MintCoins(ctx, minttypes.ModuleName, amounts); err != nil {
		return err
	}

	return s.app.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, addr, amounts)
}

func (s *IntegrationTestSuite) FundModuleAccount(ctx sdk.Context, bankKeeper bankkeeper.Keeper, recipientMod string, amounts sdk.Coins) error {
	if err := bankKeeper.MintCoins(ctx, recipientMod, amounts); err != nil {
		return err
	}

	s.Commit()

	return nil

	//return bankKeeper.SendCoinsFromModuleToAccount(ctx, "grow", s.transferPathAB.EndpointB.Chain.SenderAccount.GetAddress(), amounts)
}

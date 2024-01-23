package types

const (
	AttributeValueCategory                 = ModuleName
	AttributeKeyActionDeposit              = "grow_deposit"
	AttributeKeyActionWithdrawal           = "grow_withdrawal"
	AttributeKeyActionCreateLend           = "create_lend"
	AttributeKeyActionDeleteLend           = "delete_lend"
	AttributeKeyActionDepositColletaral    = "deposit_collateral"
	AttributeKeyActionWithdrawalColletaral = "withdrawal_collateral"
	AttributeKeyActionCreateLiqPosition    = "create_liquidation_position"
	AttributeKeyActionCloseLiqPosition     = "close_liquidation_position"

	EventRegisterLendAssetProposal                       = "register_lend_asset_proposal"
	EventRegisterGTokenPairProposal                      = "register_gToken_pair_proposal"
	EventRegisterChangeGrowYieldReserveAddressProposal   = "register_change_grow_yield_reserve_address_proposal"
	EventRegisterChangeUSQReserveAddressProposal         = "register_change_usq_reserve_address_proposal"
	EventRegisterChangeGrowStakingReserveAddressProposal = "register_change_grow_staking_reserve_address_proposal"
	EventRegisterChangeRealRateProposal                  = "register_change_real_rate_proposal"
	EventRegisterChangeBorrowRateProposal                = "register_change_borrow_rate_proposal"
	EventRegisterRemoveLendAssetProposal                 = "register_remove_lend_asset_proposal"
	EventRegisterRemoveGTokenPairProposal                = "register_remove_gToken_pair_proposal"

	EventRegisterChangeDepositMethodStatusProposal    = "register_change_deposit_method_proposal"
	EventRegisterChangeCollateralMethodStatusProposal = "register_change_collateral_method_proposal"
	EventRegisterChangeBorrowMethodStatusProposal     = "register_change_borrow_method_proposal"
)

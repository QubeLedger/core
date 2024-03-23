package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidDenom             = sdkerrors.Register(ModuleName, 2000, "invalid token denom")
	ErrInvalidHostChain         = sdkerrors.Register(ModuleName, 2001, "host chain not registered")
	ErrMinDeposit               = sdkerrors.Register(ModuleName, 2002, "deposit amount less than minimum deposit")
	ErrFailedDeposit            = sdkerrors.Register(ModuleName, 2003, "deposit failed")
	ErrMintFailed               = sdkerrors.Register(ModuleName, 2004, "minting failed")
	ErrRegisterFailed           = sdkerrors.Register(ModuleName, 2005, "host chain register failed")
	ErrDepositNotFound          = sdkerrors.Register(ModuleName, 2007, "deposit record not found")
	ErrICATxFailure             = sdkerrors.Register(ModuleName, 2008, "ica transaction failed")
	ErrInvalidMessages          = sdkerrors.Register(ModuleName, 2009, "not enough messages")
	ErrInvalidResponses         = sdkerrors.Register(ModuleName, 2010, "not enough message responses")
	ErrValidatorNotFound        = sdkerrors.Register(ModuleName, 2011, "validator not found")
	ErrNotEnoughDelegations     = sdkerrors.Register(ModuleName, 2012, "delegated amount is less than undelegation amount requested")
	ErrRedeemFailed             = sdkerrors.Register(ModuleName, 2013, "an error occurred while instant redeeming tokens")
	ErrBurnFailed               = sdkerrors.Register(ModuleName, 2014, "burn failed")
	ErrParsingAmount            = sdkerrors.Register(ModuleName, 2015, "could not parse message amount")
	ErrHostChainInactive        = sdkerrors.Register(ModuleName, 2016, "host chain is not active")
	ErrInvalidLSMDenom          = sdkerrors.Register(ModuleName, 2018, "invalid lsm token denom")
	ErrLSMNotEnabled            = sdkerrors.Register(ModuleName, 2019, "host chain has LSM staking disabled")
	ErrLSMDepositProcessing     = sdkerrors.Register(ModuleName, 2020, "already processing LSM deposit")
	ErrLSMValidatorInvalidState = sdkerrors.Register(ModuleName, 2021, "validator invalid state")
	ErrInsufficientDeposits     = sdkerrors.Register(ModuleName, 2022, "insufficient deposits")
	ErrInvalidSigner            = sdkerrors.Register(ModuleName, 13, "expected authority account as only signer for proposal message")
)

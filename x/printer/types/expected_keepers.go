package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	GetSequence(ctx sdk.Context, addr sdk.AccAddress) (uint64, error)
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
}

type StakingKeeper interface {
	UnbondingTime(ctx sdk.Context) (res time.Duration)
	Delegate(
		ctx sdk.Context, delAddr sdk.AccAddress, bondAmt sdk.Int, tokenSrc stakingtypes.BondStatus,
		validator stakingtypes.Validator, subtractAccount bool,
	) (newShares sdk.Dec, err error)
	BeginRedelegation(
		ctx sdk.Context, delAddr sdk.AccAddress, valSrcAddr, valDstAddr sdk.ValAddress, sharesAmount sdk.Dec,
	) (completionTime time.Time, err error)
	GetValidator(ctx sdk.Context, addr sdk.ValAddress) (validator stakingtypes.Validator, found bool)
	BondDenom(ctx sdk.Context) (res string)
	ValidateUnbondAmount(
		ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress, amt sdk.Int,
	) (shares sdk.Dec, err error)
	RemoveRedelegation(ctx sdk.Context, red stakingtypes.Redelegation)
	Unbond(
		ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress, shares sdk.Dec,
	) (amount sdk.Int, err error)
	GetDelegation(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) (delegation stakingtypes.Delegation, found bool)
	TotalBondedTokens(ctx sdk.Context) sdk.Int
	GetDelegatorBonded(ctx sdk.Context, delegator sdk.AccAddress) sdk.Int
	RemoveValidatorTokensAndShares(ctx sdk.Context, validator stakingtypes.Validator,
		sharesToRemove sdk.Dec,
	) (valOut stakingtypes.Validator, removedTokens sdk.Int)
	RemoveValidatorTokens(ctx sdk.Context,
		validator stakingtypes.Validator, tokensToRemove sdk.Int,
	) stakingtypes.Validator
	IterateDelegatorDelegations(ctx sdk.Context, delegator sdk.AccAddress, cb func(delegation stakingtypes.Delegation) (stop bool))
	GetAllValidators(ctx sdk.Context) (validators []stakingtypes.Validator)
}

package keeper

import (
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/QuadrateOrg/core/x/liquidstakeibc/types"
)

type delegation struct {
	validator        string
	ideal            sdk.Int
	delegation       sdk.Int
	diff             sdk.Int
	validatorDetails types.Validator
}

/* #nosec */
func (k Keeper) GenerateRedelegateMsgs(ctx sdk.Context, hc types.HostChain) []sdk.Msg {
	AcceptableDelta := hc.Params.RedelegationAcceptableDelta // #nosec
	MaxRedelegationEntries := hc.Params.MaxEntries
	sum := sdk.ZeroInt()
	for _, validator := range hc.Validators {
		sum = sum.Add(validator.DelegatedAmount)
	}

	idealDelegationList := make([]delegation, len(hc.Validators))
	sum2 := sdk.ZeroInt()
	for i, validator := range hc.Validators {
		idealAmt := validator.Weight.MulInt(sum).TruncateInt()
		// last element
		if i == len(hc.Validators)-1 {
			idealAmt = sum.Sub(sum2)
		}
		sum2 = sum2.Add(idealAmt)
		idealDelegationList[i] = delegation{
			validator:        validator.OperatorAddress,
			ideal:            idealAmt,
			delegation:       validator.DelegatedAmount,
			diff:             validator.DelegatedAmount.Sub(idealAmt),
			validatorDetails: *validator,
		}
	}
	// negative diffs first, so ascending
	idealDelegationList = sortDelegationListAsc(idealDelegationList)
	revIdealList := make([]delegation, len(idealDelegationList))
	copy(revIdealList, idealDelegationList)
	// positive diffs first (descending)
	//Reverse(revIdealList)
	for i, j := 0, len(revIdealList)-1; i < j; i, j = i+1, j-1 {
		revIdealList[i], revIdealList[j] = revIdealList[j], revIdealList[i]
	}
	redelegations, ok := k.GetRedelegations(ctx, hc.ChainId)
	if !ok {
		redelegations = &types.Redelegations{
			ChainID:       hc.ChainId,
			Redelegations: []*stakingtypes.Redelegation{},
		}
	}

	var msgs []sdk.Msg
L1:
	for i := range revIdealList {
		if revIdealList[i].diff.LT(AcceptableDelta) {
			break L1
		}
		// RedelegationExistsToValidator: This is not updated inside the loop (with newer msgs), so some ICA redelegate txns might fail, and it is ok.
		if !k.RedelegationExistsToValidator(redelegations.Redelegations, revIdealList[i].validator) {
			// re-sort idealDelegationAsc
			idealDelegationList = sortDelegationListAsc(idealDelegationList)
		L2:
			for j := range idealDelegationList {
				if revIdealList[i].validator == idealDelegationList[j].validator {
					break L1
				}
				if revIdealList[i].diff.LT(AcceptableDelta) || idealDelegationList[j].diff.IsPositive() {
					break L2
				}
				if !idealDelegationList[j].validatorDetails.Delegable || idealDelegationList[j].validatorDetails.Status != stakingtypes.Bonded.String() {
					continue L2
				}

				// RedelegationFromAToB: This is not updated inside the loop (with newer msgs), so some ICA redelegate txns might fail, and it is ok.
				_, numEntries := k.RedelegationFromAToB(redelegations.Redelegations, revIdealList[i].validator, idealDelegationList[j].validator)
				if numEntries < MaxRedelegationEntries {
					redelegationAmt := sdk.MinInt(revIdealList[i].diff.Abs(), idealDelegationList[j].diff.Abs())
					redelegateMsg := &stakingtypes.MsgBeginRedelegate{
						DelegatorAddress:    hc.DelegationAccount.Address,
						ValidatorSrcAddress: revIdealList[i].validator,
						ValidatorDstAddress: idealDelegationList[j].validator,
						Amount:              sdk.NewCoin(hc.HostDenom, redelegationAmt),
					}
					msgs = append(msgs, redelegateMsg)
					revIdealList[i].diff = revIdealList[i].diff.Sub(redelegationAmt)
					idealDelegationList[j].diff = idealDelegationList[j].diff.Add(redelegationAmt)
				}
			}
		}
	}
	return msgs
}

func (k Keeper) RedelegationExistsToValidator(redelegations []*stakingtypes.Redelegation, toValoper string) bool {
	for _, redelegation := range redelegations {
		if redelegation.ValidatorDstAddress == toValoper && len(redelegation.Entries) > 0 {
			return true
		}
	}
	return false
}

/* #nosec */
func (k Keeper) RedelegationFromAToB(redelegations []*stakingtypes.Redelegation, fromValoper, toValoper string) (bool, uint32) {
	for _, redelegation := range redelegations {
		if redelegation.ValidatorDstAddress == toValoper && redelegation.ValidatorSrcAddress == fromValoper {
			return true, uint32(len(redelegation.Entries))
		}
	}
	return false, 0
}

func sortDelegationListAsc(idealDelegationList []delegation) []delegation {
	sort.SliceStable(idealDelegationList, func(i, j int) bool {
		switch {
		case idealDelegationList[i].diff.LT(idealDelegationList[j].diff):
			return true
		case idealDelegationList[i].diff.GT(idealDelegationList[j].diff):
			return false
		default:
			return idealDelegationList[i].validator < idealDelegationList[j].validator
		}
	})
	return idealDelegationList
}

// remove when go updates to 1.21, and use slices package.
// Reverse reverses the elements of the slice in place.
/* #nosec */
func Reverse[S ~[]E, E any](s S) { // #nosec
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

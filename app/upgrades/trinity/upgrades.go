package trinity

import (
	growmodulekeeper "github.com/QuadrateOrg/core/x/grow/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

/* #nosec */
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	growkeeper growmodulekeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		migrations, err := mm.RunMigrations(ctx, configurator, fromVM)
		if err != nil {
			return nil, err
		}

		all_position := growkeeper.GetAllPosition(ctx)

		ctx.Logger().Info(`
		.:^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^:.
		.~?5GBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBG5?~.
	    .7PB#BG5J????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????J5GB#BP7.
	    ^PBBBJ^.                                                                                                                                .^JBBBP^
	    :GBBP^                                                                                                                                      ^PBBG:
	    5BBG.					   ████████╗██████╗░██╗███╗░░██╗██╗████████╗██╗░░░██╗						 .GBB5
	    5BBG.					   ╚══██╔══╝██╔══██╗██║████╗░██║██║╚══██╔══╝╚██╗░██╔╝						 .GBB5
	    5BBG.					   ░░░██║░░░██████╔╝██║██╔██╗██║██║░░░██║░░░░╚████╔╝░						 .GBB5
	    5BBG.					   ░░░██║░░░██╔══██╗██║██║╚████║██║░░░██║░░░░░╚██╔╝░░						 .GBB5
	    5BBG.					   ░░░██║░░░██║░░██║██║██║░╚███║██║░░░██║░░░░░░██║░░░						 .GBB5
            5BBG.					   ░░░╚═╝░░░╚═╝░░╚═╝╚═╝╚═╝░░╚══╝╚═╝░░░╚═╝░░░░░░╚═╝░░░						 .GBB5
	    :GBBP^                                                                                                                                      ^PBBG:
	    ^PBBBJ^.                                                                                                                                .^JBBBP^
	    .7PB#BG5J????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????J5GB#BP7.
		.~?5GBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBG5?~.
		    .::^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^::.
		`)

		for _, pos := range all_position {
			for _, lend_id := range pos.LendId {
				lend, _ := growkeeper.GetLendByLendId(ctx, lend_id)
				growkeeper.RemoveLend(ctx, lend.Id)
			}
			for _, loan_id := range pos.LendId {
				loan, _ := growkeeper.GetLoadByLoanId(ctx, loan_id)
				growkeeper.RemoveLoan(ctx, loan.Id)
			}

			ctx.Logger().Info("qLabs: Trinity: remove: position: id: %v", pos.DepositId)

			growkeeper.RemovePosition(ctx, pos.Id)
		}

		ctx.Logger().Info("qLabs: Trinity: remove: all position")

		return migrations, nil
	}
}

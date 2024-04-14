package app

import (
	"fmt"
	"io"
	stdlog "log"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authrest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/mint"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ica "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts"
	icacontroller "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/controller"
	icacontrollerkeeper "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/controller/keeper"
	icacontrollertypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/controller/types"
	icahost "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/host"
	icahostkeeper "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/host/keeper"
	icahosttypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/host/types"
	icatypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/types"
	"github.com/cosmos/ibc-go/v4/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v4/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v4/modules/core"
	ibcclient "github.com/cosmos/ibc-go/v4/modules/core/02-client"
	ibcclientclient "github.com/cosmos/ibc-go/v4/modules/core/02-client/client"
	ibcclienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	porttypes "github.com/cosmos/ibc-go/v4/modules/core/05-port/types"
	ibchost "github.com/cosmos/ibc-go/v4/modules/core/24-host"
	ibckeeper "github.com/cosmos/ibc-go/v4/modules/core/keeper"
	"github.com/gorilla/mux"
	"github.com/rakyll/statik/fs"
	"github.com/spf13/cast"
	abci "github.com/tendermint/tendermint/abci/types"
	tmjson "github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"

	packetforward "github.com/strangelove-ventures/packet-forward-middleware/v4/router"
	packetforwardkeeper "github.com/strangelove-ventures/packet-forward-middleware/v4/router/keeper"
	packetforwardtypes "github.com/strangelove-ventures/packet-forward-middleware/v4/router/types"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmclient "github.com/CosmWasm/wasmd/x/wasm/client"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/prometheus/client_golang/prometheus"

	quadrateante "github.com/QuadrateOrg/core/ante"
	quadrateappparams "github.com/QuadrateOrg/core/app/params"

	wasmbinding "github.com/QuadrateOrg/core/wasmbinding"

	tokenfactory "github.com/QuadrateOrg/core/x/tokenfactory"
	tokenfactorykeeper "github.com/QuadrateOrg/core/x/tokenfactory/keeper"
	tokenfactorytypes "github.com/QuadrateOrg/core/x/tokenfactory/types"

	tfupgrades "github.com/QuadrateOrg/core/app/upgrades/TF"
	v4 "github.com/QuadrateOrg/core/app/upgrades/v1/v4"
	v4rc0 "github.com/QuadrateOrg/core/app/upgrades/v1/v4rc0"
	v5 "github.com/QuadrateOrg/core/app/upgrades/v1/v5"
	v0 "github.com/QuadrateOrg/core/app/upgrades/v2/v0"
	v2 "github.com/QuadrateOrg/core/app/upgrades/v2/v0"
	v1 "github.com/QuadrateOrg/core/app/upgrades/v2/v1"
	v022 "github.com/QuadrateOrg/core/app/upgrades/v2/v2"
	v025 "github.com/QuadrateOrg/core/app/upgrades/v2/v5"
	v025rc0 "github.com/QuadrateOrg/core/app/upgrades/v2/v5rc0"
	v030 "github.com/QuadrateOrg/core/app/upgrades/v3/v0"

	oraclemodule "github.com/QuadrateOrg/core/x/oracle"
	oracleclient "github.com/QuadrateOrg/core/x/oracle/client"
	oraclemodulekeeper "github.com/QuadrateOrg/core/x/oracle/keeper"
	oraclemoduletypes "github.com/QuadrateOrg/core/x/oracle/types"

	stablemodule "github.com/QuadrateOrg/core/x/stable"
	stableclient "github.com/QuadrateOrg/core/x/stable/client"
	stablemodulekeeper "github.com/QuadrateOrg/core/x/stable/keeper"
	stablemoduletypes "github.com/QuadrateOrg/core/x/stable/types"

	growmodule "github.com/QuadrateOrg/core/x/grow"
	growclient "github.com/QuadrateOrg/core/x/grow/client"
	growmodulekeeper "github.com/QuadrateOrg/core/x/grow/keeper"
	growmoduletypes "github.com/QuadrateOrg/core/x/grow/types"

	dexmodule "github.com/QuadrateOrg/core/x/dex"
	dexmodulekeeper "github.com/QuadrateOrg/core/x/dex/keeper"
	dexmoduletypes "github.com/QuadrateOrg/core/x/dex/types"

	swapmiddleware "github.com/QuadrateOrg/core/x/ibcswap"
	swapkeeper "github.com/QuadrateOrg/core/x/ibcswap/keeper"
	swaptypes "github.com/QuadrateOrg/core/x/ibcswap/types"

	epochs "github.com/QuadrateOrg/core/x/epochs"
	epochskeeper "github.com/QuadrateOrg/core/x/epochs/keeper"
	epochstypes "github.com/QuadrateOrg/core/x/epochs/types"

	liquidstakeibc "github.com/QuadrateOrg/core/x/liquidstakeibc"
	liquidstakeibckeeper "github.com/QuadrateOrg/core/x/liquidstakeibc/keeper"
	liquidstakeibctypes "github.com/QuadrateOrg/core/x/liquidstakeibc/types"

	interchainquery "github.com/QuadrateOrg/core/x/interchainquery"
	interchainquerykeeper "github.com/QuadrateOrg/core/x/interchainquery/keeper"
	interchainquerytypes "github.com/QuadrateOrg/core/x/interchainquery/types"

	ibchooker "github.com/QuadrateOrg/core/x/ibchooker"
	ibchookerkeeper "github.com/QuadrateOrg/core/x/ibchooker/keeper"
	ibchookertypes "github.com/QuadrateOrg/core/x/ibchooker/types"

	perpetualmodule "github.com/QuadrateOrg/core/x/perpetual"
	perpetualmodulekeeper "github.com/QuadrateOrg/core/x/perpetual/keeper"
	perpetualmoduletypes "github.com/QuadrateOrg/core/x/perpetual/types"

	gmpmiddleware "github.com/QuadrateOrg/core/x/gmp"

	// unnamed import of statik for swagger UI support
	_ "github.com/cosmos/cosmos-sdk/client/docs/statik"
)

var (
	ProposalsEnabled        = "true"
	EnableSpecificProposals = ""
	HomeDir                 = ".quadrate"
)

// GetEnabledProposals parses the ProposalsEnabled / EnableSpecificProposals values to
// produce a list of enabled proposals to pass into wasmd app.
func GetEnabledProposals() []wasm.ProposalType {
	if EnableSpecificProposals == "" {
		if ProposalsEnabled == "true" {
			return wasm.EnableAllProposals
		}
		return wasm.DisableAllProposals
	}
	chunks := strings.Split(EnableSpecificProposals, ",")
	proposals, err := wasm.ConvertToProposals(chunks)
	if err != nil {
		panic(err)
	}
	return proposals
}

// GetWasmOpts build wasm options
func GetWasmOpts(appOpts servertypes.AppOptions) []wasm.Option {
	var wasmOpts []wasm.Option
	if cast.ToBool(appOpts.Get("telemetry.enabled")) {
		wasmOpts = append(wasmOpts, wasmkeeper.WithVMCacheMetrics(prometheus.DefaultRegisterer))
	}

	return wasmOpts
}

func getGovProposalHandlers() []govclient.ProposalHandler {
	var govProposalHandlers []govclient.ProposalHandler

	govProposalHandlers = wasmclient.ProposalHandlers

	govProposalHandlers = append(govProposalHandlers,
		paramsclient.ProposalHandler,
		distrclient.ProposalHandler,
		upgradeclient.ProposalHandler,
		upgradeclient.CancelProposalHandler,
		ibcclientclient.UpdateClientProposalHandler,
		ibcclientclient.UpgradeProposalHandler,
		stableclient.RegisterPairHandler,
		stableclient.RegisterChangeBurningFundAddressHandler,
		stableclient.RegisterChangeReserveFundAddressHandler,
		stableclient.RegisterDeletePairHandler,
		growclient.RegisterChangeGrowStakingReserveAddressProposalHandler,
		growclient.RegisterChangeGrowYieldReserveAddressProposalHandler,
		growclient.RegisterChangeRealRateProposalHandler,
		growclient.RegisterChangeBorrowRateProposalHandler,
		growclient.RegisterChangeLendRateProposalHandler,
		growclient.RegisterChangeUSQReserveAddressProposalHandler,
		growclient.RegisterGTokenPairProposalHandler,
		growclient.RegisterLendAssetProposalHandler,
		growclient.RegisterChangeDepositMethodStatusProposalHandler,
		growclient.RegisterChangeCollateralMethodStatusProposalHandler,
		growclient.RegisterChangeBorrowMethodStatusProposalHandler,
		oracleclient.RegisterAddNewDenomProposal,
	)

	return govProposalHandlers
}

var (
	// DefaultNodeHome default home directories for the application daemon
	DefaultNodeHome string

	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		genutil.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(
			getGovProposalHandlers()...,
		),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		authzmodule.AppModuleBasic{},
		ibc.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		transfer.AppModuleBasic{},
		vesting.AppModuleBasic{},
		packetforward.AppModuleBasic{},
		ica.AppModuleBasic{},
		wasm.AppModuleBasic{},
		tokenfactory.AppModuleBasic{},
		oraclemodule.AppModuleBasic{},
		stablemodule.AppModuleBasic{},
		growmodule.AppModuleBasic{},
		perpetualmodule.AppModuleBasic{},
		dexmodule.AppModuleBasic{},
		epochs.AppModuleBasic{},
		liquidstakeibc.AppModuleBasic{},
		interchainquery.AppModuleBasic{},
	)

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:                    nil,
		distrtypes.ModuleName:                         nil,
		icatypes.ModuleName:                           nil,
		minttypes.ModuleName:                          {authtypes.Minter},
		stakingtypes.BondedPoolName:                   {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName:                {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:                           {authtypes.Burner},
		ibctransfertypes.ModuleName:                   {authtypes.Minter, authtypes.Burner},
		wasm.ModuleName:                               {authtypes.Burner},
		tokenfactorytypes.ModuleName:                  {authtypes.Minter, authtypes.Burner},
		oraclemoduletypes.ModuleName:                  {authtypes.Minter, authtypes.Burner, authtypes.Staking},
		stablemoduletypes.ModuleName:                  {authtypes.Minter, authtypes.Burner},
		growmoduletypes.ModuleName:                    {authtypes.Minter, authtypes.Burner, authtypes.Staking},
		dexmoduletypes.ModuleName:                     {authtypes.Minter, authtypes.Burner, authtypes.Staking},
		liquidstakeibctypes.ModuleName:                {authtypes.Minter, authtypes.Burner},
		liquidstakeibctypes.DepositModuleAccount:      nil,
		liquidstakeibctypes.UndelegationModuleAccount: {authtypes.Burner},
		perpetualmoduletypes.ModuleName:               {authtypes.Minter, authtypes.Burner},
	}

	receiveAllowedMAcc = map[string]bool{
		liquidstakeibctypes.DepositModuleAccount:      true,
		liquidstakeibctypes.UndelegationModuleAccount: true,
	}
)

var (
	_ servertypes.Application = (*QuadrateApp)(nil)
)

// QuadrateApp extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type QuadrateApp struct { // nolint: golint
	*baseapp.BaseApp
	legacyAmino       *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry

	invCheckPeriod uint

	// keys to access the substores
	keys    map[string]*sdk.KVStoreKey
	tkeys   map[string]*sdk.TransientStoreKey
	memKeys map[string]*sdk.MemoryStoreKey

	// keepers
	AccountKeeper      authkeeper.AccountKeeper
	BankKeeper         bankkeeper.BaseKeeper
	CapabilityKeeper   *capabilitykeeper.Keeper
	StakingKeeper      stakingkeeper.Keeper
	TokenFactoryKeeper *tokenfactorykeeper.Keeper
	SlashingKeeper     slashingkeeper.Keeper
	MintKeeper         mintkeeper.Keeper
	DistrKeeper        distrkeeper.Keeper
	GovKeeper          govkeeper.Keeper
	CrisisKeeper       crisiskeeper.Keeper
	UpgradeKeeper      upgradekeeper.Keeper
	ParamsKeeper       paramskeeper.Keeper
	// IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	IBCKeeper             *ibckeeper.Keeper
	ICAHostKeeper         icahostkeeper.Keeper
	EvidenceKeeper        evidencekeeper.Keeper
	TransferKeeper        ibctransferkeeper.Keeper
	FeeGrantKeeper        feegrantkeeper.Keeper
	AuthzKeeper           authzkeeper.Keeper
	OracleKeeper          oraclemodulekeeper.Keeper
	PacketForwardKeeper   *packetforwardkeeper.Keeper
	StableKeeper          stablemodulekeeper.Keeper
	GrowKeeper            growmodulekeeper.Keeper
	DexKeeper             dexmodulekeeper.Keeper
	SwapKeeper            swapkeeper.Keeper
	EpochsKeeper          *epochskeeper.Keeper
	LiquidStakeIBCKeeper  liquidstakeibckeeper.Keeper
	ICAControllerKeeper   icacontrollerkeeper.Keeper
	InterchainQueryKeeper interchainquerykeeper.Keeper
	TransferHooksKeeper   ibchookerkeeper.Keeper
	PerpetualKeeper       perpetualmodulekeeper.Keeper

	// make scoped keepers public for test purposes
	ScopedIBCKeeper            capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper       capabilitykeeper.ScopedKeeper
	ScopedICAHostKeeper        capabilitykeeper.ScopedKeeper
	ScopedStableKeeper         capabilitykeeper.ScopedKeeper
	ScopedICAControllerKeeper  capabilitykeeper.ScopedKeeper
	ScopedLiquidStakeIBCKeeper capabilitykeeper.ScopedKeeper

	wasmKeeper       wasm.Keeper
	scopedWasmKeeper capabilitykeeper.ScopedKeeper

	// the module manager
	mm *module.Manager

	configurator module.Configurator
}

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		stdlog.Println("Failed to get home dir %2", err)
	}

	DefaultNodeHome = filepath.Join(userHomeDir, HomeDir)
	// apply custom power reduction for 'a' base denom unit 10^18
	sdk.DefaultPowerReduction = sdk.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
}

func NewQuadrateApp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	skipUpgradeHeights map[int64]bool,
	homePath string,
	invCheckPeriod uint,
	encodingConfig quadrateappparams.EncodingConfig,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *QuadrateApp {

	appCodec := encodingConfig.Marshaler
	legacyAmino := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	bApp := baseapp.NewBaseApp(appName, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)

	keys := sdk.NewKVStoreKeys(
		authtypes.StoreKey, banktypes.StoreKey,
		stakingtypes.StoreKey, minttypes.StoreKey,
		distrtypes.StoreKey, slashingtypes.StoreKey,
		govtypes.StoreKey, paramstypes.StoreKey,
		ibchost.StoreKey, upgradetypes.StoreKey,
		evidencetypes.StoreKey, ibctransfertypes.StoreKey,
		capabilitytypes.StoreKey, feegrant.StoreKey,
		authzkeeper.StoreKey, icahosttypes.StoreKey,
		wasm.StoreKey, tokenfactorytypes.StoreKey,
		oraclemoduletypes.StoreKey, packetforwardtypes.StoreKey,
		stablemoduletypes.StoreKey, growmoduletypes.StoreKey,
		dexmoduletypes.StoreKey, epochstypes.StoreKey,
		liquidstakeibctypes.StoreKey, interchainquerytypes.StoreKey,
		icacontrollertypes.StoreKey, perpetualmoduletypes.StoreKey,
	)
	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	app := &QuadrateApp{
		BaseApp:           bApp,
		legacyAmino:       legacyAmino,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		invCheckPeriod:    invCheckPeriod,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
	}

	app.ParamsKeeper = initParamsKeeper(
		appCodec,
		legacyAmino,
		keys[paramstypes.StoreKey],
		tkeys[paramstypes.TStoreKey],
	)

	// set the BaseApp's parameter store
	bApp.SetParamStore(
		app.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramskeeper.ConsensusParamsKeyTable()),
	)

	// add capability keeper and ScopeToModule for ibc module
	app.CapabilityKeeper = capabilitykeeper.NewKeeper(appCodec, keys[capabilitytypes.StoreKey], memKeys[capabilitytypes.MemStoreKey])
	scopedIBCKeeper := app.CapabilityKeeper.ScopeToModule(ibchost.ModuleName)
	scopedTransferKeeper := app.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	scopedICAHostKeeper := app.CapabilityKeeper.ScopeToModule(icahosttypes.SubModuleName)
	scopedStableKeeper := app.CapabilityKeeper.ScopeToModule(stablemoduletypes.ModuleName)
	scopedWasmKeeper := app.CapabilityKeeper.ScopeToModule(wasm.ModuleName)
	scopedICAControllerKeeper := app.CapabilityKeeper.ScopeToModule(icacontrollertypes.SubModuleName)
	scopedLiquidStakeIBCKeeper := app.CapabilityKeeper.ScopeToModule(liquidstakeibctypes.ModuleName)

	app.CapabilityKeeper.Seal()

	// add keepers
	app.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec,
		keys[authtypes.StoreKey],
		app.GetSubspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount,
		maccPerms,
	)
	app.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec,
		keys[banktypes.StoreKey],
		app.AccountKeeper,
		app.GetSubspace(banktypes.ModuleName),
		app.ModuleAccountAddrs(),
	)
	app.AuthzKeeper = authzkeeper.NewKeeper(
		keys[authzkeeper.StoreKey],
		appCodec,
		app.BaseApp.MsgServiceRouter(),
	)
	app.FeeGrantKeeper = feegrantkeeper.NewKeeper(
		appCodec,
		keys[feegrant.StoreKey],
		app.AccountKeeper,
	)
	stakingKeeper := stakingkeeper.NewKeeper(
		appCodec,
		keys[stakingtypes.StoreKey],
		app.AccountKeeper,
		app.BankKeeper,
		app.GetSubspace(stakingtypes.ModuleName),
	)
	app.MintKeeper = mintkeeper.NewKeeper(
		appCodec,
		keys[minttypes.StoreKey],
		app.GetSubspace(minttypes.ModuleName),
		&stakingKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		authtypes.FeeCollectorName,
	)
	app.DistrKeeper = distrkeeper.NewKeeper(
		appCodec,
		keys[distrtypes.StoreKey],
		app.GetSubspace(distrtypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		&stakingKeeper,
		authtypes.FeeCollectorName,
		app.ModuleAccountAddrs(),
	)
	app.SlashingKeeper = slashingkeeper.NewKeeper(
		appCodec,
		keys[slashingtypes.StoreKey],
		&stakingKeeper,
		app.GetSubspace(slashingtypes.ModuleName),
	)
	app.CrisisKeeper = crisiskeeper.NewKeeper(
		app.GetSubspace(crisistypes.ModuleName),
		invCheckPeriod,
		app.BankKeeper,
		authtypes.FeeCollectorName,
	)

	app.EpochsKeeper = epochskeeper.NewKeeper(keys[epochstypes.StoreKey])

	app.EpochsKeeper.SetHooks(
		epochstypes.NewMultiEpochHooks(
			epochstypes.NewMultiEpochHooks(),
			app.LiquidStakeIBCKeeper.NewEpochHooks(),
		),
	)

	app.UpgradeKeeper = upgradekeeper.NewKeeper(
		skipUpgradeHeights,
		keys[upgradetypes.StoreKey],
		appCodec,
		homePath,
		nil,
	)

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	app.StakingKeeper = *stakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(app.DistrKeeper.Hooks(), app.SlashingKeeper.Hooks()),
	)

	tokenFactoryKeeper := tokenfactorykeeper.NewKeeper(
		appCodec,
		app.keys[tokenfactorytypes.StoreKey],
		app.GetSubspace(tokenfactorytypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper.WithMintCoinsRestriction(tokenfactorytypes.NewTokenFactoryDenomMintCoinsRestriction()),
	)
	app.TokenFactoryKeeper = &tokenFactoryKeeper

	app.IBCKeeper = ibckeeper.NewKeeper(
		appCodec,
		keys[ibchost.StoreKey],
		app.GetSubspace(ibchost.ModuleName),
		app.StakingKeeper,
		app.UpgradeKeeper,
		scopedIBCKeeper,
	)

	app.OracleKeeper = oraclemodulekeeper.NewKeeper(
		appCodec,
		keys[oraclemoduletypes.StoreKey],
		app.GetSubspace(oraclemoduletypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		app.DistrKeeper,
		app.SlashingKeeper,
		&stakingKeeper,
		distrtypes.ModuleName,
	)
	oracleModule := oraclemodule.NewAppModule(appCodec, app.OracleKeeper, app.AccountKeeper, app.BankKeeper)

	app.StableKeeper = *stablemodulekeeper.NewKeeper(
		appCodec,
		keys[stablemoduletypes.StoreKey],
		keys[stablemoduletypes.StoreKey],
		app.GetSubspace(stablemoduletypes.ModuleName),
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		scopedStableKeeper,
		app.BankKeeper,
		app.OracleKeeper,
	)
	stableModule := stablemodule.NewAppModule(appCodec, app.StableKeeper, app.AccountKeeper, app.BankKeeper, app.OracleKeeper)
	stableIBCModule := stablemodule.NewIBCModule(app.StableKeeper)

	app.GrowKeeper = *growmodulekeeper.NewKeeper(
		appCodec,
		keys[growmoduletypes.StoreKey],
		keys[growmoduletypes.MemStoreKey],
		app.GetSubspace(growmoduletypes.ModuleName),
		app.BankKeeper,
		app.OracleKeeper,
		app.StableKeeper,
	)

	app.DexKeeper = *dexmodulekeeper.NewKeeper(
		appCodec,
		keys[dexmoduletypes.StoreKey],
		keys[dexmoduletypes.MemStoreKey],
		app.GetSubspace(dexmoduletypes.ModuleName),
		app.BankKeeper,
	)
	dexModule := dexmodule.NewAppModule(appCodec, app.DexKeeper, app.BankKeeper)

	// Create swap middleware keeper
	app.SwapKeeper = swapkeeper.NewKeeper(
		appCodec,
		app.MsgServiceRouter(),
		app.IBCKeeper.ChannelKeeper,
		app.BankKeeper,
	)
	swapModule := swapmiddleware.NewAppModule(app.SwapKeeper)

	app.PerpetualKeeper = *perpetualmodulekeeper.NewKeeper(
		appCodec,
		keys[perpetualmoduletypes.StoreKey],
		keys[perpetualmoduletypes.MemStoreKey],
		app.GetSubspace(perpetualmoduletypes.ModuleName),
		app.BankKeeper,
		app.OracleKeeper,
	)

	// register the proposal types
	govRouter := govtypes.NewRouter()
	govRouter.
		AddRoute(govtypes.RouterKey, govtypes.ProposalHandler).
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(app.ParamsKeeper)).
		AddRoute(distrtypes.RouterKey, distr.NewCommunityPoolSpendProposalHandler(app.DistrKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.UpgradeKeeper)).
		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(app.IBCKeeper.ClientKeeper)).
		AddRoute(stablemoduletypes.RouterKey, stablemodule.NewStableProposalHandler(&app.StableKeeper)).
		AddRoute(growmoduletypes.RouterKey, growmodule.NewGrowProposalHandler(&app.GrowKeeper)).
		AddRoute(oraclemoduletypes.RouterKey, oraclemodule.NewOracleProposalHandler(&app.OracleKeeper))

	wasmDir := filepath.Join(homePath, "data")
	wasmConfig, err := wasm.ReadWasmConfig(appOpts)
	if err != nil {
		panic("error while reading wasm config: " + err.Error())
	}

	// The last arguments can contain custom message handlers, and custom query handlers,
	// if we want to allow any custom callbacks
	supportedFeatures := "iterator,staking,stargate"
	wasmOpts := GetWasmOpts(appOpts)
	wasmOpts = append(wasmbinding.RegisterCustomPlugins(&app.BankKeeper, app.TokenFactoryKeeper, &app.OracleKeeper), wasmOpts...)
	app.wasmKeeper = wasm.NewKeeper(
		appCodec,
		keys[wasm.StoreKey],
		app.GetSubspace(wasm.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		app.StakingKeeper,
		app.DistrKeeper,
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		scopedWasmKeeper,
		app.TransferKeeper,
		app.MsgServiceRouter(),
		app.GRPCQueryRouter(),
		wasmDir,
		wasmConfig,
		supportedFeatures,
		wasmOpts...,
	)

	// register wasm gov proposal types
	enabledProposals := GetEnabledProposals()
	if len(enabledProposals) != 0 {
		govRouter.AddRoute(wasm.RouterKey, wasm.NewWasmProposalHandler(app.wasmKeeper, enabledProposals))
	}
	app.GovKeeper = govkeeper.NewKeeper(
		appCodec,
		keys[govtypes.StoreKey],
		app.GetSubspace(govtypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		&stakingKeeper,
		govRouter,
	)

	app.TransferKeeper = ibctransferkeeper.NewKeeper(
		appCodec,
		keys[ibctransfertypes.StoreKey],
		app.GetSubspace(ibctransfertypes.ModuleName),
		app.IBCKeeper.ChannelKeeper,
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		scopedTransferKeeper,
	)
	transferModule := transfer.NewAppModule(app.TransferKeeper)

	ibcTransferHooksKeeper := ibchookerkeeper.NewKeeper()
	app.TransferHooksKeeper = *ibcTransferHooksKeeper.SetHooks(
		ibchookertypes.NewMultiStakingHooks(
			app.LiquidStakeIBCKeeper.NewIBCTransferHooks(),
		),
	)

	app.PacketForwardKeeper = packetforwardkeeper.NewKeeper(
		appCodec,
		keys[packetforwardtypes.StoreKey],
		app.GetSubspace(packetforwardtypes.ModuleName),
		app.TransferKeeper, // will be zero-value here. reference set later on with SetTransferKeeper.
		app.IBCKeeper.ChannelKeeper,
		app.DistrKeeper,
		app.BankKeeper,
		app.IBCKeeper.ChannelKeeper,
	)

	app.ICAControllerKeeper = icacontrollerkeeper.NewKeeper(
		appCodec,
		keys[icacontrollertypes.StoreKey],
		app.GetSubspace(icacontrollertypes.SubModuleName),
		app.IBCKeeper.ChannelKeeper, // may be replaced with middleware such as ics29 fee
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		scopedICAControllerKeeper,
		app.MsgServiceRouter(),
	)

	app.ICAHostKeeper = icahostkeeper.NewKeeper(
		appCodec,
		keys[icahosttypes.StoreKey],
		app.GetSubspace(icahosttypes.SubModuleName),
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		app.AccountKeeper,
		scopedICAHostKeeper,
		app.MsgServiceRouter(),
	)

	app.InterchainQueryKeeper = interchainquerykeeper.NewKeeper(appCodec, keys[interchainquerytypes.StoreKey], app.IBCKeeper)

	app.LiquidStakeIBCKeeper = liquidstakeibckeeper.NewKeeper(
		appCodec,
		keys[liquidstakeibctypes.StoreKey],
		app.AccountKeeper,
		app.BankKeeper,
		app.EpochsKeeper,
		app.ICAControllerKeeper,
		app.IBCKeeper, // TODO: Move to module interface
		app.TransferKeeper,
		&app.InterchainQueryKeeper,
		&app.IBCKeeper.PortKeeper,
		scopedLiquidStakeIBCKeeper,
		app.GetSubspace(liquidstakeibctypes.ModuleName),
		app.MsgServiceRouter(),
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	app.LiquidStakeIBCKeeper = *app.LiquidStakeIBCKeeper.SetHooks(
		liquidstakeibctypes.NewMultiLiquidStakeIBCHooks(),
	)

	_ = app.InterchainQueryKeeper.SetCallbackHandler(liquidstakeibctypes.ModuleName, app.LiquidStakeIBCKeeper.CallbackHandler())

	liquidStakeIBCModule := liquidstakeibc.NewIBCModule(app.LiquidStakeIBCKeeper)

	var transferStack porttypes.IBCModule
	transferStack = transfer.NewIBCModule(app.TransferKeeper)
	transferStack = ibchooker.NewAppModule(app.TransferHooksKeeper, transferStack)
	transferStack = packetforward.NewIBCMiddleware(
		transferStack,
		app.PacketForwardKeeper,
		0, // TODO explore changing default values for retries and timeouts
		packetforwardkeeper.DefaultForwardTransferPacketTimeoutTimestamp,
		packetforwardkeeper.DefaultRefundTransferPacketTimeoutTimestamp,
	)
	transferStack = swapmiddleware.NewIBCMiddleware(transferStack, app.SwapKeeper)
	transferStack = gmpmiddleware.NewIBCMiddleware(transferStack)

	var icaControllerStack porttypes.IBCModule = liquidStakeIBCModule
	icaControllerStack = icacontroller.NewIBCMiddleware(icaControllerStack, app.ICAControllerKeeper)

	var icaHostStack porttypes.IBCModule = icahost.NewIBCModule(app.ICAHostKeeper)
	icaHostStack = packetforward.NewIBCMiddleware(
		icaHostStack, app.PacketForwardKeeper,
		0, // TODO explore changing default values for retries and timeouts
		packetforwardkeeper.DefaultForwardTransferPacketTimeoutTimestamp,
		packetforwardkeeper.DefaultRefundTransferPacketTimeoutTimestamp,
	)

	// routerModule := router.NewAppModule(app.RouterKeeper, transferIBCModule)
	// create static IBC router, add transfer route, then set and seal it
	ibcRouter := porttypes.NewRouter()
	ibcRouter.AddRoute(ibctransfertypes.ModuleName, transferStack).
		AddRoute(wasm.ModuleName, wasm.NewIBCHandler(app.wasmKeeper, app.IBCKeeper.ChannelKeeper, app.IBCKeeper.ChannelKeeper)).
		AddRoute(stablemoduletypes.ModuleName, stableIBCModule).
		AddRoute(icahosttypes.SubModuleName, icaHostStack).
		AddRoute(icacontrollertypes.SubModuleName, icaControllerStack).
		AddRoute(liquidstakeibctypes.ModuleName, icaControllerStack)

	app.IBCKeeper.SetRouter(ibcRouter)

	// create evidence keeper with router
	evidenceKeeper := evidencekeeper.NewKeeper(
		appCodec,
		keys[evidencetypes.StoreKey],
		&app.StakingKeeper,
		app.SlashingKeeper,
	)

	app.EvidenceKeeper = *evidenceKeeper
	skipGenesisInvariants := cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.
	app.mm = module.NewManager(
		genutil.NewAppModule(
			app.AccountKeeper,
			app.StakingKeeper,
			app.BaseApp.DeliverTx,
			encodingConfig.TxConfig,
		),
		auth.NewAppModule(appCodec, app.AccountKeeper, nil),
		vesting.NewAppModule(app.AccountKeeper, app.BankKeeper),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
		crisis.NewAppModule(&app.CrisisKeeper, skipGenesisInvariants),
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		upgrade.NewAppModule(app.UpgradeKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		feegrantmodule.NewAppModule(appCodec, app.AccountKeeper, app.BankKeeper, app.FeeGrantKeeper, app.interfaceRegistry),
		authzmodule.NewAppModule(appCodec, app.AuthzKeeper, app.AccountKeeper, app.BankKeeper, app.interfaceRegistry),
		ibc.NewAppModule(app.IBCKeeper),
		params.NewAppModule(app.ParamsKeeper),
		transferModule,
		ica.NewAppModule(nil, &app.ICAHostKeeper),
		oracleModule,
		packetforward.NewAppModule(app.PacketForwardKeeper),
		wasm.NewAppModule(appCodec, &app.wasmKeeper, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		tokenfactory.NewAppModule(appCodec, *app.TokenFactoryKeeper, app.AccountKeeper, app.BankKeeper),
		stableModule,
		growmodule.NewAppModule(appCodec, app.GrowKeeper, app.AccountKeeper, app.BankKeeper),
		dexModule,
		swapModule,
		epochs.NewAppModule(*app.EpochsKeeper),
		ica.NewAppModule(&app.ICAControllerKeeper, &app.ICAHostKeeper),
		liquidstakeibc.NewAppModule(app.LiquidStakeIBCKeeper),
		interchainquery.NewAppModule(appCodec, app.InterchainQueryKeeper),
		perpetualmodule.NewAppModule(appCodec, app.PerpetualKeeper, app.AccountKeeper, app.BankKeeper),
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	// NOTE: capability module's beginblocker must come before any modules using capabilities (e.g. IBC)
	app.mm.SetOrderBeginBlockers(
		// upgrades should be run first
		upgradetypes.ModuleName,
		capabilitytypes.ModuleName,
		crisistypes.ModuleName,
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		tokenfactorytypes.ModuleName,
		ibctransfertypes.ModuleName,
		ibchost.ModuleName,
		icatypes.ModuleName,
		packetforwardtypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		minttypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		vestingtypes.ModuleName,
		wasm.ModuleName,
		oraclemoduletypes.ModuleName,
		stablemoduletypes.ModuleName,
		growmoduletypes.ModuleName,
		dexmoduletypes.ModuleName,
		swaptypes.ModuleName,
		epochstypes.ModuleName,
		liquidstakeibctypes.ModuleName,
		interchainquerytypes.ModuleName,
		ibchookertypes.ModuleName,
		perpetualmoduletypes.ModuleName,
	)

	app.mm.SetOrderEndBlockers(
		crisistypes.ModuleName,
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		tokenfactorytypes.ModuleName,
		ibctransfertypes.ModuleName,
		ibchost.ModuleName,
		icatypes.ModuleName,
		packetforwardtypes.ModuleName,
		feegrant.ModuleName,
		authz.ModuleName,
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		slashingtypes.ModuleName,
		minttypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
		wasm.ModuleName,
		oraclemoduletypes.ModuleName,
		stablemoduletypes.ModuleName,
		growmoduletypes.ModuleName,
		dexmoduletypes.ModuleName,
		swaptypes.ModuleName,
		epochstypes.ModuleName,
		liquidstakeibctypes.ModuleName,
		interchainquerytypes.ModuleName,
		ibchookertypes.ModuleName,
		perpetualmoduletypes.ModuleName,
	)

	app.mm.SetOrderInitGenesis(
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distrtypes.ModuleName,
		stakingtypes.ModuleName,
		tokenfactorytypes.ModuleName,
		slashingtypes.ModuleName,
		govtypes.ModuleName,
		minttypes.ModuleName,
		crisistypes.ModuleName,
		ibctransfertypes.ModuleName,
		ibchost.ModuleName,
		icatypes.ModuleName,
		evidencetypes.ModuleName,
		feegrant.ModuleName,
		authz.ModuleName,
		genutiltypes.ModuleName,
		packetforwardtypes.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		vestingtypes.ModuleName,
		wasm.ModuleName,
		oraclemoduletypes.ModuleName,
		stablemoduletypes.ModuleName,
		growmoduletypes.ModuleName,
		dexmoduletypes.ModuleName,
		swaptypes.ModuleName,
		epochstypes.ModuleName,
		liquidstakeibctypes.ModuleName,
		interchainquerytypes.ModuleName,
		ibchookertypes.ModuleName,
		perpetualmoduletypes.ModuleName,
	)

	app.mm.RegisterInvariants(&app.CrisisKeeper)
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), encodingConfig.Amino)

	app.configurator = module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.mm.RegisterServices(app.configurator)

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	anteHandler, err := quadrateante.NewAnteHandler(
		quadrateante.HandlerOptions{
			AccountKeeper:     app.AccountKeeper,
			BankKeeper:        app.BankKeeper,
			FeegrantKeeper:    app.FeeGrantKeeper,
			SignModeHandler:   encodingConfig.TxConfig.SignModeHandler(),
			SigGasConsumer:    quadrateante.SigVerificationGasConsumer,
			IBCKeeper:         app.IBCKeeper,
			TxCounterStoreKey: keys[wasm.StoreKey],
			WasmConfig:        wasmConfig,
		},
	)
	if err != nil {
		panic(fmt.Errorf("failed to create AnteHandler: %s", err))
	}

	app.SetAnteHandler(anteHandler)
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)
	app.setUpgradeHandlers()

	app.ScopedIBCKeeper = scopedIBCKeeper
	app.ScopedTransferKeeper = scopedTransferKeeper
	app.scopedWasmKeeper = scopedWasmKeeper
	app.ScopedICAControllerKeeper = scopedICAControllerKeeper
	app.ScopedLiquidStakeIBCKeeper = scopedLiquidStakeIBCKeeper

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(fmt.Sprintf("failed to load latest version: %s", err))
		}

		// Initialize and seal the capability keeper so all persistent capabilities
		// are loaded in-memory and prevent any further modules from creating scoped
		// sub-keepers.
		// This must be done during creation of baseapp rather than in InitChain so
		// that in-memory capabilities get regenerated on app restart.
		// Note that since this reads from the store, we can only perform it when
		// `loadLatest` is set to true.
		//app.CapabilityKeeper.Seal()
	}

	return app
}

// Name returns the name of the App
func (app *QuadrateApp) Name() string { return app.BaseApp.Name() }

// BeginBlocker application updates every begin block
func (app *QuadrateApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker application updates every end block
func (app *QuadrateApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// InitChainer application update at chain initialization
func (app *QuadrateApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	if err := tmjson.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}

	app.UpgradeKeeper.SetModuleVersionMap(ctx, app.mm.GetVersionMap())

	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// LoadHeight loads a particular height
func (app *QuadrateApp) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// ModuleAccountAddrs returns all the app's module account addresses.
/* #nosec */
func (app *QuadrateApp) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = !receiveAllowedMAcc[acc]
	}

	return modAccAddrs
}

// LegacyAmino returns QuadrateApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *QuadrateApp) LegacyAmino() *codec.LegacyAmino {
	return app.legacyAmino
}

// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *QuadrateApp) AppCodec() codec.Codec {
	return app.appCodec
}

func (app *QuadrateApp) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *QuadrateApp) GetKey(storeKey string) *sdk.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *QuadrateApp) GetTKey(storeKey string) *sdk.TransientStoreKey {
	return app.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (app *QuadrateApp) GetMemKey(storeKey string) *sdk.MemoryStoreKey {
	return app.memKeys[storeKey]
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *QuadrateApp) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// RegisterAPIRoutes registers all application module routes with the provided
// API server.
func (app *QuadrateApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	rpc.RegisterRoutes(clientCtx, apiSvr.Router)
	// Register legacy tx routes.
	authrest.RegisterTxRoutes(clientCtx, apiSvr.Router)
	// Register new tx routes from grpc-gateway.
	authtx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register legacy and grpc-gateway routes for all modules.
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// register swagger API from root so that other applications can override easily
	if apiConfig.Swagger {
		RegisterSwaggerAPI(apiSvr.Router)
	}
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *QuadrateApp) RegisterTxService(clientCtx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *QuadrateApp) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.interfaceRegistry)
}

// RegisterSwaggerAPI registers swagger route with API Server
func RegisterSwaggerAPI(rtr *mux.Router) {
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}

	staticServer := http.FileServer(statikFS)
	rtr.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", staticServer))
}

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey sdk.StoreKey) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(minttypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govtypes.ParamKeyTable())
	paramsKeeper.Subspace(crisistypes.ModuleName)
	paramsKeeper.Subspace(ibctransfertypes.ModuleName)
	paramsKeeper.Subspace(ibchost.ModuleName)
	paramsKeeper.Subspace(packetforwardtypes.ModuleName).WithKeyTable(packetforwardtypes.ParamKeyTable())
	paramsKeeper.Subspace(icahosttypes.SubModuleName)
	paramsKeeper.Subspace(wasm.ModuleName)
	paramsKeeper.Subspace(tokenfactorytypes.ModuleName)
	paramsKeeper.Subspace(oraclemoduletypes.ModuleName)
	paramsKeeper.Subspace(stablemoduletypes.ModuleName)
	paramsKeeper.Subspace(growmoduletypes.ModuleName)
	paramsKeeper.Subspace(dexmoduletypes.ModuleName)
	paramsKeeper.Subspace(icacontrollertypes.SubModuleName)
	paramsKeeper.Subspace(perpetualmoduletypes.ModuleName)
	return paramsKeeper
}

func (app *QuadrateApp) setUpgradeHandlers() {
	app.UpgradeKeeper.SetUpgradeHandler(
		tfupgrades.UpgradeName,
		tfupgrades.CreateUpgradeHandler(
			app.mm,
			app.configurator,
			*app.TokenFactoryKeeper,
		),
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		v4.UpgradeName,
		v4.CreateUpgradeHandler(
			app.mm,
			app.configurator,
		),
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		v4rc0.UpgradeName,
		v4rc0.CreateUpgradeHandler(
			app.mm,
			app.configurator,
		),
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		v5.UpgradeName,
		v5.CreateUpgradeHandler(
			app.mm,
			app.configurator,
		),
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		v2.UpgradeName,
		v2.CreateUpgradeHandler(
			app.mm,
			app.configurator,
			app.StableKeeper,
			app.GrowKeeper,
		),
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		v0.UpgradeName,
		v0.CreateUpgradeHandler(
			app.mm,
			app.configurator,
			app.StableKeeper,
			app.GrowKeeper,
		),
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		v1.UpgradeName,
		v1.CreateUpgradeHandler(
			app.mm,
			app.configurator,
			app.GrowKeeper,
		),
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		v022.UpgradeName,
		v022.CreateUpgradeHandler(
			app.mm,
			app.configurator,
			app.StableKeeper,
		),
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		v025.UpgradeName,
		v025.CreateUpgradeHandler(
			app.mm,
			app.configurator,
			app.GrowKeeper,
		),
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		v025rc0.UpgradeName,
		v025rc0.CreateUpgradeHandler(
			app.mm,
			app.configurator,
			app.GrowKeeper,
		),
	)

	app.UpgradeKeeper.SetUpgradeHandler(
		v030.UpgradeName,
		v030.CreateUpgradeHandler(
			app.mm,
			app.configurator,
			app.StableKeeper,
			app.GrowKeeper,
			app.LiquidStakeIBCKeeper,
		),
	)

	// When a planned update height is reached, the old binary will panic
	// writing on disk the height and name of the update that triggered it
	// This will read that value, and execute the preparations for the upgrade.
	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Errorf("failed to read upgrade info from disk: %w", err))
	}

	if app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		return
	}

	var storeUpgrades []*storetypes.StoreUpgrades

	switch upgradeInfo.Name {
	case tfupgrades.UpgradeName:
	case v4.UpgradeName:
	case v4rc0.UpgradeName:
	case v5.UpgradeName:
	case v2.UpgradeName:
	case v1.UpgradeName:
		storeUpgrades = append(storeUpgrades, &v1.Upgrade.StoreUpgrades)
	case v025.UpgradeName:
	case v025rc0.UpgradeName:
	case v030.UpgradeName:
		storeUpgrades = append(storeUpgrades, &v030.Upgrade.StoreUpgrades)
	}

	for _, storeUpgrade := range storeUpgrades {
		if storeUpgrade != nil {
			// configure store loader that checks if version == upgradeHeight and applies store upgrades
			app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, storeUpgrade))
		}
	}
}

// GetBaseApp implements the TestingApp interface.
func (app *QuadrateApp) GetBaseApp() *baseapp.BaseApp {
	return app.BaseApp
}

func (app *QuadrateApp) GetIBCKeeper() *ibckeeper.Keeper {
	return app.IBCKeeper
}

// GetScopedIBCKeeper implements the TestingApp interface.
func (app *QuadrateApp) GetScopedIBCKeeper() capabilitykeeper.ScopedKeeper {
	return app.ScopedIBCKeeper
}

// GetStakingKeeper implements the TestingApp interface.
func (app *QuadrateApp) GetStakingKeeper() stakingkeeper.Keeper {
	return app.StakingKeeper
}

// GetTxConfig implements the TestingApp interface.
func (app *QuadrateApp) GetTxConfig() client.TxConfig {
	cfg := MakeEncodingConfig()
	return cfg.TxConfig
}

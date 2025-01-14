package upgrade

import (
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/DFWallet/anatha/client/context"
	"github.com/DFWallet/anatha/client/flags"
	"github.com/DFWallet/anatha/codec"
	sdk "github.com/DFWallet/anatha/types"
	"github.com/DFWallet/anatha/types/module"
	"github.com/DFWallet/project-anatha/x/upgrade/client/cli"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic implements the sdk.AppModuleBasic interface
type AppModuleBasic struct{}

// Name returns the ModuleName
func (AppModuleBasic) Name() string {
	return ModuleName
}

// RegisterCodec registers the upgrade types on the amino codec
func (AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
	RegisterCodec(cdc)
}

// RegisterRESTRoutes registers all REST query handlers
func (AppModuleBasic) RegisterRESTRoutes(ctx context.CLIContext, r *mux.Router) {}

// GetQueryCmd returns the cli query commands for this module
func (AppModuleBasic) GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	queryCmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Querying commands for the upgrade module",
	}
	queryCmd.AddCommand(flags.GetCommands(
		cli.GetPlanCmd(StoreKey, cdc),
		cli.GetAppliedHeightCmd(StoreKey, cdc),
	)...)

	return queryCmd
}

// GetTxCmd returns the transaction commands for this module
func (AppModuleBasic) GetTxCmd(cdc *codec.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Upgrade transaction subcommands",
	}
	txCmd.AddCommand(flags.PostCommands()...)
	return txCmd
}

// AppModule implements the sdk.AppModule interface
type AppModule struct {
	AppModuleBasic
	keeper Keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(keeper Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         keeper,
	}
}

// RegisterInvariants does nothing, there are no invariants to enforce
func (AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// Route is empty, as we do not handle Messages (just proposals)
func (AppModule) Route() string { return "" }

// NewHandler is empty, as we do not handle Messages (just proposals)
func (am AppModule) NewHandler() sdk.Handler { return nil }

// QuerierRoute returns the route we respond to for abci queries
func (AppModule) QuerierRoute() string { return QuerierKey }

// NewQuerierHandler registers a query handler to respond to the module-specific queries
func (am AppModule) NewQuerierHandler() sdk.Querier {
	return NewQuerier(am.keeper)
}

// InitGenesis is ignored, no sense in serializing future upgrades
func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

// DefaultGenesis is an empty object
func (AppModuleBasic) DefaultGenesis() json.RawMessage {
	return []byte("{}")
}

// ValidateGenesis is always successful, as we ignore the value
func (AppModuleBasic) ValidateGenesis(bz json.RawMessage) error {
	return nil
}

// ExportGenesis is always empty, as InitGenesis does nothing either
func (am AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	return am.DefaultGenesis()
}

// BeginBlock calls the upgrade module hooks
//
// CONTRACT: this is registered in BeginBlocker *before* all other modules' BeginBlock functions
func (am AppModule) BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock) {
	BeginBlocker(am.keeper, ctx, req)
}

// EndBlock does nothing
func (am AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

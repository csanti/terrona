package policy

import (
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	//"github.com/csanti/terrona/x/policy/client/cli"
	"github.com/csanti/terrona/x/policy/client/rest"
	"github.com/csanti/terrona/x/policy/keeper"
	"github.com/csanti/terrona/x/policy/types"
)

// Type check to ensure the interface is properly implemented
var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic defines the basic application module used by the clean module.
type AppModuleBasic struct{}

// Name returns the clean module's name.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterCodec registers the clean module's types for the given codec.
func (AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
	types.RegisterCodec(cdc)
}

// DefaultGenesis returns default genesis state as raw bytes for the clean
// module.
func (AppModuleBasic) DefaultGenesis() json.RawMessage {
	return types.ModuleCdc.MustMarshalJSON(types.DefaultGenesisState())
}

// ValidateGenesis performs genesis state validation for the clean module.
func (AppModuleBasic) ValidateGenesis(bz json.RawMessage) error {
	var data types.GenesisState
	err := types.ModuleCdc.UnmarshalJSON(bz, &data)
	if err != nil {
		return err
	}
	return types.ValidateGenesis(data)
}

// RegisterRESTRoutes registers the REST routes for the clean module.
func (AppModuleBasic) RegisterRESTRoutes(ctx context.CLIContext, rtr *mux.Router) {
	rest.RegisterRoutes(ctx, rtr)
}

// GetTxCmd returns the root tx command for the clean module.
func (AppModuleBasic) GetTxCmd(cdc *codec.Codec) *cobra.Command {
	return nil
}

// GetQueryCmd returns no root query command for the clean module.
func (AppModuleBasic) GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	return nil
}

//____________________________________________________________________________

// AppModule implements an application module for the clean module.
type AppModule struct {
	AppModuleBasic

	keeper     keeper.Keeper
	coinKeeper bank.Keeper
	// TODO: Add keepers that your application depends on

}

// NewAppModule creates a new AppModule object
func NewAppModule(k keeper.Keeper, bankKeeper bank.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         k,
		coinKeeper:     bankKeeper,
		// TODO: Add keepers that your application depends on
	}
}

// Name returns the clean module's name.
func (AppModule) Name() string {
	return types.ModuleName
}

// RegisterInvariants registers the clean module invariants.
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// Route returns the message routing key for the clean module.
func (AppModule) Route() string {
	return types.RouterKey
}

// NewHandler returns an sdk.Handler for the clean module.
func (am AppModule) NewHandler() sdk.Handler {
	//return NewHandler(am.keeper)
	// This module does not need handlers
	return nil
}

// QuerierRoute returns the clean module's querier route name.
func (AppModule) QuerierRoute() string {
	return types.QuerierRoute
}

// NewQuerierHandler returns the clean module sdk.Querier.
func (am AppModule) NewQuerierHandler() sdk.Querier {
	return keeper.NewQuerier(am.keeper)
}

// InitGenesis performs genesis initialization for the clean module. It returns
// no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	/*
	var genesisState types.GenesisState
	types.ModuleCdc.MustUnmarshalJSON(data, &genesisState)
	InitGenesis(ctx, am.keeper, genesisState)
	*/
	// no needed, just default values to the parameters
	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the exported genesis state as raw bytes for the clean
// module.
func (am AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	//gs := ExportGenesis(ctx, am.keeper)
	//return types.ModuleCdc.MustMarshalJSON(gs)
	return nil
}

// BeginBlock returns the begin blocker for the clean module.
func (am AppModule) BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock) {
	//BeginBlocker(ctx, req, am.keeper)
}

// EndBlock returns the end blocker for the clean module. It returns no validator
// updates.
func (AppModule) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

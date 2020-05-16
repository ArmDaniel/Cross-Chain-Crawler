package dcrawl

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	//abci "github.com/tendermint/tendermint/abci/types"
	"github.com/Motanovici/Cross-Chain-Crawler/x/dcrawl/types"
)

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, k Keeper /* TODO: Define what keepers the module needs */, data GenesisState) {
	// TODO: Define logic for when you would like to initalize a new genesis
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k Keeper) (data GenesisState) {
	// TODO: Define logic for exporting state
	return types.NewGenesisState()
}

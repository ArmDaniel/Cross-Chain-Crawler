package dcrawl

//path might need to be modified a little after commit to github

import (
	"github.com/Motanovici/Cross-Chain-Crawler/x/dcrawl/keeper"
	"github.com/Motanovici/Cross-Chain-Crawler/x/dcrawl/types"
)

const (
	// TODO: define constants that you would like exposed from your module

	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	DefaultParamspace = types.DefaultParamspace
	//QueryParams       = types.QueryParams
	QuerierRoute      = types.QuerierRoute
)

var (
	// functions aliases
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier
	RegisterCodec       = types.RegisterCodec
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis
	// TODO: Fill out function aliases

	// variable aliases
	ModuleCdc     = types.ModuleCdc
	// TODO: Fill out variable aliases
	NewMsgCreateScavenge = types.NewMsgCreateScavenge
	NewMsgCommitSolution = types.NewMsgCommitSolution
	NewMsgRevealSolution = types.NewMsgRevealSolution
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
	Params       = types.Params

	// TODO: Fill out module types
	MsgCreateScavenge = types.MsgCreateScavenge
	MsgCommitSolution = types.MsgCommitSolution
	MsgRevealSolution = types.MsgRevealSolution
)

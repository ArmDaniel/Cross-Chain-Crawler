package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateScavenge{},"dcrawl/CreateScavenge",nil)
	cdc.RegisterConcrete(MsgCommitSolution{},"dcrawl/CommitSolution",nil)
	cdc.RegisterConcrete(MsgRevealSolution{},"dcrawl/RevealSolution",nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}

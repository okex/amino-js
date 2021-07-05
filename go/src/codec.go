package src

import (
	"github.com/cosmos/amino-js/go/lib"
	"github.com/cosmos/amino-js/go/lib/exchain/ethtypes"
	"github.com/tendermint/go-amino"
)

var codec *amino.Codec

func init() {
	codec = amino.NewCodec()
	lib.RegisterCodec(codec)
	codec.Seal()
	ethtypes.InitCdc(codec)
}

func GetCdc() *amino.Codec {
	return codec
}
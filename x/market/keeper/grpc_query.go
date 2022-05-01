package keeper

import (
	"github.com/VelaChain/vela/x/market/types"
)

var _ types.QueryServer = Keeper{}

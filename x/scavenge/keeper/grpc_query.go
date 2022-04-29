package keeper

import (
	"github.com/carmonasl/scavenge/x/scavenge/types"
)

var _ types.QueryServer = Keeper{}

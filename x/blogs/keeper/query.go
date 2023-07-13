package keeper

import (
	"blogs/x/blogs/types"
)

var _ types.QueryServer = Keeper{}

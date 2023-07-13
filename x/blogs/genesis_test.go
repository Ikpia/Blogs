package blogs_test

import (
	"testing"

	keepertest "blogs/testutil/keeper"
	"blogs/testutil/nullify"
	"blogs/x/blogs"
	"blogs/x/blogs/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BlogsKeeper(t)
	blogs.InitGenesis(ctx, *k, genesisState)
	got := blogs.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

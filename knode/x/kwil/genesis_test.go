package kwil_test

import (
	"testing"

	"github.com/kwilteam/kwil-db/knode/internal/x/kwil/types"
	keepertest "github.com/kwilteam/kwil-db/knode/testutil/keeper"
	"github.com/kwilteam/kwil-db/knode/testutil/nullify"
	"github.com/kwilteam/kwil-db/knode/x/kwil"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		DatabasesList: []types.Databases{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		DdlList: []types.Ddl{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		DdlindexList: []types.Ddlindex{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		QueryidsList: []types.Queryids{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.KwilKeeper(t)
	kwil.InitGenesis(ctx, *k, genesisState)
	got := kwil.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.DatabasesList, got.DatabasesList)
	require.ElementsMatch(t, genesisState.DdlList, got.DdlList)
	require.ElementsMatch(t, genesisState.DdlindexList, got.DdlindexList)
	require.ElementsMatch(t, genesisState.QueryidsList, got.QueryidsList)
	// this line is used by starport scaffolding # genesis/test/assert
}

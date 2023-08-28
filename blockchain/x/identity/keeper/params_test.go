package keeper_test

import (
	"testing"

	testkeeper "github.com/qredo/fusionchain/testutil/keeper"
	"github.com/qredo/fusionchain/x/identity/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.IdentityKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

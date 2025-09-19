package keeper_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"resist/x/posts/keeper"
	"resist/x/posts/types"
)

func createNSource(keeper keeper.Keeper, ctx context.Context, n int) []types.Source {
	items := make([]types.Source, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
		items[i].Url = strconv.Itoa(i)
		items[i].Title = strconv.Itoa(i)
		items[i].Description = strconv.Itoa(i)
		items[i].CredibilityScore = int64(i)
		items[i].AnalysisSummary = strconv.Itoa(i)
		items[i].Verified = true
		_ = keeper.Source.Set(ctx, items[i].Index, items[i])
	}
	return items
}

func TestSourceQuerySingle(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNSource(f.keeper, f.ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetSourceRequest
		response *types.QueryGetSourceResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetSourceRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetSourceResponse{Source: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetSourceRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetSourceResponse{Source: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetSourceRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := qs.GetSource(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.EqualExportedValues(t, tc.response, response)
			}
		})
	}
}

func TestSourceQueryPaginated(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNSource(f.keeper, f.ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllSourceRequest {
		return &types.QueryAllSourceRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListSource(f.ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Source), step)
			require.Subset(t, msgs, resp.Source)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListSource(f.ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Source), step)
			require.Subset(t, msgs, resp.Source)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListSource(f.ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.EqualExportedValues(t, msgs, resp.Source)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListSource(f.ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

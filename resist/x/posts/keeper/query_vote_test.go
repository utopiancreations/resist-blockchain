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

func createNVote(keeper keeper.Keeper, ctx context.Context, n int) []types.Vote {
	items := make([]types.Vote, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
		items[i].VoterAddress = strconv.Itoa(i)
		items[i].PostId = uint64(i)
		items[i].VoteType = strconv.Itoa(i)
		items[i].Timestamp = int64(i)
		_ = keeper.Vote.Set(ctx, items[i].Index, items[i])
	}
	return items
}

func TestVoteQuerySingle(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNVote(f.keeper, f.ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetVoteRequest
		response *types.QueryGetVoteResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetVoteRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetVoteResponse{Vote: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetVoteRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetVoteResponse{Vote: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetVoteRequest{
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
			response, err := qs.GetVote(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.EqualExportedValues(t, tc.response, response)
			}
		})
	}
}

func TestVoteQueryPaginated(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNVote(f.keeper, f.ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllVoteRequest {
		return &types.QueryAllVoteRequest{
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
			resp, err := qs.ListVote(f.ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Vote), step)
			require.Subset(t, msgs, resp.Vote)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListVote(f.ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Vote), step)
			require.Subset(t, msgs, resp.Vote)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListVote(f.ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.EqualExportedValues(t, msgs, resp.Vote)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListVote(f.ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

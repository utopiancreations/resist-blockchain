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

func createNSocialPost(keeper keeper.Keeper, ctx context.Context, n int) []types.SocialPost {
	items := make([]types.SocialPost, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
		items[i].Title = strconv.Itoa(i)
		items[i].Content = strconv.Itoa(i)
		items[i].MediaUrl = strconv.Itoa(i)
		items[i].MediaType = strconv.Itoa(i)
		items[i].GroupId = uint64(i)
		items[i].Author = strconv.Itoa(i)
		items[i].Upvotes = uint64(i)
		items[i].Downvotes = uint64(i)
		items[i].CreatedAt = uint64(i)
		_ = keeper.SocialPost.Set(ctx, items[i].Index, items[i])
	}
	return items
}

func TestSocialPostQuerySingle(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNSocialPost(f.keeper, f.ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetSocialPostRequest
		response *types.QueryGetSocialPostResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetSocialPostRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetSocialPostResponse{SocialPost: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetSocialPostRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetSocialPostResponse{SocialPost: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetSocialPostRequest{
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
			response, err := qs.GetSocialPost(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.EqualExportedValues(t, tc.response, response)
			}
		})
	}
}

func TestSocialPostQueryPaginated(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNSocialPost(f.keeper, f.ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllSocialPostRequest {
		return &types.QueryAllSocialPostRequest{
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
			resp, err := qs.ListSocialPost(f.ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.SocialPost), step)
			require.Subset(t, msgs, resp.SocialPost)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListSocialPost(f.ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.SocialPost), step)
			require.Subset(t, msgs, resp.SocialPost)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListSocialPost(f.ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.EqualExportedValues(t, msgs, resp.SocialPost)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListSocialPost(f.ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

package keeper_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"resist/x/usergroups/keeper"
	"resist/x/usergroups/types"
)

func createNUserGroup(keeper keeper.Keeper, ctx context.Context, n int) []types.UserGroup {
	items := make([]types.UserGroup, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
		items[i].Name = strconv.Itoa(i)
		items[i].Description = strconv.Itoa(i)
		items[i].Admin = strconv.Itoa(i)
		items[i].Members = []string{`abc` + strconv.Itoa(i), `xyz` + strconv.Itoa(i)}
		items[i].VoteThreshold = uint64(i)
		items[i].CreatedAt = uint64(i)
		_ = keeper.UserGroup.Set(ctx, items[i].Index, items[i])
	}
	return items
}

func TestUserGroupQuerySingle(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNUserGroup(f.keeper, f.ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetUserGroupRequest
		response *types.QueryGetUserGroupResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetUserGroupRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetUserGroupResponse{UserGroup: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetUserGroupRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetUserGroupResponse{UserGroup: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetUserGroupRequest{
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
			response, err := qs.GetUserGroup(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.EqualExportedValues(t, tc.response, response)
			}
		})
	}
}

func TestUserGroupQueryPaginated(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNUserGroup(f.keeper, f.ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllUserGroupRequest {
		return &types.QueryAllUserGroupRequest{
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
			resp, err := qs.ListUserGroup(f.ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.UserGroup), step)
			require.Subset(t, msgs, resp.UserGroup)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListUserGroup(f.ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.UserGroup), step)
			require.Subset(t, msgs, resp.UserGroup)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListUserGroup(f.ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.EqualExportedValues(t, msgs, resp.UserGroup)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListUserGroup(f.ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

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

func createNContentReport(keeper keeper.Keeper, ctx context.Context, n int) []types.ContentReport {
	items := make([]types.ContentReport, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
		items[i].PostId = uint64(i)
		items[i].Reporter = strconv.Itoa(i)
		items[i].Reason = strconv.Itoa(i)
		items[i].Evidence = strconv.Itoa(i)
		items[i].Status = strconv.Itoa(i)
		items[i].CommunityResponse = strconv.Itoa(i)
		items[i].Resolution = strconv.Itoa(i)
		_ = keeper.ContentReport.Set(ctx, items[i].Index, items[i])
	}
	return items
}

func TestContentReportQuerySingle(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNContentReport(f.keeper, f.ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetContentReportRequest
		response *types.QueryGetContentReportResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetContentReportRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetContentReportResponse{ContentReport: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetContentReportRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetContentReportResponse{ContentReport: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetContentReportRequest{
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
			response, err := qs.GetContentReport(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.EqualExportedValues(t, tc.response, response)
			}
		})
	}
}

func TestContentReportQueryPaginated(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNContentReport(f.keeper, f.ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllContentReportRequest {
		return &types.QueryAllContentReportRequest{
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
			resp, err := qs.ListContentReport(f.ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ContentReport), step)
			require.Subset(t, msgs, resp.ContentReport)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListContentReport(f.ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ContentReport), step)
			require.Subset(t, msgs, resp.ContentReport)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListContentReport(f.ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.EqualExportedValues(t, msgs, resp.ContentReport)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListContentReport(f.ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

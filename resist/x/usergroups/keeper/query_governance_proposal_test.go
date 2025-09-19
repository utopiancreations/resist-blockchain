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

func createNGovernanceProposal(keeper keeper.Keeper, ctx context.Context, n int) []types.GovernanceProposal {
	items := make([]types.GovernanceProposal, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
		items[i].Title = strconv.Itoa(i)
		items[i].Description = strconv.Itoa(i)
		items[i].Proposer = strconv.Itoa(i)
		items[i].ProposalType = strconv.Itoa(i)
		items[i].VotingPeriodStart = int64(i)
		items[i].VotingPeriodEnd = int64(i)
		items[i].YesVotes = uint64(i)
		items[i].NoVotes = uint64(i)
		items[i].AbstainVotes = uint64(i)
		items[i].Status = strconv.Itoa(i)
		_ = keeper.GovernanceProposal.Set(ctx, items[i].Index, items[i])
	}
	return items
}

func TestGovernanceProposalQuerySingle(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNGovernanceProposal(f.keeper, f.ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetGovernanceProposalRequest
		response *types.QueryGetGovernanceProposalResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetGovernanceProposalRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetGovernanceProposalResponse{GovernanceProposal: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetGovernanceProposalRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetGovernanceProposalResponse{GovernanceProposal: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetGovernanceProposalRequest{
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
			response, err := qs.GetGovernanceProposal(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.EqualExportedValues(t, tc.response, response)
			}
		})
	}
}

func TestGovernanceProposalQueryPaginated(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNGovernanceProposal(f.keeper, f.ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllGovernanceProposalRequest {
		return &types.QueryAllGovernanceProposalRequest{
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
			resp, err := qs.ListGovernanceProposal(f.ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.GovernanceProposal), step)
			require.Subset(t, msgs, resp.GovernanceProposal)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListGovernanceProposal(f.ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.GovernanceProposal), step)
			require.Subset(t, msgs, resp.GovernanceProposal)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListGovernanceProposal(f.ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.EqualExportedValues(t, msgs, resp.GovernanceProposal)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListGovernanceProposal(f.ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

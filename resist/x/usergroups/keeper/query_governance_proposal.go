package keeper

import (
	"context"
	"errors"

	"resist/x/usergroups/types"

	"cosmossdk.io/collections"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ListGovernanceProposal(ctx context.Context, req *types.QueryAllGovernanceProposalRequest) (*types.QueryAllGovernanceProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	governanceProposals, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.GovernanceProposal,
		req.Pagination,
		func(_ string, value types.GovernanceProposal) (types.GovernanceProposal, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllGovernanceProposalResponse{GovernanceProposal: governanceProposals, Pagination: pageRes}, nil
}

func (q queryServer) GetGovernanceProposal(ctx context.Context, req *types.QueryGetGovernanceProposalRequest) (*types.QueryGetGovernanceProposalResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, err := q.k.GovernanceProposal.Get(ctx, req.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetGovernanceProposalResponse{GovernanceProposal: val}, nil
}

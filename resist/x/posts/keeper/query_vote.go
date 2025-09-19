package keeper

import (
	"context"
	"errors"

	"resist/x/posts/types"

	"cosmossdk.io/collections"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ListVote(ctx context.Context, req *types.QueryAllVoteRequest) (*types.QueryAllVoteResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	votes, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.Vote,
		req.Pagination,
		func(_ string, value types.Vote) (types.Vote, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVoteResponse{Vote: votes, Pagination: pageRes}, nil
}

func (q queryServer) GetVote(ctx context.Context, req *types.QueryGetVoteRequest) (*types.QueryGetVoteResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, err := q.k.Vote.Get(ctx, req.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetVoteResponse{Vote: val}, nil
}

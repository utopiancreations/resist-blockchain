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

func (q queryServer) ListSource(ctx context.Context, req *types.QueryAllSourceRequest) (*types.QueryAllSourceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	sources, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.Source,
		req.Pagination,
		func(_ string, value types.Source) (types.Source, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSourceResponse{Source: sources, Pagination: pageRes}, nil
}

func (q queryServer) GetSource(ctx context.Context, req *types.QueryGetSourceRequest) (*types.QueryGetSourceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, err := q.k.Source.Get(ctx, req.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetSourceResponse{Source: val}, nil
}

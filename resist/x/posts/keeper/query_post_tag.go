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

func (q queryServer) ListPostTag(ctx context.Context, req *types.QueryAllPostTagRequest) (*types.QueryAllPostTagResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	postTags, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.PostTag,
		req.Pagination,
		func(_ string, value types.PostTag) (types.PostTag, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPostTagResponse{PostTag: postTags, Pagination: pageRes}, nil
}

func (q queryServer) GetPostTag(ctx context.Context, req *types.QueryGetPostTagRequest) (*types.QueryGetPostTagResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, err := q.k.PostTag.Get(ctx, req.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetPostTagResponse{PostTag: val}, nil
}

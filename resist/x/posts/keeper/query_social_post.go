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

func (q queryServer) ListSocialPost(ctx context.Context, req *types.QueryAllSocialPostRequest) (*types.QueryAllSocialPostResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	socialPosts, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.SocialPost,
		req.Pagination,
		func(_ string, value types.SocialPost) (types.SocialPost, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSocialPostResponse{SocialPost: socialPosts, Pagination: pageRes}, nil
}

func (q queryServer) GetSocialPost(ctx context.Context, req *types.QueryGetSocialPostRequest) (*types.QueryGetSocialPostResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, err := q.k.SocialPost.Get(ctx, req.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetSocialPostResponse{SocialPost: val}, nil
}

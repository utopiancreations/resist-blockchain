package keeper

import (
	"context"
	"errors"

	"resist/x/identity/types"

	"cosmossdk.io/collections"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ListUserProfile(ctx context.Context, req *types.QueryAllUserProfileRequest) (*types.QueryAllUserProfileResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	userProfiles, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.UserProfile,
		req.Pagination,
		func(_ string, value types.UserProfile) (types.UserProfile, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUserProfileResponse{UserProfile: userProfiles, Pagination: pageRes}, nil
}

func (q queryServer) GetUserProfile(ctx context.Context, req *types.QueryGetUserProfileRequest) (*types.QueryGetUserProfileResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, err := q.k.UserProfile.Get(ctx, req.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetUserProfileResponse{UserProfile: val}, nil
}

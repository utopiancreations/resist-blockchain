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

func (q queryServer) ListUserGroup(ctx context.Context, req *types.QueryAllUserGroupRequest) (*types.QueryAllUserGroupResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	userGroups, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.UserGroup,
		req.Pagination,
		func(_ string, value types.UserGroup) (types.UserGroup, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllUserGroupResponse{UserGroup: userGroups, Pagination: pageRes}, nil
}

func (q queryServer) GetUserGroup(ctx context.Context, req *types.QueryGetUserGroupRequest) (*types.QueryGetUserGroupResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, err := q.k.UserGroup.Get(ctx, req.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetUserGroupResponse{UserGroup: val}, nil
}

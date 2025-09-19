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

func (q queryServer) ListContentReport(ctx context.Context, req *types.QueryAllContentReportRequest) (*types.QueryAllContentReportResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	contentReports, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.ContentReport,
		req.Pagination,
		func(_ string, value types.ContentReport) (types.ContentReport, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllContentReportResponse{ContentReport: contentReports, Pagination: pageRes}, nil
}

func (q queryServer) GetContentReport(ctx context.Context, req *types.QueryGetContentReportRequest) (*types.QueryGetContentReportResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, err := q.k.ContentReport.Get(ctx, req.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetContentReportResponse{ContentReport: val}, nil
}

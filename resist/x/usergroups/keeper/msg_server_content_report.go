package keeper

import (
	"context"
	"errors"
	"fmt"

	"resist/x/usergroups/types"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateContentReport(ctx context.Context, msg *types.MsgCreateContentReport) (*types.MsgCreateContentReportResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	// Check if the value already exists
	ok, err := k.ContentReport.Has(ctx, msg.Index)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	} else if ok {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var contentReport = types.ContentReport{
		Creator:           msg.Creator,
		Index:             msg.Index,
		PostId:            msg.PostId,
		Reporter:          msg.Reporter,
		Reason:            msg.Reason,
		Evidence:          msg.Evidence,
		Status:            msg.Status,
		CommunityResponse: msg.CommunityResponse,
		Resolution:        msg.Resolution,
	}

	if err := k.ContentReport.Set(ctx, contentReport.Index, contentReport); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	return &types.MsgCreateContentReportResponse{}, nil
}

func (k msgServer) UpdateContentReport(ctx context.Context, msg *types.MsgUpdateContentReport) (*types.MsgUpdateContentReportResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.ContentReport.Get(ctx, msg.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var contentReport = types.ContentReport{
		Creator:           msg.Creator,
		Index:             msg.Index,
		PostId:            msg.PostId,
		Reporter:          msg.Reporter,
		Reason:            msg.Reason,
		Evidence:          msg.Evidence,
		Status:            msg.Status,
		CommunityResponse: msg.CommunityResponse,
		Resolution:        msg.Resolution,
	}

	if err := k.ContentReport.Set(ctx, contentReport.Index, contentReport); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update contentReport")
	}

	return &types.MsgUpdateContentReportResponse{}, nil
}

func (k msgServer) DeleteContentReport(ctx context.Context, msg *types.MsgDeleteContentReport) (*types.MsgDeleteContentReportResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.ContentReport.Get(ctx, msg.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if err := k.ContentReport.Remove(ctx, msg.Index); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to remove contentReport")
	}

	return &types.MsgDeleteContentReportResponse{}, nil
}

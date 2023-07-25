package keeper

import (
	"context"

	"blogs/x/blogs/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowingPost(goCtx context.Context, req *types.QueryShowingPostRequest) (*types.QueryShowingPostResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx
	// TODO: Process the query
	
	return &types.QueryShowingPostResponse{}, nil
}

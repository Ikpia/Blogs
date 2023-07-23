package keeper

import (
	"context"

	"blogs/x/blogs/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowingPost(goCtx context.Context, req *types.QueryShowingPostRequest) (*types.QueryShowingPostResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	post, found := k.GetPost(ctx, req.Id)
	comment, _ := k.GetComment(ctx, req.Id)
	
	
	
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	var fullpost = types.FullPost{
		Title: post.Title,
		Body: post.Body,
		Creator: post.Creator, 
		Comment: comment.Comment,
		Id: req.Id,
	} 

	return &types.QueryShowingPostResponse{
		fullpost,
		}, nil
}

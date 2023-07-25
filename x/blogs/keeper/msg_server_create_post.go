package keeper

import (
	"context"

	"blogs/x/blogs/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var post = types.Title{
		Creator: msg.Creator,
		Title:   msg.Title,
		Body:    msg.Body,
	}
	id := k.AppendPost(
		ctx,
		post,
	)
	var comment = types.Comment{
		Id:id,
	}
	k.SetCommentId(ctx, comment)
	// TODO: Handling the message

	return &types.MsgCreatePostResponse{
		Id: id,
	}, nil
}

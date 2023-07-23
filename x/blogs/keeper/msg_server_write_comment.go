package keeper

import (
	"blogs/x/blogs/types"
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) WriteComment(goCtx context.Context, msg *types.MsgWriteComment) (*types.MsgWriteCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	val, found := k.GetPost(ctx, msg.Id)
	value, _ := k.GetComment(ctx, msg.Id)
	var comments []string
	for _, value := range value.Comment {
		comments = append(comments, value)
	}
	comments = append(comments, msg.Comment)
	var post = types.Comment{
		Id: msg.Id,
		Comment: comments,
	}

	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator == val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Can't comment on your blog")
	}
	k.SetComment(ctx, post)
	// TODO: Handling the message
	return &types.MsgWriteCommentResponse{}, nil
}

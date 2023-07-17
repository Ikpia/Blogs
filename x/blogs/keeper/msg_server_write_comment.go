package keeper

import (
	"context"
	"fmt"
	"blogs/x/blogs/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) WriteComment(goCtx context.Context, msg *types.MsgWriteComment) (*types.MsgWriteCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	val, found := k.GetPost(ctx, msg.Id)
	var comments []string
	for _, value := range val.Comment {
		comments = append(comments, value)
	}
	comments = append(comments, msg.Comment)
	var post = types.Title{
		Title:   val.Title,
		Body:    val.Body,
		Creator: val.Creator,
		Id:      msg.Id,
		Comment: comments,
	}
	
	if !found {
        return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
    }
    if msg.Creator == val.Creator {
        return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Can't comment on your blog")
    }
	k.SetPost(ctx, post)
	// TODO: Handling the message


	return &types.MsgWriteCommentResponse{}, nil
}

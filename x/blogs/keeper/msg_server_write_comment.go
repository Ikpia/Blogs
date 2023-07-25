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
	//Get post of comment Id, since a comment must have the id of a valid post before it can be created
	post, found := k.GetPost(ctx, msg.Id)
	
	//if post does'nt exist comment should not be created
	if !found {
		fmt.Println("Post not found")
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	//the creator of the post should not be able to make comments on his post 
	if msg.Creator == post.Creator {
		fmt.Println("Can't comment on your blog")
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "owner cannot comment on his blog")
	}
	
	var comment = types.Review{
		Creator: msg.Creator,
		Comment: msg.Comment,
	}
	var content = types.Comment{
		Id: msg.Id,
	}
	k.SetComment(ctx, content, comment)//msg.Comment will be replaced by the newly created type
	// TODO: Handling the message
	return &types.MsgWriteCommentResponse{}, nil
}


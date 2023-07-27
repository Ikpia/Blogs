package keeper

import (
	"blogs/x/blogs/types"
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AppendPost(ctx sdk.Context, post types.Title) uint64 {
	count := k.GetPostCount(ctx)
	post.Id = count
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	appendedValue := k.cdc.MustMarshal(&post)
	store.Set(GetPostIDBytes(post.Id), appendedValue)
	k.SetPostCount(ctx, count+1)
	return count
}

func (k Keeper) GetPostCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PostCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}


func GetPostIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func (k Keeper) SetPostCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.PostCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) GetPost(ctx sdk.Context, id uint64) (val types.Title, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	b := store.Get(GetPostIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetComment(ctx sdk.Context, id uint64) (value types.Comment, find bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CommentKey))
	b := store.Get(GetPostIDBytes(id))
	if b == nil {
		return value, false
	}
	k.cdc.MustUnmarshal(b, &value)
	return value, true
}

func (k Keeper) SetPost(ctx sdk.Context, post types.Title) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PostKey))
	b := k.cdc.MustMarshal(&post)
	store.Set(GetPostIDBytes(post.Id), b)
}

func (k Keeper) SetCommentId(ctx sdk.Context, comments types.Comment) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CommentKey))
	b := k.cdc.MustMarshal(&comments)
	store.Set(GetPostIDBytes(comments.Id), b)
}

func (k Keeper) SetComment(ctx sdk.Context, content types.Comment, comment types.Review) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CommentKey))
	existing_comment := store.Get(GetPostIDBytes(content.Id))
		k.cdc.MustUnmarshal(existing_comment, &content)
		content.Content = append(content.Content, &comment)
		modified_comment := k.cdc.MustMarshal(&content)
		store.Set(GetPostIDBytes(content.Id), modified_comment)
	}

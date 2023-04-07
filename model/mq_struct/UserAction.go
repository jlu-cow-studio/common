package mq_struct

type UserActionOp string

const (
	UserActionOp_AddFavorite = "add_favorite"
	UserActionOp_DelFavorite = "del_favorite"

	UserActionOp_Follow   = "follow"
	UserActionOp_UnFollow = "unfollow"

	UserActionExtraKey_UserId      = "user_id"
	UserActionExtraKey_ItemId      = "item_id"
	UserActionExtraKey_FollowerId  = "follower_id"
	UserActionExtraKey_FollowingId = "following_id"
)

type UserActionMsg struct {
	Op    UserActionOp      `json:"op"`
	Extra map[string]string `json:"extra"`
}

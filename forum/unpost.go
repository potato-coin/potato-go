package forum

import (
	potato "github.com/rise-worlds/potato-go"
)

// NewUnPost is an action undoing a post that is active
func NewUnPost(poster potato.AccountName, postUUID string) *potato.Action {
	a := &potato.Action{
		Account: ForumAN,
		Name:    ActN("unpost"),
		Authorization: []potato.PermissionLevel{
			{Actor: poster, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(UnPost{
			Poster:   poster,
			PostUUID: postUUID,
		}),
	}
	return a
}

// UnPost represents the `poc.forum::unpost` action.
type UnPost struct {
	Poster   potato.AccountName `json:"poster"`
	PostUUID string          `json:"post_uuid"`
}

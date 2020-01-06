package forum

import (
	potato "github.com/rise-worlds/potato-go"
)

// NewPost is an action representing a simple message to be posted
// through the chain network.
func NewPost(poster potato.AccountName, postUUID, content string, replyToPoster potato.AccountName, replyToPostUUID string, certify bool, jsonMetadata string) *potato.Action {
	a := &potato.Action{
		Account: ForumAN,
		Name:    ActN("post"),
		Authorization: []potato.PermissionLevel{
			{Actor: poster, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(Post{
			Poster:          poster,
			PostUUID:        postUUID,
			Content:         content,
			ReplyToPoster:   replyToPoster,
			ReplyToPostUUID: replyToPostUUID,
			Certify:         certify,
			JSONMetadata:    jsonMetadata,
		}),
	}
	return a
}

// Post represents the `poc.forum::post` action.
type Post struct {
	Poster          potato.AccountName `json:"poster"`
	PostUUID        string          `json:"post_uuid"`
	Content         string          `json:"content"`
	ReplyToPoster   potato.AccountName `json:"reply_to_poster"`
	ReplyToPostUUID string          `json:"reply_to_post_uuid"`
	Certify         bool            `json:"certify"`
	JSONMetadata    string          `json:"json_metadata"`
}

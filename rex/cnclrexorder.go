package rex

import (
	potato "github.com/rise-worlds/potato-go"
)

func NewCancelREXOrder(owner potato.AccountName) *potato.Action {
	return &potato.Action{
		Account: REXAN,
		Name:    ActN("cnclrexorder"),
		Authorization: []potato.PermissionLevel{
			{Actor: owner, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(CancelREXOrder{
			Owner: owner,
		}),
	}
}

type CancelREXOrder struct {
	Owner potato.AccountName
}

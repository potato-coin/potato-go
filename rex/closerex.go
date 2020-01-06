package rex

import (
	potato "github.com/rise-worlds/potato-go"
)

func NewCloseREX(owner potato.AccountName) *potato.Action {
	return &potato.Action{
		Account: REXAN,
		Name:    ActN("closerex"),
		Authorization: []potato.PermissionLevel{
			{Actor: owner, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(CloseREX{
			Ownwer: owner,
		}),
	}
}

type CloseREX struct {
	Ownwer potato.AccountName
}

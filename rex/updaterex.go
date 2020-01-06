package rex

import (
	potato "github.com/rise-worlds/potato-go"
)

func NewUpdateREX(owner potato.AccountName) *potato.Action {
	return &potato.Action{
		Account: REXAN,
		Name:    ActN("updaterex"),
		Authorization: []potato.PermissionLevel{
			{Actor: owner, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(UpdateREX{
			Owner: owner,
		}),
	}
}

type UpdateREX struct {
	Owner potato.AccountName
}

package rex

import (
	potato "github.com/rise-worlds/potato-go"
)

func NewConsolidate(owner potato.AccountName) *potato.Action {
	return &potato.Action{
		Account: REXAN,
		Name:    ActN("consolidate"),
		Authorization: []potato.PermissionLevel{
			{Actor: owner, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(Consolidate{
			Owner: owner,
		}),
	}
}

type Consolidate struct {
	Owner potato.AccountName
}

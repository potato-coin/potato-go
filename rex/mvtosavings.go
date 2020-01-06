package rex

import (
	potato "github.com/rise-worlds/potato-go"
)

func NewMoveToSavings(owner potato.AccountName, rex potato.Asset) *potato.Action {
	return &potato.Action{
		Account: REXAN,
		Name:    ActN("mvtosavings"),
		Authorization: []potato.PermissionLevel{
			{Actor: owner, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(MoveToSavings{
			Owner: owner,
			REX:   rex,
		}),
	}
}

type MoveToSavings struct {
	Owner potato.AccountName
	REX   potato.Asset
}

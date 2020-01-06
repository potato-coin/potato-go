package rex

import (
	potato "github.com/rise-worlds/potato-go"
)

func NewMoveFromSavings(owner potato.AccountName, rex potato.Asset) *potato.Action {
	return &potato.Action{
		Account: REXAN,
		Name:    ActN("mvfrsavings"),
		Authorization: []potato.PermissionLevel{
			{Actor: owner, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(MoveFromSavings{
			Owner: owner,
			REX:   rex,
		}),
	}
}

type MoveFromSavings struct {
	Owner potato.AccountName
	REX   potato.Asset
}

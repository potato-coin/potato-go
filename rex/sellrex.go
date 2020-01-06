package rex

import (
	potato "github.com/rise-worlds/potato-go"
)

func NewSellREX(from potato.AccountName, rex potato.Asset) *potato.Action {
	return &potato.Action{
		Account: REXAN,
		Name:    ActN("sellrex"),
		Authorization: []potato.PermissionLevel{
			{Actor: from, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(SellREX{
			From: from,
			REX:  rex,
		}),
	}
}

type SellREX struct {
	From potato.AccountName
	REX  potato.Asset
}

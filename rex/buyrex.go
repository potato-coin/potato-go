package rex

import (
	potato "github.com/rise-worlds/potato-go"
)

func NewBuyREX(from potato.AccountName, amount potato.Asset) *potato.Action {
	return &potato.Action{
		Account: REXAN,
		Name:    ActN("buyrex"),
		Authorization: []potato.PermissionLevel{
			{Actor: from, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(BuyREX{
			From:   from,
			Amount: amount,
		}),
	}
}

type BuyREX struct {
	From   potato.AccountName
	Amount potato.Asset
}

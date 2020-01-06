package rex

import (
	potato "github.com/rise-worlds/potato-go"
)

func NewDeposit(owner potato.AccountName, amount potato.Asset) *potato.Action {
	return &potato.Action{
		Account: REXAN,
		Name:    ActN("deposit"),
		Authorization: []potato.PermissionLevel{
			{Actor: owner, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(Deposit{
			Owner:  owner,
			Amount: amount,
		}),
	}
}

type Deposit struct {
	Owner  potato.AccountName
	Amount potato.Asset
}

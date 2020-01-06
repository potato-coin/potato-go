package rex

import (
	potato "github.com/rise-worlds/potato-go"
)

func NewWithdraw(owner potato.AccountName, amount potato.Asset) *potato.Action {
	return &potato.Action{
		Account: REXAN,
		Name:    ActN("withdraw"),
		Authorization: []potato.PermissionLevel{
			{Actor: owner, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(Withdraw{
			Owner:  owner,
			Amount: amount,
		}),
	}
}

type Withdraw struct {
	Owner  potato.AccountName
	Amount potato.Asset
}

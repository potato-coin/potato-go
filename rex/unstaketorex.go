package rex

import (
	potato "github.com/rise-worlds/potato-go"
)

func NewUnstakeToREX(
	owner potato.AccountName,
	receiver potato.AccountName,
	fromNet potato.Asset,
	fromCPU potato.Asset,
) *potato.Action {
	return &potato.Action{
		Account: REXAN,
		Name:    ActN("unstaketorex"),
		Authorization: []potato.PermissionLevel{
			{Actor: owner, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(UnstakeToREX{
			Owner:    owner,
			Receiver: receiver,
			FromNet:  fromNet,
			FromCPU:  fromCPU,
		}),
	}
}

type UnstakeToREX struct {
	Owner    potato.AccountName
	Receiver potato.AccountName
	FromNet  potato.Asset
	FromCPU  potato.Asset
}

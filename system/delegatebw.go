package system

import (
	potato "github.com/rise-worlds/potato-go"
)

// NewDelegateBW returns a `delegatebw` action that lives on the
// `poc.system` contract.
func NewDelegateBW(from, receiver potato.AccountName, stakeCPU, stakeNet potato.Asset, transfer bool) *potato.Action {
	return &potato.Action{
		Account: AN("potato"),
		Name:    ActN("delegatebw"),
		Authorization: []potato.PermissionLevel{
			{Actor: from, Permission: PN("active")},
		},
		ActionData: potato.NewActionData(DelegateBW{
			From:     from,
			Receiver: receiver,
			StakeNet: stakeNet,
			StakeCPU: stakeCPU,
			Transfer: potato.Bool(transfer),
		}),
	}
}

// DelegateBW represents the `poc.system::delegatebw` action.
type DelegateBW struct {
	From     potato.AccountName `json:"from"`
	Receiver potato.AccountName `json:"receiver"`
	StakeNet potato.Asset       `json:"stake_net"`
	StakeCPU potato.Asset       `json:"stake_cpu"`
	Transfer potato.Bool        `json:"transfer"`
}

package system

import (
	potato "github.com/rise-worlds/potato-go"
)

// UnDelegateBW returns a `undelegatebw` action that lives on the
// `poc.system` contract.
func NewUnDelegateBW(from, receiver potato.AccountName, unstakeCPU, unstakeNet potato.Asset) *potato.Action {
	return &potato.Action{
		Account: AN("potato"),
		Name:    ActN("undelegatebw"),
		Authorization: []potato.PermissionLevel{
			{Actor: from, Permission: PN("active")},
		},
		ActionData: potato.NewActionData(UnDelegateBW{
			From:     from,
			Receiver: receiver,
			UnstakeNet: unstakeNet,
			UnstakeCPU: unstakeCPU,
		}),
	}
}

// UnDelegateBW represents the `poc.system::undelegatebw` action.
type UnDelegateBW struct {
	From         potato.AccountName `json:"from"`
	Receiver     potato.AccountName `json:"receiver"`
	UnstakeNet   potato.Asset       `json:"unstake_net_quantity"`
	UnstakeCPU   potato.Asset       `json:"unstake_cpu_quantity"`
}

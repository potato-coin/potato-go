package system

import (
	potato "github.com/rise-worlds/potato-go"
)

// NewSetalimits sets the account limits. Requires signature from `potato@active` account.
func NewSetalimits(account potato.AccountName, ramBytes, netWeight, cpuWeight int64) *potato.Action {
	a := &potato.Action{
		Account: AN("potato"),
		Name:    ActN("setalimit"),
		Authorization: []potato.PermissionLevel{
			{Actor: potato.AccountName("potato"), Permission: PN("active")},
		},
		ActionData: potato.NewActionData(Setalimits{
			Account:   account,
			RAMBytes:  ramBytes,
			NetWeight: netWeight,
			CPUWeight: cpuWeight,
		}),
	}
	return a
}

// Setalimits represents the `poc.system::setalimit` action.
type Setalimits struct {
	Account   potato.AccountName `json:"account"`
	RAMBytes  int64           `json:"ram_bytes"`
	NetWeight int64           `json:"net_weight"`
	CPUWeight int64           `json:"cpu_weight"`
}

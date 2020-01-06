package system

import (
	potato "github.com/rise-worlds/potato-go"
)

// NewSellRAM will sell at current market price a given number of
// bytes of RAM.
func NewSellRAM(account potato.AccountName, bytes uint64) *potato.Action {
	a := &potato.Action{
		Account: AN("potato"),
		Name:    ActN("sellram"),
		Authorization: []potato.PermissionLevel{
			{Actor: account, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(SellRAM{
			Account: account,
			Bytes:   bytes,
		}),
	}
	return a
}

// SellRAM represents the `poc.system::sellram` action.
type SellRAM struct {
	Account potato.AccountName `json:"account"`
	Bytes   uint64          `json:"bytes"`
}

package system

import (
	potato "github.com/rise-worlds/potato-go"
)

// NewRemoveProducer returns a `rmvproducer` action that lives on the
// `poc.system` contract.  This is to be called by the consortium of
// BPs, to oust a BP from its place.  If you want to unregister
// yourself as a BP, use `unregproducer`.
func NewRemoveProducer(producer potato.AccountName) *potato.Action {
	return &potato.Action{
		Account: AN("potato"),
		Name:    ActN("rmvproducer"),
		Authorization: []potato.PermissionLevel{
			{Actor: AN("potato"), Permission: PN("active")},
		},
		ActionData: potato.NewActionData(RemoveProducer{
			Producer: producer,
		}),
	}
}

// RemoveProducer represents the `poc.system::rmvproducer` action
type RemoveProducer struct {
	Producer potato.AccountName `json:"producer"`
}

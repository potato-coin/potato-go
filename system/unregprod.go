package system

import (
	potato "github.com/rise-worlds/potato-go"
)

// NewUnregProducer returns a `unregproducer` action that lives on the
// `poc.system` contract.
func NewUnregProducer(producer potato.AccountName) *potato.Action {
	return &potato.Action{
		Account: AN("potato"),
		Name:    ActN("unregproducer"),
		Authorization: []potato.PermissionLevel{
			{Actor: producer, Permission: PN("active")},
		},
		ActionData: potato.NewActionData(UnregProducer{
			Producer: producer,
		}),
	}
}

// UnregProducer represents the `poc.system::unregproducer` action
type UnregProducer struct {
	Producer potato.AccountName `json:"producer"`
}

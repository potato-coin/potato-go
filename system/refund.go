package system

import (
	potato "github.com/rise-worlds/potato-go"
)

// NewRefund returns a `refund` action that lives on the
// `poc.system` contract.
func NewRefund(owner potato.AccountName) *potato.Action {
	return &potato.Action{
		Account: AN("potato"),
		Name:    ActN("refund"),
		Authorization: []potato.PermissionLevel{
			{Actor: owner, Permission: PN("active")},
		},
		ActionData: potato.NewActionData(Refund{
			Owner: owner,
		}),
	}
}

// Refund represents the `poc.system::refund` action
type Refund struct {
	Owner potato.AccountName `json:"owner"`
}

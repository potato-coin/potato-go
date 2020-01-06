package sudo

import (
	potato "github.com/rise-worlds/potato-go"
)

// NewExec creates an `exec` action, found in the `poc.wrap`
// contract.
//
// Given an `potato.Transaction`, call `potato.MarshalBinary` on it first,
// pass the resulting bytes as `potato.HexBytes` here.
func NewExec(executer potato.AccountName, transaction potato.Transaction) *potato.Action {
	a := &potato.Action{
		Account: potato.AccountName("poc.wrap"),
		Name:    potato.ActionName("exec"),
		Authorization: []potato.PermissionLevel{
			{Actor: executer, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(Exec{
			Executer:    executer,
			Transaction: transaction,
		}),
	}
	return a
}

// Exec represents the `poc.system::exec` action.
type Exec struct {
	Executer    potato.AccountName `json:"executer"`
	Transaction potato.Transaction `json:"trx"`
}

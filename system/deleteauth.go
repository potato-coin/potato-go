package system

import "github.com/rise-worlds/potato-go"

// NewDeleteAuth creates an action from the `poc.system` contract
// called `deleteauth`.
//
// You cannot delete the `owner` or `active` permissions.  Also, if a
// permission is still linked through a previous `updatelink` action,
// you will need to `unlinkauth` first.
func NewDeleteAuth(account potato.AccountName, permission potato.PermissionName) *potato.Action {
	a := &potato.Action{
		Account: AN("potato"),
		Name:    ActN("deleteauth"),
		Authorization: []potato.PermissionLevel{
			{Actor: account, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(DeleteAuth{
			Account:    account,
			Permission: permission,
		}),
	}

	return a
}

// DeleteAuth represents the native `deleteauth` action, reachable
// through the `poc.system` contract.
type DeleteAuth struct {
	Account    potato.AccountName    `json:"account"`
	Permission potato.PermissionName `json:"permission"`
}

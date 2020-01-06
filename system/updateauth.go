package system

import "github.com/rise-worlds/potato-go"

// NewUpdateAuth creates an action from the `poc.system` contract
// called `updateauth`.
//
// usingPermission needs to be `owner` if you want to modify the
// `owner` authorization, otherwise `active` will do for the rest.
func NewUpdateAuth(account potato.AccountName, permission, parent potato.PermissionName, authority potato.Authority, usingPermission potato.PermissionName) *potato.Action {
	a := &potato.Action{
		Account: AN("potato"),
		Name:    ActN("updateauth"),
		Authorization: []potato.PermissionLevel{
			{
				Actor:      account,
				Permission: usingPermission,
			},
		},
		ActionData: potato.NewActionData(UpdateAuth{
			Account:    account,
			Permission: permission,
			Parent:     parent,
			Auth:       authority,
		}),
	}

	return a
}

// UpdateAuth represents the hard-coded `updateauth` action.
//
// If you change the `active` permission, `owner` is the required parent.
//
// If you change the `owner` permission, there should be no parent.
type UpdateAuth struct {
	Account    potato.AccountName    `json:"account"`
	Permission potato.PermissionName `json:"permission"`
	Parent     potato.PermissionName `json:"parent"`
	Auth       potato.Authority      `json:"auth"`
}

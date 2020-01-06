package system

import "github.com/rise-worlds/potato-go"

// NewUnlinkAuth creates an action from the `poc.system` contract
// called `unlinkauth`.
//
// `unlinkauth` detaches a previously set permission from a
// `code::actionName`. See `linkauth`.
func NewUnlinkAuth(account, code potato.AccountName, actionName potato.ActionName) *potato.Action {
	a := &potato.Action{
		Account: AN("potato"),
		Name:    ActN("unlinkauth"),
		Authorization: []potato.PermissionLevel{
			{
				Actor:      account,
				Permission: potato.PermissionName("active"),
			},
		},
		ActionData: potato.NewActionData(UnlinkAuth{
			Account: account,
			Code:    code,
			Type:    actionName,
		}),
	}

	return a
}

// UnlinkAuth represents the native `unlinkauth` action, through the
// system contract.
type UnlinkAuth struct {
	Account potato.AccountName `json:"account"`
	Code    potato.AccountName `json:"code"`
	Type    potato.ActionName  `json:"type"`
}

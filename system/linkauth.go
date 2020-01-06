package system

import "github.com/rise-worlds/potato-go"

// NewLinkAuth creates an action from the `poc.system` contract
// called `linkauth`.
//
// `linkauth` allows you to attach certain permission to the given
// `code::actionName`. With this set on-chain, you can use the
// `requiredPermission` to sign transactions for `code::actionName`
// and not rely on your `active` (which might be more sensitive as it
// can sign anything) for the given operation.
func NewLinkAuth(account, code potato.AccountName, actionName potato.ActionName, requiredPermission potato.PermissionName) *potato.Action {
	a := &potato.Action{
		Account: AN("potato"),
		Name:    ActN("linkauth"),
		Authorization: []potato.PermissionLevel{
			{account, potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(LinkAuth{
			Account:     account,
			Code:        code,
			Type:        actionName,
			Requirement: requiredPermission,
		}),
	}

	return a
}

// LinkAuth represents the native `linkauth` action, through the
// system contract.
type LinkAuth struct {
	Account     potato.AccountName    `json:"account"`
	Code        potato.AccountName    `json:"code"`
	Type        potato.ActionName     `json:"type"`
	Requirement potato.PermissionName `json:"requirement"`
}

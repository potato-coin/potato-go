package system

import potato "github.com/rise-worlds/potato-go"

// NewSetPriv returns a `setpriv` action that lives on the
// `poc.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `potato-bios` boot process by the
// `poc.system` contract.
func NewSetPriv(account potato.AccountName) *potato.Action {
	a := &potato.Action{
		Account: AN("potato"),
		Name:    ActN("setpriv"),
		Authorization: []potato.PermissionLevel{
			{Actor: AN("potato"), Permission: PN("active")},
		},
		ActionData: potato.NewActionData(SetPriv{
			Account: account,
			IsPriv:  potato.Bool(true),
		}),
	}
	return a
}

// SetPriv sets privileged account status. Used in the bios boot mechanism.
type SetPriv struct {
	Account potato.AccountName `json:"account"`
	IsPriv  potato.Bool        `json:"is_priv"`
}

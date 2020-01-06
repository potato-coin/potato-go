package system

import "github.com/rise-worlds/potato-go"

// NewNonce returns a `nonce` action that lives on the
// `poc.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `potato-bios` boot process by the
// `poc.system` contract.
func NewNonce(nonce string) *potato.Action {
	a := &potato.Action{
		Account:       AN("potato"),
		Name:          ActN("nonce"),
		Authorization: []potato.PermissionLevel{
			//{Actor: AN("potato"), Permission: PN("active")},
		},
		ActionData: potato.NewActionData(Nonce{
			Value: nonce,
		}),
	}
	return a
}

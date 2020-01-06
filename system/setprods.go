package system

import (
	potato "github.com/rise-worlds/potato-go"
	"github.com/rise-worlds/potato-go/ecc"
)

// NewSetPriv returns a `setpriv` action that lives on the
// `poc.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `potato-bios` boot process by the
// `poc.system` contract.
func NewSetProds(producers []ProducerKey) *potato.Action {
	a := &potato.Action{
		Account: AN("potato"),
		Name:    ActN("setprods"),
		Authorization: []potato.PermissionLevel{
			{Actor: AN("potato"), Permission: PN("active")},
		},
		ActionData: potato.NewActionData(SetProds{
			Schedule: producers,
		}),
	}
	return a
}

// SetProds is present in `potato.bios` contract. Used only at boot time.
type SetProds struct {
	Schedule []ProducerKey `json:"schedule"`
}

type ProducerKey struct {
	ProducerName    potato.AccountName `json:"producer_name"`
	BlockSigningKey ecc.PublicKey   `json:"block_signing_key"`
}

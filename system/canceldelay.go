package system

import "github.com/rise-worlds/potato-go"

// NewCancelDelay creates an action from the `poc.system` contract
// called `canceldelay`.
//
// `canceldelay` allows you to cancel a deferred transaction,
// previously sent to the chain with a `delay_sec` larger than 0.  You
// need to sign with cancelingAuth, to cancel a transaction signed
// with that same authority.
func NewCancelDelay(cancelingAuth potato.PermissionLevel, transactionID potato.Checksum256) *potato.Action {
	a := &potato.Action{
		Account: AN("potato"),
		Name:    ActN("canceldelay"),
		Authorization: []potato.PermissionLevel{
			cancelingAuth,
		},
		ActionData: potato.NewActionData(CancelDelay{
			CancelingAuth: cancelingAuth,
			TransactionID: transactionID,
		}),
	}

	return a
}

// CancelDelay represents the native `canceldelay` action, through the
// system contract.
type CancelDelay struct {
	CancelingAuth potato.PermissionLevel `json:"canceling_auth"`
	TransactionID potato.Checksum256     `json:"trx_id"`
}

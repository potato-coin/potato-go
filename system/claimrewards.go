package system

import (
	potato "github.com/rise-worlds/potato-go"
)

// NewClaimRewards will buy at current market price a given number of
// bytes of RAM, and grant them to the `receiver` account.
func NewClaimRewards(owner potato.AccountName) *potato.Action {
	a := &potato.Action{
		Account: AN("potato"),
		Name:    ActN("claimwards"),
		Authorization: []potato.PermissionLevel{
			{Actor: owner, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(ClaimRewards{
			Owner: owner,
		}),
	}
	return a
}

// ClaimRewards represents the `poc.system::claimwards` action.
type ClaimRewards struct {
	Owner potato.AccountName `json:"owner"`
}

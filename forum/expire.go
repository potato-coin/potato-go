package forum

import (
	potato "github.com/rise-worlds/potato-go"
)

// NewExpire is an action to expire a proposal ahead of its natural death.
func NewExpire(proposer potato.AccountName, proposalName potato.Name) *potato.Action {
	a := &potato.Action{
		Account: ForumAN,
		Name:    ActN("expire"),
		Authorization: []potato.PermissionLevel{
			{Actor: proposer, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(Expire{
			ProposalName: proposalName,
		}),
	}
	return a
}

// Expire represents the `poc.forum::propose` action.
type Expire struct {
	ProposalName potato.Name `json:"proposal_name"`
}

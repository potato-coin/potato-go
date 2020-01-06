package forum

import (
	potato "github.com/rise-worlds/potato-go"
)

// NewUnVote is an action representing the action to undoing a current vote
func NewUnVote(voter potato.AccountName, proposalName potato.Name) *potato.Action {
	a := &potato.Action{
		Account: ForumAN,
		Name:    ActN("unvote"),
		Authorization: []potato.PermissionLevel{
			{Actor: voter, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(UnVote{
			Voter:        voter,
			ProposalName: proposalName,
		}),
	}
	return a
}

// UnVote represents the `poc.forum::unvote` action.
type UnVote struct {
	Voter        potato.AccountName `json:"voter"`
	ProposalName potato.Name        `json:"proposal_name"`
}

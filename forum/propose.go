package forum

import (
	potato "github.com/rise-worlds/potato-go"
)

// NewPropose is an action to submit a proposal for vote.
func NewPropose(proposer potato.AccountName, proposalName potato.Name, title string, proposalJSON string, expiresAt potato.JSONTime) *potato.Action {
	a := &potato.Action{
		Account: ForumAN,
		Name:    ActN("propose"),
		Authorization: []potato.PermissionLevel{
			{Actor: proposer, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(Propose{
			Proposer:     proposer,
			ProposalName: proposalName,
			Title:        title,
			ProposalJSON: proposalJSON,
			ExpiresAt:    expiresAt,
		}),
	}
	return a
}

// Propose represents the `poc.forum::propose` action.
type Propose struct {
	Proposer     potato.AccountName `json:"proposer"`
	ProposalName potato.Name        `json:"proposal_name"`
	Title        string          `json:"title"`
	ProposalJSON string          `json:"proposal_json"`
	ExpiresAt    potato.JSONTime    `json:"expires_at"`
}

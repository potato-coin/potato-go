package forum

import (
	potato "github.com/rise-worlds/potato-go"
)

// NewVote is an action representing a simple vote to be broadcast
// through the chain network.
func NewVote(voter potato.AccountName, proposalName potato.Name, voteValue uint8, voteJSON string) *potato.Action {
	a := &potato.Action{
		Account: ForumAN,
		Name:    ActN("vote"),
		Authorization: []potato.PermissionLevel{
			{Actor: voter, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(Vote{
			Voter:        voter,
			ProposalName: proposalName,
			Vote:         voteValue,
			VoteJSON:     voteJSON,
		}),
	}
	return a
}

// Vote represents the `poc.forum::vote` action.
type Vote struct {
	Voter        potato.AccountName `json:"voter"`
	ProposalName potato.Name        `json:"proposal_name"`
	Vote         uint8           `json:"vote"`
	VoteJSON     string          `json:"vote_json"`
}

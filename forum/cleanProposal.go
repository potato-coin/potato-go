package forum

import (
	potato "github.com/rise-worlds/potato-go"
)

// CleanProposal is an action to flush proposal and allow RAM used by it.
func NewCleanProposal(cleaner potato.AccountName, proposalName potato.Name, maxCount uint64) *potato.Action {
	a := &potato.Action{
		Account: ForumAN,
		Name:    ActN("clnproposal"),
		Authorization: []potato.PermissionLevel{
			{Actor: cleaner, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(CleanProposal{
			ProposalName: proposalName,
			MaxCount:     maxCount,
		}),
	}
	return a
}

// CleanProposal represents the `poc.forum::clnproposal` action.
type CleanProposal struct {
	ProposalName potato.Name `json:"proposal_name"`
	MaxCount     uint64   `json:"max_count"`
}

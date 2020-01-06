package msig

import (
	potato "github.com/rise-worlds/potato-go"
)

// NewCancel returns a `cancel` action that lives on the
// `poc.msig` contract.
func NewCancel(proposer potato.AccountName, proposalName potato.Name, canceler potato.AccountName) *potato.Action {
	return &potato.Action{
		Account: potato.AccountName("poc.msig"),
		Name:    potato.ActionName("cancel"),
		// TODO: double check in this package that the `Actor` is always the `proposer`..
		Authorization: []potato.PermissionLevel{
			{Actor: canceler, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(Cancel{proposer, proposalName, canceler}),
	}
}

type Cancel struct {
	Proposer     potato.AccountName `json:"proposer"`
	ProposalName potato.Name        `json:"proposal_name"`
	Canceler     potato.AccountName `json:"canceler"`
}

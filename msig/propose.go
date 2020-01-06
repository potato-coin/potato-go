package msig

import (
	potato "github.com/rise-worlds/potato-go"
)

// NewPropose returns a `propose` action that lives on the
// `poc.msig` contract.
func NewPropose(proposer potato.AccountName, proposalName potato.Name, requested []potato.PermissionLevel, transaction *potato.Transaction) *potato.Action {
	return &potato.Action{
		Account: potato.AccountName("poc.msig"),
		Name:    potato.ActionName("propose"),
		Authorization: []potato.PermissionLevel{
			{Actor: proposer, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(Propose{proposer, proposalName, requested, transaction}),
	}
}

type Propose struct {
	Proposer     potato.AccountName       `json:"proposer"`
	ProposalName potato.Name              `json:"proposal_name"`
	Requested    []potato.PermissionLevel `json:"requested"`
	Transaction  *potato.Transaction      `json:"trx"`
}

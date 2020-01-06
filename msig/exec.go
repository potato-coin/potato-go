package msig

import (
	potato "github.com/rise-worlds/potato-go"
)

// NewExec returns a `exec` action that lives on the
// `poc.msig` contract.
func NewExec(proposer potato.AccountName, proposalName potato.Name, executer potato.AccountName) *potato.Action {
	return &potato.Action{
		Account: potato.AccountName("poc.msig"),
		Name:    potato.ActionName("exec"),
		// TODO: double check in this package that the `Actor` is always the `proposer`..
		Authorization: []potato.PermissionLevel{
			{Actor: executer, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(Exec{proposer, proposalName, executer}),
	}
}

type Exec struct {
	Proposer     potato.AccountName `json:"proposer"`
	ProposalName potato.Name        `json:"proposal_name"`
	Executer     potato.AccountName `json:"executer"`
}

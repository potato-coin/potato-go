package msig

import (
	potato "github.com/rise-worlds/potato-go"
)

// NewApprove returns a `approve` action that lives on the
// `poc.msig` contract.
func NewApprove(proposer potato.AccountName, proposalName potato.Name, level potato.PermissionLevel) *potato.Action {
	return &potato.Action{
		Account:       potato.AccountName("poc.msig"),
		Name:          potato.ActionName("approve"),
		Authorization: []potato.PermissionLevel{level},
		ActionData:    potato.NewActionData(Approve{proposer, proposalName, level}),
	}
}

type Approve struct {
	Proposer     potato.AccountName     `json:"proposer"`
	ProposalName potato.Name            `json:"proposal_name"`
	Level        potato.PermissionLevel `json:"level"`
}

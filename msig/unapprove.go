package msig

import (
	potato "github.com/rise-worlds/potato-go"
)

// NewUnapprove returns a `unapprove` action that lives on the
// `poc.msig` contract.
func NewUnapprove(proposer potato.AccountName, proposalName potato.Name, level potato.PermissionLevel) *potato.Action {
	return &potato.Action{
		Account:       potato.AccountName("poc.msig"),
		Name:          potato.ActionName("unapprove"),
		Authorization: []potato.PermissionLevel{level},
		ActionData:    potato.NewActionData(Unapprove{proposer, proposalName, level}),
	}
}

type Unapprove struct {
	Proposer     potato.AccountName     `json:"proposer"`
	ProposalName potato.Name            `json:"proposal_name"`
	Level        potato.PermissionLevel `json:"level"`
}

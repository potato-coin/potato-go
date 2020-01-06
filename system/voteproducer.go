package system

import "github.com/rise-worlds/potato-go"

// NewNonce returns a `nonce` action that lives on the
// `poc.bios` contract. It should exist only when booting a new
// network, as it is replaced using the `potato-bios` boot process by the
// `poc.system` contract.
func NewVoteProducer(voter potato.AccountName, proxy potato.AccountName, producers ...potato.AccountName) *potato.Action {
	a := &potato.Action{
		Account: AN("potato"),
		Name:    ActN("voteproducer"),
		Authorization: []potato.PermissionLevel{
			{Actor: voter, Permission: PN("active")},
		},
		ActionData: potato.NewActionData(
			VoteProducer{
				Voter:     voter,
				Proxy:     proxy,
				Producers: producers,
			},
		),
	}
	return a
}

// VoteProducer represents the `poc.system::voteproducer` action
type VoteProducer struct {
	Voter     potato.AccountName   `json:"voter"`
	Proxy     potato.AccountName   `json:"proxy"`
	Producers []potato.AccountName `json:"producers"`
}

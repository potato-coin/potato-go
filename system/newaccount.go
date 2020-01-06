package system

import (
	"github.com/rise-worlds/potato-go"
	"github.com/rise-worlds/potato-go/ecc"
)

// NewNewAccount returns a `newaccount` action that lives on the
// `poc.system` contract.
func NewNewAccount(creator, newAccount potato.AccountName, publicKey ecc.PublicKey) *potato.Action {
	return &potato.Action{
		Account: AN("potato"),
		Name:    ActN("newaccount"),
		Authorization: []potato.PermissionLevel{
			{Actor: creator, Permission: PN("active")},
		},
		ActionData: potato.NewActionData(NewAccount{
			Creator: creator,
			Name:    newAccount,
			Owner: potato.Authority{
				Threshold: 1,
				Keys: []potato.KeyWeight{
					{
						PublicKey: publicKey,
						Weight:    1,
					},
				},
				Accounts: []potato.PermissionLevelWeight{},
			},
			Active: potato.Authority{
				Threshold: 1,
				Keys: []potato.KeyWeight{
					{
						PublicKey: publicKey,
						Weight:    1,
					},
				},
				Accounts: []potato.PermissionLevelWeight{},
			},
		}),
	}
}

// NewDelegatedNewAccount returns a `newaccount` action that lives on the
// `poc.system` contract. It is filled with an authority structure that
// delegates full control of the new account to an already existing account.
func NewDelegatedNewAccount(creator, newAccount potato.AccountName, delegatedTo potato.AccountName) *potato.Action {
	return &potato.Action{
		Account: AN("potato"),
		Name:    ActN("newaccount"),
		Authorization: []potato.PermissionLevel{
			{Actor: creator, Permission: PN("active")},
		},
		ActionData: potato.NewActionData(NewAccount{
			Creator: creator,
			Name:    newAccount,
			Owner: potato.Authority{
				Threshold: 1,
				Keys:      []potato.KeyWeight{},
				Accounts: []potato.PermissionLevelWeight{
					potato.PermissionLevelWeight{
						Permission: potato.PermissionLevel{
							Actor:      delegatedTo,
							Permission: PN("active"),
						},
						Weight: 1,
					},
				},
			},
			Active: potato.Authority{
				Threshold: 1,
				Keys:      []potato.KeyWeight{},
				Accounts: []potato.PermissionLevelWeight{
					potato.PermissionLevelWeight{
						Permission: potato.PermissionLevel{
							Actor:      delegatedTo,
							Permission: PN("active"),
						},
						Weight: 1,
					},
				},
			},
		}),
	}
}

// NewCustomNewAccount returns a `newaccount` action that lives on the
// `poc.system` contract. You can specify your own `owner` and
// `active` permissions.
func NewCustomNewAccount(creator, newAccount potato.AccountName, owner, active potato.Authority) *potato.Action {
	return &potato.Action{
		Account: AN("potato"),
		Name:    ActN("newaccount"),
		Authorization: []potato.PermissionLevel{
			{Actor: creator, Permission: PN("active")},
		},
		ActionData: potato.NewActionData(NewAccount{
			Creator: creator,
			Name:    newAccount,
			Owner:   owner,
			Active:  active,
		}),
	}
}

// NewAccount represents a `newaccount` action on the `poc.system`
// contract. It is one of the rare ones to be hard-coded into the
// blockchain.
type NewAccount struct {
	Creator potato.AccountName `json:"creator"`
	Name    potato.AccountName `json:"name"`
	Owner   potato.Authority   `json:"owner"`
	Active  potato.Authority   `json:"active"`
}

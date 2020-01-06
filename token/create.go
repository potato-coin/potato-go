package token

import potato "github.com/rise-worlds/potato-go"

func NewCreate(issuer potato.AccountName, maxSupply potato.Asset) *potato.Action {
	return &potato.Action{
		Account: AN("poc.token"),
		Name:    ActN("create"),
		Authorization: []potato.PermissionLevel{
			{Actor: AN("poc.token"), Permission: PN("active")},
		},
		ActionData: potato.NewActionData(Create{
			Issuer:        issuer,
			MaximumSupply: maxSupply,
		}),
	}
}

// Create represents the `create` struct on the `poc.token` contract.
type Create struct {
	Issuer        potato.AccountName `json:"issuer"`
	MaximumSupply potato.Asset       `json:"maximum_supply"`
}

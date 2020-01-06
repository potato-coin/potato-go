package token

import potato "github.com/rise-worlds/potato-go"

func NewIssue(to potato.AccountName, quantity potato.Asset, memo string) *potato.Action {
	return &potato.Action{
		Account: AN("poc.token"),
		Name:    ActN("issue"),
		Authorization: []potato.PermissionLevel{
			{Actor: AN("potato"), Permission: PN("active")},
		},
		ActionData: potato.NewActionData(Issue{
			To:       to,
			Quantity: quantity,
			Memo:     memo,
		}),
	}
}

// Issue represents the `issue` struct on the `poc.token` contract.
type Issue struct {
	To       potato.AccountName `json:"to"`
	Quantity potato.Asset       `json:"quantity"`
	Memo     string          `json:"memo"`
}

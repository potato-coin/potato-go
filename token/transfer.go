package token

import potato "github.com/rise-worlds/potato-go"

func NewTransfer(from, to potato.AccountName, quantity potato.Asset, memo string) *potato.Action {
	return &potato.Action{
		Account: AN("poc.token"),
		Name:    ActN("transfer"),
		Authorization: []potato.PermissionLevel{
			{Actor: from, Permission: PN("active")},
		},
		ActionData: potato.NewActionData(Transfer{
			From:     from,
			To:       to,
			Quantity: quantity,
			Memo:     memo,
		}),
	}
}

// Transfer represents the `transfer` struct on `poc.token` contract.
type Transfer struct {
	From     potato.AccountName `json:"from"`
	To       potato.AccountName `json:"to"`
	Quantity potato.Asset       `json:"quantity"`
	Memo     string          `json:"memo"`
}

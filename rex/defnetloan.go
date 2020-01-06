package rex

import (
	potato "github.com/rise-worlds/potato-go"
)

func NewDefundNetLoan(from potato.AccountName, loanNumber uint64, amount potato.Asset) *potato.Action {
	return &potato.Action{
		Account: REXAN,
		Name:    ActN("defnetloan"),
		Authorization: []potato.PermissionLevel{
			{Actor: from, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(DefundNetLoan{
			From:       from,
			LoanNumber: loanNumber,
			Amount:     amount,
		}),
	}
}

type DefundNetLoan struct {
	From       potato.AccountName
	LoanNumber uint64
	Amount     potato.Asset
}

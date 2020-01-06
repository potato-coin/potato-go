package rex

import (
	potato "github.com/rise-worlds/potato-go"
)

func NewDefundCPULoan(from potato.AccountName, loanNumber uint64, amount potato.Asset) *potato.Action {
	return &potato.Action{
		Account: REXAN,
		Name:    ActN("defcpuloan"),
		Authorization: []potato.PermissionLevel{
			{Actor: from, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(DefundCPULoan{
			From:       from,
			LoanNumber: loanNumber,
			Amount:     amount,
		}),
	}
}

type DefundCPULoan struct {
	From       potato.AccountName
	LoanNumber uint64
	Amount     potato.Asset
}

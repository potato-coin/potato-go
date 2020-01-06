package rex

import (
	potato "github.com/rise-worlds/potato-go"
)

func NewFundCPULoan(from potato.AccountName, loanNumber uint64, payment potato.Asset) *potato.Action {
	return &potato.Action{
		Account: REXAN,
		Name:    ActN("fundcpuloan"),
		Authorization: []potato.PermissionLevel{
			{Actor: from, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(FundCPULoan{
			From:       from,
			LoanNumber: loanNumber,
			Payment:    payment,
		}),
	}
}

type FundCPULoan struct {
	From       potato.AccountName
	LoanNumber uint64
	Payment    potato.Asset
}

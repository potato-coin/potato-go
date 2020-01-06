package rex

import (
	potato "github.com/rise-worlds/potato-go"
)

func NewFundNetLoan(from potato.AccountName, loanNumber uint64, payment potato.Asset) *potato.Action {
	return &potato.Action{
		Account: REXAN,
		Name:    ActN("fundnetloan"),
		Authorization: []potato.PermissionLevel{
			{Actor: from, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(FundNetLoan{
			From:       from,
			LoanNumber: loanNumber,
			Payment:    payment,
		}),
	}
}

type FundNetLoan struct {
	From       potato.AccountName
	LoanNumber uint64
	Payment    potato.Asset
}

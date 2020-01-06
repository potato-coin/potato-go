package rex

import (
	potato "github.com/rise-worlds/potato-go"
)

func NewRentCPU(
	from potato.AccountName,
	receiver potato.AccountName,
	loanPayment potato.Asset,
	loanFund potato.Asset,
) *potato.Action {
	return &potato.Action{
		Account: REXAN,
		Name:    ActN("rentcpu"),
		Authorization: []potato.PermissionLevel{
			{Actor: from, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(RentCPU{
			From:        from,
			Receiver:    receiver,
			LoanPayment: loanPayment,
			LoanFund:    loanFund,
		}),
	}
}

type RentCPU struct {
	From        potato.AccountName
	Receiver    potato.AccountName
	LoanPayment potato.Asset
	LoanFund    potato.Asset
}

package rex

import (
	potato "github.com/rise-worlds/potato-go"
)

func NewRentNet(
	from potato.AccountName,
	receiver potato.AccountName,
	loanPayment potato.Asset,
	loanFund potato.Asset,
) *potato.Action {
	return &potato.Action{
		Account: REXAN,
		Name:    ActN("rentnet"),
		Authorization: []potato.PermissionLevel{
			{Actor: from, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(RentNet{
			From:        from,
			Receiver:    receiver,
			LoanPayment: loanPayment,
			LoanFund:    loanFund,
		}),
	}
}

type RentNet struct {
	From        potato.AccountName
	Receiver    potato.AccountName
	LoanPayment potato.Asset
	LoanFund    potato.Asset
}

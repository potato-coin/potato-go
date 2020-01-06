package system

import (
	potato "github.com/rise-worlds/potato-go"
)

// NewBuyRAMBytes will buy at current market price a given number of
// bytes of RAM, and grant them to the `receiver` account.
func NewBuyRAMBytes(payer, receiver potato.AccountName, bytes uint32) *potato.Action {
	a := &potato.Action{
		Account: AN("potato"),
		Name:    ActN("buyrambytes"),
		Authorization: []potato.PermissionLevel{
			{Actor: payer, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(BuyRAMBytes{
			Payer:    payer,
			Receiver: receiver,
			Bytes:    bytes,
		}),
	}
	return a
}

// BuyRAMBytes represents the `potato.system::buyrambytes` action.
type BuyRAMBytes struct {
	Payer    potato.AccountName `json:"payer"`
	Receiver potato.AccountName `json:"receiver"`
	Bytes    uint32          `json:"bytes"`
}

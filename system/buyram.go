package system

import (
	potato "github.com/rise-worlds/potato-go"
)

func NewBuyRAM(payer, receiver potato.AccountName, pocQuantity uint64) *potato.Action {
	a := &potato.Action{
		Account: AN("potato"),
		Name:    ActN("buyram"),
		Authorization: []potato.PermissionLevel{
			{Actor: payer, Permission: PN("active")},
		},
		ActionData: potato.NewActionData(BuyRAM{
			Payer:    payer,
			Receiver: receiver,
			Quantity: potato.NewPOCAsset(int64(pocQuantity)),
		}),
	}
	return a
}

// BuyRAM represents the `poc.system::buyram` action.
type BuyRAM struct {
	Payer    potato.AccountName `json:"payer"`
	Receiver potato.AccountName `json:"receiver"`
	Quantity potato.Asset       `json:"quant"` // specified in POC
}

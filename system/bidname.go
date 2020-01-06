package system

import (
	potato "github.com/rise-worlds/potato-go"
)

func NewBidname(bidder, newname potato.AccountName, bid potato.Asset) *potato.Action {
	a := &potato.Action{
		Account: AN("potato"),
		Name:    ActN("bidname"),
		Authorization: []potato.PermissionLevel{
			{Actor: bidder, Permission: PN("active")},
		},
		ActionData: potato.NewActionData(Bidname{
			Bidder:  bidder,
			Newname: newname,
			Bid:     bid,
		}),
	}
	return a
}

// Bidname represents the `potato.system_contract::bidname` action.
type Bidname struct {
	Bidder  potato.AccountName `json:"bidder"`
	Newname potato.AccountName `json:"newname"`
	Bid     potato.Asset       `json:"bid"` // specified in potato
}

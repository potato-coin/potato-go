package system

import (
	potato "github.com/rise-worlds/potato-go"
	"github.com/rise-worlds/potato-go/ecc"
)

// NewRegProducer returns a `regproducer` action that lives on the
// `poc.system` contract.
func NewRegProducer(producer potato.AccountName, producerKey ecc.PublicKey, url string, location uint16) *potato.Action {
	return &potato.Action{
		Account: AN("potato"),
		Name:    ActN("regproducer"),
		Authorization: []potato.PermissionLevel{
			{Actor: producer, Permission: PN("active")},
		},
		ActionData: potato.NewActionData(RegProducer{
			Producer:    producer,
			ProducerKey: producerKey,
			URL:         url,
			Location:    location,
		}),
	}
}

// RegProducer represents the `poc.system::regproducer` action
type RegProducer struct {
	Producer    potato.AccountName `json:"producer"`
	ProducerKey ecc.PublicKey   `json:"producer_key"`
	URL         string          `json:"url"`
	Location    uint16          `json:"location"` // what,s the meaning of that anyway ?
}

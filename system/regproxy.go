package system

import (
	potato "github.com/rise-worlds/potato-go"
)

// NewRegProxy returns a `regproxy` action that lives on the
// `poc.system` contract.
func NewRegProxy(proxy potato.AccountName, isProxy bool) *potato.Action {
	return &potato.Action{
		Account: AN("potato"),
		Name:    ActN("regproxy"),
		Authorization: []potato.PermissionLevel{
			{Actor: proxy, Permission: PN("active")},
		},
		ActionData: potato.NewActionData(RegProxy{
			Proxy:   proxy,
			IsProxy: isProxy,
		}),
	}
}

// RegProxy represents the `poc.system::regproxy` action
type RegProxy struct {
	Proxy   potato.AccountName `json:"proxy"`
	IsProxy bool            `json:"isproxy"`
}

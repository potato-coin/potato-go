package forum

import (
	potato "github.com/rise-worlds/potato-go"
)

// Status is an action to set a status update for a given account on the forum contract.
func NewStatus(account potato.AccountName, content string) *potato.Action {
	a := &potato.Action{
		Account: ForumAN,
		Name:    ActN("status"),
		Authorization: []potato.PermissionLevel{
			{Actor: account, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(Status{
			Account: account,
			Content: content,
		}),
	}
	return a
}

// Status represents the `poc.forum::status` action.
type Status struct {
	Account potato.AccountName `json:"account_name"`
	Content string          `json:"content"`
}

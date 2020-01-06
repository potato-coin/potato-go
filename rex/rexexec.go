package rex

import (
	potato "github.com/rise-worlds/potato-go"
)

func NewREXExec(user potato.AccountName, max uint16) *potato.Action {
	return &potato.Action{
		Account: REXAN,
		Name:    ActN("rexexec"),
		Authorization: []potato.PermissionLevel{
			{Actor: user, Permission: potato.PermissionName("active")},
		},
		ActionData: potato.NewActionData(REXExec{
			User: user,
			Max:  max,
		}),
	}
}

type REXExec struct {
	User potato.AccountName
	Max  uint16
}

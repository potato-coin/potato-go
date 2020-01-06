package system

import (
	potato "github.com/rise-worlds/potato-go"
)

func NewSetRAMRate(bytesPerBlock uint16) *potato.Action {
	a := &potato.Action{
		Account: AN("potato"),
		Name:    ActN("setram"),
		Authorization: []potato.PermissionLevel{
			{
				Actor:      AN("potato"),
				Permission: potato.PermissionName("active"),
			},
		},
		ActionData: potato.NewActionData(SetRAMRate{
			BytesPerBlock: bytesPerBlock,
		}),
	}
	return a
}

// SetRAMRate represents the system contract's `setramrate` action.
type SetRAMRate struct {
	BytesPerBlock uint16 `json:"bytes_per_block"`
}

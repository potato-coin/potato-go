package system

import (
	potato "github.com/rise-worlds/potato-go"
)

func NewSetRAM(maxRAMSize uint64) *potato.Action {
	a := &potato.Action{
		Account: AN("potato"),
		Name:    ActN("setram"),
		Authorization: []potato.PermissionLevel{
			{
				Actor:      AN("potato"),
				Permission: potato.PermissionName("active"),
			},
		},
		ActionData: potato.NewActionData(SetRAM{
			MaxRAMSize: maxRAMSize,
		}),
	}
	return a
}

// SetRAM represents the hard-coded `setram` action.
type SetRAM struct {
	MaxRAMSize potato.Uint64 `json:"max_ram_size"`
}

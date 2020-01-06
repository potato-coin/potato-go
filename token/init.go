package token

import "github.com/rise-worlds/potato-go"

func init() {
	potato.RegisterAction(AN("poc.token"), ActN("transfer"), Transfer{})
	potato.RegisterAction(AN("poc.token"), ActN("issue"), Issue{})
	potato.RegisterAction(AN("poc.token"), ActN("create"), Create{})
}

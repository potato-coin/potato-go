package msig

import (
	"github.com/rise-worlds/potato-go"
)

func init() {
	potato.RegisterAction(AN("poc.msig"), ActN("propose"), &Propose{})
	potato.RegisterAction(AN("poc.msig"), ActN("approve"), &Approve{})
	potato.RegisterAction(AN("poc.msig"), ActN("unapprove"), &Unapprove{})
	potato.RegisterAction(AN("poc.msig"), ActN("cancel"), &Cancel{})
	potato.RegisterAction(AN("poc.msig"), ActN("exec"), &Exec{})
}

var AN = potato.AN
var PN = potato.PN
var ActN = potato.ActN

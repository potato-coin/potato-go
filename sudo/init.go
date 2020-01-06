package sudo

import potato "github.com/rise-worlds/potato-go"

func init() {
	potato.RegisterAction(AN("poc.wrap"), ActN("exec"), Exec{})
}

var AN = potato.AN
var ActN = potato.ActN

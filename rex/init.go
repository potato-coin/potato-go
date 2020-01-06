package rex

import potato "github.com/rise-worlds/potato-go"

func init() {
	potato.RegisterAction(REXAN, ActN("buyrex"), BuyREX{})
	potato.RegisterAction(REXAN, ActN("closerex"), CloseREX{})
	potato.RegisterAction(REXAN, ActN("cnclrexorder"), CancelREXOrder{})
	potato.RegisterAction(REXAN, ActN("consolidate"), Consolidate{})
	potato.RegisterAction(REXAN, ActN("defcpuloan"), DefundCPULoan{})
	potato.RegisterAction(REXAN, ActN("defnetloan"), DefundNetLoan{})
	potato.RegisterAction(REXAN, ActN("deposit"), Deposit{})
	potato.RegisterAction(REXAN, ActN("fundcpuloan"), FundCPULoan{})
	potato.RegisterAction(REXAN, ActN("fundnetloan"), FundNetLoan{})
	potato.RegisterAction(REXAN, ActN("mvfrsavings"), MoveFromSavings{})
	potato.RegisterAction(REXAN, ActN("mvtosavings"), MoveToSavings{})
	potato.RegisterAction(REXAN, ActN("rentcpu"), RentCPU{})
	potato.RegisterAction(REXAN, ActN("rentnet"), RentNet{})
	potato.RegisterAction(REXAN, ActN("rexexec"), REXExec{})
	potato.RegisterAction(REXAN, ActN("sellrex"), SellREX{})
	potato.RegisterAction(REXAN, ActN("unstaketorex"), UnstakeToREX{})
	potato.RegisterAction(REXAN, ActN("updaterex"), UpdateREX{})
	potato.RegisterAction(REXAN, ActN("withdraw"), Withdraw{})
}

var AN = potato.AN
var PN = potato.PN
var ActN = potato.ActN

var REXAN = AN("potato")

package system

import (
	"github.com/rise-worlds/potato-go"
)

func init() {
	potato.RegisterAction(AN("potato"), ActN("setcode"), SetCode{})
	potato.RegisterAction(AN("potato"), ActN("setabi"), SetABI{})
	potato.RegisterAction(AN("potato"), ActN("newaccount"), NewAccount{})
	potato.RegisterAction(AN("potato"), ActN("delegatebw"), DelegateBW{})
	potato.RegisterAction(AN("potato"), ActN("undelegatebw"), UnDelegateBW{})
	potato.RegisterAction(AN("potato"), ActN("refund"), Refund{})
	potato.RegisterAction(AN("potato"), ActN("regproducer"), RegProducer{})
	potato.RegisterAction(AN("potato"), ActN("unregproducer"), UnregProducer{})
	potato.RegisterAction(AN("potato"), ActN("regproxy"), RegProxy{})
	potato.RegisterAction(AN("potato"), ActN("voteproducer"), VoteProducer{})
	potato.RegisterAction(AN("potato"), ActN("claimwards"), ClaimRewards{})
	potato.RegisterAction(AN("potato"), ActN("buyram"), BuyRAM{})
	potato.RegisterAction(AN("potato"), ActN("buyrambytes"), BuyRAMBytes{})
	potato.RegisterAction(AN("potato"), ActN("linkauth"), LinkAuth{})
	potato.RegisterAction(AN("potato"), ActN("unlinkauth"), UnlinkAuth{})
	potato.RegisterAction(AN("potato"), ActN("deleteauth"), DeleteAuth{})
	potato.RegisterAction(AN("potato"), ActN("rmvproducer"), RemoveProducer{})
	potato.RegisterAction(AN("potato"), ActN("setprods"), SetProds{})
	potato.RegisterAction(AN("potato"), ActN("setpriv"), SetPriv{})
	potato.RegisterAction(AN("potato"), ActN("canceldelay"), CancelDelay{})
	potato.RegisterAction(AN("potato"), ActN("bidname"), Bidname{})
	// potato.RegisterAction(AN("potato"), ActN("nonce"), &Nonce{})
	potato.RegisterAction(AN("potato"), ActN("sellram"), SellRAM{})
	potato.RegisterAction(AN("potato"), ActN("updateauth"), UpdateAuth{})
	potato.RegisterAction(AN("potato"), ActN("setramrate"), SetRAMRate{})
	potato.RegisterAction(AN("potato"), ActN("setalimits"), Setalimits{})
}

var AN = potato.AN
var PN = potato.PN
var ActN = potato.ActN

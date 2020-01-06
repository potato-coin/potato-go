package forum

import potato "github.com/rise-worlds/potato-go"

func init() {
	potato.RegisterAction(ForumAN, ActN("clnproposal"), CleanProposal{})
	potato.RegisterAction(ForumAN, ActN("expire"), Expire{})
	potato.RegisterAction(ForumAN, ActN("post"), Post{})
	potato.RegisterAction(ForumAN, ActN("propose"), Propose{})
	potato.RegisterAction(ForumAN, ActN("status"), Status{})
	potato.RegisterAction(ForumAN, ActN("unpost"), UnPost{})
	potato.RegisterAction(ForumAN, ActN("unvote"), UnVote{})
	potato.RegisterAction(ForumAN, ActN("vote"), Vote{})
}

var AN = potato.AN
var PN = potato.PN
var ActN = potato.ActN

var ForumAN = AN("potato.forum")

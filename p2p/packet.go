package p2p

import (
	"github.com/rise-worlds/potato-go"
)

type Envelope struct {
	Sender   *Peer
	Receiver *Peer
	Packet   *potato.Packet `json:"envelope"`
}

func NewEnvelope(sender *Peer, receiver *Peer, packet *potato.Packet) *Envelope {
	return &Envelope{
		Sender:   sender,
		Receiver: receiver,
		Packet:   packet,
	}
}

package dht

import (
	"time"

	"github.com/anacrolix/dht/krpc"
)

type node struct {
	addr          Addr
	id            int160
	announceToken string

	lastGotQuery    time.Time
	lastGotResponse time.Time
	lastSentQuery   time.Time

	consecutiveFailures int
}

func (n *node) IsSecure() bool {
	return NodeIdSecure(n.id.AsByteArray(), n.addr.UDPAddr().IP)
}

func (n *node) idString() string {
	return n.id.ByteString()
}

func (n *node) NodeInfo() (ret krpc.NodeInfo) {
	ret.Addr = n.addr.UDPAddr()
	if n := copy(ret.ID[:], n.idString()); n != 20 {
		panic(n)
	}
	return
}

// Per the spec in BEP 5.
func (n *node) IsGood() bool {
	if n.id.IsZero() {
		return false
	}
	if time.Since(n.lastGotResponse) < 15*time.Minute {
		return true
	}
	if !n.lastGotResponse.IsZero() && time.Since(n.lastGotQuery) < 15*time.Minute {
		return true
	}
	return false
}
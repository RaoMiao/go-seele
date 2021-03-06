/**
*  @file
*  @copyright defined in go-seele/LICENSE
 */

package seele

import (
	"sync"

	"github.com/seeleteam/go-seele/common"
)

var (
	maxPeerConnected = 1024
)

type peerSet struct {
	peers map[common.Address]*peer
	lock  sync.RWMutex
}

func newPeerSet() *peerSet {
	return &peerSet{
		peers: make(map[common.Address]*peer),
		lock:  sync.RWMutex{},
	}
}

func (p *peerSet) bestPeer() *peer {
	// TODO get best peer
	return nil
}

func (p *peerSet) Find(address common.Address) *peer {
	p.lock.RLock()
	defer p.lock.RUnlock()

	return p.peers[address]
}

func (p *peerSet) Remove(address common.Address) {
	p.lock.Lock()
	defer p.lock.Unlock()

	delete(p.peers, address)
}

func (p *peerSet) Add(pe *peer) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if len(p.peers) == maxPeerConnected {
		var k common.Address
		for k = range p.peers {
			break
		}

		delete(p.peers, k)
	}

	p.peers[pe.Node.ID] = pe
}

func (p *peerSet) ForEach(handle func(*peer) bool) {
	p.lock.RLock()
	defer p.lock.RUnlock()

	for _, v := range p.peers {
		if !handle(v) {
			break
		}
	}
}

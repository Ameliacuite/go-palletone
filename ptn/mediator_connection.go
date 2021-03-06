/*
   This file is part of go-palletone.
   go-palletone is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   go-palletone is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.
   You should have received a copy of the GNU General Public License
   along with go-palletone.  If not, see <http://www.gnu.org/licenses/>.
*/
/*
 * @author PalletOne core developer Albert·Gou <dev@pallet.one>
 * @date 2018
 */

package ptn

import (
	"time"

	"github.com/palletone/go-palletone/common/event"
	"github.com/palletone/go-palletone/common/p2p/discover"
	mp "github.com/palletone/go-palletone/consensus/mediatorplugin"
	"github.com/palletone/go-palletone/dag/modules"
)

// @author Albert·Gou
type producer interface {
	// SubscribeNewUnitEvent should return an event subscription of
	// NewUnitEvent and send events to the given channel.
	SubscribeNewProducedUnitEvent(ch chan<- mp.NewProducedUnitEvent) event.Subscription
	// UnitBLSSign is to TBLS sign the unit
	ToUnitTBLSSign(newUnit *modules.Unit) error

	SubscribeSigShareEvent(ch chan<- mp.SigShareEvent) event.Subscription
	ToTBLSRecover(sigShare *mp.SigShareEvent) error

	SubscribeVSSDealEvent(ch chan<- mp.VSSDealEvent) event.Subscription
	ToProcessDeal(deal *mp.VSSDealEvent) error

	SubscribeVSSResponseEvent(ch chan<- mp.VSSResponseEvent) event.Subscription
	ToProcessResponse(resp *mp.VSSResponseEvent) error

	LocalHaveActiveMediator() bool
	LocalHavePrecedingMediator() bool

	SubscribeGroupSigEvent(ch chan<- mp.GroupSigEvent) event.Subscription
	UpdateMediatorsDKG()
}

func (pm *ProtocolManager) activeMediatorsUpdatedEventRecvLoop() {
	return
	for {
		select {
		case <-pm.activeMediatorsUpdatedCh:
			go pm.switchMediatorConnect()

			// Err() channel will be closed when unsubscribing.
		case <-pm.activeMediatorsUpdatedSub.Err():
			return
		}
	}
}

func (pm *ProtocolManager) switchMediatorConnect() {
	// 1. 若干数据还没同步完成，则忽略本次切换，继续同步
	if !pm.dag.IsSynced() {
		return
	}

	// 2. 和新的活跃mediator节点相连
	go pm.connectWitchActiveMediators()

	// 3. 检查是否连接完成，并发送事件
	//go pm.checkActiveMediatorConnection()

	// 4. 延迟关闭和旧活跃mediator节点的连接
	go pm.delayDiscPrecedingMediator()
}

func (pm *ProtocolManager) connectWitchActiveMediators() {
	// 1. 判断本节点是否是活跃mediator
	if !pm.producer.LocalHaveActiveMediator() {
		return
	}

	// 2. 和其他活跃mediator节点相连
	peers := pm.dag.GetActiveMediatorNodes()
	for id, peer := range peers {
		// 仅当不是本节点，并还未连接时，才进行连接
		if peer.ID != pm.srvr.Self().ID && pm.peers.Peer(id) == nil {
			pm.srvr.AddTrustedPeer(peer)
		}
	}
}

func (pm *ProtocolManager) checkActiveMediatorConnection() {
	// 2. 是否和所有其他活跃mediator节点相连完成
	checkFn := func() bool {
		peers := pm.dag.GetActiveMediatorNodes()
		for id, peer := range peers {
			// 仅当不是本节点，并还未连接完成时，返回false
			if peer.ID != pm.srvr.Self().ID && pm.peers.Peer(id) == nil {
				return false
			}
		}

		return true
	}

	// 3. 发送mediator连接完成的事件
	processFn := func() {
		go pm.producer.UpdateMediatorsDKG()
	}

	// 1. 设置Ticker, 每隔一段时间检查一次
	checkTick := time.NewTicker(50 * time.Millisecond)
	for {
		select {
		case <-pm.quitSync:
			return
		case <-checkTick.C:
			if checkFn() {
				checkTick.Stop()
				processFn()
				return
			}
		}
	}
}

func (pm *ProtocolManager) delayDiscPrecedingMediator() {
	// 1. 判断当前节点是否是上一届活跃mediator
	if !pm.producer.LocalHavePrecedingMediator() {
		return
	}

	// 2. 统计出需要断开连接的mediator节点
	delayDiscNodes := make(map[string]*discover.Node, 0)

	activePeers := pm.dag.GetActiveMediatorNodes()
	precedingPeers := pm.dag.GetPrecedingMediatorNodes()
	for id, peer := range precedingPeers {
		// 仅当上一届mediator 不是本届活跃mediator，并且已连接时，才断开连接
		if _, ok := activePeers[id]; !ok && pm.peers.Peer(id) != nil {
			delayDiscNodes[id] = peer
		}
	}

	// 3. 设置定时器延迟 断开连接
	discconnectFn := func() {
		for _, peer := range delayDiscNodes {
			pm.srvr.RemoveTrustedPeer(peer)
		}
	}

	expiration := pm.dag.UnitIrreversibleTime()
	delayDisc := time.NewTimer(time.Duration(expiration))

	select {
	case <-pm.quitSync:
		return
	case <-delayDisc.C:
		discconnectFn()
	}
}

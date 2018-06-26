// Copyright 2015 The go-palletone Authors
// This file is part of the go-palletone library.
//
// The go-palletone library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-palletone library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-palletone library. If not, see <http://www.gnu.org/licenses/>.

package ptn

import (
	"context"
	"math/big"

	"github.com/palletone/go-palletone/common"
	"github.com/palletone/go-palletone/core/accounts"
	//"github.com/palletone/go-palletone/common/math"
	"github.com/palletone/go-palletone/common/bloombits"
	"github.com/palletone/go-palletone/common/event"
	"github.com/palletone/go-palletone/common/ptndb"
	"github.com/palletone/go-palletone/common/rpc"
	"github.com/palletone/go-palletone/configure"
	"github.com/palletone/go-palletone/contracts/types"
	"github.com/palletone/go-palletone/dag/coredata"
	"github.com/palletone/go-palletone/dag/state"
	"github.com/palletone/go-palletone/ptn/downloader"
	"github.com/palletone/go-palletone/ptn/gasprice"
	"github.com/palletone/go-palletone/vm"
)

// EthApiBackend implements ethapi.Backend for full nodes
type EthApiBackend struct {
	eth *Ethereum
	gpo *gasprice.Oracle
}

func (b *EthApiBackend) ChainConfig() *configure.ChainConfig {
	return nil
}

func (b *EthApiBackend) CurrentBlock() *types.Block {
	//return b.eth.blockchain.CurrentBlock()
	return &types.Block{}
}

func (b *EthApiBackend) SetHead(number uint64) {
	b.eth.protocolManager.downloader.Cancel()
	//b.eth.blockchain.SetHead(number)//wangjiyou
}

func (b *EthApiBackend) HeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Header, error) {
	// Pending block is only known by the miner
	/*
		if blockNr == rpc.PendingBlockNumber {
			block := b.eth.miner.PendingBlock()
			return block.Header(), nil
		}
		// Otherwise resolve and return the block
		if blockNr == rpc.LatestBlockNumber {
			return b.eth.blockchain.CurrentBlock().Header(), nil
		}
		return b.eth.blockchain.GetHeaderByNumber(uint64(blockNr)), nil
	*/
	return &types.Header{}, nil
}

func (b *EthApiBackend) BlockByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*types.Block, error) {
	/*
		// Pending block is only known by the miner
		if blockNr == rpc.PendingBlockNumber {
			block := b.eth.miner.PendingBlock()
			return block, nil
		}
		// Otherwise resolve and return the block
		if blockNr == rpc.LatestBlockNumber {
			return b.eth.blockchain.CurrentBlock(), nil
		}
		return b.eth.blockchain.GetBlockByNumber(uint64(blockNr)), nil
	*/
	return &types.Block{}, nil
}

func (b *EthApiBackend) StateAndHeaderByNumber(ctx context.Context, blockNr rpc.BlockNumber) (*state.StateDB, *types.Header, error) {
	/*
		// Pending state is only known by the miner
		if blockNr == rpc.PendingBlockNumber {
			block, state := b.eth.miner.Pending()
			return state, block.Header(), nil
		}
		// Otherwise resolve the block number and return its state
		header, err := b.HeaderByNumber(ctx, blockNr)
		if header == nil || err != nil {
			return nil, nil, err
		}
		stateDb, err := b.eth.BlockChain().StateAt(header.Root)
		return stateDb, header, err
	*/
	return &state.StateDB{}, &types.Header{}, nil
}

func (b *EthApiBackend) GetBlock(ctx context.Context, blockHash common.Hash) (*types.Block, error) {
	//return b.eth.blockchain.GetBlockByHash(blockHash), nil
	return &types.Block{}, nil
}

func (b *EthApiBackend) GetReceipts(ctx context.Context, blockHash common.Hash) (types.Receipts, error) {
	return coredata.GetBlockReceipts(b.eth.chainDb, blockHash, coredata.GetBlockNumber(b.eth.chainDb, blockHash)), nil
}

func (b *EthApiBackend) GetLogs(ctx context.Context, blockHash common.Hash) ([][]*types.Log, error) {
	receipts := coredata.GetBlockReceipts(b.eth.chainDb, blockHash, coredata.GetBlockNumber(b.eth.chainDb, blockHash))
	if receipts == nil {
		return nil, nil
	}
	logs := make([][]*types.Log, len(receipts))
	for i, receipt := range receipts {
		logs[i] = receipt.Logs
	}
	return logs, nil
}

func (b *EthApiBackend) GetTd(blockHash common.Hash) *big.Int {
	//return b.eth.blockchain.GetTdByHash(blockHash)
	return &big.Int{}
}

func (b *EthApiBackend) GetEVM(ctx context.Context, msg coredata.Message, state *state.StateDB, header *types.Header, vmCfg vm.Config) (*vm.EVM, func() error, error) {
	/*
		state.SetBalance(msg.From(), math.MaxBig256)
		vmError := func() error { return nil }

		//context := core.NewEVMContext(msg, header, b.eth.BlockChain(), nil)//wangjiyou
		context := vm.Context{}
		return vm.NewEVM(context, state, b.eth.chainConfig, vmCfg), vmError, nil
	*/
	return &vm.EVM{}, func() error { return nil }, nil
}

func (b *EthApiBackend) SubscribeRemovedLogsEvent(ch chan<- coredata.RemovedLogsEvent) event.Subscription {
	//return b.eth.BlockChain().SubscribeRemovedLogsEvent(ch)
	return nil
}

func (b *EthApiBackend) SubscribeChainEvent(ch chan<- coredata.ChainEvent) event.Subscription {
	//return b.eth.BlockChain().SubscribeChainEvent(ch)
	return nil
}

func (b *EthApiBackend) SubscribeChainHeadEvent(ch chan<- coredata.ChainHeadEvent) event.Subscription {
	//return b.eth.BlockChain().SubscribeChainHeadEvent(ch)
	return nil
}

func (b *EthApiBackend) SubscribeChainSideEvent(ch chan<- coredata.ChainSideEvent) event.Subscription {
	//return b.eth.BlockChain().SubscribeChainSideEvent(ch)
	return nil
}

func (b *EthApiBackend) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription {
	//return b.eth.BlockChain().SubscribeLogsEvent(ch)
	return nil
}

func (b *EthApiBackend) SendConsensus(ctx context.Context) error {
	b.eth.Engine().Engine()
	return nil
}

func (b *EthApiBackend) SendTx(ctx context.Context, signedTx *types.Transaction) error {
	return b.eth.txPool.AddLocal(signedTx)
}

func (b *EthApiBackend) GetPoolTransactions() (types.Transactions, error) {
	pending, err := b.eth.txPool.Pending()
	if err != nil {
		return nil, err
	}
	var txs types.Transactions
	for _, batch := range pending {
		txs = append(txs, batch...)
	}
	return txs, nil
}

func (b *EthApiBackend) GetPoolTransaction(hash common.Hash) *types.Transaction {
	return b.eth.txPool.Get(hash)
}

func (b *EthApiBackend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	return b.eth.txPool.State().GetNonce(addr), nil
}

func (b *EthApiBackend) Stats() (pending int, queued int) {
	return b.eth.txPool.Stats()
}

func (b *EthApiBackend) TxPoolContent() (map[common.Address]types.Transactions, map[common.Address]types.Transactions) {
	return b.eth.TxPool().Content()
}

func (b *EthApiBackend) SubscribeTxPreEvent(ch chan<- coredata.TxPreEvent) event.Subscription {
	return b.eth.TxPool().SubscribeTxPreEvent(ch)
}

func (b *EthApiBackend) Downloader() *downloader.Downloader {
	return b.eth.Downloader()
}

func (b *EthApiBackend) ProtocolVersion() int {
	return b.eth.EthVersion()
}

func (b *EthApiBackend) SuggestPrice(ctx context.Context) (*big.Int, error) {
	return b.gpo.SuggestPrice(ctx)
}

func (b *EthApiBackend) ChainDb() ptndb.Database {
	return b.eth.ChainDb()
}

func (b *EthApiBackend) EventMux() *event.TypeMux {
	return b.eth.EventMux()
}

func (b *EthApiBackend) AccountManager() *accounts.Manager {
	return b.eth.AccountManager()
}

func (b *EthApiBackend) BloomStatus() (uint64, uint64) {
	sections, _, _ := b.eth.bloomIndexer.Sections()
	return configure.BloomBitsBlocks, sections
}

func (b *EthApiBackend) ServiceFilter(ctx context.Context, session *bloombits.MatcherSession) {
	for i := 0; i < bloomFilterThreads; i++ {
		go session.Multiplex(bloomRetrievalBatch, bloomRetrievalWait, b.eth.bloomRequests)
	}
}
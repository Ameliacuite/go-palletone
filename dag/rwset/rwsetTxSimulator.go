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
 * @author PalletOne core developers <dev@pallet.one>
 * @date 2018
 */

package rwset

import (
	"errors"
	"github.com/palletone/go-palletone/dag"
)

type RwSetTxSimulator struct {
	txid                    string
	rwsetBuilder            *RWSetBuilder
	state                   dag.IDag
	writePerformed          bool
	pvtdataQueriesPerformed bool
	doneInvoked             bool
}

type VersionedValue struct {
	Value   []byte
	Version *Version
}

func NewBasedTxSimulator(idag dag.IDag, txid string) *RwSetTxSimulator {
	rwsetBuilder := NewRWSetBuilder()

	logger.Debugf("constructing new tx simulator txid = [%s]", txid)
	return &RwSetTxSimulator{txid, rwsetBuilder, idag, false, false, false}
}

// GetState implements method in interface `ledger.TxSimulator`
func (s *RwSetTxSimulator) GetState(cid[]byte, ns string, key string) ([]byte, error) {
	//testValue := []byte("abc")
	if err := s.CheckDone(); err != nil {
		return nil, err
	}
	//TODO Devin
	//
	ver, val := s.state.GetContractState(ns, key)
	//fmt.Println("xiaozhi:--", ver, val)
	//TODO 这里证明数据库里面没有该账户信息，需要返回nil,nil
	if val == nil {
		logger.Errorf("get value from db[%s] failed", ns)
		return nil, nil
		//errstr := fmt.Sprintf("GetContractState [%s]-[%s] failed", ns, key)
		//		//return nil, errors.New(errstr)
	}
	//val, ver := decomposeVersionedValue(versionedValue)
	if s.rwsetBuilder != nil {
		s.rwsetBuilder.AddToReadSet(ns, key, ver)
	}

	logger.Debugf("RW:GetState,ns[%s]--key[%s]---value[%s]", ns, key, val)

	//todo change.
	//return testValue, nil
	return val, nil
}

func (s *RwSetTxSimulator) SetState(ns string, key string, value []byte) error {
	logger.Debugf("RW:SetState,ns[%s]--key[%s]---value[%s]", ns, key, value)

	if err := s.CheckDone(); err != nil {
		return err
	}
	if s.pvtdataQueriesPerformed {
		return errors.New("pvtdata Queries Performed")
	}
	//todo ValidateKeyValue
	s.rwsetBuilder.AddToWriteSet(ns, key, value)
	return nil
}

// DeleteState implements method in interface `ledger.TxSimulator`
func (s *RwSetTxSimulator) DeleteState(ns string, key string) error {
	return s.SetState(ns, key, nil)
}

func (s *RwSetTxSimulator) GetRwData(ns string) (map[string]*KVRead, map[string]*KVWrite, error) {
	var rd map[string]*KVRead
	var wt map[string]*KVWrite

	logger.Infof("ns=%s", ns)

	if s.rwsetBuilder != nil {
		if s.rwsetBuilder.pubRwBuilderMap != nil {
			if s.rwsetBuilder.pubRwBuilderMap[ns] != nil {
				if s.rwsetBuilder.pubRwBuilderMap[ns].readMap != nil {
					rd = s.rwsetBuilder.pubRwBuilderMap[ns].readMap
				}
				if s.rwsetBuilder.pubRwBuilderMap[ns].writeMap != nil {
					wt = s.rwsetBuilder.pubRwBuilderMap[ns].writeMap
				}
				pubRwBuilderMap, ok := s.rwsetBuilder.pubRwBuilderMap[ns]
				if ok {
					rd = pubRwBuilderMap.readMap
					wt = pubRwBuilderMap.writeMap
				} else {
					rd = nil
					wt = nil
				}
			}
		}
	}

	return rd, wt, nil
}

func (h *RwSetTxSimulator) CheckDone() error {
	if h.doneInvoked {
		return errors.New("This instance should not be used after calling Done()")
	}
	return nil
}

func decomposeVersionedValue(versionedValue *VersionedValue) ([]byte, *Version) {
	var value []byte
	var ver *Version
	if versionedValue != nil {
		value = versionedValue.Value
		ver = versionedValue.Version
	}
	return value, ver
}

func (h *RwSetTxSimulator) Done() {
	if h.doneInvoked {
		return
	}
	//todo
}
func (h *RwSetTxSimulator) GetTxSimulationResults() ([]byte, error) {

	return nil, nil
}

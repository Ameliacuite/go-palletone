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

package manger

import (
	"github.com/golang/protobuf/proto"
	"github.com/palletone/go-palletone/contracts/scc"
	"github.com/palletone/go-palletone/core/vmContractPub/ccprovider"
	"github.com/palletone/go-palletone/dag"
	"github.com/palletone/go-palletone/dag/rwset"
	"golang.org/x/net/context"
	"time"

	chaincode "github.com/palletone/go-palletone/contracts/core"
	"github.com/palletone/go-palletone/contracts/modules"
	pb "github.com/palletone/go-palletone/core/vmContractPub/protos/peer"
	unit "github.com/palletone/go-palletone/dag/modules"
	ut "github.com/palletone/go-palletone/dag/modules"
)

// SupportImpl provides an implementation of the endorser.Support interface
// issuing calls to various static methods of the peer
type SupportImpl struct {
}

// IsSysCCAndNotInvokableExternal returns true if the supplied chaincode is
// ia system chaincode and it NOT invokable
func (s *SupportImpl) IsSysCCAndNotInvokableExternal(name string) bool {
	return scc.IsSysCCAndNotInvokableExternal(name)
}

// GetTxSimulator returns the transaction simulator for the specified ledger
// a client may obtain more than one such simulator; they are made unique
// by way of the supplied txid
func (s *SupportImpl) GetTxSimulator(idag dag.IDag, chainid string, txid string) (rwset.TxSimulator, error) {
	return rwM.NewTxSimulator(idag, chainid, txid)
}

//IsSysCC returns true if the name matches a system chaincode's
//system chaincode names are system, chain wide
func (s *SupportImpl) IsSysCC(name string) bool {
	return scc.IsSysCC(name)
}

//Execute - execute proposal, return original response of chaincode
func (s *SupportImpl) Execute(contractid []byte, ctxt context.Context, cid, name, version, txid string, syscc bool, signedProp *pb.SignedProposal, prop *pb.Proposal, spec interface{}, timeout time.Duration) (*pb.Response, *pb.ChaincodeEvent, error) {
	cccid := ccprovider.NewCCContext(contractid, cid, name, version, txid, syscc, signedProp, prop)

	switch spec.(type) {
	case *pb.ChaincodeDeploymentSpec:
		return chaincode.Execute(ctxt, cccid, spec, timeout)
	case *pb.ChaincodeInvocationSpec:
		cis := spec.(*pb.ChaincodeInvocationSpec)

		//logger.Infof("cis:%v", cis)
		//decorate the chaincode input

		//decorators := library.InitRegistry(library.Config{}).Lookup(library.Decoration).([]decoration.Decorator)
		//cis.ChaincodeSpec.Input.Decorations = make(map[string][]byte)
		//cis.ChaincodeSpec.Input = decoration.Apply(prop, cis.ChaincodeSpec.Input, decorators...)
		//cccid.ProposalDecorations = cis.ChaincodeSpec.Input.Decorations
		return chaincode.ExecuteChaincode(ctxt, cccid, cis.ChaincodeSpec.Input.Args, timeout)
	default:
		panic("programming error, unkwnown spec type")
	}
}

// shorttxid replicates the chaincode package function to shorten txids.
func shorttxid(txid string) string {
	if len(txid) < 8 {
		return txid
	}
	return txid[0:8]
}

func GetBytesChaincodeEvent(event *pb.ChaincodeEvent) ([]byte, error) {
	eventBytes, err := proto.Marshal(event)
	return eventBytes, err
}

func RwTxResult2DagInvokeUnit(tx rwset.TxSimulator, txid string, nm string, deployId []byte, args [][]byte, timeout time.Duration) (*modules.ContractInvokeResult, error) {
	logger.Debug("enter")
	//invokeData := ut.ContractInvokePayload{}
	//invokeData.ContractId = []byte(txid)

	rd, wt, err := tx.GetRwData(nm)
	if err != nil {
		return nil, err
	}
	tokenPay, err := tx.GetPayOutData(nm)
	if err != nil {
		return nil, err
	}
	tokenDefine, _ := tx.GetTokenDefineData(nm)
	tokenSupply, _ := tx.GetTokenSupplyData(nm)
	logger.Infof("txid=%s, nm=%s, rd=%v, wt=%v", txid, nm, rd, wt)
	invoke := &modules.ContractInvokeResult{
		FunctionName:  string(args[1]),
		ContractId:    deployId,
		Args:          args,
		ExecutionTime: timeout,
		ReadSet:       make([]unit.ContractReadSet, 0),
		WriteSet:      make([]unit.ContractWriteSet, 0),
		TokenPayOut:   tokenPay,
		TokenDefine:   tokenDefine,
		TokenSupply:   tokenSupply,
	}

	for idx, val := range rd {
		rd := unit.ContractReadSet{
			Key:     val.GetKey(),
			Version: val.GetVersion(),
		}
		invoke.ReadSet = append(invoke.ReadSet, rd)
		logger.Infof("ReadSet: idx[%s], fun[%s], key[%s], val[%v]", idx, args[1], val.GetKey(), *val.GetVersion())
	}
	for idx, val := range wt {
		rd := unit.ContractWriteSet{
			Key:      val.GetKey(),
			Value:    val.GetValue(),
			IsDelete: val.GetIsDelete(),
		}
		invoke.WriteSet = append(invoke.WriteSet, rd)
		logger.Infof("WriteSet: idx[%s], fun[%s], key[%s], val[%v], delete[%v]", idx, args[1], val.GetKey(), val.GetValue(), val.GetIsDelete())
	}

	return invoke, nil
}

//func RwTxResult2DagDeployUnit(tx rwset.TxSimulator, txid string, nm string, fun []byte) (*pb.ContractDeployPayload, error) {
func RwTxResult2DagDeployUnit(tx rwset.TxSimulator, templateId []byte, txid string, nm string, deployId []byte, args [][]byte, timeout time.Duration) (*unit.ContractDeployPayload, error) {
	logger.Debug("enter")
	data := ut.ContractDeployPayload{}
	data.ContractId = []byte(txid)

	rd, wt, err := tx.GetRwData(nm)
	if err != nil {
		return nil, err
	}
	logger.Infof("txid=%s, nm=%s, rd=%v, wt=%v", txid, nm, rd, wt)
	deploy := &unit.ContractDeployPayload{
		TemplateId:    templateId,
		ContractId:    deployId,
		Name:          nm,
		Args:          args,
		ExecutionTime: timeout,
		ReadSet:       make([]unit.ContractReadSet, 0),
		WriteSet:      make([]unit.ContractWriteSet, 0),
	}

	for idx, val := range rd {
		rd := unit.ContractReadSet{
			Key:     val.GetKey(),
			Version: val.GetVersion(),
		}
		deploy.ReadSet = append(deploy.ReadSet, rd)
		logger.Infof("ReadSet: idx[%s], fun[%s], key[%s], val[%v]", idx, args[1], val.GetKey(), *val.GetVersion())
	}
	for idx, val := range wt {
		rd := unit.ContractWriteSet{
			Key:      val.GetKey(),
			Value:    val.GetValue(),
			IsDelete: val.GetIsDelete(),
		}
		deploy.WriteSet = append(deploy.WriteSet, rd)
		logger.Infof("WriteSet: idx[%s], fun[%s], key[%s], val[%v], delete[%v]", idx, args[1], val.GetKey(), val.GetValue(), val.GetIsDelete())
	}

	return deploy, nil
}

var rwM *rwset.RwSetTxMgr

func init() {
	var err error
	rwM, err = rwset.NewRwSetMgr("default")
	if err != nil {
		logger.Error("fail!")
	}
}

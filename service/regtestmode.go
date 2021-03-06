package service

import (
	"hyperledger.abchain.org/applications/util/blockchain"
	evt "hyperledger.abchain.org/chaincode/lib/events"
	txhandle "hyperledger.abchain.org/chaincode/lib/txhandle"
	"hyperledger.abchain.org/client/local"
	"hyperledger.abchain.org/example/ae/chaincode/cc"
)

func init() {
	aecc := chaincode.NewChaincode(true)
	client.AddChaincode("local", aecc)

	//also build txparser from chaincode ...
	parser := txhandle.GenerateTxArgParser(aecc.CollectiveTxs.Map())
	evt.MustMergeTxEventParsers(parser)
	parser[chaincode.CC_BATCH+"@"+chaincode.CC_NAME] = txhandle.BatchArgParser(chaincode.CC_NAME, parser)
	parser[client.TxErrorEventName] = client.TxErrorParser("<Parser> ")
	blockchain.SetParsers(parser)
}

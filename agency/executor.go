package agency

import (
	"github.com/theneverse/neverse-kit/types"
	"github.com/theneverse/neverse-model/pb"
)

type TxsExecutor interface {
	ApplyTransactions(txs []pb.Transaction, invalidTxs map[int]InvalidReason) []*pb.Receipt

	GetBoltContracts() map[string]Contract

	AddNormalTx(hash *types.Hash)

	GetNormalTxs() []*types.Hash

	AddInterchainCounter(to string, index *pb.VerifiedIndex)

	GetInterchainCounter() map[string][]*pb.VerifiedIndex

	GetDescription() string
}

type TxOpt struct {
	Contracts map[string]Contract
}

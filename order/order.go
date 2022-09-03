package order

import (
	"github.com/ethereum/go-ethereum/event"
	"github.com/theneverse/neverse-kit/types"
	"github.com/theneverse/neverse-model/pb"
)

//go:generate mockgen -destination mock_order/mock_order.go -package mock_order -source order.go
type Order interface {
	// Start the order service.
	Start() error

	// Stop means frees the resources which were allocated for this service.
	Stop()

	// Prepare means send transaction to the consensus engine
	Prepare(tx pb.Transaction) error

	// Commit recv blocks form Order and commit it by order
	Commit() chan *pb.CommitEvent

	// Step send msg to the consensus engine
	Step(msg []byte) error

	// Ready means whether order has finished electing leader
	Ready() error

	// ReportState means block was persisted and report it to the consensus engine
	ReportState(height uint64, blockHash *types.Hash, txHashList []*types.Hash)

	// Quorum means minimum number of nodes in the cluster that can work
	Quorum() uint64

	// GetPendingNonceByAccount will return the latest pending nonce of a given account
	GetPendingNonceByAccount(account string) uint64

	GetPendingTxByHash(hash *types.Hash) pb.Transaction

	// DelNode sends a delete vp request by given id.
	DelNode(delID uint64) error

	SubscribeTxEvent(events chan<- pb.Transactions) event.Subscription
}

package hedera

import (
	"time"

	"github.com/hashgraph/hedera-sdk-go/proto"
)

// A ConsensusTopicDeleteTransaction is for deleting a topic on HCS.
type ConsensusTopicDeleteTransaction struct {
	TransactionBuilder
	pb *proto.ConsensusDeleteTopicTransactionBody
}

// NewConsensusTopicDeleteTransaction creates a ConsensusTopicDeleteTransaction builder which can be used to construct
// and execute a Consensus Delete Topic Transaction.
func NewConsensusTopicDeleteTransaction() ConsensusTopicDeleteTransaction {
	pb := &proto.ConsensusDeleteTopicTransactionBody{}

	inner := newTransactionBuilder()
	inner.pb.Data = &proto.TransactionBody_ConsensusDeleteTopic{ConsensusDeleteTopic: pb}

	builder := ConsensusTopicDeleteTransaction{inner, pb}

	return builder
}

// SetTopicID sets the topic identifier.
func (builder ConsensusTopicDeleteTransaction) SetTopicID(id ConsensusTopicID) ConsensusTopicDeleteTransaction {
	builder.pb.TopicID = id.toProto()
	return builder
}

//
// The following _5_ must be copy-pasted at the bottom of **every** _transaction.go file
// We override the embedded fluent setter methods to return the outer type
//

// SetMaxTransactionFee sets the max transaction fee for this Transaction.
func (builder ConsensusTopicDeleteTransaction) SetMaxTransactionFee(maxTransactionFee Hbar) ConsensusTopicDeleteTransaction {
	return ConsensusTopicDeleteTransaction{builder.TransactionBuilder.SetMaxTransactionFee(maxTransactionFee), builder.pb}
}

// SetTransactionMemo sets the memo for this Transaction.
func (builder ConsensusTopicDeleteTransaction) SetTransactionMemo(memo string) ConsensusTopicDeleteTransaction {
	return ConsensusTopicDeleteTransaction{builder.TransactionBuilder.SetTransactionMemo(memo), builder.pb}
}

// SetTransactionValidDuration sets the valid duration for this Transaction.
func (builder ConsensusTopicDeleteTransaction) SetTransactionValidDuration(validDuration time.Duration) ConsensusTopicDeleteTransaction {
	return ConsensusTopicDeleteTransaction{builder.TransactionBuilder.SetTransactionValidDuration(validDuration), builder.pb}
}

// SetTransactionID sets the TransactionID for this Transaction.
func (builder ConsensusTopicDeleteTransaction) SetTransactionID(transactionID TransactionID) ConsensusTopicDeleteTransaction {
	return ConsensusTopicDeleteTransaction{builder.TransactionBuilder.SetTransactionID(transactionID), builder.pb}
}

// SetNodeAccountID sets the node AccountID for this Transaction.
func (builder ConsensusTopicDeleteTransaction) SetNodeAccountID(nodeAccountID AccountID) ConsensusTopicDeleteTransaction {
	return ConsensusTopicDeleteTransaction{builder.TransactionBuilder.SetNodeAccountID(nodeAccountID), builder.pb}
}

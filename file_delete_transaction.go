package hedera

import (
	"time"

	"github.com/hashgraph/hedera-sdk-go/proto"
)

type FileDeleteTransaction struct {
	TransactionBuilder
	pb *proto.FileDeleteTransactionBody
}

func NewFileDeleteTransaction() FileDeleteTransaction {
	pb := &proto.FileDeleteTransactionBody{}

	inner := newTransactionBuilder()
	inner.pb.Data = &proto.TransactionBody_FileDelete{FileDelete: pb}

	builder := FileDeleteTransaction{inner, pb}

	return builder
}

func (builder FileDeleteTransaction) SetFileID(id FileID) FileDeleteTransaction {
	builder.pb.FileID = id.toProto()
	return builder
}

func (builder FileDeleteTransaction) Build(client *Client) (Transaction, error) {
	return builder.TransactionBuilder.Build(client)
}

//
// The following _5_ must be copy-pasted at the bottom of **every** _transaction.go file
// We override the embedded fluent setter methods to return the outer type
//

// SetMaxTransactionFee sets the max transaction fee for this Transaction.
func (builder FileDeleteTransaction) SetMaxTransactionFee(maxTransactionFee Hbar) FileDeleteTransaction {
	return FileDeleteTransaction{builder.TransactionBuilder.SetMaxTransactionFee(maxTransactionFee), builder.pb}
}

// SetTransactionMemo sets the memo for this Transaction.
func (builder FileDeleteTransaction) SetTransactionMemo(memo string) FileDeleteTransaction {
	return FileDeleteTransaction{builder.TransactionBuilder.SetTransactionMemo(memo), builder.pb}
}

// SetTransactionValidDuration sets the valid duration for this Transaction.
func (builder FileDeleteTransaction) SetTransactionValidDuration(validDuration time.Duration) FileDeleteTransaction {
	return FileDeleteTransaction{builder.TransactionBuilder.SetTransactionValidDuration(validDuration), builder.pb}
}

// SetTransactionID sets the TransactionID for this Transaction.
func (builder FileDeleteTransaction) SetTransactionID(transactionID TransactionID) FileDeleteTransaction {
	return FileDeleteTransaction{builder.TransactionBuilder.SetTransactionID(transactionID), builder.pb}
}

// SetNodeAccountID sets the node AccountID for this Transaction.
func (builder FileDeleteTransaction) SetNodeAccountID(nodeAccountID AccountID) FileDeleteTransaction {
	return FileDeleteTransaction{builder.TransactionBuilder.SetNodeAccountID(nodeAccountID), builder.pb}
}

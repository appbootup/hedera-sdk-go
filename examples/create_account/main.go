package main

import "os"
import "github.com/hashgraph/hedera-sdk-go"

func main() {
	client, err := hedera.NewClient(
		// Node ID
		AccountID { account: 3 },
		// Node Address
		"0.testnet.hedera.com:50211",
	)

	if err != nil {
		panic(err)
	}

	operatorPrivateKey, err := hedera.Ed25519PrivateKeyFromString(os.Getenv("OPERATOR_KEY"))
	
	if err != nil {
		panic(err)
	}

	client.SetOperator(
		// Operator Account ID
		AccountID { account: 2 },
		// Operator Private Key
		operatorPrivateKey,
	)

	newKey := hedera.NewEd25519PrivateKey()
	newPublicKey := newKey.PublicKey()

	tx := hedera.NewAccountCreateTransaction(client).
		SetKey(newPublicKey).
		SetInitialBalance(1000).
		SetMaxTransactionFee(10000000)

	receipt, err := tx.ExecuteForReceipt()

	if err != nil {
		panic(err)
	}

	newAccountID := receipt.AccountID()
}
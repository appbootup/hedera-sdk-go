package hedera

type TransactionList struct {
	List []Transaction
}

// Sign uses the provided privateKey to sign the list of transactionas.
func (list TransactionList) Sign(privateKey Ed25519PrivateKey) TransactionList {
	for _, tx := range list.List {
		tx.Sign(privateKey)
	}

	return list
}

// SignWith executes the TransactionSigner and adds the resulting signature data to each Transaction's signature map
// with the publicKey as the map key.
func (list TransactionList) SignWith(publicKey Ed25519PublicKey, signer TransactionSigner) TransactionList {
	for _, tx := range list.List {
		tx.SignWith(publicKey, signer)
	}

	return list
}

func (list TransactionList) ExecuteAll(client *Client) ([]TransactionID, error) {
	ids := make([]TransactionID, len(list.List))
	for i, tx := range list.List {
		id, err := tx.Execute(client)
		if err != nil {
			return nil, err
		}

		ids[i] = id
	}

	return ids, nil
}

func (list TransactionList) Execute(client *Client) (TransactionID, error) {
	ids, err := list.ExecuteAll(client)
	if err != nil {
		return TransactionID{}, err
	}

	return ids[0], nil
}

package admin

import (
	"log"
	"math/big"

	"github.com/KibetBrian/geth/election"
	"github.com/KibetBrian/geth/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)


func GetRegistrationPhase() bool{
	Refresh();

	election, err := election.NewElection(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to create election instance: %v\n", err)
	}

	state, err := election.GetRegisrationPhase((&bind.CallOpts{From: wallet.PublicAddress, Context: ctx}));		
	if err != nil {
		log.Fatalf("Failed to get regisration phase: %v\n", err)
	}
	return state;
}

func RegisterCandidate(address, name, position string) TransactionResult {
	Refresh()


	election, err := election.NewElection(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to create election instance: %v\n", err)
	}

	//Transaction signer
	tx, err := bind.NewKeyedTransactorWithChainID(wallet.PrivateKey, netId)
	if err != nil {
		log.Fatalf("Failed to create Transaction Signer: %v\n", err)
	}

	tx.GasLimit = 3000000
	tx.GasPrice = gasPrice
	tx.Nonce = big.NewInt(int64(nonce))

	//Transactor
	candidatesAddress := utils.StringToAddress(address)
	transaction, err := election.RegisterCandidate(tx, candidatesAddress, name, position)
	if err != nil {
		log.Fatalf("Failed to add voter: %v", err)
	}
	res := TransactionResult{
		TransactionHash: transaction.Hash(),
		TransactionCost: transaction.Cost(),
		Success: true,
	}
	return res
}

func ChangeRegistrationPhase() *types.Transaction{
	Refresh();
	election, err := election.NewElection(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to create election instance: %v\n", err)
	}

	//Transaction signer
	tx, err := bind.NewKeyedTransactorWithChainID(wallet.PrivateKey, netId)
	if err != nil {
		log.Fatalf("Failed to create Transaction Signer: %v\n", err)
	}

	tx.GasLimit = 3000000
	tx.GasPrice = gasPrice
	tx.Nonce = big.NewInt(int64(nonce))

	transaction, err := election.ChangeRegistrationPhase(tx)
	if err != nil {
		log.Fatalf("Failed to change registration phase: %v", err);
	}
	return transaction;
}


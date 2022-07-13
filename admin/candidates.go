package admin

import (
	"log"

	"github.com/KibetBrian/geth/election"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func GetGorvernorCandidates() []election.ElectionCandidate {
	Refresh()

	election, err := election.NewElection(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to create election instance: %v\n", err)
	}

	gorvenorCandidates, err := election.GetGovernorCandidates((&bind.CallOpts{From: wallet.PublicAddress, Context: ctx}))
	if err != nil {
		log.Fatalf("Failed to retrieve candidates from the blockchain. Error : %v\n", err)
	}
	return gorvenorCandidates;
}

func GetPresidentCandidates() []election.ElectionCandidate{
	Refresh();
	election, err := election.NewElection(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to create election instance: %v\n", err)
	}

	presidentCandidates, err := election.GetPresidentCandidates((&bind.CallOpts{From: wallet.PublicAddress, Context: ctx}))
	if err != nil {
		log.Fatalf("Failed to retrieve candidates from the blockchain. Error : %v\n", err)
	}
	
	return presidentCandidates;
}


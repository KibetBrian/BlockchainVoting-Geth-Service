package admin

import (
	"log"

	"github.com/KibetBrian/geth/election"
	"github.com/KibetBrian/geth/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func GetGubenatorial() []election.ElectionCandidate {
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

func GetSpecificCandidate(address string) election.ElectionCandidate{
	Refresh();

	addr := utils.StringToAddress(address)

	election, err := election.NewElection(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to create election instance: %v\n", err)
	}

	candidate, err := election.GetSpecificCandidate((&bind.CallOpts{From: wallet.PublicAddress, Context: ctx}), addr)
	
	return candidate
}


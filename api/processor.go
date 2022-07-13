package api

import (
	"github.com/KibetBrian/geth/election"
	"github.com/KibetBrian/geth/utils"
)

//Checks if voters is already registered on the blockchain
func VoterRegistered ( voters []election.ElectionVoter, address string) bool {
	addr := utils.StringToAddress(address)

	for _, v := range voters {
		if v.Voter==addr && v.Registered{
			return true
		}
	}
	return false
}

//Checks if candidate is already registered in the database
func CandidateRegistered (p, g []election.ElectionCandidate, address string) bool{
	addr := utils.StringToAddress(address)

	for _, v := range p {
		if v.Addr==addr{
			return true
		}
	}
	for _, v := range g {
		if v.Addr==addr{
			return true
		}
	}
	return false
}
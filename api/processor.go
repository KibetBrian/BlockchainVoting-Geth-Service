package api

import (
	"github.com/KibetBrian/geth/election"
	"github.com/KibetBrian/geth/utils"
)


func VoterRegistered ( voters []election.ElectionVoter, address string) bool {
	addr := utils.StringToAddress(address)

	for _, v := range voters {
		if v.Addr==addr{
			return true
		}
	}
	return false
}
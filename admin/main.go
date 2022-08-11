package admin

import (
	"log"
	"math/big"
	"os"

	"github.com/KibetBrian/geth/core"
	"github.com/KibetBrian/geth/election"
	"github.com/KibetBrian/geth/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var wallet = core.DecryptKeystore(os.Getenv("wallet1passphrase"), 1)

var client, ctx = core.Connect()

var netId = utils.GetNetworkId(client, ctx)

var nonce = utils.GetPendingNonce(client, ctx, wallet.PublicAddress)

var gasPrice = utils.SuggestGasPrice(client, ctx)

var contractAddress = utils.StringToAddress("0xd0F0af13462445A9053C88a512Acb27f9fB6BFB8")

const Admin = "brian"

func Refresh() {
	utils.ConfigureEnv()

	wallet = core.DecryptKeystore(os.Getenv("wallet1passphrase"), 1)

	client, ctx := core.Connect()

	netId = utils.GetNetworkId(client, ctx)

	nonce = utils.GetPendingNonce(client, ctx, wallet.PublicAddress)
	gasPrice = utils.SuggestGasPrice(client, ctx)

}

func GetVoters() []election.ElectionVoter {
	Refresh()

	election, err := election.NewElection(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to create election instance: %v\n", err)
	}

	//Caller
	voters, err := election.GetAllVoters(&bind.CallOpts{From: wallet.PublicAddress, Context: ctx})
	if err != nil {
		log.Printf("Failed to get voters. Err: %v\n", err)
	}
	return voters
}

type TransactionResult struct {
	TransactionHash common.Hash
	TransactionCost *big.Int
	Success         bool
}

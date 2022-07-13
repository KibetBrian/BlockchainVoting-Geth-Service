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

var contractAddress = utils.StringToAddress("0x9f9B7Ca6703901a17454Ea2DD3Af98B107d472D8")

const Admin = "brian"

func Refresh() {
	utils.ConfigureEnv()

	wallet = core.DecryptKeystore(os.Getenv("wallet1passphrase"), 1)

	client, ctx := core.Connect()

	netId = utils.GetNetworkId(client, ctx)

	nonce = utils.GetPendingNonce(client, ctx, wallet.PublicAddress)
	gasPrice = utils.SuggestGasPrice(client, ctx)

	contractAddress = utils.StringToAddress("0x3311797f5BE82be3550dd3d22BF1AC76A6118C4F")
}

func main() {

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
	voterAddress := utils.StringToAddress("c50cd042edf5cf1d31cd5f9f04b499acd7e33a67")
	transaction, err := election.RegisterVoter(tx, voterAddress)
	if err != nil {
		log.Fatalf("Failed to add voter: %v", err)
	}
	log.Println("Transaction: ", transaction)
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


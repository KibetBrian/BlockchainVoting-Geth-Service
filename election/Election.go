// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package election

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ElectionCandidate is an auto generated low-level Go binding around an user-defined struct.
type ElectionCandidate struct {
	Name     string
	Addr     common.Address
	Position string
	Votes    *big.Int
}

// ElectionVoter is an auto generated low-level Go binding around an user-defined struct.
type ElectionVoter struct {
	Voter          common.Address
	President      common.Address
	Governor       common.Address
	VotedPresident bool
	VotedGovernor  bool
	Registered     bool
}

// ElectionMetaData contains all meta data concerning the Election contract.
var ElectionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"CastGovernor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"CastPresident\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ChangeRegistrationPhase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ChangeVotingPhase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetAllVoters\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"president\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"votedPresident\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"votedGovernor\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"}],\"internalType\":\"structElection.Voter[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"GetCandidateVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetGovernorCandidates\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"position\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"internalType\":\"structElection.Candidate[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetPresidentCandidates\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"position\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"internalType\":\"structElection.Candidate[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetRegisrationPhase\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetVotingPhase\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"position\",\"type\":\"string\"}],\"name\":\"RegisterCandidate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"RegisterVoter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260008060146101000a81548160ff02191690831515021790555060008060156101000a81548160ff02191690831515021790555034801561004457600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506125a1806100946000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c806387fb6de61161007157806387fb6de61461018f57806398c03ec4146101ad578063bcb962fb146101dd578063c1d128ba146101e7578063c4f1205b14610217578063f0887e4f14610235576100b4565b80630205ff5b146100b9578063112575d6146100c357806315fe6f1b146100e15780631e3ebc4c1461011157806330dae1671461012f578063744512f51461015f575b600080fd5b6100c1610253565b005b6100cb61030d565b6040516100d8919061191b565b60405180910390f35b6100fb60048036038101906100f691906119a8565b610323565b604051610108919061191b565b60405180910390f35b6101196107d0565b6040516101269190611bc2565b60405180910390f35b61014960048036038101906101449190611d19565b6109b3565b604051610156919061191b565b60405180910390f35b610179600480360381019061017491906119a8565b610ef7565b604051610186919061191b565b60405180910390f35b6101976111a3565b6040516101a4919061191b565b60405180910390f35b6101c760048036038101906101c291906119a8565b6111b9565b6040516101d49190611db3565b60405180910390f35b6101e5611205565b005b61020160048036038101906101fc91906119a8565b6112bf565b60405161020e919061191b565b60405180910390f35b61021f61156b565b60405161022c9190611bc2565b60405180910390f35b61023d61174e565b60405161024a9190611f07565b60405180910390f35b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102d890611fac565b60405180910390fd5b600060159054906101000a900460ff1615600060156101000a81548160ff021916908315150217905550565b60008060149054906101000a900460ff16905090565b600060011515600060149054906101000a900460ff1615151461037b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161037290612018565b60405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610409576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161040090611fac565b60405180910390fd5b60006040518060c001604052808473ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600015158152602001600015158152602001600115158152509050600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160169054906101000a900460ff166107bb5780600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160020160146101000a81548160ff02191690831515021790555060808201518160020160156101000a81548160ff02191690831515021790555060a08201518160020160166101000a81548160ff021916908315150217905550905050600581908060018154018082558091505060019003906000526020600020906003020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160020160146101000a81548160ff02191690831515021790555060808201518160020160156101000a81548160ff02191690831515021790555060a08201518160020160166101000a81548160ff02191690831515021790555050506107c5565b60009150506107cb565b60019150505b919050565b60606002805480602002602001604051908101604052809291908181526020016000905b828210156109aa578382906000526020600020906004020160405180608001604052908160008201805461082790612067565b80601f016020809104026020016040519081016040528092919081815260200182805461085390612067565b80156108a05780601f10610875576101008083540402835291602001916108a0565b820191906000526020600020905b81548152906001019060200180831161088357829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201805461090f90612067565b80601f016020809104026020016040519081016040528092919081815260200182805461093b90612067565b80156109885780601f1061095d57610100808354040283529160200191610988565b820191906000526020600020905b81548152906001019060200180831161096b57829003601f168201915b50505050508152602001600382015481525050815260200190600101906107f4565b50505050905090565b600060011515600060149054906101000a900460ff16151514610a0b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a0290612018565b60405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a99576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a9090611fac565b60405180910390fd5b60006040518060400160405280600981526020017f707265736964656e740000000000000000000000000000000000000000000000815250905060006040518060400160405280600881526020017f676f7665726e6f72000000000000000000000000000000000000000000000000815250905081604051602001610b1e91906120d4565b6040516020818303038152906040528051906020012084604051602001610b4591906120d4565b6040516020818303038152906040528051906020012003610cfb5760005b600380549050811015610c07578673ffffffffffffffffffffffffffffffffffffffff1660038281548110610b9b57610b9a6120eb565b5b906000526020600020906004020160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610bf45760009350505050610ef0565b8080610bff90612149565b915050610b63565b50600060405180608001604052808781526020018873ffffffffffffffffffffffffffffffffffffffff1681526020018681526020016000815250905060038190806001815401808255809150506001900390600052602060002090600402016000909190919091506000820151816000019081610c85919061233d565b5060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002019081610ce2919061233d565b5060608201518160030155505060019350505050610ef0565b80604051602001610d0c91906120d4565b6040516020818303038152906040528051906020012084604051602001610d3391906120d4565b6040516020818303038152906040528051906020012003610ee95760005b600280549050811015610df5578673ffffffffffffffffffffffffffffffffffffffff1660028281548110610d8957610d886120eb565b5b906000526020600020906004020160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610de25760009350505050610ef0565b8080610ded90612149565b915050610d51565b50600060405180608001604052808781526020018873ffffffffffffffffffffffffffffffffffffffff1681526020018681526020016000815250905060028190806001815401808255809150506001900390600052602060002090600402016000909190919091506000820151816000019081610e73919061233d565b5060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002019081610ed0919061233d565b5060608201518160030155505060019350505050610ef0565b6000925050505b9392505050565b600060011515600060159054906101000a900460ff16151514610f4f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f469061245b565b60405180910390fd5b60006040518060400160405280600981526020017f707265736964656e7400000000000000000000000000000000000000000000008152509050600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160149054906101000a900460ff16158015611072575080604051602001610ff291906120d4565b60405160208183030381529060405280519060200120600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160405160200161105a91906124fe565b60405160208183030381529060405280519060200120145b15611198576006339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600301600082825461112c9190612515565b925050819055506001600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160146101000a81548160ff021916908315150217905550600191505061119e565b60009150505b919050565b60008060159054906101000a900460ff16905090565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600301549050919050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611293576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161128a90611fac565b60405180910390fd5b600060149054906101000a900460ff1615600060146101000a81548160ff021916908315150217905550565b600060011515600060159054906101000a900460ff16151514611317576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161130e9061245b565b60405180910390fd5b60006040518060400160405280600881526020017f676f7276656e6f720000000000000000000000000000000000000000000000008152509050600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160159054906101000a900460ff1615801561143a5750806040516020016113ba91906120d4565b60405160208183030381529060405280519060200120600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160405160200161142291906124fe565b60405160208183030381529060405280519060200120145b15611560576007339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160008282546114f49190612515565b925050819055506001600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160156101000a81548160ff0219169083151502179055506001915050611566565b60009150505b919050565b60606003805480602002602001604051908101604052809291908181526020016000905b8282101561174557838290600052602060002090600402016040518060800160405290816000820180546115c290612067565b80601f01602080910402602001604051908101604052809291908181526020018280546115ee90612067565b801561163b5780601f106116105761010080835404028352916020019161163b565b820191906000526020600020905b81548152906001019060200180831161161e57829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820180546116aa90612067565b80601f01602080910402602001604051908101604052809291908181526020018280546116d690612067565b80156117235780601f106116f857610100808354040283529160200191611723565b820191906000526020600020905b81548152906001019060200180831161170657829003601f168201915b505050505081526020016003820154815250508152602001906001019061158f565b50505050905090565b60606005805480602002602001604051908101604052809291908181526020016000905b828210156118f757838290600052602060002090600302016040518060c00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160149054906101000a900460ff161515151581526020016002820160159054906101000a900460ff161515151581526020016002820160169054906101000a900460ff16151515158152505081526020019060010190611772565b50505050905090565b60008115159050919050565b61191581611900565b82525050565b6000602082019050611930600083018461190c565b92915050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006119758261194a565b9050919050565b6119858161196a565b811461199057600080fd5b50565b6000813590506119a28161197c565b92915050565b6000602082840312156119be576119bd611940565b5b60006119cc84828501611993565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b83811015611a3b578082015181840152602081019050611a20565b83811115611a4a576000848401525b50505050565b6000601f19601f8301169050919050565b6000611a6c82611a01565b611a768185611a0c565b9350611a86818560208601611a1d565b611a8f81611a50565b840191505092915050565b611aa38161196a565b82525050565b6000819050919050565b611abc81611aa9565b82525050565b60006080830160008301518482036000860152611adf8282611a61565b9150506020830151611af46020860182611a9a565b5060408301518482036040860152611b0c8282611a61565b9150506060830151611b216060860182611ab3565b508091505092915050565b6000611b388383611ac2565b905092915050565b6000602082019050919050565b6000611b58826119d5565b611b6281856119e0565b935083602082028501611b74856119f1565b8060005b85811015611bb05784840389528151611b918582611b2c565b9450611b9c83611b40565b925060208a01995050600181019050611b78565b50829750879550505050505092915050565b60006020820190508181036000830152611bdc8184611b4d565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b611c2682611a50565b810181811067ffffffffffffffff82111715611c4557611c44611bee565b5b80604052505050565b6000611c58611936565b9050611c648282611c1d565b919050565b600067ffffffffffffffff821115611c8457611c83611bee565b5b611c8d82611a50565b9050602081019050919050565b82818337600083830152505050565b6000611cbc611cb784611c69565b611c4e565b905082815260208101848484011115611cd857611cd7611be9565b5b611ce3848285611c9a565b509392505050565b600082601f830112611d0057611cff611be4565b5b8135611d10848260208601611ca9565b91505092915050565b600080600060608486031215611d3257611d31611940565b5b6000611d4086828701611993565b935050602084013567ffffffffffffffff811115611d6157611d60611945565b5b611d6d86828701611ceb565b925050604084013567ffffffffffffffff811115611d8e57611d8d611945565b5b611d9a86828701611ceb565b9150509250925092565b611dad81611aa9565b82525050565b6000602082019050611dc86000830184611da4565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b611e0381611900565b82525050565b60c082016000820151611e1f6000850182611a9a565b506020820151611e326020850182611a9a565b506040820151611e456040850182611a9a565b506060820151611e586060850182611dfa565b506080820151611e6b6080850182611dfa565b5060a0820151611e7e60a0850182611dfa565b50505050565b6000611e908383611e09565b60c08301905092915050565b6000602082019050919050565b6000611eb482611dce565b611ebe8185611dd9565b9350611ec983611dea565b8060005b83811015611efa578151611ee18882611e84565b9750611eec83611e9c565b925050600181019050611ecd565b5085935050505092915050565b60006020820190508181036000830152611f218184611ea9565b905092915050565b600082825260208201905092915050565b7f4f6e6c792061646d696e7320617265207065726d697474656420746f20646f2060008201527f7468697300000000000000000000000000000000000000000000000000000000602082015250565b6000611f96602483611f29565b9150611fa182611f3a565b604082019050919050565b60006020820190508181036000830152611fc581611f89565b9050919050565b7f526567697374726174696f6e207068617365206973206e6f74206f70656e0000600082015250565b6000612002601e83611f29565b915061200d82611fcc565b602082019050919050565b6000602082019050818103600083015261203181611ff5565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061207f57607f821691505b60208210810361209257612091612038565b5b50919050565b600081905092915050565b60006120ae82611a01565b6120b88185612098565b93506120c8818560208601611a1d565b80840191505092915050565b60006120e082846120a3565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061215482611aa9565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036121865761218561211a565b5b600182019050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026121f37fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826121b6565b6121fd86836121b6565b95508019841693508086168417925050509392505050565b6000819050919050565b600061223a61223561223084611aa9565b612215565b611aa9565b9050919050565b6000819050919050565b6122548361221f565b61226861226082612241565b8484546121c3565b825550505050565b600090565b61227d612270565b61228881848461224b565b505050565b5b818110156122ac576122a1600082612275565b60018101905061228e565b5050565b601f8211156122f1576122c281612191565b6122cb846121a6565b810160208510156122da578190505b6122ee6122e6856121a6565b83018261228d565b50505b505050565b600082821c905092915050565b6000612314600019846008026122f6565b1980831691505092915050565b600061232d8383612303565b9150826002028217905092915050565b61234682611a01565b67ffffffffffffffff81111561235f5761235e611bee565b5b6123698254612067565b6123748282856122b0565b600060209050601f8311600181146123a75760008415612395578287015190505b61239f8582612321565b865550612407565b601f1984166123b586612191565b60005b828110156123dd578489015182556001820191506020850194506020810190506123b8565b868310156123fa57848901516123f6601f891682612303565b8355505b6001600288020188555050505b505050505050565b7f566f74696e67207068617365206e6f74206f70656e0000000000000000000000600082015250565b6000612445601583611f29565b91506124508261240f565b602082019050919050565b6000602082019050818103600083015261247481612438565b9050919050565b6000815461248881612067565b6124928186612098565b945060018216600081146124ad57600181146124c2576124f5565b60ff19831686528115158202860193506124f5565b6124cb85612191565b60005b838110156124ed578154818901526001820191506020810190506124ce565b838801955050505b50505092915050565b600061250a828461247b565b915081905092915050565b600061252082611aa9565b915061252b83611aa9565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156125605761255f61211a565b5b82820190509291505056fea2646970667358221220adc355aa3b7cf674d077db43b2188b3823dcfc6f247cb8b01b7613f25dee6d0864736f6c634300080f0033",
}

// ElectionABI is the input ABI used to generate the binding from.
// Deprecated: Use ElectionMetaData.ABI instead.
var ElectionABI = ElectionMetaData.ABI

// ElectionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ElectionMetaData.Bin instead.
var ElectionBin = ElectionMetaData.Bin

// DeployElection deploys a new Ethereum contract, binding an instance of Election to it.
func DeployElection(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Election, error) {
	parsed, err := ElectionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ElectionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Election{ElectionCaller: ElectionCaller{contract: contract}, ElectionTransactor: ElectionTransactor{contract: contract}, ElectionFilterer: ElectionFilterer{contract: contract}}, nil
}

// Election is an auto generated Go binding around an Ethereum contract.
type Election struct {
	ElectionCaller     // Read-only binding to the contract
	ElectionTransactor // Write-only binding to the contract
	ElectionFilterer   // Log filterer for contract events
}

// ElectionCaller is an auto generated read-only Go binding around an Ethereum contract.
type ElectionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElectionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ElectionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElectionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ElectionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElectionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ElectionSession struct {
	Contract     *Election         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ElectionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ElectionCallerSession struct {
	Contract *ElectionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ElectionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ElectionTransactorSession struct {
	Contract     *ElectionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ElectionRaw is an auto generated low-level Go binding around an Ethereum contract.
type ElectionRaw struct {
	Contract *Election // Generic contract binding to access the raw methods on
}

// ElectionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ElectionCallerRaw struct {
	Contract *ElectionCaller // Generic read-only contract binding to access the raw methods on
}

// ElectionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ElectionTransactorRaw struct {
	Contract *ElectionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewElection creates a new instance of Election, bound to a specific deployed contract.
func NewElection(address common.Address, backend bind.ContractBackend) (*Election, error) {
	contract, err := bindElection(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Election{ElectionCaller: ElectionCaller{contract: contract}, ElectionTransactor: ElectionTransactor{contract: contract}, ElectionFilterer: ElectionFilterer{contract: contract}}, nil
}

// NewElectionCaller creates a new read-only instance of Election, bound to a specific deployed contract.
func NewElectionCaller(address common.Address, caller bind.ContractCaller) (*ElectionCaller, error) {
	contract, err := bindElection(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ElectionCaller{contract: contract}, nil
}

// NewElectionTransactor creates a new write-only instance of Election, bound to a specific deployed contract.
func NewElectionTransactor(address common.Address, transactor bind.ContractTransactor) (*ElectionTransactor, error) {
	contract, err := bindElection(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ElectionTransactor{contract: contract}, nil
}

// NewElectionFilterer creates a new log filterer instance of Election, bound to a specific deployed contract.
func NewElectionFilterer(address common.Address, filterer bind.ContractFilterer) (*ElectionFilterer, error) {
	contract, err := bindElection(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ElectionFilterer{contract: contract}, nil
}

// bindElection binds a generic wrapper to an already deployed contract.
func bindElection(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ElectionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Election *ElectionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Election.Contract.ElectionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Election *ElectionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Election.Contract.ElectionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Election *ElectionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Election.Contract.ElectionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Election *ElectionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Election.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Election *ElectionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Election.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Election *ElectionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Election.Contract.contract.Transact(opts, method, params...)
}

// GetAllVoters is a free data retrieval call binding the contract method 0xf0887e4f.
//
// Solidity: function GetAllVoters() view returns((address,address,address,bool,bool,bool)[])
func (_Election *ElectionCaller) GetAllVoters(opts *bind.CallOpts) ([]ElectionVoter, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "GetAllVoters")

	if err != nil {
		return *new([]ElectionVoter), err
	}

	out0 := *abi.ConvertType(out[0], new([]ElectionVoter)).(*[]ElectionVoter)

	return out0, err

}

// GetAllVoters is a free data retrieval call binding the contract method 0xf0887e4f.
//
// Solidity: function GetAllVoters() view returns((address,address,address,bool,bool,bool)[])
func (_Election *ElectionSession) GetAllVoters() ([]ElectionVoter, error) {
	return _Election.Contract.GetAllVoters(&_Election.CallOpts)
}

// GetAllVoters is a free data retrieval call binding the contract method 0xf0887e4f.
//
// Solidity: function GetAllVoters() view returns((address,address,address,bool,bool,bool)[])
func (_Election *ElectionCallerSession) GetAllVoters() ([]ElectionVoter, error) {
	return _Election.Contract.GetAllVoters(&_Election.CallOpts)
}

// GetCandidateVotes is a free data retrieval call binding the contract method 0x98c03ec4.
//
// Solidity: function GetCandidateVotes(address _address) view returns(uint256)
func (_Election *ElectionCaller) GetCandidateVotes(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "GetCandidateVotes", _address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCandidateVotes is a free data retrieval call binding the contract method 0x98c03ec4.
//
// Solidity: function GetCandidateVotes(address _address) view returns(uint256)
func (_Election *ElectionSession) GetCandidateVotes(_address common.Address) (*big.Int, error) {
	return _Election.Contract.GetCandidateVotes(&_Election.CallOpts, _address)
}

// GetCandidateVotes is a free data retrieval call binding the contract method 0x98c03ec4.
//
// Solidity: function GetCandidateVotes(address _address) view returns(uint256)
func (_Election *ElectionCallerSession) GetCandidateVotes(_address common.Address) (*big.Int, error) {
	return _Election.Contract.GetCandidateVotes(&_Election.CallOpts, _address)
}

// GetGovernorCandidates is a free data retrieval call binding the contract method 0x1e3ebc4c.
//
// Solidity: function GetGovernorCandidates() view returns((string,address,string,uint256)[])
func (_Election *ElectionCaller) GetGovernorCandidates(opts *bind.CallOpts) ([]ElectionCandidate, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "GetGovernorCandidates")

	if err != nil {
		return *new([]ElectionCandidate), err
	}

	out0 := *abi.ConvertType(out[0], new([]ElectionCandidate)).(*[]ElectionCandidate)

	return out0, err

}

// GetGovernorCandidates is a free data retrieval call binding the contract method 0x1e3ebc4c.
//
// Solidity: function GetGovernorCandidates() view returns((string,address,string,uint256)[])
func (_Election *ElectionSession) GetGovernorCandidates() ([]ElectionCandidate, error) {
	return _Election.Contract.GetGovernorCandidates(&_Election.CallOpts)
}

// GetGovernorCandidates is a free data retrieval call binding the contract method 0x1e3ebc4c.
//
// Solidity: function GetGovernorCandidates() view returns((string,address,string,uint256)[])
func (_Election *ElectionCallerSession) GetGovernorCandidates() ([]ElectionCandidate, error) {
	return _Election.Contract.GetGovernorCandidates(&_Election.CallOpts)
}

// GetPresidentCandidates is a free data retrieval call binding the contract method 0xc4f1205b.
//
// Solidity: function GetPresidentCandidates() view returns((string,address,string,uint256)[])
func (_Election *ElectionCaller) GetPresidentCandidates(opts *bind.CallOpts) ([]ElectionCandidate, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "GetPresidentCandidates")

	if err != nil {
		return *new([]ElectionCandidate), err
	}

	out0 := *abi.ConvertType(out[0], new([]ElectionCandidate)).(*[]ElectionCandidate)

	return out0, err

}

// GetPresidentCandidates is a free data retrieval call binding the contract method 0xc4f1205b.
//
// Solidity: function GetPresidentCandidates() view returns((string,address,string,uint256)[])
func (_Election *ElectionSession) GetPresidentCandidates() ([]ElectionCandidate, error) {
	return _Election.Contract.GetPresidentCandidates(&_Election.CallOpts)
}

// GetPresidentCandidates is a free data retrieval call binding the contract method 0xc4f1205b.
//
// Solidity: function GetPresidentCandidates() view returns((string,address,string,uint256)[])
func (_Election *ElectionCallerSession) GetPresidentCandidates() ([]ElectionCandidate, error) {
	return _Election.Contract.GetPresidentCandidates(&_Election.CallOpts)
}

// GetRegisrationPhase is a free data retrieval call binding the contract method 0x112575d6.
//
// Solidity: function GetRegisrationPhase() view returns(bool)
func (_Election *ElectionCaller) GetRegisrationPhase(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "GetRegisrationPhase")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetRegisrationPhase is a free data retrieval call binding the contract method 0x112575d6.
//
// Solidity: function GetRegisrationPhase() view returns(bool)
func (_Election *ElectionSession) GetRegisrationPhase() (bool, error) {
	return _Election.Contract.GetRegisrationPhase(&_Election.CallOpts)
}

// GetRegisrationPhase is a free data retrieval call binding the contract method 0x112575d6.
//
// Solidity: function GetRegisrationPhase() view returns(bool)
func (_Election *ElectionCallerSession) GetRegisrationPhase() (bool, error) {
	return _Election.Contract.GetRegisrationPhase(&_Election.CallOpts)
}

// GetVotingPhase is a free data retrieval call binding the contract method 0x87fb6de6.
//
// Solidity: function GetVotingPhase() view returns(bool)
func (_Election *ElectionCaller) GetVotingPhase(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "GetVotingPhase")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetVotingPhase is a free data retrieval call binding the contract method 0x87fb6de6.
//
// Solidity: function GetVotingPhase() view returns(bool)
func (_Election *ElectionSession) GetVotingPhase() (bool, error) {
	return _Election.Contract.GetVotingPhase(&_Election.CallOpts)
}

// GetVotingPhase is a free data retrieval call binding the contract method 0x87fb6de6.
//
// Solidity: function GetVotingPhase() view returns(bool)
func (_Election *ElectionCallerSession) GetVotingPhase() (bool, error) {
	return _Election.Contract.GetVotingPhase(&_Election.CallOpts)
}

// CastGovernor is a paid mutator transaction binding the contract method 0xc1d128ba.
//
// Solidity: function CastGovernor(address _address) returns(bool)
func (_Election *ElectionTransactor) CastGovernor(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Election.contract.Transact(opts, "CastGovernor", _address)
}

// CastGovernor is a paid mutator transaction binding the contract method 0xc1d128ba.
//
// Solidity: function CastGovernor(address _address) returns(bool)
func (_Election *ElectionSession) CastGovernor(_address common.Address) (*types.Transaction, error) {
	return _Election.Contract.CastGovernor(&_Election.TransactOpts, _address)
}

// CastGovernor is a paid mutator transaction binding the contract method 0xc1d128ba.
//
// Solidity: function CastGovernor(address _address) returns(bool)
func (_Election *ElectionTransactorSession) CastGovernor(_address common.Address) (*types.Transaction, error) {
	return _Election.Contract.CastGovernor(&_Election.TransactOpts, _address)
}

// CastPresident is a paid mutator transaction binding the contract method 0x744512f5.
//
// Solidity: function CastPresident(address _address) returns(bool)
func (_Election *ElectionTransactor) CastPresident(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Election.contract.Transact(opts, "CastPresident", _address)
}

// CastPresident is a paid mutator transaction binding the contract method 0x744512f5.
//
// Solidity: function CastPresident(address _address) returns(bool)
func (_Election *ElectionSession) CastPresident(_address common.Address) (*types.Transaction, error) {
	return _Election.Contract.CastPresident(&_Election.TransactOpts, _address)
}

// CastPresident is a paid mutator transaction binding the contract method 0x744512f5.
//
// Solidity: function CastPresident(address _address) returns(bool)
func (_Election *ElectionTransactorSession) CastPresident(_address common.Address) (*types.Transaction, error) {
	return _Election.Contract.CastPresident(&_Election.TransactOpts, _address)
}

// ChangeRegistrationPhase is a paid mutator transaction binding the contract method 0xbcb962fb.
//
// Solidity: function ChangeRegistrationPhase() returns()
func (_Election *ElectionTransactor) ChangeRegistrationPhase(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Election.contract.Transact(opts, "ChangeRegistrationPhase")
}

// ChangeRegistrationPhase is a paid mutator transaction binding the contract method 0xbcb962fb.
//
// Solidity: function ChangeRegistrationPhase() returns()
func (_Election *ElectionSession) ChangeRegistrationPhase() (*types.Transaction, error) {
	return _Election.Contract.ChangeRegistrationPhase(&_Election.TransactOpts)
}

// ChangeRegistrationPhase is a paid mutator transaction binding the contract method 0xbcb962fb.
//
// Solidity: function ChangeRegistrationPhase() returns()
func (_Election *ElectionTransactorSession) ChangeRegistrationPhase() (*types.Transaction, error) {
	return _Election.Contract.ChangeRegistrationPhase(&_Election.TransactOpts)
}

// ChangeVotingPhase is a paid mutator transaction binding the contract method 0x0205ff5b.
//
// Solidity: function ChangeVotingPhase() returns()
func (_Election *ElectionTransactor) ChangeVotingPhase(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Election.contract.Transact(opts, "ChangeVotingPhase")
}

// ChangeVotingPhase is a paid mutator transaction binding the contract method 0x0205ff5b.
//
// Solidity: function ChangeVotingPhase() returns()
func (_Election *ElectionSession) ChangeVotingPhase() (*types.Transaction, error) {
	return _Election.Contract.ChangeVotingPhase(&_Election.TransactOpts)
}

// ChangeVotingPhase is a paid mutator transaction binding the contract method 0x0205ff5b.
//
// Solidity: function ChangeVotingPhase() returns()
func (_Election *ElectionTransactorSession) ChangeVotingPhase() (*types.Transaction, error) {
	return _Election.Contract.ChangeVotingPhase(&_Election.TransactOpts)
}

// RegisterCandidate is a paid mutator transaction binding the contract method 0x30dae167.
//
// Solidity: function RegisterCandidate(address _address, string _name, string position) returns(bool)
func (_Election *ElectionTransactor) RegisterCandidate(opts *bind.TransactOpts, _address common.Address, _name string, position string) (*types.Transaction, error) {
	return _Election.contract.Transact(opts, "RegisterCandidate", _address, _name, position)
}

// RegisterCandidate is a paid mutator transaction binding the contract method 0x30dae167.
//
// Solidity: function RegisterCandidate(address _address, string _name, string position) returns(bool)
func (_Election *ElectionSession) RegisterCandidate(_address common.Address, _name string, position string) (*types.Transaction, error) {
	return _Election.Contract.RegisterCandidate(&_Election.TransactOpts, _address, _name, position)
}

// RegisterCandidate is a paid mutator transaction binding the contract method 0x30dae167.
//
// Solidity: function RegisterCandidate(address _address, string _name, string position) returns(bool)
func (_Election *ElectionTransactorSession) RegisterCandidate(_address common.Address, _name string, position string) (*types.Transaction, error) {
	return _Election.Contract.RegisterCandidate(&_Election.TransactOpts, _address, _name, position)
}

// RegisterVoter is a paid mutator transaction binding the contract method 0x15fe6f1b.
//
// Solidity: function RegisterVoter(address _address) returns(bool)
func (_Election *ElectionTransactor) RegisterVoter(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Election.contract.Transact(opts, "RegisterVoter", _address)
}

// RegisterVoter is a paid mutator transaction binding the contract method 0x15fe6f1b.
//
// Solidity: function RegisterVoter(address _address) returns(bool)
func (_Election *ElectionSession) RegisterVoter(_address common.Address) (*types.Transaction, error) {
	return _Election.Contract.RegisterVoter(&_Election.TransactOpts, _address)
}

// RegisterVoter is a paid mutator transaction binding the contract method 0x15fe6f1b.
//
// Solidity: function RegisterVoter(address _address) returns(bool)
func (_Election *ElectionTransactorSession) RegisterVoter(_address common.Address) (*types.Transaction, error) {
	return _Election.Contract.RegisterVoter(&_Election.TransactOpts, _address)
}

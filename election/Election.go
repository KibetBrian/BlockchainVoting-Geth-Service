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
	VotedPresident bool
	VotedGovernor  bool
	Registered     bool
}

// ElectionMetaData contains all meta data concerning the Election contract.
var ElectionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"CastGovernor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"CastPresident\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ChangeRegistrationPhase\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ChangeVotingPhase\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetAllVoters\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"votedPresident\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"votedGovernor\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"}],\"internalType\":\"structElection.Voter[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"GetCandidateVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetGovernorCandidates\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"position\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"internalType\":\"structElection.Candidate[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetGubernatorial\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetPresidentCandidates\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"position\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"internalType\":\"structElection.Candidate[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetPresidentialVotes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetRegisrationPhase\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"GetSpecificCandidate\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"position\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"internalType\":\"structElection.Candidate\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetVotingPhase\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"position\",\"type\":\"string\"}],\"name\":\"RegisterCandidate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"RegisterVoter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260008060146101000a81548160ff02191690831515021790555060008060156101000a81548160ff02191690831515021790555034801561004457600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550612a5a806100946000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c8063744512f511610097578063bcb962fb11610066578063bcb962fb1461029e578063c1d128ba146102bc578063c4f1205b146102ec578063f0887e4f1461030a576100f5565b8063744512f51461020257806387fb6de61461023257806398c03ec4146102505780639ceb907c14610280576100f5565b80631e3ebc4c116100d35780631e3ebc4c1461016657806330dae1671461018457806340794885146101b45780634ddbf6e1146101d2576100f5565b80630205ff5b146100fa578063112575d61461011857806315fe6f1b14610136575b600080fd5b610102610328565b60405161010f9190611cc8565b60405180910390f35b6101206103f8565b60405161012d9190611cc8565b60405180910390f35b610150600480360381019061014b9190611d55565b61040e565b60405161015d9190611cc8565b60405180910390f35b61016e610761565b60405161017b9190611f66565b60405180910390f35b61019e600480360381019061019991906120bd565b610944565b6040516101ab9190611cc8565b60405180910390f35b6101bc611008565b6040516101c991906121f7565b60405180910390f35b6101ec60048036038101906101e79190611d55565b611096565b6040516101f99190612283565b60405180910390f35b61021c60048036038101906102179190611d55565b611273565b6040516102299190611cc8565b60405180910390f35b61023a61151e565b6040516102479190611cc8565b60405180910390f35b61026a60048036038101906102659190611d55565b611534565b60405161027791906122b4565b60405180910390f35b610288611580565b60405161029591906121f7565b60405180910390f35b6102a661160e565b6040516102b39190611cc8565b60405180910390f35b6102d660048036038101906102d19190611d55565b6116de565b6040516102e39190611cc8565b60405180910390f35b6102f461198a565b6040516103019190611f66565b60405180910390f35b610312611b6d565b60405161031f91906123e2565b60405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103b090612487565b60405180910390fd5b600060159054906101000a900460ff1615600060156101000a81548160ff021916908315150217905550600060159054906101000a900460ff16905090565b60008060149054906101000a900460ff16905090565b600060011515600060149054906101000a900460ff16151514610466576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161045d906124f3565b60405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104f4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104eb90612487565b60405180910390fd5b600060405180608001604052808473ffffffffffffffffffffffffffffffffffffffff168152602001600015158152602001600015158152602001600115158152509050600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160169054906101000a900460ff1661074c5780600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff02191690831515021790555060408201518160000160156101000a81548160ff02191690831515021790555060608201518160000160166101000a81548160ff0219169083151502179055509050506005819080600181540180825580915050600190039060005260206000200160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff02191690831515021790555060408201518160000160156101000a81548160ff02191690831515021790555060608201518160000160166101000a81548160ff0219169083151502179055505050610756565b600091505061075c565b60019150505b919050565b60606002805480602002602001604051908101604052809291908181526020016000905b8282101561093b57838290600052602060002090600402016040518060800160405290816000820180546107b890612542565b80601f01602080910402602001604051908101604052809291908181526020018280546107e490612542565b80156108315780601f1061080657610100808354040283529160200191610831565b820191906000526020600020905b81548152906001019060200180831161081457829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820180546108a090612542565b80601f01602080910402602001604051908101604052809291908181526020018280546108cc90612542565b80156109195780601f106108ee57610100808354040283529160200191610919565b820191906000526020600020905b8154815290600101906020018083116108fc57829003601f168201915b5050505050815260200160038201548152505081526020019060010190610785565b50505050905090565b600060011515600060149054906101000a900460ff1615151461099c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610993906124f3565b60405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a2a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a2190612487565b60405180910390fd5b60006040518060400160405280600981526020017f707265736964656e740000000000000000000000000000000000000000000000815250905060006040518060400160405280600881526020017f676f7665726e6f72000000000000000000000000000000000000000000000000815250905081604051602001610aaf91906125af565b6040516020818303038152906040528051906020012084604051602001610ad691906125af565b6040516020818303038152906040528051906020012003610d4c5760005b600380549050811015610b98578673ffffffffffffffffffffffffffffffffffffffff1660038281548110610b2c57610b2b6125c6565b5b906000526020600020906004020160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610b855760009350505050611001565b8080610b9090612624565b915050610af4565b50600060405180608001604052808781526020018873ffffffffffffffffffffffffffffffffffffffff1681526020018681526020016000815250905080600160008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000019081610c2a9190612818565b5060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002019081610c879190612818565b506060820151816003015590505060038190806001815401808255809150506001900390600052602060002090600402016000909190919091506000820151816000019081610cd69190612818565b5060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002019081610d339190612818565b5060608201518160030155505060019350505050611001565b80604051602001610d5d91906125af565b6040516020818303038152906040528051906020012084604051602001610d8491906125af565b6040516020818303038152906040528051906020012003610ffa5760005b600280549050811015610e46578673ffffffffffffffffffffffffffffffffffffffff1660028281548110610dda57610dd96125c6565b5b906000526020600020906004020160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610e335760009350505050611001565b8080610e3e90612624565b915050610da2565b50600060405180608001604052808781526020018873ffffffffffffffffffffffffffffffffffffffff1681526020018681526020016000815250905080600160008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000019081610ed89190612818565b5060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002019081610f359190612818565b506060820151816003015590505060028190806001815401808255809150506001900390600052602060002090600402016000909190919091506000820151816000019081610f849190612818565b5060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002019081610fe19190612818565b5060608201518160030155505060019350505050611001565b6000925050505b9392505050565b6060600680548060200260200160405190810160405280929190818152602001828054801561108c57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611042575b5050505050905090565b61109e611c6f565b600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060800160405290816000820180546110f890612542565b80601f016020809104026020016040519081016040528092919081815260200182805461112490612542565b80156111715780601f1061114657610100808354040283529160200191611171565b820191906000526020600020905b81548152906001019060200180831161115457829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820180546111e090612542565b80601f016020809104026020016040519081016040528092919081815260200182805461120c90612542565b80156112595780601f1061122e57610100808354040283529160200191611259565b820191906000526020600020905b81548152906001019060200180831161123c57829003601f168201915b505050505081526020016003820154815250509050919050565b600060011515600060159054906101000a900460ff161515146112cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112c290612936565b60405180910390fd5b60006040518060400160405280600981526020017f707265736964656e7400000000000000000000000000000000000000000000008152509050600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160149054906101000a900460ff161580156113ee57508060405160200161136e91906125af565b60405160208183030381529060405280519060200120600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002016040516020016113d691906129d9565b60405160208183030381529060405280519060200120145b15611513576006339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160008154809291906114a990612624565b91905055506001600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160146101000a81548160ff0219169083151502179055506001915050611519565b60009150505b919050565b60008060159054906101000a900460ff16905090565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600301549050919050565b6060600780548060200260200160405190810160405280929190818152602001828054801561160457602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116115ba575b5050505050905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461169f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161169690612487565b60405180910390fd5b600060149054906101000a900460ff1615600060146101000a81548160ff021916908315150217905550600060149054906101000a900460ff16905090565b600060011515600060159054906101000a900460ff16151514611736576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161172d90612936565b60405180910390fd5b60006040518060400160405280600881526020017f676f7276656e6f720000000000000000000000000000000000000000000000008152509050600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160159054906101000a900460ff161580156118595750806040516020016117d991906125af565b60405160208183030381529060405280519060200120600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160405160200161184191906129d9565b60405160208183030381529060405280519060200120145b1561197f576007339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600301600082825461191391906129f0565b925050819055506001600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160156101000a81548160ff0219169083151502179055506001915050611985565b60009150505b919050565b60606003805480602002602001604051908101604052809291908181526020016000905b82821015611b6457838290600052602060002090600402016040518060800160405290816000820180546119e190612542565b80601f0160208091040260200160405190810160405280929190818152602001828054611a0d90612542565b8015611a5a5780601f10611a2f57610100808354040283529160200191611a5a565b820191906000526020600020905b815481529060010190602001808311611a3d57829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282018054611ac990612542565b80601f0160208091040260200160405190810160405280929190818152602001828054611af590612542565b8015611b425780601f10611b1757610100808354040283529160200191611b42565b820191906000526020600020905b815481529060010190602001808311611b2557829003601f168201915b50505050508152602001600382015481525050815260200190600101906119ae565b50505050905090565b60606005805480602002602001604051908101604052809291908181526020016000905b82821015611c66578382906000526020600020016040518060800160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff161515151581526020016000820160159054906101000a900460ff161515151581526020016000820160169054906101000a900460ff16151515158152505081526020019060010190611b91565b50505050905090565b604051806080016040528060608152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001600081525090565b60008115159050919050565b611cc281611cad565b82525050565b6000602082019050611cdd6000830184611cb9565b92915050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000611d2282611cf7565b9050919050565b611d3281611d17565b8114611d3d57600080fd5b50565b600081359050611d4f81611d29565b92915050565b600060208284031215611d6b57611d6a611ced565b5b6000611d7984828501611d40565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b83811015611de8578082015181840152602081019050611dcd565b60008484015250505050565b6000601f19601f8301169050919050565b6000611e1082611dae565b611e1a8185611db9565b9350611e2a818560208601611dca565b611e3381611df4565b840191505092915050565b611e4781611d17565b82525050565b6000819050919050565b611e6081611e4d565b82525050565b60006080830160008301518482036000860152611e838282611e05565b9150506020830151611e986020860182611e3e565b5060408301518482036040860152611eb08282611e05565b9150506060830151611ec56060860182611e57565b508091505092915050565b6000611edc8383611e66565b905092915050565b6000602082019050919050565b6000611efc82611d82565b611f068185611d8d565b935083602082028501611f1885611d9e565b8060005b85811015611f545784840389528151611f358582611ed0565b9450611f4083611ee4565b925060208a01995050600181019050611f1c565b50829750879550505050505092915050565b60006020820190508181036000830152611f808184611ef1565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b611fca82611df4565b810181811067ffffffffffffffff82111715611fe957611fe8611f92565b5b80604052505050565b6000611ffc611ce3565b90506120088282611fc1565b919050565b600067ffffffffffffffff82111561202857612027611f92565b5b61203182611df4565b9050602081019050919050565b82818337600083830152505050565b600061206061205b8461200d565b611ff2565b90508281526020810184848401111561207c5761207b611f8d565b5b61208784828561203e565b509392505050565b600082601f8301126120a4576120a3611f88565b5b81356120b484826020860161204d565b91505092915050565b6000806000606084860312156120d6576120d5611ced565b5b60006120e486828701611d40565b935050602084013567ffffffffffffffff81111561210557612104611cf2565b5b6121118682870161208f565b925050604084013567ffffffffffffffff81111561213257612131611cf2565b5b61213e8682870161208f565b9150509250925092565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b60006121808383611e3e565b60208301905092915050565b6000602082019050919050565b60006121a482612148565b6121ae8185612153565b93506121b983612164565b8060005b838110156121ea5781516121d18882612174565b97506121dc8361218c565b9250506001810190506121bd565b5085935050505092915050565b600060208201905081810360008301526122118184612199565b905092915050565b600060808301600083015184820360008601526122368282611e05565b915050602083015161224b6020860182611e3e565b50604083015184820360408601526122638282611e05565b91505060608301516122786060860182611e57565b508091505092915050565b6000602082019050818103600083015261229d8184612219565b905092915050565b6122ae81611e4d565b82525050565b60006020820190506122c960008301846122a5565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61230481611cad565b82525050565b6080820160008201516123206000850182611e3e565b50602082015161233360208501826122fb565b50604082015161234660408501826122fb565b50606082015161235960608501826122fb565b50505050565b600061236b838361230a565b60808301905092915050565b6000602082019050919050565b600061238f826122cf565b61239981856122da565b93506123a4836122eb565b8060005b838110156123d55781516123bc888261235f565b97506123c783612377565b9250506001810190506123a8565b5085935050505092915050565b600060208201905081810360008301526123fc8184612384565b905092915050565b600082825260208201905092915050565b7f4f6e6c792061646d696e7320617265207065726d697474656420746f20646f2060008201527f7468697300000000000000000000000000000000000000000000000000000000602082015250565b6000612471602483612404565b915061247c82612415565b604082019050919050565b600060208201905081810360008301526124a081612464565b9050919050565b7f526567697374726174696f6e207068617365206973206e6f74206f70656e0000600082015250565b60006124dd601e83612404565b91506124e8826124a7565b602082019050919050565b6000602082019050818103600083015261250c816124d0565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061255a57607f821691505b60208210810361256d5761256c612513565b5b50919050565b600081905092915050565b600061258982611dae565b6125938185612573565b93506125a3818560208601611dca565b80840191505092915050565b60006125bb828461257e565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061262f82611e4d565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203612661576126606125f5565b5b600182019050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026126ce7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82612691565b6126d88683612691565b95508019841693508086168417925050509392505050565b6000819050919050565b600061271561271061270b84611e4d565b6126f0565b611e4d565b9050919050565b6000819050919050565b61272f836126fa565b61274361273b8261271c565b84845461269e565b825550505050565b600090565b61275861274b565b612763818484612726565b505050565b5b818110156127875761277c600082612750565b600181019050612769565b5050565b601f8211156127cc5761279d8161266c565b6127a684612681565b810160208510156127b5578190505b6127c96127c185612681565b830182612768565b50505b505050565b600082821c905092915050565b60006127ef600019846008026127d1565b1980831691505092915050565b600061280883836127de565b9150826002028217905092915050565b61282182611dae565b67ffffffffffffffff81111561283a57612839611f92565b5b6128448254612542565b61284f82828561278b565b600060209050601f8311600181146128825760008415612870578287015190505b61287a85826127fc565b8655506128e2565b601f1984166128908661266c565b60005b828110156128b857848901518255600182019150602085019450602081019050612893565b868310156128d557848901516128d1601f8916826127de565b8355505b6001600288020188555050505b505050505050565b7f566f74696e67207068617365206e6f74206f70656e0000000000000000000000600082015250565b6000612920601583612404565b915061292b826128ea565b602082019050919050565b6000602082019050818103600083015261294f81612913565b9050919050565b6000815461296381612542565b61296d8186612573565b94506001821660008114612988576001811461299d576129d0565b60ff19831686528115158202860193506129d0565b6129a68561266c565b60005b838110156129c8578154818901526001820191506020810190506129a9565b838801955050505b50505092915050565b60006129e58284612956565b915081905092915050565b60006129fb82611e4d565b9150612a0683611e4d565b9250828201905080821115612a1e57612a1d6125f5565b5b9291505056fea2646970667358221220fe2e757493f3a447e434fa7e0603389378058cde41fe628d72ecbe88eb11c6da64736f6c63430008100033",
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
// Solidity: function GetAllVoters() view returns((address,bool,bool,bool)[])
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
// Solidity: function GetAllVoters() view returns((address,bool,bool,bool)[])
func (_Election *ElectionSession) GetAllVoters() ([]ElectionVoter, error) {
	return _Election.Contract.GetAllVoters(&_Election.CallOpts)
}

// GetAllVoters is a free data retrieval call binding the contract method 0xf0887e4f.
//
// Solidity: function GetAllVoters() view returns((address,bool,bool,bool)[])
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

// GetGubernatorial is a free data retrieval call binding the contract method 0x9ceb907c.
//
// Solidity: function GetGubernatorial() view returns(address[])
func (_Election *ElectionCaller) GetGubernatorial(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "GetGubernatorial")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetGubernatorial is a free data retrieval call binding the contract method 0x9ceb907c.
//
// Solidity: function GetGubernatorial() view returns(address[])
func (_Election *ElectionSession) GetGubernatorial() ([]common.Address, error) {
	return _Election.Contract.GetGubernatorial(&_Election.CallOpts)
}

// GetGubernatorial is a free data retrieval call binding the contract method 0x9ceb907c.
//
// Solidity: function GetGubernatorial() view returns(address[])
func (_Election *ElectionCallerSession) GetGubernatorial() ([]common.Address, error) {
	return _Election.Contract.GetGubernatorial(&_Election.CallOpts)
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

// GetPresidentialVotes is a free data retrieval call binding the contract method 0x40794885.
//
// Solidity: function GetPresidentialVotes() view returns(address[])
func (_Election *ElectionCaller) GetPresidentialVotes(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "GetPresidentialVotes")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPresidentialVotes is a free data retrieval call binding the contract method 0x40794885.
//
// Solidity: function GetPresidentialVotes() view returns(address[])
func (_Election *ElectionSession) GetPresidentialVotes() ([]common.Address, error) {
	return _Election.Contract.GetPresidentialVotes(&_Election.CallOpts)
}

// GetPresidentialVotes is a free data retrieval call binding the contract method 0x40794885.
//
// Solidity: function GetPresidentialVotes() view returns(address[])
func (_Election *ElectionCallerSession) GetPresidentialVotes() ([]common.Address, error) {
	return _Election.Contract.GetPresidentialVotes(&_Election.CallOpts)
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

// GetSpecificCandidate is a free data retrieval call binding the contract method 0x4ddbf6e1.
//
// Solidity: function GetSpecificCandidate(address _addr) view returns((string,address,string,uint256))
func (_Election *ElectionCaller) GetSpecificCandidate(opts *bind.CallOpts, _addr common.Address) (ElectionCandidate, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "GetSpecificCandidate", _addr)

	if err != nil {
		return *new(ElectionCandidate), err
	}

	out0 := *abi.ConvertType(out[0], new(ElectionCandidate)).(*ElectionCandidate)

	return out0, err

}

// GetSpecificCandidate is a free data retrieval call binding the contract method 0x4ddbf6e1.
//
// Solidity: function GetSpecificCandidate(address _addr) view returns((string,address,string,uint256))
func (_Election *ElectionSession) GetSpecificCandidate(_addr common.Address) (ElectionCandidate, error) {
	return _Election.Contract.GetSpecificCandidate(&_Election.CallOpts, _addr)
}

// GetSpecificCandidate is a free data retrieval call binding the contract method 0x4ddbf6e1.
//
// Solidity: function GetSpecificCandidate(address _addr) view returns((string,address,string,uint256))
func (_Election *ElectionCallerSession) GetSpecificCandidate(_addr common.Address) (ElectionCandidate, error) {
	return _Election.Contract.GetSpecificCandidate(&_Election.CallOpts, _addr)
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
// Solidity: function ChangeRegistrationPhase() returns(bool)
func (_Election *ElectionTransactor) ChangeRegistrationPhase(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Election.contract.Transact(opts, "ChangeRegistrationPhase")
}

// ChangeRegistrationPhase is a paid mutator transaction binding the contract method 0xbcb962fb.
//
// Solidity: function ChangeRegistrationPhase() returns(bool)
func (_Election *ElectionSession) ChangeRegistrationPhase() (*types.Transaction, error) {
	return _Election.Contract.ChangeRegistrationPhase(&_Election.TransactOpts)
}

// ChangeRegistrationPhase is a paid mutator transaction binding the contract method 0xbcb962fb.
//
// Solidity: function ChangeRegistrationPhase() returns(bool)
func (_Election *ElectionTransactorSession) ChangeRegistrationPhase() (*types.Transaction, error) {
	return _Election.Contract.ChangeRegistrationPhase(&_Election.TransactOpts)
}

// ChangeVotingPhase is a paid mutator transaction binding the contract method 0x0205ff5b.
//
// Solidity: function ChangeVotingPhase() returns(bool)
func (_Election *ElectionTransactor) ChangeVotingPhase(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Election.contract.Transact(opts, "ChangeVotingPhase")
}

// ChangeVotingPhase is a paid mutator transaction binding the contract method 0x0205ff5b.
//
// Solidity: function ChangeVotingPhase() returns(bool)
func (_Election *ElectionSession) ChangeVotingPhase() (*types.Transaction, error) {
	return _Election.Contract.ChangeVotingPhase(&_Election.TransactOpts)
}

// ChangeVotingPhase is a paid mutator transaction binding the contract method 0x0205ff5b.
//
// Solidity: function ChangeVotingPhase() returns(bool)
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

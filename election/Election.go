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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"CastGovernor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"CastPresident\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ChangeRegistrationPhase\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ChangeVotingPhase\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetAllVoters\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"president\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"votedPresident\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"votedGovernor\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"}],\"internalType\":\"structElection.Voter[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"GetCandidateVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetGovernorCandidates\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"position\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"internalType\":\"structElection.Candidate[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetPresidentCandidates\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"position\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"internalType\":\"structElection.Candidate[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetRegisrationPhase\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetVotingPhase\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"position\",\"type\":\"string\"}],\"name\":\"RegisterCandidate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"RegisterVoter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260008060146101000a81548160ff02191690831515021790555060008060156101000a81548160ff02191690831515021790555034801561004457600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506125ca806100946000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c806387fb6de61161007157806387fb6de6146101a357806398c03ec4146101c1578063bcb962fb146101f1578063c1d128ba1461020f578063c4f1205b1461023f578063f0887e4f1461025d576100b4565b80630205ff5b146100b9578063112575d6146100d757806315fe6f1b146100f55780631e3ebc4c1461012557806330dae16714610143578063744512f514610173575b600080fd5b6100c161027b565b6040516100ce919061196f565b60405180910390f35b6100df61034b565b6040516100ec919061196f565b60405180910390f35b61010f600480360381019061010a91906119fc565b610361565b60405161011c919061196f565b60405180910390f35b61012d61080e565b60405161013a9190611c0d565b60405180910390f35b61015d60048036038101906101589190611d64565b6109f1565b60405161016a919061196f565b60405180910390f35b61018d600480360381019061018891906119fc565b610f35565b60405161019a919061196f565b60405180910390f35b6101ab6111e1565b6040516101b8919061196f565b60405180910390f35b6101db60048036038101906101d691906119fc565b6111f7565b6040516101e89190611dfe565b60405180910390f35b6101f9611243565b604051610206919061196f565b60405180910390f35b610229600480360381019061022491906119fc565b611313565b604051610236919061196f565b60405180910390f35b6102476115bf565b6040516102549190611c0d565b60405180910390f35b6102656117a2565b6040516102729190611f52565b60405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461030c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161030390611ff7565b60405180910390fd5b600060159054906101000a900460ff1615600060156101000a81548160ff021916908315150217905550600060159054906101000a900460ff16905090565b60008060149054906101000a900460ff16905090565b600060011515600060149054906101000a900460ff161515146103b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103b090612063565b60405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610447576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161043e90611ff7565b60405180910390fd5b60006040518060c001604052808473ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600015158152602001600015158152602001600115158152509050600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160169054906101000a900460ff166107f95780600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160020160146101000a81548160ff02191690831515021790555060808201518160020160156101000a81548160ff02191690831515021790555060a08201518160020160166101000a81548160ff021916908315150217905550905050600581908060018154018082558091505060019003906000526020600020906003020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160020160146101000a81548160ff02191690831515021790555060808201518160020160156101000a81548160ff02191690831515021790555060a08201518160020160166101000a81548160ff0219169083151502179055505050610803565b6000915050610809565b60019150505b919050565b60606002805480602002602001604051908101604052809291908181526020016000905b828210156109e85783829060005260206000209060040201604051806080016040529081600082018054610865906120b2565b80601f0160208091040260200160405190810160405280929190818152602001828054610891906120b2565b80156108de5780601f106108b3576101008083540402835291602001916108de565b820191906000526020600020905b8154815290600101906020018083116108c157829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201805461094d906120b2565b80601f0160208091040260200160405190810160405280929190818152602001828054610979906120b2565b80156109c65780601f1061099b576101008083540402835291602001916109c6565b820191906000526020600020905b8154815290600101906020018083116109a957829003601f168201915b5050505050815260200160038201548152505081526020019060010190610832565b50505050905090565b600060011515600060149054906101000a900460ff16151514610a49576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a4090612063565b60405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ad7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ace90611ff7565b60405180910390fd5b60006040518060400160405280600981526020017f707265736964656e740000000000000000000000000000000000000000000000815250905060006040518060400160405280600881526020017f676f7665726e6f72000000000000000000000000000000000000000000000000815250905081604051602001610b5c919061211f565b6040516020818303038152906040528051906020012084604051602001610b83919061211f565b6040516020818303038152906040528051906020012003610d395760005b600380549050811015610c45578673ffffffffffffffffffffffffffffffffffffffff1660038281548110610bd957610bd8612136565b5b906000526020600020906004020160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610c325760009350505050610f2e565b8080610c3d90612194565b915050610ba1565b50600060405180608001604052808781526020018873ffffffffffffffffffffffffffffffffffffffff1681526020018681526020016000815250905060038190806001815401808255809150506001900390600052602060002090600402016000909190919091506000820151816000019081610cc39190612388565b5060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002019081610d209190612388565b5060608201518160030155505060019350505050610f2e565b80604051602001610d4a919061211f565b6040516020818303038152906040528051906020012084604051602001610d71919061211f565b6040516020818303038152906040528051906020012003610f275760005b600280549050811015610e33578673ffffffffffffffffffffffffffffffffffffffff1660028281548110610dc757610dc6612136565b5b906000526020600020906004020160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610e205760009350505050610f2e565b8080610e2b90612194565b915050610d8f565b50600060405180608001604052808781526020018873ffffffffffffffffffffffffffffffffffffffff1681526020018681526020016000815250905060028190806001815401808255809150506001900390600052602060002090600402016000909190919091506000820151816000019081610eb19190612388565b5060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002019081610f0e9190612388565b5060608201518160030155505060019350505050610f2e565b6000925050505b9392505050565b600060011515600060159054906101000a900460ff16151514610f8d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f84906124a6565b60405180910390fd5b60006040518060400160405280600981526020017f707265736964656e7400000000000000000000000000000000000000000000008152509050600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160149054906101000a900460ff161580156110b0575080604051602001611030919061211f565b60405160208183030381529060405280519060200120600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002016040516020016110989190612549565b60405160208183030381529060405280519060200120145b156111d6576006339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600301600082825461116a9190612560565b925050819055506001600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160146101000a81548160ff02191690831515021790555060019150506111dc565b60009150505b919050565b60008060159054906101000a900460ff16905090565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600301549050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146112d4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112cb90611ff7565b60405180910390fd5b600060149054906101000a900460ff1615600060146101000a81548160ff021916908315150217905550600060149054906101000a900460ff16905090565b600060011515600060159054906101000a900460ff1615151461136b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611362906124a6565b60405180910390fd5b60006040518060400160405280600881526020017f676f7276656e6f720000000000000000000000000000000000000000000000008152509050600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160159054906101000a900460ff1615801561148e57508060405160200161140e919061211f565b60405160208183030381529060405280519060200120600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002016040516020016114769190612549565b60405160208183030381529060405280519060200120145b156115b4576007339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160008282546115489190612560565b925050819055506001600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160156101000a81548160ff02191690831515021790555060019150506115ba565b60009150505b919050565b60606003805480602002602001604051908101604052809291908181526020016000905b828210156117995783829060005260206000209060040201604051806080016040529081600082018054611616906120b2565b80601f0160208091040260200160405190810160405280929190818152602001828054611642906120b2565b801561168f5780601f106116645761010080835404028352916020019161168f565b820191906000526020600020905b81548152906001019060200180831161167257829003601f168201915b505050505081526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820180546116fe906120b2565b80601f016020809104026020016040519081016040528092919081815260200182805461172a906120b2565b80156117775780601f1061174c57610100808354040283529160200191611777565b820191906000526020600020905b81548152906001019060200180831161175a57829003601f168201915b50505050508152602001600382015481525050815260200190600101906115e3565b50505050905090565b60606005805480602002602001604051908101604052809291908181526020016000905b8282101561194b57838290600052602060002090600302016040518060c00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160149054906101000a900460ff161515151581526020016002820160159054906101000a900460ff161515151581526020016002820160169054906101000a900460ff161515151581525050815260200190600101906117c6565b50505050905090565b60008115159050919050565b61196981611954565b82525050565b60006020820190506119846000830184611960565b92915050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006119c98261199e565b9050919050565b6119d9816119be565b81146119e457600080fd5b50565b6000813590506119f6816119d0565b92915050565b600060208284031215611a1257611a11611994565b5b6000611a20848285016119e7565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b83811015611a8f578082015181840152602081019050611a74565b60008484015250505050565b6000601f19601f8301169050919050565b6000611ab782611a55565b611ac18185611a60565b9350611ad1818560208601611a71565b611ada81611a9b565b840191505092915050565b611aee816119be565b82525050565b6000819050919050565b611b0781611af4565b82525050565b60006080830160008301518482036000860152611b2a8282611aac565b9150506020830151611b3f6020860182611ae5565b5060408301518482036040860152611b578282611aac565b9150506060830151611b6c6060860182611afe565b508091505092915050565b6000611b838383611b0d565b905092915050565b6000602082019050919050565b6000611ba382611a29565b611bad8185611a34565b935083602082028501611bbf85611a45565b8060005b85811015611bfb5784840389528151611bdc8582611b77565b9450611be783611b8b565b925060208a01995050600181019050611bc3565b50829750879550505050505092915050565b60006020820190508181036000830152611c278184611b98565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b611c7182611a9b565b810181811067ffffffffffffffff82111715611c9057611c8f611c39565b5b80604052505050565b6000611ca361198a565b9050611caf8282611c68565b919050565b600067ffffffffffffffff821115611ccf57611cce611c39565b5b611cd882611a9b565b9050602081019050919050565b82818337600083830152505050565b6000611d07611d0284611cb4565b611c99565b905082815260208101848484011115611d2357611d22611c34565b5b611d2e848285611ce5565b509392505050565b600082601f830112611d4b57611d4a611c2f565b5b8135611d5b848260208601611cf4565b91505092915050565b600080600060608486031215611d7d57611d7c611994565b5b6000611d8b868287016119e7565b935050602084013567ffffffffffffffff811115611dac57611dab611999565b5b611db886828701611d36565b925050604084013567ffffffffffffffff811115611dd957611dd8611999565b5b611de586828701611d36565b9150509250925092565b611df881611af4565b82525050565b6000602082019050611e136000830184611def565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b611e4e81611954565b82525050565b60c082016000820151611e6a6000850182611ae5565b506020820151611e7d6020850182611ae5565b506040820151611e906040850182611ae5565b506060820151611ea36060850182611e45565b506080820151611eb66080850182611e45565b5060a0820151611ec960a0850182611e45565b50505050565b6000611edb8383611e54565b60c08301905092915050565b6000602082019050919050565b6000611eff82611e19565b611f098185611e24565b9350611f1483611e35565b8060005b83811015611f45578151611f2c8882611ecf565b9750611f3783611ee7565b925050600181019050611f18565b5085935050505092915050565b60006020820190508181036000830152611f6c8184611ef4565b905092915050565b600082825260208201905092915050565b7f4f6e6c792061646d696e7320617265207065726d697474656420746f20646f2060008201527f7468697300000000000000000000000000000000000000000000000000000000602082015250565b6000611fe1602483611f74565b9150611fec82611f85565b604082019050919050565b6000602082019050818103600083015261201081611fd4565b9050919050565b7f526567697374726174696f6e207068617365206973206e6f74206f70656e0000600082015250565b600061204d601e83611f74565b915061205882612017565b602082019050919050565b6000602082019050818103600083015261207c81612040565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806120ca57607f821691505b6020821081036120dd576120dc612083565b5b50919050565b600081905092915050565b60006120f982611a55565b61210381856120e3565b9350612113818560208601611a71565b80840191505092915050565b600061212b82846120ee565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061219f82611af4565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036121d1576121d0612165565b5b600182019050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b60006008830261223e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82612201565b6122488683612201565b95508019841693508086168417925050509392505050565b6000819050919050565b600061228561228061227b84611af4565b612260565b611af4565b9050919050565b6000819050919050565b61229f8361226a565b6122b36122ab8261228c565b84845461220e565b825550505050565b600090565b6122c86122bb565b6122d3818484612296565b505050565b5b818110156122f7576122ec6000826122c0565b6001810190506122d9565b5050565b601f82111561233c5761230d816121dc565b612316846121f1565b81016020851015612325578190505b612339612331856121f1565b8301826122d8565b50505b505050565b600082821c905092915050565b600061235f60001984600802612341565b1980831691505092915050565b6000612378838361234e565b9150826002028217905092915050565b61239182611a55565b67ffffffffffffffff8111156123aa576123a9611c39565b5b6123b482546120b2565b6123bf8282856122fb565b600060209050601f8311600181146123f257600084156123e0578287015190505b6123ea858261236c565b865550612452565b601f198416612400866121dc565b60005b8281101561242857848901518255600182019150602085019450602081019050612403565b868310156124455784890151612441601f89168261234e565b8355505b6001600288020188555050505b505050505050565b7f566f74696e67207068617365206e6f74206f70656e0000000000000000000000600082015250565b6000612490601583611f74565b915061249b8261245a565b602082019050919050565b600060208201905081810360008301526124bf81612483565b9050919050565b600081546124d3816120b2565b6124dd81866120e3565b945060018216600081146124f8576001811461250d57612540565b60ff1983168652811515820286019350612540565b612516856121dc565b60005b8381101561253857815481890152600182019150602081019050612519565b838801955050505b50505092915050565b600061255582846124c6565b915081905092915050565b600061256b82611af4565b915061257683611af4565b925082820190508082111561258e5761258d612165565b5b9291505056fea26469706673582212202507abfc2a562447dad1753ada7cd49056651aef60f5d79a3714a1403556aff164736f6c63430008100033",
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

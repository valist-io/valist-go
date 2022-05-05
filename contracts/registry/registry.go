// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registry

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

// RegistryMetaData contains all meta data concerning the Registry contract.
var RegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_forwarder\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_accountID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_metaURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"AccountCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_accountID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"AccountMemberAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_accountID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"AccountMemberRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_accountID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_metaURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"AccountUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_accountID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_projectID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_metaURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"ProjectCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_projectID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"ProjectMemberAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_projectID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_member\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"ProjectMemberRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_projectID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_metaURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"ProjectUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_releaseID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"ReleaseApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_projectID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_releaseID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_metaURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"ReleaseCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_releaseID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"ReleaseRevoked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_accountID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addAccountMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addProjectMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_releaseID\",\"type\":\"uint256\"}],\"name\":\"approveRelease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_metaURI\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"_members\",\"type\":\"address[]\"}],\"name\":\"createAccount\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_accountID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_metaURI\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"_members\",\"type\":\"address[]\"}],\"name\":\"createProject\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_metaURI\",\"type\":\"string\"}],\"name\":\"createRelease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_parentID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"generateID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_accountID\",\"type\":\"uint256\"}],\"name\":\"getAccountMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectID\",\"type\":\"uint256\"}],\"name\":\"getLatestReleaseID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_releaseID\",\"type\":\"uint256\"}],\"name\":\"getPreviousReleaseID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectID\",\"type\":\"uint256\"}],\"name\":\"getProjectAccountID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectID\",\"type\":\"uint256\"}],\"name\":\"getProjectMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_releaseID\",\"type\":\"uint256\"}],\"name\":\"getReleaseProjectID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_releaseID\",\"type\":\"uint256\"}],\"name\":\"getReleaseSigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_accountID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_member\",\"type\":\"address\"}],\"name\":\"isAccountMember\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_member\",\"type\":\"address\"}],\"name\":\"isProjectMember\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_releaseID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"isReleaseSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"metaByID\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_accountID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"removeAccountMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"removeProjectMember\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_releaseID\",\"type\":\"uint256\"}],\"name\":\"revokeRelease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_accountID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_metaURI\",\"type\":\"string\"}],\"name\":\"setAccountMetaURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_claimFee\",\"type\":\"uint256\"}],\"name\":\"setClaimFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_projectID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_metaURI\",\"type\":\"string\"}],\"name\":\"setProjectMetaURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_forwarder\",\"type\":\"address\"}],\"name\":\"setTrustedForwarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedForwarder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"versionRecipient\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60c06040526005608081905264322e322e3360d81b60a09081526200002691908162000097565b503480156200003457600080fd5b50604051620022403803806200224083398101604081905262000057916200013d565b600680546001600160a01b031916331790556200009081600080546001600160a01b0319166001600160a01b0392909216919091179055565b50620001aa565b828054620000a5906200016d565b90600052602060002090601f016020900481019282620000c9576000855562000114565b82601f10620000e457805160ff191683800117855562000114565b8280016001018555821562000114579182015b8281111562000114578251825591602001919060010190620000f7565b506200012292915062000126565b5090565b5b8082111562000122576000815560010162000127565b6000602082840312156200014f578081fd5b81516001600160a01b038116811462000166578182fd5b9392505050565b600181811c908216806200018257607f821691505b60208210811415620001a457634e487b7160e01b600052602260045260246000fd5b50919050565b61208680620001ba6000396000f3fe6080604052600436106101d85760003560e01c80637da0a87711610102578063bb11d22a11610095578063da74222811610064578063da742228146105b5578063db5954f4146105d5578063ec95ecb6146105f5578063ef83d3ab1461061557600080fd5b8063bb11d22a14610532578063c25cf2d514610545578063d515d9d914610565578063d77302571461058557600080fd5b806399d32fc4116100d157806399d32fc4146104af5780639f4cfef9146104c5578063a2f6ce9b146104e5578063aef28ac31461050557600080fd5b80637da0a87714610410578063806553f31461044257806383b1fc41146104625780638da5cb5b1461048f57600080fd5b8063416cab0c1161017a578063600c391111610149578063600c39111461038057806368704f02146103a05780636ad0a5ee146103d05780637c869665146103f057600080fd5b8063416cab0c146102d257806343417c06146102ff578063486ff0cd1461032c578063572b6c051461034157600080fd5b80631bd2cea1116101b65780631bd2cea114610252578063250b72de1461027257806328a10ec6146102925780632e75ab50146102b257600080fd5b80630859bbcd146101dd57806313af403514610210578063176037f514610232575b600080fd5b3480156101e957600080fd5b506101fd6101f8366004611ba7565b610635565b6040519081526020015b60405180910390f35b34801561021c57600080fd5b5061023061022b366004611ac0565b61067c565b005b34801561023e57600080fd5b5061023061024d366004611b78565b6106dc565b34801561025e57600080fd5b5061023061026d366004611b60565b61079c565b34801561027e57600080fd5b5061023061028d366004611c4c565b610883565b34801561029e57600080fd5b506102306102ad366004611bec565b610a98565b3480156102be57600080fd5b506102306102cd366004611b60565b610c3b565b3480156102de57600080fd5b506101fd6102ed366004611b60565b60009081526002602052604090205490565b34801561030b57600080fd5b5061031f61031a366004611b60565b610c75565b6040516102079190611d72565b34801561033857600080fd5b5061031f610d0f565b34801561034d57600080fd5b5061037061035c366004611ac0565b6000546001600160a01b0391821691161490565b6040519015158152602001610207565b34801561038c57600080fd5b5061037061039b366004611b78565b610d1c565b3480156103ac57600080fd5b506101fd6103bb366004611b60565b60009081526003602052604090206001015490565b3480156103dc57600080fd5b506103706103eb366004611b78565b610d3e565b3480156103fc57600080fd5b5061037061040b366004611b78565b610d56565b34801561041c57600080fd5b506000546001600160a01b03165b6040516001600160a01b039091168152602001610207565b34801561044e57600080fd5b5061023061045d366004611b60565b610d71565b34801561046e57600080fd5b5061048261047d366004611b60565b610e1d565b6040516102079190611d25565b34801561049b57600080fd5b5060065461042a906001600160a01b031681565b3480156104bb57600080fd5b506101fd60075481565b3480156104d157600080fd5b506102306104e0366004611ba7565b610e3a565b3480156104f157600080fd5b50610482610500366004611b60565b610f15565b34801561051157600080fd5b506101fd610520366004611b60565b60009081526003602052604090205490565b610230610540366004611adc565b610f2f565b34801561055157600080fd5b50610230610560366004611b78565b611198565b34801561057157600080fd5b50610482610580366004611b60565b6112a4565b34801561059157600080fd5b506101fd6105a0366004611b60565b60009081526002602052604090206001015490565b3480156105c157600080fd5b506102306105d0366004611ac0565b6112c4565b3480156105e157600080fd5b506102306105f0366004611ba7565b61131a565b34801561060157600080fd5b50610230610610366004611b78565b611404565b34801561062157600080fd5b50610230610630366004611b78565b611508565b600082828051906020012060405160200161065a929190918252602082015260400190565b6040516020818303038152906040528051906020012060001c90505b92915050565b6106846115bd565b6006546001600160a01b039081169116146106ba5760405162461bcd60e51b81526004016106b190611dd9565b60405180910390fd5b600680546001600160a01b0319166001600160a01b0392909216919091179055565b6106e8826103eb6115bd565b6107045760405162461bcd60e51b81526004016106b190611e60565b61070e8282610d3e565b1561072b5760405162461bcd60e51b81526004016106b190611d85565b600082815260016020526040902061074390826115f1565b507fda06031a12c2b6322a6a2a29b8e7ec92a209562f0995acbf9e5fe8cce89617b3828261076f6115bd565b604080519384526001600160a01b0392831660208501529116908201526060015b60405180910390a15050565b600081815260046020526040812080546107b590611fb9565b9050116107d45760405162461bcd60e51b81526004016106b190611e88565b6107f66107df6115bd565b600083815260036020526040902060020190611606565b156108135760405162461bcd60e51b81526004016106b190611d85565b61083561081e6115bd565b6000838152600360205260409020600201906115f1565b507f492f404869834bbfb3dfd354d8ebd27ca988aa59eda4724d2e604400e1635df0816108606115bd565b604080519283526001600160a01b0390911660208301520160405180910390a150565b60008251116108a45760405162461bcd60e51b81526004016106b190611e10565b60008351116108c55760405162461bcd60e51b81526004016106b190611e38565b60006108d18585610635565b90506108df856103eb6115bd565b6108fb5760405162461bcd60e51b81526004016106b190611e60565b6000818152600460205260409020805461091490611fb9565b1590506109335760405162461bcd60e51b81526004016106b190611daf565b6000818152600460209081526040909120845161095292860190611930565b5060008181526002602052604090208590557fdce0cbd02a04a917caf7c163eeb153eee961368abd158a17a8d17298afc2a62f858286866109916115bd565b6040516109a2959493929190611f25565b60405180910390a160005b8251811015610a9057610a058382815181106109d957634e487b7160e01b600052603260045260246000fd5b6020026020010151600260008581526020019081526020016000206002016115f190919063ffffffff16565b507fa641673cc2c2925ec554d696d7ecb92ad62fd37a878bb62194e01d15af4ecbfe82848381518110610a4857634e487b7160e01b600052603260045260246000fd5b6020026020010151610a586115bd565b604080519384526001600160a01b03928316602085015291169082015260600160405180910390a1610a8981611ff4565b90506109ad565b505050505050565b6000825111610ab95760405162461bcd60e51b81526004016106b190611e38565b6000815111610ada5760405162461bcd60e51b81526004016106b190611e10565b60008381526004602052604081208054610af390611fb9565b905011610b125760405162461bcd60e51b81526004016106b190611e88565b6000610b1e8484610635565b6000818152600460205260409020805491925090610b3b90611fb9565b159050610b5a5760405162461bcd60e51b81526004016106b190611daf565b600084815260026020526040902054610b758561039b6115bd565b80610b875750610b87816103eb6115bd565b610ba35760405162461bcd60e51b81526004016106b190611e60565b60008581526002602090815260408083206001018054908690558584526004835292208551610bd492870190611930565b506000838152600360205260409020600181018290558690557f28f16f2c48f7434e8ee65a00bca94583249ae3dc79ebd4c8fac3ce1cc41227a986848787610c1a6115bd565b604051610c2b959493929190611f25565b60405180910390a1505050505050565b610c436115bd565b6006546001600160a01b03908116911614610c705760405162461bcd60e51b81526004016106b190611dd9565b600755565b60046020526000908152604090208054610c8e90611fb9565b80601f0160208091040260200160405190810160405280929190818152602001828054610cba90611fb9565b8015610d075780601f10610cdc57610100808354040283529160200191610d07565b820191906000526020600020905b815481529060010190602001808311610cea57829003601f168201915b505050505081565b60058054610c8e90611fb9565b60008281526002602081905260408220610d37910183611606565b9392505050565b6000828152600160205260408120610d379083611606565b6000828152600360205260408120610d379060020183611606565b60008181526004602052604081208054610d8a90611fb9565b905011610da95760405162461bcd60e51b81526004016106b190611e88565b610db46107df6115bd565b610dd05760405162461bcd60e51b81526004016106b190611d85565b610df2610ddb6115bd565b600083815260036020526040902060020190611628565b507f6709c97fa27e5e2954e292620faa1a587d1f25adb1bef740c229b7d947b41a99816108606115bd565b60008181526003602052604090206060906106769060020161163d565b6000815111610e5b5760405162461bcd60e51b81526004016106b190611e10565b610e67826103eb6115bd565b610e835760405162461bcd60e51b81526004016106b190611e60565b60008281526004602052604081208054610e9c90611fb9565b905011610ebb5760405162461bcd60e51b81526004016106b190611e88565b60008281526004602090815260409091208251610eda92840190611930565b507fdc2b36fe3e9f9129790c8dc1fefe27a260e2ab12bbd7d2503bf3c701b7a6b1078282610f066115bd565b60405161079093929190611eaf565b60008181526001602052604090206060906106769061163d565b600754341015610f6d5760405162461bcd60e51b81526020600482015260096024820152686572722d76616c756560b81b60448201526064016106b1565b6000825111610f8e5760405162461bcd60e51b81526004016106b190611e10565b6000835111610faf5760405162461bcd60e51b81526004016106b190611e38565b6000815111610ff45760405162461bcd60e51b81526020600482015260116024820152706572722d656d7074792d6d656d6265727360781b60448201526064016106b1565b60006110004685610635565b600081815260046020526040902080549192509061101d90611fb9565b15905061103c5760405162461bcd60e51b81526004016106b190611daf565b6000818152600460209081526040909120845161105b92860190611930565b507f7bab27b7daa748e5bd855a000a2445c6e1d22fc5b2a6c4df5c96f63508afc7b98185856110886115bd565b6040516110989493929190611ee0565b60405180910390a160005b825181101561117b576110f08382815181106110cf57634e487b7160e01b600052603260045260246000fd5b602090810291909101810151600085815260019092526040909120906115f1565b507fda06031a12c2b6322a6a2a29b8e7ec92a209562f0995acbf9e5fe8cce89617b38284838151811061113357634e487b7160e01b600052603260045260246000fd5b60200260200101516111436115bd565b604080519384526001600160a01b03928316602085015291169082015260600160405180910390a161117481611ff4565b90506110a3565b50600654611192906001600160a01b03163461164a565b50505050565b600082815260046020526040812080546111b190611fb9565b9050116111d05760405162461bcd60e51b81526004016106b190611e88565b6111da8282610d1c565b156111f75760405162461bcd60e51b81526004016106b190611d85565b600082815260026020526040902054611212816103eb6115bd565b61122e5760405162461bcd60e51b81526004016106b190611e60565b600083815260026020819052604090912061124a9101836115f1565b507fa641673cc2c2925ec554d696d7ecb92ad62fd37a878bb62194e01d15af4ecbfe83836112766115bd565b604080519384526001600160a01b0392831660208501529116908201526060015b60405180910390a1505050565b60606106766002600084815260200190815260200160002060020161163d565b6112cc6115bd565b6006546001600160a01b039081169116146112f95760405162461bcd60e51b81526004016106b190611dd9565b600080546001600160a01b0319166001600160a01b03831617905550565b50565b600081511161133b5760405162461bcd60e51b81526004016106b190611e10565b6000828152600460205260408120805461135490611fb9565b9050116113735760405162461bcd60e51b81526004016106b190611e88565b60008281526002602052604090205461138e816103eb6115bd565b6113aa5760405162461bcd60e51b81526004016106b190611e60565b600083815260046020908152604090912083516113c992850190611930565b507fb1c971f708cf6cac0bf1b55c79db4b7e06ed48aa4e24618f1e40cbebecc5b06383836113f56115bd565b60405161129793929190611eaf565b6000828152600460205260408120805461141d90611fb9565b90501161143c5760405162461bcd60e51b81526004016106b190611e88565b6114468282610d1c565b6114895760405162461bcd60e51b8152602060048201526014602482015273195c9c8b5b595b58995c8b5b9bdd0b595e1a5cdd60621b60448201526064016106b1565b6000828152600260205260409020546114a4816103eb6115bd565b6114c05760405162461bcd60e51b81526004016106b190611e60565b60008381526002602081905260409091206114dc910183611628565b507f3c2f57383ed351452c76a1869a6abf17e5521b3de241aa8d31d1769d4e193c2483836112766115bd565b611514826103eb6115bd565b6115305760405162461bcd60e51b81526004016106b190611e60565b61153a8282610d3e565b61157d5760405162461bcd60e51b8152602060048201526014602482015273195c9c8b5b595b58995c8b5b9bdd0b595e1a5cdd60621b60448201526064016106b1565b60008281526001602052604090206115959082611628565b507febdaf29c191f35cacd52dad5ad5e63fe0a2806c921f5f362dfcc9ded63874550828261076f5b6000601436108015906115da57506000546001600160a01b031633145b156115ec575060131936013560601c90565b503390565b6000610d37836001600160a01b038416611768565b6001600160a01b03811660009081526001830160205260408120541515610d37565b6000610d37836001600160a01b0384166117b7565b60606000610d37836118d4565b8047101561169a5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e636500000060448201526064016106b1565b6000826001600160a01b03168260405160006040518083038185875af1925050503d80600081146116e7576040519150601f19603f3d011682016040523d82523d6000602084013e6116ec565b606091505b50509050806117635760405162461bcd60e51b815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d6179206861766520726576657274656400000000000060648201526084016106b1565b505050565b60008181526001830160205260408120546117af57508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610676565b506000610676565b600081815260018301602052604081205480156118ca5760006117db600183611fa2565b85549091506000906117ef90600190611fa2565b905081811461187057600086600001828154811061181d57634e487b7160e01b600052603260045260246000fd5b906000526020600020015490508087600001848154811061184e57634e487b7160e01b600052603260045260246000fd5b6000918252602080832090910192909255918252600188019052604090208390555b855486908061188f57634e487b7160e01b600052603160045260246000fd5b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610676565b6000915050610676565b60608160000180548060200260200160405190810160405280929190818152602001828054801561192457602002820191906000526020600020905b815481526020019060010190808311611910575b50505050509050919050565b82805461193c90611fb9565b90600052602060002090601f01602090048101928261195e57600085556119a4565b82601f1061197757805160ff19168380011785556119a4565b828001600101855582156119a4579182015b828111156119a4578251825591602001919060010190611989565b506119b09291506119b4565b5090565b5b808211156119b057600081556001016119b5565b600082601f8301126119d9578081fd5b8135602067ffffffffffffffff8211156119f5576119f5612025565b8160051b611a04828201611f71565b838152828101908684018388018501891015611a1e578687fd5b8693505b85841015611a49578035611a358161203b565b835260019390930192918401918401611a22565b50979650505050505050565b600082601f830112611a65578081fd5b813567ffffffffffffffff811115611a7f57611a7f612025565b611a92601f8201601f1916602001611f71565b818152846020838601011115611aa6578283fd5b816020850160208301379081016020019190915292915050565b600060208284031215611ad1578081fd5b8135610d378161203b565b600080600060608486031215611af0578182fd5b833567ffffffffffffffff80821115611b07578384fd5b611b1387838801611a55565b94506020860135915080821115611b28578384fd5b611b3487838801611a55565b93506040860135915080821115611b49578283fd5b50611b56868287016119c9565b9150509250925092565b600060208284031215611b71578081fd5b5035919050565b60008060408385031215611b8a578182fd5b823591506020830135611b9c8161203b565b809150509250929050565b60008060408385031215611bb9578182fd5b82359150602083013567ffffffffffffffff811115611bd6578182fd5b611be285828601611a55565b9150509250929050565b600080600060608486031215611c00578283fd5b83359250602084013567ffffffffffffffff80821115611c1e578384fd5b611c2a87838801611a55565b93506040860135915080821115611c3f578283fd5b50611b5686828701611a55565b60008060008060808587031215611c61578081fd5b84359350602085013567ffffffffffffffff80821115611c7f578283fd5b611c8b88838901611a55565b94506040870135915080821115611ca0578283fd5b611cac88838901611a55565b93506060870135915080821115611cc1578283fd5b50611cce878288016119c9565b91505092959194509250565b60008151808452815b81811015611cff57602081850181015186830182015201611ce3565b81811115611d105782602083870101525b50601f01601f19169290920160200192915050565b6020808252825182820181905260009190848201906040850190845b81811015611d665783516001600160a01b031683529284019291840191600101611d41565b50909695505050505050565b602081526000610d376020830184611cda565b60208082526010908201526f195c9c8b5b595b58995c8b595e1a5cdd60821b604082015260600190565b60208082526010908201526f195c9c8b5b985b594b58db185a5b595960821b604082015260600190565b60208082526017908201527f63616c6c6572206973206e6f7420746865206f776e6572000000000000000000604082015260600190565b6020808252600e908201526d6572722d656d7074792d6d65746160901b604082015260600190565b6020808252600e908201526d6572722d656d7074792d6e616d6560901b604082015260600190565b6020808252600e908201526d32b93916b737ba16b6b2b6b132b960911b604082015260600190565b6020808252600d908201526c195c9c8b5b9bdd0b595e1a5cdd609a1b604082015260600190565b838152606060208201526000611ec86060830185611cda565b905060018060a01b0383166040830152949350505050565b848152608060208201526000611ef96080830186611cda565b8281036040840152611f0b8186611cda565b91505060018060a01b038316606083015295945050505050565b85815284602082015260a060408201526000611f4460a0830186611cda565b8281036060840152611f568186611cda565b91505060018060a01b03831660808301529695505050505050565b604051601f8201601f1916810167ffffffffffffffff81118282101715611f9a57611f9a612025565b604052919050565b600082821015611fb457611fb461200f565b500390565b600181811c90821680611fcd57607f821691505b60208210811415611fee57634e487b7160e01b600052602260045260246000fd5b50919050565b60006000198214156120085761200861200f565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b038116811461131757600080fdfea2646970667358221220312ff2177abae76444f49a87fc70a4042f16d55792741e871f48441a7bcb1bc364736f6c63430008040033",
}

// RegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistryMetaData.ABI instead.
var RegistryABI = RegistryMetaData.ABI

// RegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RegistryMetaData.Bin instead.
var RegistryBin = RegistryMetaData.Bin

// DeployRegistry deploys a new Ethereum contract, binding an instance of Registry to it.
func DeployRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _forwarder common.Address) (common.Address, *types.Transaction, *Registry, error) {
	parsed, err := RegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RegistryBin), backend, _forwarder)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// ClaimFee is a free data retrieval call binding the contract method 0x99d32fc4.
//
// Solidity: function claimFee() view returns(uint256)
func (_Registry *RegistryCaller) ClaimFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "claimFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimFee is a free data retrieval call binding the contract method 0x99d32fc4.
//
// Solidity: function claimFee() view returns(uint256)
func (_Registry *RegistrySession) ClaimFee() (*big.Int, error) {
	return _Registry.Contract.ClaimFee(&_Registry.CallOpts)
}

// ClaimFee is a free data retrieval call binding the contract method 0x99d32fc4.
//
// Solidity: function claimFee() view returns(uint256)
func (_Registry *RegistryCallerSession) ClaimFee() (*big.Int, error) {
	return _Registry.Contract.ClaimFee(&_Registry.CallOpts)
}

// GenerateID is a free data retrieval call binding the contract method 0x0859bbcd.
//
// Solidity: function generateID(uint256 _parentID, string _name) pure returns(uint256)
func (_Registry *RegistryCaller) GenerateID(opts *bind.CallOpts, _parentID *big.Int, _name string) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "generateID", _parentID, _name)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GenerateID is a free data retrieval call binding the contract method 0x0859bbcd.
//
// Solidity: function generateID(uint256 _parentID, string _name) pure returns(uint256)
func (_Registry *RegistrySession) GenerateID(_parentID *big.Int, _name string) (*big.Int, error) {
	return _Registry.Contract.GenerateID(&_Registry.CallOpts, _parentID, _name)
}

// GenerateID is a free data retrieval call binding the contract method 0x0859bbcd.
//
// Solidity: function generateID(uint256 _parentID, string _name) pure returns(uint256)
func (_Registry *RegistryCallerSession) GenerateID(_parentID *big.Int, _name string) (*big.Int, error) {
	return _Registry.Contract.GenerateID(&_Registry.CallOpts, _parentID, _name)
}

// GetAccountMembers is a free data retrieval call binding the contract method 0xa2f6ce9b.
//
// Solidity: function getAccountMembers(uint256 _accountID) view returns(address[])
func (_Registry *RegistryCaller) GetAccountMembers(opts *bind.CallOpts, _accountID *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getAccountMembers", _accountID)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAccountMembers is a free data retrieval call binding the contract method 0xa2f6ce9b.
//
// Solidity: function getAccountMembers(uint256 _accountID) view returns(address[])
func (_Registry *RegistrySession) GetAccountMembers(_accountID *big.Int) ([]common.Address, error) {
	return _Registry.Contract.GetAccountMembers(&_Registry.CallOpts, _accountID)
}

// GetAccountMembers is a free data retrieval call binding the contract method 0xa2f6ce9b.
//
// Solidity: function getAccountMembers(uint256 _accountID) view returns(address[])
func (_Registry *RegistryCallerSession) GetAccountMembers(_accountID *big.Int) ([]common.Address, error) {
	return _Registry.Contract.GetAccountMembers(&_Registry.CallOpts, _accountID)
}

// GetLatestReleaseID is a free data retrieval call binding the contract method 0xd7730257.
//
// Solidity: function getLatestReleaseID(uint256 _projectID) view returns(uint256)
func (_Registry *RegistryCaller) GetLatestReleaseID(opts *bind.CallOpts, _projectID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getLatestReleaseID", _projectID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestReleaseID is a free data retrieval call binding the contract method 0xd7730257.
//
// Solidity: function getLatestReleaseID(uint256 _projectID) view returns(uint256)
func (_Registry *RegistrySession) GetLatestReleaseID(_projectID *big.Int) (*big.Int, error) {
	return _Registry.Contract.GetLatestReleaseID(&_Registry.CallOpts, _projectID)
}

// GetLatestReleaseID is a free data retrieval call binding the contract method 0xd7730257.
//
// Solidity: function getLatestReleaseID(uint256 _projectID) view returns(uint256)
func (_Registry *RegistryCallerSession) GetLatestReleaseID(_projectID *big.Int) (*big.Int, error) {
	return _Registry.Contract.GetLatestReleaseID(&_Registry.CallOpts, _projectID)
}

// GetPreviousReleaseID is a free data retrieval call binding the contract method 0x68704f02.
//
// Solidity: function getPreviousReleaseID(uint256 _releaseID) view returns(uint256)
func (_Registry *RegistryCaller) GetPreviousReleaseID(opts *bind.CallOpts, _releaseID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getPreviousReleaseID", _releaseID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPreviousReleaseID is a free data retrieval call binding the contract method 0x68704f02.
//
// Solidity: function getPreviousReleaseID(uint256 _releaseID) view returns(uint256)
func (_Registry *RegistrySession) GetPreviousReleaseID(_releaseID *big.Int) (*big.Int, error) {
	return _Registry.Contract.GetPreviousReleaseID(&_Registry.CallOpts, _releaseID)
}

// GetPreviousReleaseID is a free data retrieval call binding the contract method 0x68704f02.
//
// Solidity: function getPreviousReleaseID(uint256 _releaseID) view returns(uint256)
func (_Registry *RegistryCallerSession) GetPreviousReleaseID(_releaseID *big.Int) (*big.Int, error) {
	return _Registry.Contract.GetPreviousReleaseID(&_Registry.CallOpts, _releaseID)
}

// GetProjectAccountID is a free data retrieval call binding the contract method 0x416cab0c.
//
// Solidity: function getProjectAccountID(uint256 _projectID) view returns(uint256)
func (_Registry *RegistryCaller) GetProjectAccountID(opts *bind.CallOpts, _projectID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getProjectAccountID", _projectID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProjectAccountID is a free data retrieval call binding the contract method 0x416cab0c.
//
// Solidity: function getProjectAccountID(uint256 _projectID) view returns(uint256)
func (_Registry *RegistrySession) GetProjectAccountID(_projectID *big.Int) (*big.Int, error) {
	return _Registry.Contract.GetProjectAccountID(&_Registry.CallOpts, _projectID)
}

// GetProjectAccountID is a free data retrieval call binding the contract method 0x416cab0c.
//
// Solidity: function getProjectAccountID(uint256 _projectID) view returns(uint256)
func (_Registry *RegistryCallerSession) GetProjectAccountID(_projectID *big.Int) (*big.Int, error) {
	return _Registry.Contract.GetProjectAccountID(&_Registry.CallOpts, _projectID)
}

// GetProjectMembers is a free data retrieval call binding the contract method 0xd515d9d9.
//
// Solidity: function getProjectMembers(uint256 _projectID) view returns(address[])
func (_Registry *RegistryCaller) GetProjectMembers(opts *bind.CallOpts, _projectID *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getProjectMembers", _projectID)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetProjectMembers is a free data retrieval call binding the contract method 0xd515d9d9.
//
// Solidity: function getProjectMembers(uint256 _projectID) view returns(address[])
func (_Registry *RegistrySession) GetProjectMembers(_projectID *big.Int) ([]common.Address, error) {
	return _Registry.Contract.GetProjectMembers(&_Registry.CallOpts, _projectID)
}

// GetProjectMembers is a free data retrieval call binding the contract method 0xd515d9d9.
//
// Solidity: function getProjectMembers(uint256 _projectID) view returns(address[])
func (_Registry *RegistryCallerSession) GetProjectMembers(_projectID *big.Int) ([]common.Address, error) {
	return _Registry.Contract.GetProjectMembers(&_Registry.CallOpts, _projectID)
}

// GetReleaseProjectID is a free data retrieval call binding the contract method 0xaef28ac3.
//
// Solidity: function getReleaseProjectID(uint256 _releaseID) view returns(uint256)
func (_Registry *RegistryCaller) GetReleaseProjectID(opts *bind.CallOpts, _releaseID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getReleaseProjectID", _releaseID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReleaseProjectID is a free data retrieval call binding the contract method 0xaef28ac3.
//
// Solidity: function getReleaseProjectID(uint256 _releaseID) view returns(uint256)
func (_Registry *RegistrySession) GetReleaseProjectID(_releaseID *big.Int) (*big.Int, error) {
	return _Registry.Contract.GetReleaseProjectID(&_Registry.CallOpts, _releaseID)
}

// GetReleaseProjectID is a free data retrieval call binding the contract method 0xaef28ac3.
//
// Solidity: function getReleaseProjectID(uint256 _releaseID) view returns(uint256)
func (_Registry *RegistryCallerSession) GetReleaseProjectID(_releaseID *big.Int) (*big.Int, error) {
	return _Registry.Contract.GetReleaseProjectID(&_Registry.CallOpts, _releaseID)
}

// GetReleaseSigners is a free data retrieval call binding the contract method 0x83b1fc41.
//
// Solidity: function getReleaseSigners(uint256 _releaseID) view returns(address[])
func (_Registry *RegistryCaller) GetReleaseSigners(opts *bind.CallOpts, _releaseID *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getReleaseSigners", _releaseID)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetReleaseSigners is a free data retrieval call binding the contract method 0x83b1fc41.
//
// Solidity: function getReleaseSigners(uint256 _releaseID) view returns(address[])
func (_Registry *RegistrySession) GetReleaseSigners(_releaseID *big.Int) ([]common.Address, error) {
	return _Registry.Contract.GetReleaseSigners(&_Registry.CallOpts, _releaseID)
}

// GetReleaseSigners is a free data retrieval call binding the contract method 0x83b1fc41.
//
// Solidity: function getReleaseSigners(uint256 _releaseID) view returns(address[])
func (_Registry *RegistryCallerSession) GetReleaseSigners(_releaseID *big.Int) ([]common.Address, error) {
	return _Registry.Contract.GetReleaseSigners(&_Registry.CallOpts, _releaseID)
}

// IsAccountMember is a free data retrieval call binding the contract method 0x6ad0a5ee.
//
// Solidity: function isAccountMember(uint256 _accountID, address _member) view returns(bool)
func (_Registry *RegistryCaller) IsAccountMember(opts *bind.CallOpts, _accountID *big.Int, _member common.Address) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "isAccountMember", _accountID, _member)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAccountMember is a free data retrieval call binding the contract method 0x6ad0a5ee.
//
// Solidity: function isAccountMember(uint256 _accountID, address _member) view returns(bool)
func (_Registry *RegistrySession) IsAccountMember(_accountID *big.Int, _member common.Address) (bool, error) {
	return _Registry.Contract.IsAccountMember(&_Registry.CallOpts, _accountID, _member)
}

// IsAccountMember is a free data retrieval call binding the contract method 0x6ad0a5ee.
//
// Solidity: function isAccountMember(uint256 _accountID, address _member) view returns(bool)
func (_Registry *RegistryCallerSession) IsAccountMember(_accountID *big.Int, _member common.Address) (bool, error) {
	return _Registry.Contract.IsAccountMember(&_Registry.CallOpts, _accountID, _member)
}

// IsProjectMember is a free data retrieval call binding the contract method 0x600c3911.
//
// Solidity: function isProjectMember(uint256 _projectID, address _member) view returns(bool)
func (_Registry *RegistryCaller) IsProjectMember(opts *bind.CallOpts, _projectID *big.Int, _member common.Address) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "isProjectMember", _projectID, _member)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProjectMember is a free data retrieval call binding the contract method 0x600c3911.
//
// Solidity: function isProjectMember(uint256 _projectID, address _member) view returns(bool)
func (_Registry *RegistrySession) IsProjectMember(_projectID *big.Int, _member common.Address) (bool, error) {
	return _Registry.Contract.IsProjectMember(&_Registry.CallOpts, _projectID, _member)
}

// IsProjectMember is a free data retrieval call binding the contract method 0x600c3911.
//
// Solidity: function isProjectMember(uint256 _projectID, address _member) view returns(bool)
func (_Registry *RegistryCallerSession) IsProjectMember(_projectID *big.Int, _member common.Address) (bool, error) {
	return _Registry.Contract.IsProjectMember(&_Registry.CallOpts, _projectID, _member)
}

// IsReleaseSigner is a free data retrieval call binding the contract method 0x7c869665.
//
// Solidity: function isReleaseSigner(uint256 _releaseID, address _signer) view returns(bool)
func (_Registry *RegistryCaller) IsReleaseSigner(opts *bind.CallOpts, _releaseID *big.Int, _signer common.Address) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "isReleaseSigner", _releaseID, _signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReleaseSigner is a free data retrieval call binding the contract method 0x7c869665.
//
// Solidity: function isReleaseSigner(uint256 _releaseID, address _signer) view returns(bool)
func (_Registry *RegistrySession) IsReleaseSigner(_releaseID *big.Int, _signer common.Address) (bool, error) {
	return _Registry.Contract.IsReleaseSigner(&_Registry.CallOpts, _releaseID, _signer)
}

// IsReleaseSigner is a free data retrieval call binding the contract method 0x7c869665.
//
// Solidity: function isReleaseSigner(uint256 _releaseID, address _signer) view returns(bool)
func (_Registry *RegistryCallerSession) IsReleaseSigner(_releaseID *big.Int, _signer common.Address) (bool, error) {
	return _Registry.Contract.IsReleaseSigner(&_Registry.CallOpts, _releaseID, _signer)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Registry *RegistryCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Registry *RegistrySession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Registry.Contract.IsTrustedForwarder(&_Registry.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Registry *RegistryCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Registry.Contract.IsTrustedForwarder(&_Registry.CallOpts, forwarder)
}

// MetaByID is a free data retrieval call binding the contract method 0x43417c06.
//
// Solidity: function metaByID(uint256 ) view returns(string)
func (_Registry *RegistryCaller) MetaByID(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "metaByID", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MetaByID is a free data retrieval call binding the contract method 0x43417c06.
//
// Solidity: function metaByID(uint256 ) view returns(string)
func (_Registry *RegistrySession) MetaByID(arg0 *big.Int) (string, error) {
	return _Registry.Contract.MetaByID(&_Registry.CallOpts, arg0)
}

// MetaByID is a free data retrieval call binding the contract method 0x43417c06.
//
// Solidity: function metaByID(uint256 ) view returns(string)
func (_Registry *RegistryCallerSession) MetaByID(arg0 *big.Int) (string, error) {
	return _Registry.Contract.MetaByID(&_Registry.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistrySession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistryCallerSession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_Registry *RegistryCaller) TrustedForwarder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "trustedForwarder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_Registry *RegistrySession) TrustedForwarder() (common.Address, error) {
	return _Registry.Contract.TrustedForwarder(&_Registry.CallOpts)
}

// TrustedForwarder is a free data retrieval call binding the contract method 0x7da0a877.
//
// Solidity: function trustedForwarder() view returns(address)
func (_Registry *RegistryCallerSession) TrustedForwarder() (common.Address, error) {
	return _Registry.Contract.TrustedForwarder(&_Registry.CallOpts)
}

// VersionRecipient is a free data retrieval call binding the contract method 0x486ff0cd.
//
// Solidity: function versionRecipient() view returns(string)
func (_Registry *RegistryCaller) VersionRecipient(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "versionRecipient")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VersionRecipient is a free data retrieval call binding the contract method 0x486ff0cd.
//
// Solidity: function versionRecipient() view returns(string)
func (_Registry *RegistrySession) VersionRecipient() (string, error) {
	return _Registry.Contract.VersionRecipient(&_Registry.CallOpts)
}

// VersionRecipient is a free data retrieval call binding the contract method 0x486ff0cd.
//
// Solidity: function versionRecipient() view returns(string)
func (_Registry *RegistryCallerSession) VersionRecipient() (string, error) {
	return _Registry.Contract.VersionRecipient(&_Registry.CallOpts)
}

// AddAccountMember is a paid mutator transaction binding the contract method 0x176037f5.
//
// Solidity: function addAccountMember(uint256 _accountID, address _address) returns()
func (_Registry *RegistryTransactor) AddAccountMember(opts *bind.TransactOpts, _accountID *big.Int, _address common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "addAccountMember", _accountID, _address)
}

// AddAccountMember is a paid mutator transaction binding the contract method 0x176037f5.
//
// Solidity: function addAccountMember(uint256 _accountID, address _address) returns()
func (_Registry *RegistrySession) AddAccountMember(_accountID *big.Int, _address common.Address) (*types.Transaction, error) {
	return _Registry.Contract.AddAccountMember(&_Registry.TransactOpts, _accountID, _address)
}

// AddAccountMember is a paid mutator transaction binding the contract method 0x176037f5.
//
// Solidity: function addAccountMember(uint256 _accountID, address _address) returns()
func (_Registry *RegistryTransactorSession) AddAccountMember(_accountID *big.Int, _address common.Address) (*types.Transaction, error) {
	return _Registry.Contract.AddAccountMember(&_Registry.TransactOpts, _accountID, _address)
}

// AddProjectMember is a paid mutator transaction binding the contract method 0xc25cf2d5.
//
// Solidity: function addProjectMember(uint256 _projectID, address _address) returns()
func (_Registry *RegistryTransactor) AddProjectMember(opts *bind.TransactOpts, _projectID *big.Int, _address common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "addProjectMember", _projectID, _address)
}

// AddProjectMember is a paid mutator transaction binding the contract method 0xc25cf2d5.
//
// Solidity: function addProjectMember(uint256 _projectID, address _address) returns()
func (_Registry *RegistrySession) AddProjectMember(_projectID *big.Int, _address common.Address) (*types.Transaction, error) {
	return _Registry.Contract.AddProjectMember(&_Registry.TransactOpts, _projectID, _address)
}

// AddProjectMember is a paid mutator transaction binding the contract method 0xc25cf2d5.
//
// Solidity: function addProjectMember(uint256 _projectID, address _address) returns()
func (_Registry *RegistryTransactorSession) AddProjectMember(_projectID *big.Int, _address common.Address) (*types.Transaction, error) {
	return _Registry.Contract.AddProjectMember(&_Registry.TransactOpts, _projectID, _address)
}

// ApproveRelease is a paid mutator transaction binding the contract method 0x1bd2cea1.
//
// Solidity: function approveRelease(uint256 _releaseID) returns()
func (_Registry *RegistryTransactor) ApproveRelease(opts *bind.TransactOpts, _releaseID *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "approveRelease", _releaseID)
}

// ApproveRelease is a paid mutator transaction binding the contract method 0x1bd2cea1.
//
// Solidity: function approveRelease(uint256 _releaseID) returns()
func (_Registry *RegistrySession) ApproveRelease(_releaseID *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.ApproveRelease(&_Registry.TransactOpts, _releaseID)
}

// ApproveRelease is a paid mutator transaction binding the contract method 0x1bd2cea1.
//
// Solidity: function approveRelease(uint256 _releaseID) returns()
func (_Registry *RegistryTransactorSession) ApproveRelease(_releaseID *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.ApproveRelease(&_Registry.TransactOpts, _releaseID)
}

// CreateAccount is a paid mutator transaction binding the contract method 0xbb11d22a.
//
// Solidity: function createAccount(string _name, string _metaURI, address[] _members) payable returns()
func (_Registry *RegistryTransactor) CreateAccount(opts *bind.TransactOpts, _name string, _metaURI string, _members []common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "createAccount", _name, _metaURI, _members)
}

// CreateAccount is a paid mutator transaction binding the contract method 0xbb11d22a.
//
// Solidity: function createAccount(string _name, string _metaURI, address[] _members) payable returns()
func (_Registry *RegistrySession) CreateAccount(_name string, _metaURI string, _members []common.Address) (*types.Transaction, error) {
	return _Registry.Contract.CreateAccount(&_Registry.TransactOpts, _name, _metaURI, _members)
}

// CreateAccount is a paid mutator transaction binding the contract method 0xbb11d22a.
//
// Solidity: function createAccount(string _name, string _metaURI, address[] _members) payable returns()
func (_Registry *RegistryTransactorSession) CreateAccount(_name string, _metaURI string, _members []common.Address) (*types.Transaction, error) {
	return _Registry.Contract.CreateAccount(&_Registry.TransactOpts, _name, _metaURI, _members)
}

// CreateProject is a paid mutator transaction binding the contract method 0x250b72de.
//
// Solidity: function createProject(uint256 _accountID, string _name, string _metaURI, address[] _members) returns()
func (_Registry *RegistryTransactor) CreateProject(opts *bind.TransactOpts, _accountID *big.Int, _name string, _metaURI string, _members []common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "createProject", _accountID, _name, _metaURI, _members)
}

// CreateProject is a paid mutator transaction binding the contract method 0x250b72de.
//
// Solidity: function createProject(uint256 _accountID, string _name, string _metaURI, address[] _members) returns()
func (_Registry *RegistrySession) CreateProject(_accountID *big.Int, _name string, _metaURI string, _members []common.Address) (*types.Transaction, error) {
	return _Registry.Contract.CreateProject(&_Registry.TransactOpts, _accountID, _name, _metaURI, _members)
}

// CreateProject is a paid mutator transaction binding the contract method 0x250b72de.
//
// Solidity: function createProject(uint256 _accountID, string _name, string _metaURI, address[] _members) returns()
func (_Registry *RegistryTransactorSession) CreateProject(_accountID *big.Int, _name string, _metaURI string, _members []common.Address) (*types.Transaction, error) {
	return _Registry.Contract.CreateProject(&_Registry.TransactOpts, _accountID, _name, _metaURI, _members)
}

// CreateRelease is a paid mutator transaction binding the contract method 0x28a10ec6.
//
// Solidity: function createRelease(uint256 _projectID, string _name, string _metaURI) returns()
func (_Registry *RegistryTransactor) CreateRelease(opts *bind.TransactOpts, _projectID *big.Int, _name string, _metaURI string) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "createRelease", _projectID, _name, _metaURI)
}

// CreateRelease is a paid mutator transaction binding the contract method 0x28a10ec6.
//
// Solidity: function createRelease(uint256 _projectID, string _name, string _metaURI) returns()
func (_Registry *RegistrySession) CreateRelease(_projectID *big.Int, _name string, _metaURI string) (*types.Transaction, error) {
	return _Registry.Contract.CreateRelease(&_Registry.TransactOpts, _projectID, _name, _metaURI)
}

// CreateRelease is a paid mutator transaction binding the contract method 0x28a10ec6.
//
// Solidity: function createRelease(uint256 _projectID, string _name, string _metaURI) returns()
func (_Registry *RegistryTransactorSession) CreateRelease(_projectID *big.Int, _name string, _metaURI string) (*types.Transaction, error) {
	return _Registry.Contract.CreateRelease(&_Registry.TransactOpts, _projectID, _name, _metaURI)
}

// RemoveAccountMember is a paid mutator transaction binding the contract method 0xef83d3ab.
//
// Solidity: function removeAccountMember(uint256 _accountID, address _address) returns()
func (_Registry *RegistryTransactor) RemoveAccountMember(opts *bind.TransactOpts, _accountID *big.Int, _address common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "removeAccountMember", _accountID, _address)
}

// RemoveAccountMember is a paid mutator transaction binding the contract method 0xef83d3ab.
//
// Solidity: function removeAccountMember(uint256 _accountID, address _address) returns()
func (_Registry *RegistrySession) RemoveAccountMember(_accountID *big.Int, _address common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RemoveAccountMember(&_Registry.TransactOpts, _accountID, _address)
}

// RemoveAccountMember is a paid mutator transaction binding the contract method 0xef83d3ab.
//
// Solidity: function removeAccountMember(uint256 _accountID, address _address) returns()
func (_Registry *RegistryTransactorSession) RemoveAccountMember(_accountID *big.Int, _address common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RemoveAccountMember(&_Registry.TransactOpts, _accountID, _address)
}

// RemoveProjectMember is a paid mutator transaction binding the contract method 0xec95ecb6.
//
// Solidity: function removeProjectMember(uint256 _projectID, address _address) returns()
func (_Registry *RegistryTransactor) RemoveProjectMember(opts *bind.TransactOpts, _projectID *big.Int, _address common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "removeProjectMember", _projectID, _address)
}

// RemoveProjectMember is a paid mutator transaction binding the contract method 0xec95ecb6.
//
// Solidity: function removeProjectMember(uint256 _projectID, address _address) returns()
func (_Registry *RegistrySession) RemoveProjectMember(_projectID *big.Int, _address common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RemoveProjectMember(&_Registry.TransactOpts, _projectID, _address)
}

// RemoveProjectMember is a paid mutator transaction binding the contract method 0xec95ecb6.
//
// Solidity: function removeProjectMember(uint256 _projectID, address _address) returns()
func (_Registry *RegistryTransactorSession) RemoveProjectMember(_projectID *big.Int, _address common.Address) (*types.Transaction, error) {
	return _Registry.Contract.RemoveProjectMember(&_Registry.TransactOpts, _projectID, _address)
}

// RevokeRelease is a paid mutator transaction binding the contract method 0x806553f3.
//
// Solidity: function revokeRelease(uint256 _releaseID) returns()
func (_Registry *RegistryTransactor) RevokeRelease(opts *bind.TransactOpts, _releaseID *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "revokeRelease", _releaseID)
}

// RevokeRelease is a paid mutator transaction binding the contract method 0x806553f3.
//
// Solidity: function revokeRelease(uint256 _releaseID) returns()
func (_Registry *RegistrySession) RevokeRelease(_releaseID *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.RevokeRelease(&_Registry.TransactOpts, _releaseID)
}

// RevokeRelease is a paid mutator transaction binding the contract method 0x806553f3.
//
// Solidity: function revokeRelease(uint256 _releaseID) returns()
func (_Registry *RegistryTransactorSession) RevokeRelease(_releaseID *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.RevokeRelease(&_Registry.TransactOpts, _releaseID)
}

// SetAccountMetaURI is a paid mutator transaction binding the contract method 0x9f4cfef9.
//
// Solidity: function setAccountMetaURI(uint256 _accountID, string _metaURI) returns()
func (_Registry *RegistryTransactor) SetAccountMetaURI(opts *bind.TransactOpts, _accountID *big.Int, _metaURI string) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setAccountMetaURI", _accountID, _metaURI)
}

// SetAccountMetaURI is a paid mutator transaction binding the contract method 0x9f4cfef9.
//
// Solidity: function setAccountMetaURI(uint256 _accountID, string _metaURI) returns()
func (_Registry *RegistrySession) SetAccountMetaURI(_accountID *big.Int, _metaURI string) (*types.Transaction, error) {
	return _Registry.Contract.SetAccountMetaURI(&_Registry.TransactOpts, _accountID, _metaURI)
}

// SetAccountMetaURI is a paid mutator transaction binding the contract method 0x9f4cfef9.
//
// Solidity: function setAccountMetaURI(uint256 _accountID, string _metaURI) returns()
func (_Registry *RegistryTransactorSession) SetAccountMetaURI(_accountID *big.Int, _metaURI string) (*types.Transaction, error) {
	return _Registry.Contract.SetAccountMetaURI(&_Registry.TransactOpts, _accountID, _metaURI)
}

// SetClaimFee is a paid mutator transaction binding the contract method 0x2e75ab50.
//
// Solidity: function setClaimFee(uint256 _claimFee) returns()
func (_Registry *RegistryTransactor) SetClaimFee(opts *bind.TransactOpts, _claimFee *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setClaimFee", _claimFee)
}

// SetClaimFee is a paid mutator transaction binding the contract method 0x2e75ab50.
//
// Solidity: function setClaimFee(uint256 _claimFee) returns()
func (_Registry *RegistrySession) SetClaimFee(_claimFee *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.SetClaimFee(&_Registry.TransactOpts, _claimFee)
}

// SetClaimFee is a paid mutator transaction binding the contract method 0x2e75ab50.
//
// Solidity: function setClaimFee(uint256 _claimFee) returns()
func (_Registry *RegistryTransactorSession) SetClaimFee(_claimFee *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.SetClaimFee(&_Registry.TransactOpts, _claimFee)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_Registry *RegistryTransactor) SetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setOwner", _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_Registry *RegistrySession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetOwner(&_Registry.TransactOpts, _owner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _owner) returns()
func (_Registry *RegistryTransactorSession) SetOwner(_owner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetOwner(&_Registry.TransactOpts, _owner)
}

// SetProjectMetaURI is a paid mutator transaction binding the contract method 0xdb5954f4.
//
// Solidity: function setProjectMetaURI(uint256 _projectID, string _metaURI) returns()
func (_Registry *RegistryTransactor) SetProjectMetaURI(opts *bind.TransactOpts, _projectID *big.Int, _metaURI string) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setProjectMetaURI", _projectID, _metaURI)
}

// SetProjectMetaURI is a paid mutator transaction binding the contract method 0xdb5954f4.
//
// Solidity: function setProjectMetaURI(uint256 _projectID, string _metaURI) returns()
func (_Registry *RegistrySession) SetProjectMetaURI(_projectID *big.Int, _metaURI string) (*types.Transaction, error) {
	return _Registry.Contract.SetProjectMetaURI(&_Registry.TransactOpts, _projectID, _metaURI)
}

// SetProjectMetaURI is a paid mutator transaction binding the contract method 0xdb5954f4.
//
// Solidity: function setProjectMetaURI(uint256 _projectID, string _metaURI) returns()
func (_Registry *RegistryTransactorSession) SetProjectMetaURI(_projectID *big.Int, _metaURI string) (*types.Transaction, error) {
	return _Registry.Contract.SetProjectMetaURI(&_Registry.TransactOpts, _projectID, _metaURI)
}

// SetTrustedForwarder is a paid mutator transaction binding the contract method 0xda742228.
//
// Solidity: function setTrustedForwarder(address _forwarder) returns()
func (_Registry *RegistryTransactor) SetTrustedForwarder(opts *bind.TransactOpts, _forwarder common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setTrustedForwarder", _forwarder)
}

// SetTrustedForwarder is a paid mutator transaction binding the contract method 0xda742228.
//
// Solidity: function setTrustedForwarder(address _forwarder) returns()
func (_Registry *RegistrySession) SetTrustedForwarder(_forwarder common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetTrustedForwarder(&_Registry.TransactOpts, _forwarder)
}

// SetTrustedForwarder is a paid mutator transaction binding the contract method 0xda742228.
//
// Solidity: function setTrustedForwarder(address _forwarder) returns()
func (_Registry *RegistryTransactorSession) SetTrustedForwarder(_forwarder common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetTrustedForwarder(&_Registry.TransactOpts, _forwarder)
}

// RegistryAccountCreatedIterator is returned from FilterAccountCreated and is used to iterate over the raw logs and unpacked data for AccountCreated events raised by the Registry contract.
type RegistryAccountCreatedIterator struct {
	Event *RegistryAccountCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryAccountCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryAccountCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryAccountCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryAccountCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryAccountCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryAccountCreated represents a AccountCreated event raised by the Registry contract.
type RegistryAccountCreated struct {
	AccountID *big.Int
	Name      string
	MetaURI   string
	Sender    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAccountCreated is a free log retrieval operation binding the contract event 0x7bab27b7daa748e5bd855a000a2445c6e1d22fc5b2a6c4df5c96f63508afc7b9.
//
// Solidity: event AccountCreated(uint256 _accountID, string _name, string _metaURI, address _sender)
func (_Registry *RegistryFilterer) FilterAccountCreated(opts *bind.FilterOpts) (*RegistryAccountCreatedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "AccountCreated")
	if err != nil {
		return nil, err
	}
	return &RegistryAccountCreatedIterator{contract: _Registry.contract, event: "AccountCreated", logs: logs, sub: sub}, nil
}

// WatchAccountCreated is a free log subscription operation binding the contract event 0x7bab27b7daa748e5bd855a000a2445c6e1d22fc5b2a6c4df5c96f63508afc7b9.
//
// Solidity: event AccountCreated(uint256 _accountID, string _name, string _metaURI, address _sender)
func (_Registry *RegistryFilterer) WatchAccountCreated(opts *bind.WatchOpts, sink chan<- *RegistryAccountCreated) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "AccountCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryAccountCreated)
				if err := _Registry.contract.UnpackLog(event, "AccountCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccountCreated is a log parse operation binding the contract event 0x7bab27b7daa748e5bd855a000a2445c6e1d22fc5b2a6c4df5c96f63508afc7b9.
//
// Solidity: event AccountCreated(uint256 _accountID, string _name, string _metaURI, address _sender)
func (_Registry *RegistryFilterer) ParseAccountCreated(log types.Log) (*RegistryAccountCreated, error) {
	event := new(RegistryAccountCreated)
	if err := _Registry.contract.UnpackLog(event, "AccountCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryAccountMemberAddedIterator is returned from FilterAccountMemberAdded and is used to iterate over the raw logs and unpacked data for AccountMemberAdded events raised by the Registry contract.
type RegistryAccountMemberAddedIterator struct {
	Event *RegistryAccountMemberAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryAccountMemberAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryAccountMemberAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryAccountMemberAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryAccountMemberAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryAccountMemberAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryAccountMemberAdded represents a AccountMemberAdded event raised by the Registry contract.
type RegistryAccountMemberAdded struct {
	AccountID *big.Int
	Member    common.Address
	Sender    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAccountMemberAdded is a free log retrieval operation binding the contract event 0xda06031a12c2b6322a6a2a29b8e7ec92a209562f0995acbf9e5fe8cce89617b3.
//
// Solidity: event AccountMemberAdded(uint256 _accountID, address _member, address _sender)
func (_Registry *RegistryFilterer) FilterAccountMemberAdded(opts *bind.FilterOpts) (*RegistryAccountMemberAddedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "AccountMemberAdded")
	if err != nil {
		return nil, err
	}
	return &RegistryAccountMemberAddedIterator{contract: _Registry.contract, event: "AccountMemberAdded", logs: logs, sub: sub}, nil
}

// WatchAccountMemberAdded is a free log subscription operation binding the contract event 0xda06031a12c2b6322a6a2a29b8e7ec92a209562f0995acbf9e5fe8cce89617b3.
//
// Solidity: event AccountMemberAdded(uint256 _accountID, address _member, address _sender)
func (_Registry *RegistryFilterer) WatchAccountMemberAdded(opts *bind.WatchOpts, sink chan<- *RegistryAccountMemberAdded) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "AccountMemberAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryAccountMemberAdded)
				if err := _Registry.contract.UnpackLog(event, "AccountMemberAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccountMemberAdded is a log parse operation binding the contract event 0xda06031a12c2b6322a6a2a29b8e7ec92a209562f0995acbf9e5fe8cce89617b3.
//
// Solidity: event AccountMemberAdded(uint256 _accountID, address _member, address _sender)
func (_Registry *RegistryFilterer) ParseAccountMemberAdded(log types.Log) (*RegistryAccountMemberAdded, error) {
	event := new(RegistryAccountMemberAdded)
	if err := _Registry.contract.UnpackLog(event, "AccountMemberAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryAccountMemberRemovedIterator is returned from FilterAccountMemberRemoved and is used to iterate over the raw logs and unpacked data for AccountMemberRemoved events raised by the Registry contract.
type RegistryAccountMemberRemovedIterator struct {
	Event *RegistryAccountMemberRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryAccountMemberRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryAccountMemberRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryAccountMemberRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryAccountMemberRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryAccountMemberRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryAccountMemberRemoved represents a AccountMemberRemoved event raised by the Registry contract.
type RegistryAccountMemberRemoved struct {
	AccountID *big.Int
	Member    common.Address
	Sender    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAccountMemberRemoved is a free log retrieval operation binding the contract event 0xebdaf29c191f35cacd52dad5ad5e63fe0a2806c921f5f362dfcc9ded63874550.
//
// Solidity: event AccountMemberRemoved(uint256 _accountID, address _member, address _sender)
func (_Registry *RegistryFilterer) FilterAccountMemberRemoved(opts *bind.FilterOpts) (*RegistryAccountMemberRemovedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "AccountMemberRemoved")
	if err != nil {
		return nil, err
	}
	return &RegistryAccountMemberRemovedIterator{contract: _Registry.contract, event: "AccountMemberRemoved", logs: logs, sub: sub}, nil
}

// WatchAccountMemberRemoved is a free log subscription operation binding the contract event 0xebdaf29c191f35cacd52dad5ad5e63fe0a2806c921f5f362dfcc9ded63874550.
//
// Solidity: event AccountMemberRemoved(uint256 _accountID, address _member, address _sender)
func (_Registry *RegistryFilterer) WatchAccountMemberRemoved(opts *bind.WatchOpts, sink chan<- *RegistryAccountMemberRemoved) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "AccountMemberRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryAccountMemberRemoved)
				if err := _Registry.contract.UnpackLog(event, "AccountMemberRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccountMemberRemoved is a log parse operation binding the contract event 0xebdaf29c191f35cacd52dad5ad5e63fe0a2806c921f5f362dfcc9ded63874550.
//
// Solidity: event AccountMemberRemoved(uint256 _accountID, address _member, address _sender)
func (_Registry *RegistryFilterer) ParseAccountMemberRemoved(log types.Log) (*RegistryAccountMemberRemoved, error) {
	event := new(RegistryAccountMemberRemoved)
	if err := _Registry.contract.UnpackLog(event, "AccountMemberRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryAccountUpdatedIterator is returned from FilterAccountUpdated and is used to iterate over the raw logs and unpacked data for AccountUpdated events raised by the Registry contract.
type RegistryAccountUpdatedIterator struct {
	Event *RegistryAccountUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryAccountUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryAccountUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryAccountUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryAccountUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryAccountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryAccountUpdated represents a AccountUpdated event raised by the Registry contract.
type RegistryAccountUpdated struct {
	AccountID *big.Int
	MetaURI   string
	Sender    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAccountUpdated is a free log retrieval operation binding the contract event 0xdc2b36fe3e9f9129790c8dc1fefe27a260e2ab12bbd7d2503bf3c701b7a6b107.
//
// Solidity: event AccountUpdated(uint256 _accountID, string _metaURI, address _sender)
func (_Registry *RegistryFilterer) FilterAccountUpdated(opts *bind.FilterOpts) (*RegistryAccountUpdatedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "AccountUpdated")
	if err != nil {
		return nil, err
	}
	return &RegistryAccountUpdatedIterator{contract: _Registry.contract, event: "AccountUpdated", logs: logs, sub: sub}, nil
}

// WatchAccountUpdated is a free log subscription operation binding the contract event 0xdc2b36fe3e9f9129790c8dc1fefe27a260e2ab12bbd7d2503bf3c701b7a6b107.
//
// Solidity: event AccountUpdated(uint256 _accountID, string _metaURI, address _sender)
func (_Registry *RegistryFilterer) WatchAccountUpdated(opts *bind.WatchOpts, sink chan<- *RegistryAccountUpdated) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "AccountUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryAccountUpdated)
				if err := _Registry.contract.UnpackLog(event, "AccountUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccountUpdated is a log parse operation binding the contract event 0xdc2b36fe3e9f9129790c8dc1fefe27a260e2ab12bbd7d2503bf3c701b7a6b107.
//
// Solidity: event AccountUpdated(uint256 _accountID, string _metaURI, address _sender)
func (_Registry *RegistryFilterer) ParseAccountUpdated(log types.Log) (*RegistryAccountUpdated, error) {
	event := new(RegistryAccountUpdated)
	if err := _Registry.contract.UnpackLog(event, "AccountUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryProjectCreatedIterator is returned from FilterProjectCreated and is used to iterate over the raw logs and unpacked data for ProjectCreated events raised by the Registry contract.
type RegistryProjectCreatedIterator struct {
	Event *RegistryProjectCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryProjectCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryProjectCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryProjectCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryProjectCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryProjectCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryProjectCreated represents a ProjectCreated event raised by the Registry contract.
type RegistryProjectCreated struct {
	AccountID *big.Int
	ProjectID *big.Int
	Name      string
	MetaURI   string
	Sender    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProjectCreated is a free log retrieval operation binding the contract event 0xdce0cbd02a04a917caf7c163eeb153eee961368abd158a17a8d17298afc2a62f.
//
// Solidity: event ProjectCreated(uint256 _accountID, uint256 _projectID, string _name, string _metaURI, address _sender)
func (_Registry *RegistryFilterer) FilterProjectCreated(opts *bind.FilterOpts) (*RegistryProjectCreatedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ProjectCreated")
	if err != nil {
		return nil, err
	}
	return &RegistryProjectCreatedIterator{contract: _Registry.contract, event: "ProjectCreated", logs: logs, sub: sub}, nil
}

// WatchProjectCreated is a free log subscription operation binding the contract event 0xdce0cbd02a04a917caf7c163eeb153eee961368abd158a17a8d17298afc2a62f.
//
// Solidity: event ProjectCreated(uint256 _accountID, uint256 _projectID, string _name, string _metaURI, address _sender)
func (_Registry *RegistryFilterer) WatchProjectCreated(opts *bind.WatchOpts, sink chan<- *RegistryProjectCreated) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ProjectCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryProjectCreated)
				if err := _Registry.contract.UnpackLog(event, "ProjectCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProjectCreated is a log parse operation binding the contract event 0xdce0cbd02a04a917caf7c163eeb153eee961368abd158a17a8d17298afc2a62f.
//
// Solidity: event ProjectCreated(uint256 _accountID, uint256 _projectID, string _name, string _metaURI, address _sender)
func (_Registry *RegistryFilterer) ParseProjectCreated(log types.Log) (*RegistryProjectCreated, error) {
	event := new(RegistryProjectCreated)
	if err := _Registry.contract.UnpackLog(event, "ProjectCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryProjectMemberAddedIterator is returned from FilterProjectMemberAdded and is used to iterate over the raw logs and unpacked data for ProjectMemberAdded events raised by the Registry contract.
type RegistryProjectMemberAddedIterator struct {
	Event *RegistryProjectMemberAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryProjectMemberAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryProjectMemberAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryProjectMemberAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryProjectMemberAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryProjectMemberAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryProjectMemberAdded represents a ProjectMemberAdded event raised by the Registry contract.
type RegistryProjectMemberAdded struct {
	ProjectID *big.Int
	Member    common.Address
	Sender    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProjectMemberAdded is a free log retrieval operation binding the contract event 0xa641673cc2c2925ec554d696d7ecb92ad62fd37a878bb62194e01d15af4ecbfe.
//
// Solidity: event ProjectMemberAdded(uint256 _projectID, address _member, address _sender)
func (_Registry *RegistryFilterer) FilterProjectMemberAdded(opts *bind.FilterOpts) (*RegistryProjectMemberAddedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ProjectMemberAdded")
	if err != nil {
		return nil, err
	}
	return &RegistryProjectMemberAddedIterator{contract: _Registry.contract, event: "ProjectMemberAdded", logs: logs, sub: sub}, nil
}

// WatchProjectMemberAdded is a free log subscription operation binding the contract event 0xa641673cc2c2925ec554d696d7ecb92ad62fd37a878bb62194e01d15af4ecbfe.
//
// Solidity: event ProjectMemberAdded(uint256 _projectID, address _member, address _sender)
func (_Registry *RegistryFilterer) WatchProjectMemberAdded(opts *bind.WatchOpts, sink chan<- *RegistryProjectMemberAdded) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ProjectMemberAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryProjectMemberAdded)
				if err := _Registry.contract.UnpackLog(event, "ProjectMemberAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProjectMemberAdded is a log parse operation binding the contract event 0xa641673cc2c2925ec554d696d7ecb92ad62fd37a878bb62194e01d15af4ecbfe.
//
// Solidity: event ProjectMemberAdded(uint256 _projectID, address _member, address _sender)
func (_Registry *RegistryFilterer) ParseProjectMemberAdded(log types.Log) (*RegistryProjectMemberAdded, error) {
	event := new(RegistryProjectMemberAdded)
	if err := _Registry.contract.UnpackLog(event, "ProjectMemberAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryProjectMemberRemovedIterator is returned from FilterProjectMemberRemoved and is used to iterate over the raw logs and unpacked data for ProjectMemberRemoved events raised by the Registry contract.
type RegistryProjectMemberRemovedIterator struct {
	Event *RegistryProjectMemberRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryProjectMemberRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryProjectMemberRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryProjectMemberRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryProjectMemberRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryProjectMemberRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryProjectMemberRemoved represents a ProjectMemberRemoved event raised by the Registry contract.
type RegistryProjectMemberRemoved struct {
	ProjectID *big.Int
	Member    common.Address
	Sender    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProjectMemberRemoved is a free log retrieval operation binding the contract event 0x3c2f57383ed351452c76a1869a6abf17e5521b3de241aa8d31d1769d4e193c24.
//
// Solidity: event ProjectMemberRemoved(uint256 _projectID, address _member, address _sender)
func (_Registry *RegistryFilterer) FilterProjectMemberRemoved(opts *bind.FilterOpts) (*RegistryProjectMemberRemovedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ProjectMemberRemoved")
	if err != nil {
		return nil, err
	}
	return &RegistryProjectMemberRemovedIterator{contract: _Registry.contract, event: "ProjectMemberRemoved", logs: logs, sub: sub}, nil
}

// WatchProjectMemberRemoved is a free log subscription operation binding the contract event 0x3c2f57383ed351452c76a1869a6abf17e5521b3de241aa8d31d1769d4e193c24.
//
// Solidity: event ProjectMemberRemoved(uint256 _projectID, address _member, address _sender)
func (_Registry *RegistryFilterer) WatchProjectMemberRemoved(opts *bind.WatchOpts, sink chan<- *RegistryProjectMemberRemoved) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ProjectMemberRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryProjectMemberRemoved)
				if err := _Registry.contract.UnpackLog(event, "ProjectMemberRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProjectMemberRemoved is a log parse operation binding the contract event 0x3c2f57383ed351452c76a1869a6abf17e5521b3de241aa8d31d1769d4e193c24.
//
// Solidity: event ProjectMemberRemoved(uint256 _projectID, address _member, address _sender)
func (_Registry *RegistryFilterer) ParseProjectMemberRemoved(log types.Log) (*RegistryProjectMemberRemoved, error) {
	event := new(RegistryProjectMemberRemoved)
	if err := _Registry.contract.UnpackLog(event, "ProjectMemberRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryProjectUpdatedIterator is returned from FilterProjectUpdated and is used to iterate over the raw logs and unpacked data for ProjectUpdated events raised by the Registry contract.
type RegistryProjectUpdatedIterator struct {
	Event *RegistryProjectUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryProjectUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryProjectUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryProjectUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryProjectUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryProjectUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryProjectUpdated represents a ProjectUpdated event raised by the Registry contract.
type RegistryProjectUpdated struct {
	ProjectID *big.Int
	MetaURI   string
	Sender    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProjectUpdated is a free log retrieval operation binding the contract event 0xb1c971f708cf6cac0bf1b55c79db4b7e06ed48aa4e24618f1e40cbebecc5b063.
//
// Solidity: event ProjectUpdated(uint256 _projectID, string _metaURI, address _sender)
func (_Registry *RegistryFilterer) FilterProjectUpdated(opts *bind.FilterOpts) (*RegistryProjectUpdatedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ProjectUpdated")
	if err != nil {
		return nil, err
	}
	return &RegistryProjectUpdatedIterator{contract: _Registry.contract, event: "ProjectUpdated", logs: logs, sub: sub}, nil
}

// WatchProjectUpdated is a free log subscription operation binding the contract event 0xb1c971f708cf6cac0bf1b55c79db4b7e06ed48aa4e24618f1e40cbebecc5b063.
//
// Solidity: event ProjectUpdated(uint256 _projectID, string _metaURI, address _sender)
func (_Registry *RegistryFilterer) WatchProjectUpdated(opts *bind.WatchOpts, sink chan<- *RegistryProjectUpdated) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ProjectUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryProjectUpdated)
				if err := _Registry.contract.UnpackLog(event, "ProjectUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProjectUpdated is a log parse operation binding the contract event 0xb1c971f708cf6cac0bf1b55c79db4b7e06ed48aa4e24618f1e40cbebecc5b063.
//
// Solidity: event ProjectUpdated(uint256 _projectID, string _metaURI, address _sender)
func (_Registry *RegistryFilterer) ParseProjectUpdated(log types.Log) (*RegistryProjectUpdated, error) {
	event := new(RegistryProjectUpdated)
	if err := _Registry.contract.UnpackLog(event, "ProjectUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryReleaseApprovedIterator is returned from FilterReleaseApproved and is used to iterate over the raw logs and unpacked data for ReleaseApproved events raised by the Registry contract.
type RegistryReleaseApprovedIterator struct {
	Event *RegistryReleaseApproved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryReleaseApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryReleaseApproved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryReleaseApproved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryReleaseApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryReleaseApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryReleaseApproved represents a ReleaseApproved event raised by the Registry contract.
type RegistryReleaseApproved struct {
	ReleaseID *big.Int
	Sender    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReleaseApproved is a free log retrieval operation binding the contract event 0x492f404869834bbfb3dfd354d8ebd27ca988aa59eda4724d2e604400e1635df0.
//
// Solidity: event ReleaseApproved(uint256 _releaseID, address _sender)
func (_Registry *RegistryFilterer) FilterReleaseApproved(opts *bind.FilterOpts) (*RegistryReleaseApprovedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ReleaseApproved")
	if err != nil {
		return nil, err
	}
	return &RegistryReleaseApprovedIterator{contract: _Registry.contract, event: "ReleaseApproved", logs: logs, sub: sub}, nil
}

// WatchReleaseApproved is a free log subscription operation binding the contract event 0x492f404869834bbfb3dfd354d8ebd27ca988aa59eda4724d2e604400e1635df0.
//
// Solidity: event ReleaseApproved(uint256 _releaseID, address _sender)
func (_Registry *RegistryFilterer) WatchReleaseApproved(opts *bind.WatchOpts, sink chan<- *RegistryReleaseApproved) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ReleaseApproved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryReleaseApproved)
				if err := _Registry.contract.UnpackLog(event, "ReleaseApproved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReleaseApproved is a log parse operation binding the contract event 0x492f404869834bbfb3dfd354d8ebd27ca988aa59eda4724d2e604400e1635df0.
//
// Solidity: event ReleaseApproved(uint256 _releaseID, address _sender)
func (_Registry *RegistryFilterer) ParseReleaseApproved(log types.Log) (*RegistryReleaseApproved, error) {
	event := new(RegistryReleaseApproved)
	if err := _Registry.contract.UnpackLog(event, "ReleaseApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryReleaseCreatedIterator is returned from FilterReleaseCreated and is used to iterate over the raw logs and unpacked data for ReleaseCreated events raised by the Registry contract.
type RegistryReleaseCreatedIterator struct {
	Event *RegistryReleaseCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryReleaseCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryReleaseCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryReleaseCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryReleaseCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryReleaseCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryReleaseCreated represents a ReleaseCreated event raised by the Registry contract.
type RegistryReleaseCreated struct {
	ProjectID *big.Int
	ReleaseID *big.Int
	Name      string
	MetaURI   string
	Sender    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReleaseCreated is a free log retrieval operation binding the contract event 0x28f16f2c48f7434e8ee65a00bca94583249ae3dc79ebd4c8fac3ce1cc41227a9.
//
// Solidity: event ReleaseCreated(uint256 _projectID, uint256 _releaseID, string _name, string _metaURI, address _sender)
func (_Registry *RegistryFilterer) FilterReleaseCreated(opts *bind.FilterOpts) (*RegistryReleaseCreatedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ReleaseCreated")
	if err != nil {
		return nil, err
	}
	return &RegistryReleaseCreatedIterator{contract: _Registry.contract, event: "ReleaseCreated", logs: logs, sub: sub}, nil
}

// WatchReleaseCreated is a free log subscription operation binding the contract event 0x28f16f2c48f7434e8ee65a00bca94583249ae3dc79ebd4c8fac3ce1cc41227a9.
//
// Solidity: event ReleaseCreated(uint256 _projectID, uint256 _releaseID, string _name, string _metaURI, address _sender)
func (_Registry *RegistryFilterer) WatchReleaseCreated(opts *bind.WatchOpts, sink chan<- *RegistryReleaseCreated) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ReleaseCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryReleaseCreated)
				if err := _Registry.contract.UnpackLog(event, "ReleaseCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReleaseCreated is a log parse operation binding the contract event 0x28f16f2c48f7434e8ee65a00bca94583249ae3dc79ebd4c8fac3ce1cc41227a9.
//
// Solidity: event ReleaseCreated(uint256 _projectID, uint256 _releaseID, string _name, string _metaURI, address _sender)
func (_Registry *RegistryFilterer) ParseReleaseCreated(log types.Log) (*RegistryReleaseCreated, error) {
	event := new(RegistryReleaseCreated)
	if err := _Registry.contract.UnpackLog(event, "ReleaseCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryReleaseRevokedIterator is returned from FilterReleaseRevoked and is used to iterate over the raw logs and unpacked data for ReleaseRevoked events raised by the Registry contract.
type RegistryReleaseRevokedIterator struct {
	Event *RegistryReleaseRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryReleaseRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryReleaseRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RegistryReleaseRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RegistryReleaseRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryReleaseRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryReleaseRevoked represents a ReleaseRevoked event raised by the Registry contract.
type RegistryReleaseRevoked struct {
	ReleaseID *big.Int
	Sender    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReleaseRevoked is a free log retrieval operation binding the contract event 0x6709c97fa27e5e2954e292620faa1a587d1f25adb1bef740c229b7d947b41a99.
//
// Solidity: event ReleaseRevoked(uint256 _releaseID, address _sender)
func (_Registry *RegistryFilterer) FilterReleaseRevoked(opts *bind.FilterOpts) (*RegistryReleaseRevokedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ReleaseRevoked")
	if err != nil {
		return nil, err
	}
	return &RegistryReleaseRevokedIterator{contract: _Registry.contract, event: "ReleaseRevoked", logs: logs, sub: sub}, nil
}

// WatchReleaseRevoked is a free log subscription operation binding the contract event 0x6709c97fa27e5e2954e292620faa1a587d1f25adb1bef740c229b7d947b41a99.
//
// Solidity: event ReleaseRevoked(uint256 _releaseID, address _sender)
func (_Registry *RegistryFilterer) WatchReleaseRevoked(opts *bind.WatchOpts, sink chan<- *RegistryReleaseRevoked) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ReleaseRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryReleaseRevoked)
				if err := _Registry.contract.UnpackLog(event, "ReleaseRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReleaseRevoked is a log parse operation binding the contract event 0x6709c97fa27e5e2954e292620faa1a587d1f25adb1bef740c229b7d947b41a99.
//
// Solidity: event ReleaseRevoked(uint256 _releaseID, address _sender)
func (_Registry *RegistryFilterer) ParseReleaseRevoked(log types.Log) (*RegistryReleaseRevoked, error) {
	event := new(RegistryReleaseRevoked)
	if err := _Registry.contract.UnpackLog(event, "ReleaseRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

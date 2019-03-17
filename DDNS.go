// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eddns

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// EddnsABI is the input ABI used to generate the binding from.
const EddnsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"paymentReceipts\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DOMAIN_NAME_COST_SHORT_ADDITION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DOMAIN_NAME_COST\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"domain\",\"type\":\"bytes\"},{\"name\":\"tld\",\"type\":\"bytes12\"},{\"name\":\"ip\",\"type\":\"bytes15\"}],\"name\":\"register\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"domainNames\",\"outputs\":[{\"name\":\"name\",\"type\":\"bytes\"},{\"name\":\"tld\",\"type\":\"bytes12\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"expires\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"domain\",\"type\":\"bytes\"},{\"name\":\"tld\",\"type\":\"bytes12\"}],\"name\":\"getIp\",\"outputs\":[{\"name\":\"outArr\",\"type\":\"bytes15[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DOMAIN_NAME_EXPENSIVE_LENGTH\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DOMAIN_NAME_MIN_LENGTH\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"domain\",\"type\":\"bytes\"},{\"name\":\"topLevel\",\"type\":\"bytes12\"}],\"name\":\"getDomainHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"domain\",\"type\":\"bytes\"}],\"name\":\"getPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DOMAIN_EXPIRATION_DATE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"domain\",\"type\":\"bytes\"},{\"name\":\"tld\",\"type\":\"bytes12\"},{\"name\":\"newIp\",\"type\":\"bytes15\"}],\"name\":\"edit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"receiptDetails\",\"outputs\":[{\"name\":\"amountPaid\",\"type\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\"},{\"name\":\"expires\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BYTES_DEFAULT_VALUE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes1\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TLD_MIN_LENGTH\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"domainName\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"tld\",\"type\":\"bytes12\"}],\"name\":\"DomainNameRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"domainName\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"tld\",\"type\":\"bytes12\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"DomainNameRenewed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"domainName\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"tld\",\"type\":\"bytes12\"},{\"indexed\":false,\"name\":\"newIp\",\"type\":\"bytes15\"}],\"name\":\"DomainNameEdited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"domainName\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"tld\",\"type\":\"bytes12\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"DomainNameTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PurchaseChangeReturned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"domainName\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"amountInWei\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"ReceiptSaved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// EddnsBin is the compiled bytecode used for deploying new contracts.
const EddnsBin = `60806040523480156200001157600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620000da6040518060400160405280600681526020017f676f6f676c6500000000000000000000000000000000000000000000000000008152507f636f6d00000000000000000000000000000000000000000000000000000000007f392e392e392e3900000000000000000000000000000000000000000000000000620000e060201b60201c565b62000700565b82826000620000f683836200043460201b60201c565b905042600160008381526020019081526020016000206003015410151562000186576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f446f6d61696e204e616d65206973206e6f7420617661696c61626c650000000081525060200191505060405180910390fd5b60006200019a87876200043460201b60201c565b905060606001604051908082528060200260200182016040528015620001cf5781602001602082028038833980820191505090505b50905085816000815181101515620001e357fe5b9060200190602002019070ffffffffffffffffffffffffffffffffff1916908170ffffffffffffffffffffffffffffffffff19168152505062000225620004ea565b6040518060a001604052808a81526020018973ffffffffffffffffffffffffffffffffffffffff191681526020013373ffffffffffffffffffffffffffffffffffffffff1681526020018381526020016301e133804201815250905080600160008581526020019081526020016000206000820151816000019080519060200190620002b392919062000546565b5060208201518160010160006101000a8154816bffffffffffffffffffffffff021916908360a01c0217905550604082015181600101600c6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550606082015181600201908051906020019062000345929190620005cd565b5060808201518160030155905050427f832f37c89efadc9f90d7ab1795bd40ec811a854e2185331f3ebd563110a16a148a8a60405180806020018373ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff19168152602001828103825284818151815260200191508051906020019080838360005b83811015620003ed578082015181840152602081019050620003d0565b50505050905090810190601f1680156200041b5780820380516001836020036101000a031916815260200191505b50935050505060405180910390a2505050505050505050565b600082826040516020018083805190602001908083835b6020831015156200047257805182526020820191506020810190506020830392506200044b565b6001836020036101000a0380198251168184511680821785525050505050509050018273ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff19168152600c019250505060405160208183030381529060405280519060200120905092915050565b6040518060a0016040528060608152602001600073ffffffffffffffffffffffffffffffffffffffff19168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200058957805160ff1916838001178555620005ba565b82800160010185558215620005ba579182015b82811115620005b95782518255916020019190600101906200059c565b5b509050620005c9919062000697565b5090565b82805482825590600052602060002090600101600290048101928215620006845791602002820160005b838211156200064557835183826101000a8154816effffffffffffffffffffffffffffff021916908360881c02179055509260200192600f01602081600e01049283019260010302620005f7565b8015620006825782816101000a8154906effffffffffffffffffffffffffffff0219169055600f01602081600e0104928301926001030262000645565b505b509050620006939190620006bf565b5090565b620006bc91905b80821115620006b85760008160009055506001016200069e565b5090565b90565b620006fd91905b80821115620006f957600081816101000a8154906effffffffffffffffffffffffffffff021916905550600101620006c6565b5090565b90565b6117e780620007106000396000f3fe6080604052600436106101145760003560e01c80636b4dd158116100a05780638f32d59b116100645780638f32d59b14610914578063b7a1ff8f14610943578063d8d57733146109a0578063f2fde38b14610a0f578063f7a37c4714610a6057610114565b80636b4dd15814610698578063715018a61461077457806373ff42a21461078b57806378129e19146107b65780638da5cb5b146108bd57610114565b8063316d2d78116100e7578063316d2d78146102d857806346aca76f146103fb5780634e6fcc1414610539578063666916f81461056a5780636ac98cdf1461059b57610114565b806304f9c16d14610119578063080f81ba14610188578063148d9417146101b35780631e932159146101de575b600080fd5b34801561012557600080fd5b506101726004803603604081101561013c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610a91565b6040518082815260200191505060405180910390f35b34801561019457600080fd5b5061019d610ac1565b6040518082815260200191505060405180910390f35b3480156101bf57600080fd5b506101c8610acd565b6040518082815260200191505060405180910390f35b6102d6600480360360608110156101f457600080fd5b810190808035906020019064010000000081111561021157600080fd5b82018360208201111561022357600080fd5b8035906020019184600183028401116401000000008311171561024557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff19169060200190929190803570ffffffffffffffffffffffffffffffffff19169060200190929190505050610ad9565b005b3480156102e457600080fd5b50610311600480360360208110156102fb57600080fd5b8101908080359060200190929190505050610e11565b60405180806020018573ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff191681526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828103825286818151815260200191508051906020019080838360005b838110156103bd5780820151818401526020810190506103a2565b50505050905090810190601f1680156103ea5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b34801561040757600080fd5b506104e26004803603604081101561041e57600080fd5b810190808035906020019064010000000081111561043b57600080fd5b82018360208201111561044d57600080fd5b8035906020019184600183028401116401000000008311171561046f57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff19169060200190929190505050610f06565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561052557808201518184015260208101905061050a565b505050509050019250505060405180910390f35b34801561054557600080fd5b5061054e610fbb565b604051808260ff1660ff16815260200191505060405180910390f35b34801561057657600080fd5b5061057f610fc0565b604051808260ff1660ff16815260200191505060405180910390f35b3480156105a757600080fd5b50610682600480360360408110156105be57600080fd5b81019080803590602001906401000000008111156105db57600080fd5b8201836020820111156105ed57600080fd5b8035906020019184600183028401116401000000008311171561060f57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff19169060200190929190505050610fc5565b6040518082815260200191505060405180910390f35b3480156106a457600080fd5b5061075e600480360360208110156106bb57600080fd5b81019080803590602001906401000000008111156106d857600080fd5b8201836020820111156106ea57600080fd5b8035906020019184600183028401116401000000008311171561070c57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050611079565b6040518082815260200191505060405180910390f35b34801561078057600080fd5b506107896110ab565b005b34801561079757600080fd5b506107a0611165565b6040518082815260200191505060405180910390f35b3480156107c257600080fd5b506108bb600480360360608110156107d957600080fd5b81019080803590602001906401000000008111156107f657600080fd5b82018360208201111561080857600080fd5b8035906020019184600183028401116401000000008311171561082a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff19169060200190929190803570ffffffffffffffffffffffffffffffffff1916906020019092919050505061116d565b005b3480156108c957600080fd5b506108d26113c0565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561092057600080fd5b506109296113e9565b604051808215151515815260200191505060405180910390f35b34801561094f57600080fd5b5061097c6004803603602081101561096657600080fd5b8101908080359060200190929190505050611440565b60405180848152602001838152602001828152602001935050505060405180910390f35b3480156109ac57600080fd5b506109b561146a565b60405180827effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200191505060405180910390f35b348015610a1b57600080fd5b50610a5e60048036036020811015610a3257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611472565b005b348015610a6c57600080fd5b50610a75611491565b604051808260ff1660ff16815260200191505060405180910390f35b600260205281600052604060002081815481101515610aac57fe5b90600052602060002001600091509150505481565b670de0b6b3a764000081565b670de0b6b3a764000081565b82826000610ae78383610fc5565b9050426001600083815260200190815260200160002060030154101515610b76576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f446f6d61696e204e616d65206973206e6f7420617661696c61626c650000000081525060200191505060405180910390fd5b6000610b828787610fc5565b905060606001604051908082528060200260200182016040528015610bb65781602001602082028038833980820191505090505b50905085816000815181101515610bc957fe5b9060200190602002019070ffffffffffffffffffffffffffffffffff1916908170ffffffffffffffffffffffffffffffffff191681525050610c09611590565b6040518060a001604052808a81526020018973ffffffffffffffffffffffffffffffffffffffff191681526020013373ffffffffffffffffffffffffffffffffffffffff1681526020018381526020016301e133804201815250905080600160008581526020019081526020016000206000820151816000019080519060200190610c959291906115ec565b5060208201518160010160006101000a8154816bffffffffffffffffffffffff021916908360a01c0217905550604082015181600101600c6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816002019080519060200190610d2592919061166c565b5060808201518160030155905050427f832f37c89efadc9f90d7ab1795bd40ec811a854e2185331f3ebd563110a16a148a8a60405180806020018373ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff19168152602001828103825284818151815260200191508051906020019080838360005b83811015610dcb578082015181840152602081019050610db0565b50505050905090810190601f168015610df85780820380516001836020036101000a031916815260200191505b50935050505060405180910390a2505050505050505050565b6001602052806000526040600020600091509050806000018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610ebd5780601f10610e9257610100808354040283529160200191610ebd565b820191906000526020600020905b815481529060010190602001808311610ea057829003601f168201915b5050505050908060010160009054906101000a900460a01b9080600101600c9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b60606000610f148484610fc5565b905060016000828152602001908152602001600020600201805480602002602001604051908101604052809291908181526020018280548015610fad57602002820191906000526020600020906000905b82829054906101000a900460881b70ffffffffffffffffffffffffffffffffff1916815260200190600f0190602082600e01049283019260010382029150808411610f655790505b505050505091505092915050565b600881565b600581565b600082826040516020018083805190602001908083835b6020831015156110015780518252602082019150602081019050602083039250610fdc565b6001836020036101000a0380198251168184511680821785525050505050509050018273ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff19168152600c019250505060405160208183030381529060405280519060200120905092915050565b6000600860ff168251101561109a57670de0b6b3a7640000800190506110a6565b670de0b6b3a764000090505b919050565b6110b36113e9565b15156110be57600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a260008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6301e1338081565b8282600061117b8383610fc5565b90503373ffffffffffffffffffffffffffffffffffffffff1660016000838152602001908152602001600020600101600c9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515611239576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001806117936029913960400191505060405180910390fd5b60006112458787610fc5565b90506001600082815260200190815260200160002060020185908060018154018082558091505090600182039060005260206000209060029182820401919006600f02909192909190916101000a8154816effffffffffffffffffffffffffffff021916908360881c021790555050427f7b95dee59bbf82f78c8cb2ee4c0db3dc14b02b55c6608b7861b4213ee25f40f188888860405180806020018473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff191681526020018370ffffffffffffffffffffffffffffffffff191670ffffffffffffffffffffffffffffffffff19168152602001828103825285818151815260200191508051906020019080838360005b8381101561137b578082015181840152602081019050611360565b50505050905090810190601f1680156113a85780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a250505050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b60036020528060005260406000206000915090508060000154908060010154908060020154905083565b600060f81b81565b61147a6113e9565b151561148557600080fd5b61148e81611496565b50565b600181565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156114d257600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6040518060a0016040528060608152602001600073ffffffffffffffffffffffffffffffffffffffff19168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061162d57805160ff191683800117855561165b565b8280016001018555821561165b579182015b8281111561165a57825182559160200191906001019061163f565b5b509050611668919061172f565b5090565b8280548282559060005260206000209060010160029004810192821561171e5791602002820160005b838211156116e157835183826101000a8154816effffffffffffffffffffffffffffff021916908360881c02179055509260200192600f01602081600e01049283019260010302611695565b801561171c5782816101000a8154906effffffffffffffffffffffffffffff0219169055600f01602081600e010492830192600103026116e1565b505b50905061172b9190611754565b5090565b61175191905b8082111561174d576000816000905550600101611735565b5090565b90565b61178f91905b8082111561178b57600081816101000a8154906effffffffffffffffffffffffffffff02191690555060010161175a565b5090565b9056fe5472616e73616374696f6e20696e69746961746f72206973206e6f7420646f6d61696e206f776e6572a165627a7a72305820190021a72091f4f651f913478dfaea7cfad14f4294ed94fd9be7042bda9659ba0029`

// DeployEddns deploys a new Ethereum contract, binding an instance of Eddns to it.
func DeployEddns(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Eddns, error) {
	parsed, err := abi.JSON(strings.NewReader(EddnsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EddnsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Eddns{EddnsCaller: EddnsCaller{contract: contract}, EddnsTransactor: EddnsTransactor{contract: contract}, EddnsFilterer: EddnsFilterer{contract: contract}}, nil
}

// Eddns is an auto generated Go binding around an Ethereum contract.
type Eddns struct {
	EddnsCaller     // Read-only binding to the contract
	EddnsTransactor // Write-only binding to the contract
	EddnsFilterer   // Log filterer for contract events
}

// EddnsCaller is an auto generated read-only Go binding around an Ethereum contract.
type EddnsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EddnsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EddnsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EddnsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EddnsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EddnsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EddnsSession struct {
	Contract     *Eddns            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EddnsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EddnsCallerSession struct {
	Contract *EddnsCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EddnsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EddnsTransactorSession struct {
	Contract     *EddnsTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EddnsRaw is an auto generated low-level Go binding around an Ethereum contract.
type EddnsRaw struct {
	Contract *Eddns // Generic contract binding to access the raw methods on
}

// EddnsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EddnsCallerRaw struct {
	Contract *EddnsCaller // Generic read-only contract binding to access the raw methods on
}

// EddnsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EddnsTransactorRaw struct {
	Contract *EddnsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEddns creates a new instance of Eddns, bound to a specific deployed contract.
func NewEddns(address common.Address, backend bind.ContractBackend) (*Eddns, error) {
	contract, err := bindEddns(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Eddns{EddnsCaller: EddnsCaller{contract: contract}, EddnsTransactor: EddnsTransactor{contract: contract}, EddnsFilterer: EddnsFilterer{contract: contract}}, nil
}

// NewEddnsCaller creates a new read-only instance of Eddns, bound to a specific deployed contract.
func NewEddnsCaller(address common.Address, caller bind.ContractCaller) (*EddnsCaller, error) {
	contract, err := bindEddns(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EddnsCaller{contract: contract}, nil
}

// NewEddnsTransactor creates a new write-only instance of Eddns, bound to a specific deployed contract.
func NewEddnsTransactor(address common.Address, transactor bind.ContractTransactor) (*EddnsTransactor, error) {
	contract, err := bindEddns(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EddnsTransactor{contract: contract}, nil
}

// NewEddnsFilterer creates a new log filterer instance of Eddns, bound to a specific deployed contract.
func NewEddnsFilterer(address common.Address, filterer bind.ContractFilterer) (*EddnsFilterer, error) {
	contract, err := bindEddns(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EddnsFilterer{contract: contract}, nil
}

// bindEddns binds a generic wrapper to an already deployed contract.
func bindEddns(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EddnsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eddns *EddnsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Eddns.Contract.EddnsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eddns *EddnsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eddns.Contract.EddnsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eddns *EddnsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eddns.Contract.EddnsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eddns *EddnsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Eddns.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eddns *EddnsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eddns.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eddns *EddnsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eddns.Contract.contract.Transact(opts, method, params...)
}

// BYTESDEFAULTVALUE is a free data retrieval call binding the contract method 0xd8d57733.
//
// Solidity: function BYTES_DEFAULT_VALUE() constant returns(bytes1)
func (_Eddns *EddnsCaller) BYTESDEFAULTVALUE(opts *bind.CallOpts) ([1]byte, error) {
	var (
		ret0 = new([1]byte)
	)
	out := ret0
	err := _Eddns.contract.Call(opts, out, "BYTES_DEFAULT_VALUE")
	return *ret0, err
}

// BYTESDEFAULTVALUE is a free data retrieval call binding the contract method 0xd8d57733.
//
// Solidity: function BYTES_DEFAULT_VALUE() constant returns(bytes1)
func (_Eddns *EddnsSession) BYTESDEFAULTVALUE() ([1]byte, error) {
	return _Eddns.Contract.BYTESDEFAULTVALUE(&_Eddns.CallOpts)
}

// BYTESDEFAULTVALUE is a free data retrieval call binding the contract method 0xd8d57733.
//
// Solidity: function BYTES_DEFAULT_VALUE() constant returns(bytes1)
func (_Eddns *EddnsCallerSession) BYTESDEFAULTVALUE() ([1]byte, error) {
	return _Eddns.Contract.BYTESDEFAULTVALUE(&_Eddns.CallOpts)
}

// DOMAINEXPIRATIONDATE is a free data retrieval call binding the contract method 0x73ff42a2.
//
// Solidity: function DOMAIN_EXPIRATION_DATE() constant returns(uint256)
func (_Eddns *EddnsCaller) DOMAINEXPIRATIONDATE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Eddns.contract.Call(opts, out, "DOMAIN_EXPIRATION_DATE")
	return *ret0, err
}

// DOMAINEXPIRATIONDATE is a free data retrieval call binding the contract method 0x73ff42a2.
//
// Solidity: function DOMAIN_EXPIRATION_DATE() constant returns(uint256)
func (_Eddns *EddnsSession) DOMAINEXPIRATIONDATE() (*big.Int, error) {
	return _Eddns.Contract.DOMAINEXPIRATIONDATE(&_Eddns.CallOpts)
}

// DOMAINEXPIRATIONDATE is a free data retrieval call binding the contract method 0x73ff42a2.
//
// Solidity: function DOMAIN_EXPIRATION_DATE() constant returns(uint256)
func (_Eddns *EddnsCallerSession) DOMAINEXPIRATIONDATE() (*big.Int, error) {
	return _Eddns.Contract.DOMAINEXPIRATIONDATE(&_Eddns.CallOpts)
}

// DOMAINNAMECOST is a free data retrieval call binding the contract method 0x148d9417.
//
// Solidity: function DOMAIN_NAME_COST() constant returns(uint256)
func (_Eddns *EddnsCaller) DOMAINNAMECOST(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Eddns.contract.Call(opts, out, "DOMAIN_NAME_COST")
	return *ret0, err
}

// DOMAINNAMECOST is a free data retrieval call binding the contract method 0x148d9417.
//
// Solidity: function DOMAIN_NAME_COST() constant returns(uint256)
func (_Eddns *EddnsSession) DOMAINNAMECOST() (*big.Int, error) {
	return _Eddns.Contract.DOMAINNAMECOST(&_Eddns.CallOpts)
}

// DOMAINNAMECOST is a free data retrieval call binding the contract method 0x148d9417.
//
// Solidity: function DOMAIN_NAME_COST() constant returns(uint256)
func (_Eddns *EddnsCallerSession) DOMAINNAMECOST() (*big.Int, error) {
	return _Eddns.Contract.DOMAINNAMECOST(&_Eddns.CallOpts)
}

// DOMAINNAMECOSTSHORTADDITION is a free data retrieval call binding the contract method 0x080f81ba.
//
// Solidity: function DOMAIN_NAME_COST_SHORT_ADDITION() constant returns(uint256)
func (_Eddns *EddnsCaller) DOMAINNAMECOSTSHORTADDITION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Eddns.contract.Call(opts, out, "DOMAIN_NAME_COST_SHORT_ADDITION")
	return *ret0, err
}

// DOMAINNAMECOSTSHORTADDITION is a free data retrieval call binding the contract method 0x080f81ba.
//
// Solidity: function DOMAIN_NAME_COST_SHORT_ADDITION() constant returns(uint256)
func (_Eddns *EddnsSession) DOMAINNAMECOSTSHORTADDITION() (*big.Int, error) {
	return _Eddns.Contract.DOMAINNAMECOSTSHORTADDITION(&_Eddns.CallOpts)
}

// DOMAINNAMECOSTSHORTADDITION is a free data retrieval call binding the contract method 0x080f81ba.
//
// Solidity: function DOMAIN_NAME_COST_SHORT_ADDITION() constant returns(uint256)
func (_Eddns *EddnsCallerSession) DOMAINNAMECOSTSHORTADDITION() (*big.Int, error) {
	return _Eddns.Contract.DOMAINNAMECOSTSHORTADDITION(&_Eddns.CallOpts)
}

// DOMAINNAMEEXPENSIVELENGTH is a free data retrieval call binding the contract method 0x4e6fcc14.
//
// Solidity: function DOMAIN_NAME_EXPENSIVE_LENGTH() constant returns(uint8)
func (_Eddns *EddnsCaller) DOMAINNAMEEXPENSIVELENGTH(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Eddns.contract.Call(opts, out, "DOMAIN_NAME_EXPENSIVE_LENGTH")
	return *ret0, err
}

// DOMAINNAMEEXPENSIVELENGTH is a free data retrieval call binding the contract method 0x4e6fcc14.
//
// Solidity: function DOMAIN_NAME_EXPENSIVE_LENGTH() constant returns(uint8)
func (_Eddns *EddnsSession) DOMAINNAMEEXPENSIVELENGTH() (uint8, error) {
	return _Eddns.Contract.DOMAINNAMEEXPENSIVELENGTH(&_Eddns.CallOpts)
}

// DOMAINNAMEEXPENSIVELENGTH is a free data retrieval call binding the contract method 0x4e6fcc14.
//
// Solidity: function DOMAIN_NAME_EXPENSIVE_LENGTH() constant returns(uint8)
func (_Eddns *EddnsCallerSession) DOMAINNAMEEXPENSIVELENGTH() (uint8, error) {
	return _Eddns.Contract.DOMAINNAMEEXPENSIVELENGTH(&_Eddns.CallOpts)
}

// DOMAINNAMEMINLENGTH is a free data retrieval call binding the contract method 0x666916f8.
//
// Solidity: function DOMAIN_NAME_MIN_LENGTH() constant returns(uint8)
func (_Eddns *EddnsCaller) DOMAINNAMEMINLENGTH(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Eddns.contract.Call(opts, out, "DOMAIN_NAME_MIN_LENGTH")
	return *ret0, err
}

// DOMAINNAMEMINLENGTH is a free data retrieval call binding the contract method 0x666916f8.
//
// Solidity: function DOMAIN_NAME_MIN_LENGTH() constant returns(uint8)
func (_Eddns *EddnsSession) DOMAINNAMEMINLENGTH() (uint8, error) {
	return _Eddns.Contract.DOMAINNAMEMINLENGTH(&_Eddns.CallOpts)
}

// DOMAINNAMEMINLENGTH is a free data retrieval call binding the contract method 0x666916f8.
//
// Solidity: function DOMAIN_NAME_MIN_LENGTH() constant returns(uint8)
func (_Eddns *EddnsCallerSession) DOMAINNAMEMINLENGTH() (uint8, error) {
	return _Eddns.Contract.DOMAINNAMEMINLENGTH(&_Eddns.CallOpts)
}

// TLDMINLENGTH is a free data retrieval call binding the contract method 0xf7a37c47.
//
// Solidity: function TLD_MIN_LENGTH() constant returns(uint8)
func (_Eddns *EddnsCaller) TLDMINLENGTH(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Eddns.contract.Call(opts, out, "TLD_MIN_LENGTH")
	return *ret0, err
}

// TLDMINLENGTH is a free data retrieval call binding the contract method 0xf7a37c47.
//
// Solidity: function TLD_MIN_LENGTH() constant returns(uint8)
func (_Eddns *EddnsSession) TLDMINLENGTH() (uint8, error) {
	return _Eddns.Contract.TLDMINLENGTH(&_Eddns.CallOpts)
}

// TLDMINLENGTH is a free data retrieval call binding the contract method 0xf7a37c47.
//
// Solidity: function TLD_MIN_LENGTH() constant returns(uint8)
func (_Eddns *EddnsCallerSession) TLDMINLENGTH() (uint8, error) {
	return _Eddns.Contract.TLDMINLENGTH(&_Eddns.CallOpts)
}

// DomainNames is a free data retrieval call binding the contract method 0x316d2d78.
//
// Solidity: function domainNames(bytes32 ) constant returns(bytes name, bytes12 tld, address owner, uint256 expires)
func (_Eddns *EddnsCaller) DomainNames(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Name    []byte
	Tld     [12]byte
	Owner   common.Address
	Expires *big.Int
}, error) {
	ret := new(struct {
		Name    []byte
		Tld     [12]byte
		Owner   common.Address
		Expires *big.Int
	})
	out := ret
	err := _Eddns.contract.Call(opts, out, "domainNames", arg0)
	return *ret, err
}

// DomainNames is a free data retrieval call binding the contract method 0x316d2d78.
//
// Solidity: function domainNames(bytes32 ) constant returns(bytes name, bytes12 tld, address owner, uint256 expires)
func (_Eddns *EddnsSession) DomainNames(arg0 [32]byte) (struct {
	Name    []byte
	Tld     [12]byte
	Owner   common.Address
	Expires *big.Int
}, error) {
	return _Eddns.Contract.DomainNames(&_Eddns.CallOpts, arg0)
}

// DomainNames is a free data retrieval call binding the contract method 0x316d2d78.
//
// Solidity: function domainNames(bytes32 ) constant returns(bytes name, bytes12 tld, address owner, uint256 expires)
func (_Eddns *EddnsCallerSession) DomainNames(arg0 [32]byte) (struct {
	Name    []byte
	Tld     [12]byte
	Owner   common.Address
	Expires *big.Int
}, error) {
	return _Eddns.Contract.DomainNames(&_Eddns.CallOpts, arg0)
}

// GetDomainHash is a free data retrieval call binding the contract method 0x6ac98cdf.
//
// Solidity: function getDomainHash(bytes domain, bytes12 topLevel) constant returns(bytes32)
func (_Eddns *EddnsCaller) GetDomainHash(opts *bind.CallOpts, domain []byte, topLevel [12]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Eddns.contract.Call(opts, out, "getDomainHash", domain, topLevel)
	return *ret0, err
}

// GetDomainHash is a free data retrieval call binding the contract method 0x6ac98cdf.
//
// Solidity: function getDomainHash(bytes domain, bytes12 topLevel) constant returns(bytes32)
func (_Eddns *EddnsSession) GetDomainHash(domain []byte, topLevel [12]byte) ([32]byte, error) {
	return _Eddns.Contract.GetDomainHash(&_Eddns.CallOpts, domain, topLevel)
}

// GetDomainHash is a free data retrieval call binding the contract method 0x6ac98cdf.
//
// Solidity: function getDomainHash(bytes domain, bytes12 topLevel) constant returns(bytes32)
func (_Eddns *EddnsCallerSession) GetDomainHash(domain []byte, topLevel [12]byte) ([32]byte, error) {
	return _Eddns.Contract.GetDomainHash(&_Eddns.CallOpts, domain, topLevel)
}

// GetIp is a free data retrieval call binding the contract method 0x46aca76f.
//
// Solidity: function getIp(bytes domain, bytes12 tld) constant returns(bytes15[] outArr)
func (_Eddns *EddnsCaller) GetIp(opts *bind.CallOpts, domain []byte, tld [12]byte) ([][15]byte, error) {
	var (
		ret0 = new([][15]byte)
	)
	out := ret0
	err := _Eddns.contract.Call(opts, out, "getIp", domain, tld)
	return *ret0, err
}

// GetIp is a free data retrieval call binding the contract method 0x46aca76f.
//
// Solidity: function getIp(bytes domain, bytes12 tld) constant returns(bytes15[] outArr)
func (_Eddns *EddnsSession) GetIp(domain []byte, tld [12]byte) ([][15]byte, error) {
	return _Eddns.Contract.GetIp(&_Eddns.CallOpts, domain, tld)
}

// GetIp is a free data retrieval call binding the contract method 0x46aca76f.
//
// Solidity: function getIp(bytes domain, bytes12 tld) constant returns(bytes15[] outArr)
func (_Eddns *EddnsCallerSession) GetIp(domain []byte, tld [12]byte) ([][15]byte, error) {
	return _Eddns.Contract.GetIp(&_Eddns.CallOpts, domain, tld)
}

// GetPrice is a free data retrieval call binding the contract method 0x6b4dd158.
//
// Solidity: function getPrice(bytes domain) constant returns(uint256)
func (_Eddns *EddnsCaller) GetPrice(opts *bind.CallOpts, domain []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Eddns.contract.Call(opts, out, "getPrice", domain)
	return *ret0, err
}

// GetPrice is a free data retrieval call binding the contract method 0x6b4dd158.
//
// Solidity: function getPrice(bytes domain) constant returns(uint256)
func (_Eddns *EddnsSession) GetPrice(domain []byte) (*big.Int, error) {
	return _Eddns.Contract.GetPrice(&_Eddns.CallOpts, domain)
}

// GetPrice is a free data retrieval call binding the contract method 0x6b4dd158.
//
// Solidity: function getPrice(bytes domain) constant returns(uint256)
func (_Eddns *EddnsCallerSession) GetPrice(domain []byte) (*big.Int, error) {
	return _Eddns.Contract.GetPrice(&_Eddns.CallOpts, domain)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Eddns *EddnsCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Eddns.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Eddns *EddnsSession) IsOwner() (bool, error) {
	return _Eddns.Contract.IsOwner(&_Eddns.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Eddns *EddnsCallerSession) IsOwner() (bool, error) {
	return _Eddns.Contract.IsOwner(&_Eddns.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Eddns *EddnsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Eddns.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Eddns *EddnsSession) Owner() (common.Address, error) {
	return _Eddns.Contract.Owner(&_Eddns.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Eddns *EddnsCallerSession) Owner() (common.Address, error) {
	return _Eddns.Contract.Owner(&_Eddns.CallOpts)
}

// PaymentReceipts is a free data retrieval call binding the contract method 0x04f9c16d.
//
// Solidity: function paymentReceipts(address , uint256 ) constant returns(bytes32)
func (_Eddns *EddnsCaller) PaymentReceipts(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Eddns.contract.Call(opts, out, "paymentReceipts", arg0, arg1)
	return *ret0, err
}

// PaymentReceipts is a free data retrieval call binding the contract method 0x04f9c16d.
//
// Solidity: function paymentReceipts(address , uint256 ) constant returns(bytes32)
func (_Eddns *EddnsSession) PaymentReceipts(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _Eddns.Contract.PaymentReceipts(&_Eddns.CallOpts, arg0, arg1)
}

// PaymentReceipts is a free data retrieval call binding the contract method 0x04f9c16d.
//
// Solidity: function paymentReceipts(address , uint256 ) constant returns(bytes32)
func (_Eddns *EddnsCallerSession) PaymentReceipts(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _Eddns.Contract.PaymentReceipts(&_Eddns.CallOpts, arg0, arg1)
}

// ReceiptDetails is a free data retrieval call binding the contract method 0xb7a1ff8f.
//
// Solidity: function receiptDetails(bytes32 ) constant returns(uint256 amountPaid, uint256 timestamp, uint256 expires)
func (_Eddns *EddnsCaller) ReceiptDetails(opts *bind.CallOpts, arg0 [32]byte) (struct {
	AmountPaid *big.Int
	Timestamp  *big.Int
	Expires    *big.Int
}, error) {
	ret := new(struct {
		AmountPaid *big.Int
		Timestamp  *big.Int
		Expires    *big.Int
	})
	out := ret
	err := _Eddns.contract.Call(opts, out, "receiptDetails", arg0)
	return *ret, err
}

// ReceiptDetails is a free data retrieval call binding the contract method 0xb7a1ff8f.
//
// Solidity: function receiptDetails(bytes32 ) constant returns(uint256 amountPaid, uint256 timestamp, uint256 expires)
func (_Eddns *EddnsSession) ReceiptDetails(arg0 [32]byte) (struct {
	AmountPaid *big.Int
	Timestamp  *big.Int
	Expires    *big.Int
}, error) {
	return _Eddns.Contract.ReceiptDetails(&_Eddns.CallOpts, arg0)
}

// ReceiptDetails is a free data retrieval call binding the contract method 0xb7a1ff8f.
//
// Solidity: function receiptDetails(bytes32 ) constant returns(uint256 amountPaid, uint256 timestamp, uint256 expires)
func (_Eddns *EddnsCallerSession) ReceiptDetails(arg0 [32]byte) (struct {
	AmountPaid *big.Int
	Timestamp  *big.Int
	Expires    *big.Int
}, error) {
	return _Eddns.Contract.ReceiptDetails(&_Eddns.CallOpts, arg0)
}

// Edit is a paid mutator transaction binding the contract method 0x78129e19.
//
// Solidity: function edit(bytes domain, bytes12 tld, bytes15 newIp) returns()
func (_Eddns *EddnsTransactor) Edit(opts *bind.TransactOpts, domain []byte, tld [12]byte, newIp [15]byte) (*types.Transaction, error) {
	return _Eddns.contract.Transact(opts, "edit", domain, tld, newIp)
}

// Edit is a paid mutator transaction binding the contract method 0x78129e19.
//
// Solidity: function edit(bytes domain, bytes12 tld, bytes15 newIp) returns()
func (_Eddns *EddnsSession) Edit(domain []byte, tld [12]byte, newIp [15]byte) (*types.Transaction, error) {
	return _Eddns.Contract.Edit(&_Eddns.TransactOpts, domain, tld, newIp)
}

// Edit is a paid mutator transaction binding the contract method 0x78129e19.
//
// Solidity: function edit(bytes domain, bytes12 tld, bytes15 newIp) returns()
func (_Eddns *EddnsTransactorSession) Edit(domain []byte, tld [12]byte, newIp [15]byte) (*types.Transaction, error) {
	return _Eddns.Contract.Edit(&_Eddns.TransactOpts, domain, tld, newIp)
}

// Register is a paid mutator transaction binding the contract method 0x1e932159.
//
// Solidity: function register(bytes domain, bytes12 tld, bytes15 ip) returns()
func (_Eddns *EddnsTransactor) Register(opts *bind.TransactOpts, domain []byte, tld [12]byte, ip [15]byte) (*types.Transaction, error) {
	return _Eddns.contract.Transact(opts, "register", domain, tld, ip)
}

// Register is a paid mutator transaction binding the contract method 0x1e932159.
//
// Solidity: function register(bytes domain, bytes12 tld, bytes15 ip) returns()
func (_Eddns *EddnsSession) Register(domain []byte, tld [12]byte, ip [15]byte) (*types.Transaction, error) {
	return _Eddns.Contract.Register(&_Eddns.TransactOpts, domain, tld, ip)
}

// Register is a paid mutator transaction binding the contract method 0x1e932159.
//
// Solidity: function register(bytes domain, bytes12 tld, bytes15 ip) returns()
func (_Eddns *EddnsTransactorSession) Register(domain []byte, tld [12]byte, ip [15]byte) (*types.Transaction, error) {
	return _Eddns.Contract.Register(&_Eddns.TransactOpts, domain, tld, ip)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Eddns *EddnsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eddns.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Eddns *EddnsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Eddns.Contract.RenounceOwnership(&_Eddns.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Eddns *EddnsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Eddns.Contract.RenounceOwnership(&_Eddns.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Eddns *EddnsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Eddns.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Eddns *EddnsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Eddns.Contract.TransferOwnership(&_Eddns.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Eddns *EddnsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Eddns.Contract.TransferOwnership(&_Eddns.TransactOpts, newOwner)
}

// EddnsDomainNameEditedIterator is returned from FilterDomainNameEdited and is used to iterate over the raw logs and unpacked data for DomainNameEdited events raised by the Eddns contract.
type EddnsDomainNameEditedIterator struct {
	Event *EddnsDomainNameEdited // Event containing the contract specifics and raw log

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
func (it *EddnsDomainNameEditedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EddnsDomainNameEdited)
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
		it.Event = new(EddnsDomainNameEdited)
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
func (it *EddnsDomainNameEditedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EddnsDomainNameEditedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EddnsDomainNameEdited represents a DomainNameEdited event raised by the Eddns contract.
type EddnsDomainNameEdited struct {
	Timestamp  *big.Int
	DomainName []byte
	Tld        [12]byte
	NewIp      [15]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDomainNameEdited is a free log retrieval operation binding the contract event 0x7b95dee59bbf82f78c8cb2ee4c0db3dc14b02b55c6608b7861b4213ee25f40f1.
//
// Solidity: event DomainNameEdited(uint256 indexed timestamp, bytes domainName, bytes12 tld, bytes15 newIp)
func (_Eddns *EddnsFilterer) FilterDomainNameEdited(opts *bind.FilterOpts, timestamp []*big.Int) (*EddnsDomainNameEditedIterator, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _Eddns.contract.FilterLogs(opts, "DomainNameEdited", timestampRule)
	if err != nil {
		return nil, err
	}
	return &EddnsDomainNameEditedIterator{contract: _Eddns.contract, event: "DomainNameEdited", logs: logs, sub: sub}, nil
}

// WatchDomainNameEdited is a free log subscription operation binding the contract event 0x7b95dee59bbf82f78c8cb2ee4c0db3dc14b02b55c6608b7861b4213ee25f40f1.
//
// Solidity: event DomainNameEdited(uint256 indexed timestamp, bytes domainName, bytes12 tld, bytes15 newIp)
func (_Eddns *EddnsFilterer) WatchDomainNameEdited(opts *bind.WatchOpts, sink chan<- *EddnsDomainNameEdited, timestamp []*big.Int) (event.Subscription, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _Eddns.contract.WatchLogs(opts, "DomainNameEdited", timestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EddnsDomainNameEdited)
				if err := _Eddns.contract.UnpackLog(event, "DomainNameEdited", log); err != nil {
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

// EddnsDomainNameRegisteredIterator is returned from FilterDomainNameRegistered and is used to iterate over the raw logs and unpacked data for DomainNameRegistered events raised by the Eddns contract.
type EddnsDomainNameRegisteredIterator struct {
	Event *EddnsDomainNameRegistered // Event containing the contract specifics and raw log

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
func (it *EddnsDomainNameRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EddnsDomainNameRegistered)
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
		it.Event = new(EddnsDomainNameRegistered)
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
func (it *EddnsDomainNameRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EddnsDomainNameRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EddnsDomainNameRegistered represents a DomainNameRegistered event raised by the Eddns contract.
type EddnsDomainNameRegistered struct {
	Timestamp  *big.Int
	DomainName []byte
	Tld        [12]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDomainNameRegistered is a free log retrieval operation binding the contract event 0x832f37c89efadc9f90d7ab1795bd40ec811a854e2185331f3ebd563110a16a14.
//
// Solidity: event DomainNameRegistered(uint256 indexed timestamp, bytes domainName, bytes12 tld)
func (_Eddns *EddnsFilterer) FilterDomainNameRegistered(opts *bind.FilterOpts, timestamp []*big.Int) (*EddnsDomainNameRegisteredIterator, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _Eddns.contract.FilterLogs(opts, "DomainNameRegistered", timestampRule)
	if err != nil {
		return nil, err
	}
	return &EddnsDomainNameRegisteredIterator{contract: _Eddns.contract, event: "DomainNameRegistered", logs: logs, sub: sub}, nil
}

// WatchDomainNameRegistered is a free log subscription operation binding the contract event 0x832f37c89efadc9f90d7ab1795bd40ec811a854e2185331f3ebd563110a16a14.
//
// Solidity: event DomainNameRegistered(uint256 indexed timestamp, bytes domainName, bytes12 tld)
func (_Eddns *EddnsFilterer) WatchDomainNameRegistered(opts *bind.WatchOpts, sink chan<- *EddnsDomainNameRegistered, timestamp []*big.Int) (event.Subscription, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _Eddns.contract.WatchLogs(opts, "DomainNameRegistered", timestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EddnsDomainNameRegistered)
				if err := _Eddns.contract.UnpackLog(event, "DomainNameRegistered", log); err != nil {
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

// EddnsDomainNameRenewedIterator is returned from FilterDomainNameRenewed and is used to iterate over the raw logs and unpacked data for DomainNameRenewed events raised by the Eddns contract.
type EddnsDomainNameRenewedIterator struct {
	Event *EddnsDomainNameRenewed // Event containing the contract specifics and raw log

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
func (it *EddnsDomainNameRenewedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EddnsDomainNameRenewed)
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
		it.Event = new(EddnsDomainNameRenewed)
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
func (it *EddnsDomainNameRenewedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EddnsDomainNameRenewedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EddnsDomainNameRenewed represents a DomainNameRenewed event raised by the Eddns contract.
type EddnsDomainNameRenewed struct {
	Timestamp  *big.Int
	DomainName []byte
	Tld        [12]byte
	Owner      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDomainNameRenewed is a free log retrieval operation binding the contract event 0x3d2e3c19c0f8c7f676dc851894f9803941d5c72860636059e2fdb8349ece3184.
//
// Solidity: event DomainNameRenewed(uint256 indexed timestamp, bytes domainName, bytes12 tld, address indexed owner)
func (_Eddns *EddnsFilterer) FilterDomainNameRenewed(opts *bind.FilterOpts, timestamp []*big.Int, owner []common.Address) (*EddnsDomainNameRenewedIterator, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Eddns.contract.FilterLogs(opts, "DomainNameRenewed", timestampRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &EddnsDomainNameRenewedIterator{contract: _Eddns.contract, event: "DomainNameRenewed", logs: logs, sub: sub}, nil
}

// WatchDomainNameRenewed is a free log subscription operation binding the contract event 0x3d2e3c19c0f8c7f676dc851894f9803941d5c72860636059e2fdb8349ece3184.
//
// Solidity: event DomainNameRenewed(uint256 indexed timestamp, bytes domainName, bytes12 tld, address indexed owner)
func (_Eddns *EddnsFilterer) WatchDomainNameRenewed(opts *bind.WatchOpts, sink chan<- *EddnsDomainNameRenewed, timestamp []*big.Int, owner []common.Address) (event.Subscription, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Eddns.contract.WatchLogs(opts, "DomainNameRenewed", timestampRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EddnsDomainNameRenewed)
				if err := _Eddns.contract.UnpackLog(event, "DomainNameRenewed", log); err != nil {
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

// EddnsDomainNameTransferredIterator is returned from FilterDomainNameTransferred and is used to iterate over the raw logs and unpacked data for DomainNameTransferred events raised by the Eddns contract.
type EddnsDomainNameTransferredIterator struct {
	Event *EddnsDomainNameTransferred // Event containing the contract specifics and raw log

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
func (it *EddnsDomainNameTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EddnsDomainNameTransferred)
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
		it.Event = new(EddnsDomainNameTransferred)
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
func (it *EddnsDomainNameTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EddnsDomainNameTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EddnsDomainNameTransferred represents a DomainNameTransferred event raised by the Eddns contract.
type EddnsDomainNameTransferred struct {
	Timestamp  *big.Int
	DomainName []byte
	Tld        [12]byte
	Owner      common.Address
	NewOwner   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDomainNameTransferred is a free log retrieval operation binding the contract event 0x279a08cd5150365d776640b24b483dabf0a2a84d7948a9c777b62c43451978fc.
//
// Solidity: event DomainNameTransferred(uint256 indexed timestamp, bytes domainName, bytes12 tld, address indexed owner, address newOwner)
func (_Eddns *EddnsFilterer) FilterDomainNameTransferred(opts *bind.FilterOpts, timestamp []*big.Int, owner []common.Address) (*EddnsDomainNameTransferredIterator, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Eddns.contract.FilterLogs(opts, "DomainNameTransferred", timestampRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &EddnsDomainNameTransferredIterator{contract: _Eddns.contract, event: "DomainNameTransferred", logs: logs, sub: sub}, nil
}

// WatchDomainNameTransferred is a free log subscription operation binding the contract event 0x279a08cd5150365d776640b24b483dabf0a2a84d7948a9c777b62c43451978fc.
//
// Solidity: event DomainNameTransferred(uint256 indexed timestamp, bytes domainName, bytes12 tld, address indexed owner, address newOwner)
func (_Eddns *EddnsFilterer) WatchDomainNameTransferred(opts *bind.WatchOpts, sink chan<- *EddnsDomainNameTransferred, timestamp []*big.Int, owner []common.Address) (event.Subscription, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Eddns.contract.WatchLogs(opts, "DomainNameTransferred", timestampRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EddnsDomainNameTransferred)
				if err := _Eddns.contract.UnpackLog(event, "DomainNameTransferred", log); err != nil {
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

// EddnsOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Eddns contract.
type EddnsOwnershipRenouncedIterator struct {
	Event *EddnsOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *EddnsOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EddnsOwnershipRenounced)
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
		it.Event = new(EddnsOwnershipRenounced)
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
func (it *EddnsOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EddnsOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EddnsOwnershipRenounced represents a OwnershipRenounced event raised by the Eddns contract.
type EddnsOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Eddns *EddnsFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*EddnsOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Eddns.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EddnsOwnershipRenouncedIterator{contract: _Eddns.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Eddns *EddnsFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *EddnsOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Eddns.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EddnsOwnershipRenounced)
				if err := _Eddns.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// EddnsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Eddns contract.
type EddnsOwnershipTransferredIterator struct {
	Event *EddnsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EddnsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EddnsOwnershipTransferred)
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
		it.Event = new(EddnsOwnershipTransferred)
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
func (it *EddnsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EddnsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EddnsOwnershipTransferred represents a OwnershipTransferred event raised by the Eddns contract.
type EddnsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Eddns *EddnsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EddnsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Eddns.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EddnsOwnershipTransferredIterator{contract: _Eddns.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Eddns *EddnsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EddnsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Eddns.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EddnsOwnershipTransferred)
				if err := _Eddns.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// EddnsPurchaseChangeReturnedIterator is returned from FilterPurchaseChangeReturned and is used to iterate over the raw logs and unpacked data for PurchaseChangeReturned events raised by the Eddns contract.
type EddnsPurchaseChangeReturnedIterator struct {
	Event *EddnsPurchaseChangeReturned // Event containing the contract specifics and raw log

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
func (it *EddnsPurchaseChangeReturnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EddnsPurchaseChangeReturned)
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
		it.Event = new(EddnsPurchaseChangeReturned)
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
func (it *EddnsPurchaseChangeReturnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EddnsPurchaseChangeReturnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EddnsPurchaseChangeReturned represents a PurchaseChangeReturned event raised by the Eddns contract.
type EddnsPurchaseChangeReturned struct {
	Timestamp *big.Int
	Owner     common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPurchaseChangeReturned is a free log retrieval operation binding the contract event 0x143928567f98bc30842393622febbc6cb1974def90f176b87d47138d99ffea16.
//
// Solidity: event PurchaseChangeReturned(uint256 indexed timestamp, address indexed _owner, uint256 amount)
func (_Eddns *EddnsFilterer) FilterPurchaseChangeReturned(opts *bind.FilterOpts, timestamp []*big.Int, _owner []common.Address) (*EddnsPurchaseChangeReturnedIterator, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}
	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}

	logs, sub, err := _Eddns.contract.FilterLogs(opts, "PurchaseChangeReturned", timestampRule, _ownerRule)
	if err != nil {
		return nil, err
	}
	return &EddnsPurchaseChangeReturnedIterator{contract: _Eddns.contract, event: "PurchaseChangeReturned", logs: logs, sub: sub}, nil
}

// WatchPurchaseChangeReturned is a free log subscription operation binding the contract event 0x143928567f98bc30842393622febbc6cb1974def90f176b87d47138d99ffea16.
//
// Solidity: event PurchaseChangeReturned(uint256 indexed timestamp, address indexed _owner, uint256 amount)
func (_Eddns *EddnsFilterer) WatchPurchaseChangeReturned(opts *bind.WatchOpts, sink chan<- *EddnsPurchaseChangeReturned, timestamp []*big.Int, _owner []common.Address) (event.Subscription, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}
	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}

	logs, sub, err := _Eddns.contract.WatchLogs(opts, "PurchaseChangeReturned", timestampRule, _ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EddnsPurchaseChangeReturned)
				if err := _Eddns.contract.UnpackLog(event, "PurchaseChangeReturned", log); err != nil {
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

// EddnsReceiptSavedIterator is returned from FilterReceiptSaved and is used to iterate over the raw logs and unpacked data for ReceiptSaved events raised by the Eddns contract.
type EddnsReceiptSavedIterator struct {
	Event *EddnsReceiptSaved // Event containing the contract specifics and raw log

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
func (it *EddnsReceiptSavedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EddnsReceiptSaved)
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
		it.Event = new(EddnsReceiptSaved)
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
func (it *EddnsReceiptSavedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EddnsReceiptSavedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EddnsReceiptSaved represents a ReceiptSaved event raised by the Eddns contract.
type EddnsReceiptSaved struct {
	Timestamp   *big.Int
	DomainName  []byte
	AmountInWei *big.Int
	Expires     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReceiptSaved is a free log retrieval operation binding the contract event 0x96aac10a352e1dd8f5c7f3d2128c2de9990c4b1288970a3f2222ef93660a000e.
//
// Solidity: event ReceiptSaved(uint256 indexed timestamp, bytes domainName, uint256 amountInWei, uint256 expires)
func (_Eddns *EddnsFilterer) FilterReceiptSaved(opts *bind.FilterOpts, timestamp []*big.Int) (*EddnsReceiptSavedIterator, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _Eddns.contract.FilterLogs(opts, "ReceiptSaved", timestampRule)
	if err != nil {
		return nil, err
	}
	return &EddnsReceiptSavedIterator{contract: _Eddns.contract, event: "ReceiptSaved", logs: logs, sub: sub}, nil
}

// WatchReceiptSaved is a free log subscription operation binding the contract event 0x96aac10a352e1dd8f5c7f3d2128c2de9990c4b1288970a3f2222ef93660a000e.
//
// Solidity: event ReceiptSaved(uint256 indexed timestamp, bytes domainName, uint256 amountInWei, uint256 expires)
func (_Eddns *EddnsFilterer) WatchReceiptSaved(opts *bind.WatchOpts, sink chan<- *EddnsReceiptSaved, timestamp []*big.Int) (event.Subscription, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}

	logs, sub, err := _Eddns.contract.WatchLogs(opts, "ReceiptSaved", timestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EddnsReceiptSaved)
				if err := _Eddns.contract.UnpackLog(event, "ReceiptSaved", log); err != nil {
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

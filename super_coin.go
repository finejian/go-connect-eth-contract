// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package coin

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

// SuperCoinABI is the input ABI used to generate the binding from.
const SuperCoinABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"mintingFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CAP_SUPPLY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INITIAL_SUPPLY\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishMinting\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeFromWhitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"existedOnWhitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ONCE_MAX_MINT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addToWhitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"WhitelistedAddressAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"WhitelistedAddressRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// SuperCoinBin is the compiled bytecode used for deploying new contracts.
const SuperCoinBin = `60806040526000600360146101000a81548160ff0219169083151502179055506040805190810160405280600981526020017f5375706572436f696e0000000000000000000000000000000000000000000000815250600690805190602001906200006c92919062000212565b506040805190810160405280600281526020017f534300000000000000000000000000000000000000000000000000000000000081525060079080519060200190620000ba92919062000212565b506012600860006101000a81548160ff021916908360ff1602179055506a56da9d67d20d77090000006009556aadb53acfa41aee12000000600a55683782dace9d90000000600b553480156200010f57600080fd5b50600a5433600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000811115156200016457600080fd5b80600481905550506009546001819055506009546000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055503373ffffffffffffffffffffffffffffffffffffffff1660007fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef6009546040518082815260200191505060405180910390a3620002c1565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200025557805160ff191683800117855562000286565b8280016001018555821562000286579182015b828111156200028557825182559160200191906001019062000268565b5b50905062000295919062000299565b5090565b620002be91905b80821115620002ba576000816000905550600101620002a0565b5090565b90565b6120cb80620002d16000396000f30060806040526004361061013e576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806305d2035b1461014357806306fdde0314610172578063070c0c4e14610202578063095ea7b31461022d57806318160ddd1461029257806323b872dd146102bd5780632ff2e9dc14610342578063313ce5671461036d578063355274ea1461039e57806340c10f19146103c9578063661884631461042e57806370a08231146104935780637d64bcb4146104ea5780638ab1d681146105195780638ab905d4146105745780638da5cb5b146105cf57806395d89b41146106265780639aabb18c146106b65780639b19251a146106e1578063a9059cbb1461073c578063d73dd623146107a1578063dd62ed3e14610806578063e43252d71461087d578063f2fde38b146108d8575b600080fd5b34801561014f57600080fd5b5061015861091b565b604051808215151515815260200191505060405180910390f35b34801561017e57600080fd5b5061018761092e565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101c75780820151818401526020810190506101ac565b50505050905090810190601f1680156101f45780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561020e57600080fd5b506102176109cc565b6040518082815260200191505060405180910390f35b34801561023957600080fd5b50610278600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506109d2565b604051808215151515815260200191505060405180910390f35b34801561029e57600080fd5b506102a7610ac4565b6040518082815260200191505060405180910390f35b3480156102c957600080fd5b50610328600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610ace565b604051808215151515815260200191505060405180910390f35b34801561034e57600080fd5b50610357610e88565b6040518082815260200191505060405180910390f35b34801561037957600080fd5b50610382610e8e565b604051808260ff1660ff16815260200191505060405180910390f35b3480156103aa57600080fd5b506103b3610ea1565b6040518082815260200191505060405180910390f35b3480156103d557600080fd5b50610414600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610ea7565b604051808215151515815260200191505060405180910390f35b34801561043a57600080fd5b50610479600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610f9d565b604051808215151515815260200191505060405180910390f35b34801561049f57600080fd5b506104d4600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061122e565b6040518082815260200191505060405180910390f35b3480156104f657600080fd5b506104ff611276565b604051808215151515815260200191505060405180910390f35b34801561052557600080fd5b5061055a600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061133e565b604051808215151515815260200191505060405180910390f35b34801561058057600080fd5b506105b5600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506114bd565b604051808215151515815260200191505060405180910390f35b3480156105db57600080fd5b506105e461156f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561063257600080fd5b5061063b611595565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561067b578082015181840152602081019050610660565b50505050905090810190601f1680156106a85780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156106c257600080fd5b506106cb611633565b6040518082815260200191505060405180910390f35b3480156106ed57600080fd5b50610722600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611639565b604051808215151515815260200191505060405180910390f35b34801561074857600080fd5b50610787600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611659565b604051808215151515815260200191505060405180910390f35b3480156107ad57600080fd5b506107ec600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611878565b604051808215151515815260200191505060405180910390f35b34801561081257600080fd5b50610867600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611a74565b6040518082815260200191505060405180910390f35b34801561088957600080fd5b506108be600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611afb565b604051808215151515815260200191505060405180910390f35b3480156108e457600080fd5b50610919600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611c79565b005b600360149054906101000a900460ff1681565b60068054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109c45780601f10610999576101008083540402835291602001916109c4565b820191906000526020600020905b8154815290600101906020018083116109a757829003601f168201915b505050505081565b600a5481565b600081600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040518082815260200191505060405180910390a36001905092915050565b6000600154905090565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614151515610b0b57600080fd5b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548211151515610b5857600080fd5b600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548211151515610be357600080fd5b610c34826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611dd190919063ffffffff16565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610cc7826000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611dea90919063ffffffff16565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610d9882600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611dd190919063ffffffff16565b600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a3600190509392505050565b60095481565b600860009054906101000a900460ff1681565b60045481565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610f0557600080fd5b600360149054906101000a900460ff16151515610f2157600080fd5b82600560008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515610f7a57600080fd5b600b54831115610f8a57600b5492505b610f948484611e08565b91505092915050565b600080600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050808311156110ae576000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611142565b6110c18382611dd190919063ffffffff16565b600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b8373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546040518082815260200191505060405180910390a3600191505092915050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156112d457600080fd5b600360149054906101000a900460ff161515156112f057600080fd5b6001600360146101000a81548160ff0219169083151502179055507fae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa0860405160405180910390a16001905090565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561139c57600080fd5b600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615156113f857600090506114b8565b6000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507ff1abf01a1043b7c244d128e8595cf0c1d10743b022b03a02dffd8ca3bf729f5a82604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1600190505b919050565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561151b57600080fd5b600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60078054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561162b5780601f106116005761010080835404028352916020019161162b565b820191906000526020600020905b81548152906001019060200180831161160e57829003601f168201915b505050505081565b600b5481565b60056020528060005260406000206000915054906101000a900460ff1681565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415151561169657600080fd5b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482111515156116e357600080fd5b611734826000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611dd190919063ffffffff16565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506117c7826000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611dea90919063ffffffff16565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a36001905092915050565b600061190982600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611dea90919063ffffffff16565b600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546040518082815260200191505060405180910390a36001905092915050565b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611b5957600080fd5b600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615611bb45760009050611c74565b6001600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507fd1bba68c128cc3f427e5831b3c6f99f480b6efa6b9e80c757768f6124158cc3f82604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1600190505b919050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611cd557600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515611d1157600080fd5b8073ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000828211151515611ddf57fe5b818303905092915050565b6000808284019050838110151515611dfe57fe5b8091505092915050565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611e6657600080fd5b600360149054906101000a900460ff16151515611e8257600080fd5b600454611e9a83600154611dea90919063ffffffff16565b11151515611ea757600080fd5b611eb18383611eb9565b905092915050565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611f1757600080fd5b600360149054906101000a900460ff16151515611f3357600080fd5b611f4882600154611dea90919063ffffffff16565b600181905550611f9f826000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611dea90919063ffffffff16565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885836040518082815260200191505060405180910390a28273ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a360019050929150505600a165627a7a72305820a346d15aeb4c885ec4d7a531c4d9dcce7dd623320dc220a5aa458efa03efa79c0029`

// DeploySuperCoin deploys a new Ethereum contract, binding an instance of SuperCoin to it.
func DeploySuperCoin(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SuperCoin, error) {
	parsed, err := abi.JSON(strings.NewReader(SuperCoinABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SuperCoinBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SuperCoin{SuperCoinCaller: SuperCoinCaller{contract: contract}, SuperCoinTransactor: SuperCoinTransactor{contract: contract}, SuperCoinFilterer: SuperCoinFilterer{contract: contract}}, nil
}

// SuperCoin is an auto generated Go binding around an Ethereum contract.
type SuperCoin struct {
	SuperCoinCaller     // Read-only binding to the contract
	SuperCoinTransactor // Write-only binding to the contract
	SuperCoinFilterer   // Log filterer for contract events
}

// SuperCoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type SuperCoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SuperCoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SuperCoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SuperCoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SuperCoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SuperCoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SuperCoinSession struct {
	Contract     *SuperCoin        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SuperCoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SuperCoinCallerSession struct {
	Contract *SuperCoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SuperCoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SuperCoinTransactorSession struct {
	Contract     *SuperCoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SuperCoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type SuperCoinRaw struct {
	Contract *SuperCoin // Generic contract binding to access the raw methods on
}

// SuperCoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SuperCoinCallerRaw struct {
	Contract *SuperCoinCaller // Generic read-only contract binding to access the raw methods on
}

// SuperCoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SuperCoinTransactorRaw struct {
	Contract *SuperCoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSuperCoin creates a new instance of SuperCoin, bound to a specific deployed contract.
func NewSuperCoin(address common.Address, backend bind.ContractBackend) (*SuperCoin, error) {
	contract, err := bindSuperCoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SuperCoin{SuperCoinCaller: SuperCoinCaller{contract: contract}, SuperCoinTransactor: SuperCoinTransactor{contract: contract}, SuperCoinFilterer: SuperCoinFilterer{contract: contract}}, nil
}

// NewSuperCoinCaller creates a new read-only instance of SuperCoin, bound to a specific deployed contract.
func NewSuperCoinCaller(address common.Address, caller bind.ContractCaller) (*SuperCoinCaller, error) {
	contract, err := bindSuperCoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SuperCoinCaller{contract: contract}, nil
}

// NewSuperCoinTransactor creates a new write-only instance of SuperCoin, bound to a specific deployed contract.
func NewSuperCoinTransactor(address common.Address, transactor bind.ContractTransactor) (*SuperCoinTransactor, error) {
	contract, err := bindSuperCoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SuperCoinTransactor{contract: contract}, nil
}

// NewSuperCoinFilterer creates a new log filterer instance of SuperCoin, bound to a specific deployed contract.
func NewSuperCoinFilterer(address common.Address, filterer bind.ContractFilterer) (*SuperCoinFilterer, error) {
	contract, err := bindSuperCoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SuperCoinFilterer{contract: contract}, nil
}

// bindSuperCoin binds a generic wrapper to an already deployed contract.
func bindSuperCoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SuperCoinABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SuperCoin *SuperCoinRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SuperCoin.Contract.SuperCoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SuperCoin *SuperCoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SuperCoin.Contract.SuperCoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SuperCoin *SuperCoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SuperCoin.Contract.SuperCoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SuperCoin *SuperCoinCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SuperCoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SuperCoin *SuperCoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SuperCoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SuperCoin *SuperCoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SuperCoin.Contract.contract.Transact(opts, method, params...)
}

// CAPSUPPLY is a free data retrieval call binding the contract method 0x070c0c4e.
//
// Solidity: function CAP_SUPPLY() constant returns(uint256)
func (_SuperCoin *SuperCoinCaller) CAPSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SuperCoin.contract.Call(opts, out, "CAP_SUPPLY")
	return *ret0, err
}

// CAPSUPPLY is a free data retrieval call binding the contract method 0x070c0c4e.
//
// Solidity: function CAP_SUPPLY() constant returns(uint256)
func (_SuperCoin *SuperCoinSession) CAPSUPPLY() (*big.Int, error) {
	return _SuperCoin.Contract.CAPSUPPLY(&_SuperCoin.CallOpts)
}

// CAPSUPPLY is a free data retrieval call binding the contract method 0x070c0c4e.
//
// Solidity: function CAP_SUPPLY() constant returns(uint256)
func (_SuperCoin *SuperCoinCallerSession) CAPSUPPLY() (*big.Int, error) {
	return _SuperCoin.Contract.CAPSUPPLY(&_SuperCoin.CallOpts)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_SuperCoin *SuperCoinCaller) INITIALSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SuperCoin.contract.Call(opts, out, "INITIAL_SUPPLY")
	return *ret0, err
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_SuperCoin *SuperCoinSession) INITIALSUPPLY() (*big.Int, error) {
	return _SuperCoin.Contract.INITIALSUPPLY(&_SuperCoin.CallOpts)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() constant returns(uint256)
func (_SuperCoin *SuperCoinCallerSession) INITIALSUPPLY() (*big.Int, error) {
	return _SuperCoin.Contract.INITIALSUPPLY(&_SuperCoin.CallOpts)
}

// ONCEMAXMINT is a free data retrieval call binding the contract method 0x9aabb18c.
//
// Solidity: function ONCE_MAX_MINT() constant returns(uint256)
func (_SuperCoin *SuperCoinCaller) ONCEMAXMINT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SuperCoin.contract.Call(opts, out, "ONCE_MAX_MINT")
	return *ret0, err
}

// ONCEMAXMINT is a free data retrieval call binding the contract method 0x9aabb18c.
//
// Solidity: function ONCE_MAX_MINT() constant returns(uint256)
func (_SuperCoin *SuperCoinSession) ONCEMAXMINT() (*big.Int, error) {
	return _SuperCoin.Contract.ONCEMAXMINT(&_SuperCoin.CallOpts)
}

// ONCEMAXMINT is a free data retrieval call binding the contract method 0x9aabb18c.
//
// Solidity: function ONCE_MAX_MINT() constant returns(uint256)
func (_SuperCoin *SuperCoinCallerSession) ONCEMAXMINT() (*big.Int, error) {
	return _SuperCoin.Contract.ONCEMAXMINT(&_SuperCoin.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_SuperCoin *SuperCoinCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SuperCoin.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_SuperCoin *SuperCoinSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _SuperCoin.Contract.Allowance(&_SuperCoin.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_SuperCoin *SuperCoinCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _SuperCoin.Contract.Allowance(&_SuperCoin.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_SuperCoin *SuperCoinCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SuperCoin.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_SuperCoin *SuperCoinSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _SuperCoin.Contract.BalanceOf(&_SuperCoin.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_SuperCoin *SuperCoinCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _SuperCoin.Contract.BalanceOf(&_SuperCoin.CallOpts, _owner)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_SuperCoin *SuperCoinCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SuperCoin.contract.Call(opts, out, "cap")
	return *ret0, err
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_SuperCoin *SuperCoinSession) Cap() (*big.Int, error) {
	return _SuperCoin.Contract.Cap(&_SuperCoin.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() constant returns(uint256)
func (_SuperCoin *SuperCoinCallerSession) Cap() (*big.Int, error) {
	return _SuperCoin.Contract.Cap(&_SuperCoin.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_SuperCoin *SuperCoinCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _SuperCoin.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_SuperCoin *SuperCoinSession) Decimals() (uint8, error) {
	return _SuperCoin.Contract.Decimals(&_SuperCoin.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_SuperCoin *SuperCoinCallerSession) Decimals() (uint8, error) {
	return _SuperCoin.Contract.Decimals(&_SuperCoin.CallOpts)
}

// ExistedOnWhitelist is a free data retrieval call binding the contract method 0x8ab905d4.
//
// Solidity: function existedOnWhitelist(addr address) constant returns(bool)
func (_SuperCoin *SuperCoinCaller) ExistedOnWhitelist(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SuperCoin.contract.Call(opts, out, "existedOnWhitelist", addr)
	return *ret0, err
}

// ExistedOnWhitelist is a free data retrieval call binding the contract method 0x8ab905d4.
//
// Solidity: function existedOnWhitelist(addr address) constant returns(bool)
func (_SuperCoin *SuperCoinSession) ExistedOnWhitelist(addr common.Address) (bool, error) {
	return _SuperCoin.Contract.ExistedOnWhitelist(&_SuperCoin.CallOpts, addr)
}

// ExistedOnWhitelist is a free data retrieval call binding the contract method 0x8ab905d4.
//
// Solidity: function existedOnWhitelist(addr address) constant returns(bool)
func (_SuperCoin *SuperCoinCallerSession) ExistedOnWhitelist(addr common.Address) (bool, error) {
	return _SuperCoin.Contract.ExistedOnWhitelist(&_SuperCoin.CallOpts, addr)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_SuperCoin *SuperCoinCaller) MintingFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SuperCoin.contract.Call(opts, out, "mintingFinished")
	return *ret0, err
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_SuperCoin *SuperCoinSession) MintingFinished() (bool, error) {
	return _SuperCoin.Contract.MintingFinished(&_SuperCoin.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_SuperCoin *SuperCoinCallerSession) MintingFinished() (bool, error) {
	return _SuperCoin.Contract.MintingFinished(&_SuperCoin.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SuperCoin *SuperCoinCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SuperCoin.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SuperCoin *SuperCoinSession) Name() (string, error) {
	return _SuperCoin.Contract.Name(&_SuperCoin.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SuperCoin *SuperCoinCallerSession) Name() (string, error) {
	return _SuperCoin.Contract.Name(&_SuperCoin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SuperCoin *SuperCoinCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SuperCoin.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SuperCoin *SuperCoinSession) Owner() (common.Address, error) {
	return _SuperCoin.Contract.Owner(&_SuperCoin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SuperCoin *SuperCoinCallerSession) Owner() (common.Address, error) {
	return _SuperCoin.Contract.Owner(&_SuperCoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SuperCoin *SuperCoinCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SuperCoin.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SuperCoin *SuperCoinSession) Symbol() (string, error) {
	return _SuperCoin.Contract.Symbol(&_SuperCoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SuperCoin *SuperCoinCallerSession) Symbol() (string, error) {
	return _SuperCoin.Contract.Symbol(&_SuperCoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SuperCoin *SuperCoinCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SuperCoin.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SuperCoin *SuperCoinSession) TotalSupply() (*big.Int, error) {
	return _SuperCoin.Contract.TotalSupply(&_SuperCoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SuperCoin *SuperCoinCallerSession) TotalSupply() (*big.Int, error) {
	return _SuperCoin.Contract.TotalSupply(&_SuperCoin.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_SuperCoin *SuperCoinCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SuperCoin.contract.Call(opts, out, "whitelist", arg0)
	return *ret0, err
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_SuperCoin *SuperCoinSession) Whitelist(arg0 common.Address) (bool, error) {
	return _SuperCoin.Contract.Whitelist(&_SuperCoin.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist( address) constant returns(bool)
func (_SuperCoin *SuperCoinCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _SuperCoin.Contract.Whitelist(&_SuperCoin.CallOpts, arg0)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(addr address) returns(bool)
func (_SuperCoin *SuperCoinTransactor) AddToWhitelist(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _SuperCoin.contract.Transact(opts, "addToWhitelist", addr)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(addr address) returns(bool)
func (_SuperCoin *SuperCoinSession) AddToWhitelist(addr common.Address) (*types.Transaction, error) {
	return _SuperCoin.Contract.AddToWhitelist(&_SuperCoin.TransactOpts, addr)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(addr address) returns(bool)
func (_SuperCoin *SuperCoinTransactorSession) AddToWhitelist(addr common.Address) (*types.Transaction, error) {
	return _SuperCoin.Contract.AddToWhitelist(&_SuperCoin.TransactOpts, addr)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_SuperCoin *SuperCoinTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SuperCoin.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_SuperCoin *SuperCoinSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SuperCoin.Contract.Approve(&_SuperCoin.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_SuperCoin *SuperCoinTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SuperCoin.Contract.Approve(&_SuperCoin.TransactOpts, _spender, _value)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_SuperCoin *SuperCoinTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _SuperCoin.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_SuperCoin *SuperCoinSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _SuperCoin.Contract.DecreaseApproval(&_SuperCoin.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_SuperCoin *SuperCoinTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _SuperCoin.Contract.DecreaseApproval(&_SuperCoin.TransactOpts, _spender, _subtractedValue)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_SuperCoin *SuperCoinTransactor) FinishMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SuperCoin.contract.Transact(opts, "finishMinting")
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_SuperCoin *SuperCoinSession) FinishMinting() (*types.Transaction, error) {
	return _SuperCoin.Contract.FinishMinting(&_SuperCoin.TransactOpts)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_SuperCoin *SuperCoinTransactorSession) FinishMinting() (*types.Transaction, error) {
	return _SuperCoin.Contract.FinishMinting(&_SuperCoin.TransactOpts)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_SuperCoin *SuperCoinTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _SuperCoin.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_SuperCoin *SuperCoinSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _SuperCoin.Contract.IncreaseApproval(&_SuperCoin.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_SuperCoin *SuperCoinTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _SuperCoin.Contract.IncreaseApproval(&_SuperCoin.TransactOpts, _spender, _addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(to address, amount uint256) returns(bool)
func (_SuperCoin *SuperCoinTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SuperCoin.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(to address, amount uint256) returns(bool)
func (_SuperCoin *SuperCoinSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SuperCoin.Contract.Mint(&_SuperCoin.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(to address, amount uint256) returns(bool)
func (_SuperCoin *SuperCoinTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SuperCoin.Contract.Mint(&_SuperCoin.TransactOpts, to, amount)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(addr address) returns(bool)
func (_SuperCoin *SuperCoinTransactor) RemoveFromWhitelist(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _SuperCoin.contract.Transact(opts, "removeFromWhitelist", addr)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(addr address) returns(bool)
func (_SuperCoin *SuperCoinSession) RemoveFromWhitelist(addr common.Address) (*types.Transaction, error) {
	return _SuperCoin.Contract.RemoveFromWhitelist(&_SuperCoin.TransactOpts, addr)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(addr address) returns(bool)
func (_SuperCoin *SuperCoinTransactorSession) RemoveFromWhitelist(addr common.Address) (*types.Transaction, error) {
	return _SuperCoin.Contract.RemoveFromWhitelist(&_SuperCoin.TransactOpts, addr)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_SuperCoin *SuperCoinTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SuperCoin.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_SuperCoin *SuperCoinSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SuperCoin.Contract.Transfer(&_SuperCoin.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_SuperCoin *SuperCoinTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SuperCoin.Contract.Transfer(&_SuperCoin.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_SuperCoin *SuperCoinTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SuperCoin.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_SuperCoin *SuperCoinSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SuperCoin.Contract.TransferFrom(&_SuperCoin.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_SuperCoin *SuperCoinTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SuperCoin.Contract.TransferFrom(&_SuperCoin.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SuperCoin *SuperCoinTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SuperCoin.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SuperCoin *SuperCoinSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SuperCoin.Contract.TransferOwnership(&_SuperCoin.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SuperCoin *SuperCoinTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SuperCoin.Contract.TransferOwnership(&_SuperCoin.TransactOpts, newOwner)
}

// SuperCoinApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SuperCoin contract.
type SuperCoinApprovalIterator struct {
	Event *SuperCoinApproval // Event containing the contract specifics and raw log

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
func (it *SuperCoinApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperCoinApproval)
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
		it.Event = new(SuperCoinApproval)
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
func (it *SuperCoinApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperCoinApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperCoinApproval represents a Approval event raised by the SuperCoin contract.
type SuperCoinApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_SuperCoin *SuperCoinFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SuperCoinApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SuperCoin.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SuperCoinApprovalIterator{contract: _SuperCoin.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_SuperCoin *SuperCoinFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SuperCoinApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SuperCoin.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperCoinApproval)
				if err := _SuperCoin.contract.UnpackLog(event, "Approval", log); err != nil {
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

// SuperCoinMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the SuperCoin contract.
type SuperCoinMintIterator struct {
	Event *SuperCoinMint // Event containing the contract specifics and raw log

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
func (it *SuperCoinMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperCoinMint)
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
		it.Event = new(SuperCoinMint)
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
func (it *SuperCoinMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperCoinMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperCoinMint represents a Mint event raised by the SuperCoin contract.
type SuperCoinMint struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(to indexed address, amount uint256)
func (_SuperCoin *SuperCoinFilterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*SuperCoinMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SuperCoin.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &SuperCoinMintIterator{contract: _SuperCoin.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(to indexed address, amount uint256)
func (_SuperCoin *SuperCoinFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *SuperCoinMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SuperCoin.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperCoinMint)
				if err := _SuperCoin.contract.UnpackLog(event, "Mint", log); err != nil {
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

// SuperCoinMintFinishedIterator is returned from FilterMintFinished and is used to iterate over the raw logs and unpacked data for MintFinished events raised by the SuperCoin contract.
type SuperCoinMintFinishedIterator struct {
	Event *SuperCoinMintFinished // Event containing the contract specifics and raw log

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
func (it *SuperCoinMintFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperCoinMintFinished)
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
		it.Event = new(SuperCoinMintFinished)
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
func (it *SuperCoinMintFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperCoinMintFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperCoinMintFinished represents a MintFinished event raised by the SuperCoin contract.
type SuperCoinMintFinished struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMintFinished is a free log retrieval operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: event MintFinished()
func (_SuperCoin *SuperCoinFilterer) FilterMintFinished(opts *bind.FilterOpts) (*SuperCoinMintFinishedIterator, error) {

	logs, sub, err := _SuperCoin.contract.FilterLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return &SuperCoinMintFinishedIterator{contract: _SuperCoin.contract, event: "MintFinished", logs: logs, sub: sub}, nil
}

// WatchMintFinished is a free log subscription operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: event MintFinished()
func (_SuperCoin *SuperCoinFilterer) WatchMintFinished(opts *bind.WatchOpts, sink chan<- *SuperCoinMintFinished) (event.Subscription, error) {

	logs, sub, err := _SuperCoin.contract.WatchLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperCoinMintFinished)
				if err := _SuperCoin.contract.UnpackLog(event, "MintFinished", log); err != nil {
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

// SuperCoinOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SuperCoin contract.
type SuperCoinOwnershipTransferredIterator struct {
	Event *SuperCoinOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SuperCoinOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperCoinOwnershipTransferred)
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
		it.Event = new(SuperCoinOwnershipTransferred)
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
func (it *SuperCoinOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperCoinOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperCoinOwnershipTransferred represents a OwnershipTransferred event raised by the SuperCoin contract.
type SuperCoinOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_SuperCoin *SuperCoinFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SuperCoinOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SuperCoin.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SuperCoinOwnershipTransferredIterator{contract: _SuperCoin.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_SuperCoin *SuperCoinFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SuperCoinOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SuperCoin.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperCoinOwnershipTransferred)
				if err := _SuperCoin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// SuperCoinTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SuperCoin contract.
type SuperCoinTransferIterator struct {
	Event *SuperCoinTransfer // Event containing the contract specifics and raw log

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
func (it *SuperCoinTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperCoinTransfer)
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
		it.Event = new(SuperCoinTransfer)
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
func (it *SuperCoinTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperCoinTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperCoinTransfer represents a Transfer event raised by the SuperCoin contract.
type SuperCoinTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_SuperCoin *SuperCoinFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SuperCoinTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SuperCoin.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SuperCoinTransferIterator{contract: _SuperCoin.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_SuperCoin *SuperCoinFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SuperCoinTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SuperCoin.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperCoinTransfer)
				if err := _SuperCoin.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// SuperCoinWhitelistedAddressAddedIterator is returned from FilterWhitelistedAddressAdded and is used to iterate over the raw logs and unpacked data for WhitelistedAddressAdded events raised by the SuperCoin contract.
type SuperCoinWhitelistedAddressAddedIterator struct {
	Event *SuperCoinWhitelistedAddressAdded // Event containing the contract specifics and raw log

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
func (it *SuperCoinWhitelistedAddressAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperCoinWhitelistedAddressAdded)
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
		it.Event = new(SuperCoinWhitelistedAddressAdded)
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
func (it *SuperCoinWhitelistedAddressAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperCoinWhitelistedAddressAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperCoinWhitelistedAddressAdded represents a WhitelistedAddressAdded event raised by the SuperCoin contract.
type SuperCoinWhitelistedAddressAdded struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedAddressAdded is a free log retrieval operation binding the contract event 0xd1bba68c128cc3f427e5831b3c6f99f480b6efa6b9e80c757768f6124158cc3f.
//
// Solidity: event WhitelistedAddressAdded(addr address)
func (_SuperCoin *SuperCoinFilterer) FilterWhitelistedAddressAdded(opts *bind.FilterOpts) (*SuperCoinWhitelistedAddressAddedIterator, error) {

	logs, sub, err := _SuperCoin.contract.FilterLogs(opts, "WhitelistedAddressAdded")
	if err != nil {
		return nil, err
	}
	return &SuperCoinWhitelistedAddressAddedIterator{contract: _SuperCoin.contract, event: "WhitelistedAddressAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistedAddressAdded is a free log subscription operation binding the contract event 0xd1bba68c128cc3f427e5831b3c6f99f480b6efa6b9e80c757768f6124158cc3f.
//
// Solidity: event WhitelistedAddressAdded(addr address)
func (_SuperCoin *SuperCoinFilterer) WatchWhitelistedAddressAdded(opts *bind.WatchOpts, sink chan<- *SuperCoinWhitelistedAddressAdded) (event.Subscription, error) {

	logs, sub, err := _SuperCoin.contract.WatchLogs(opts, "WhitelistedAddressAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperCoinWhitelistedAddressAdded)
				if err := _SuperCoin.contract.UnpackLog(event, "WhitelistedAddressAdded", log); err != nil {
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

// SuperCoinWhitelistedAddressRemovedIterator is returned from FilterWhitelistedAddressRemoved and is used to iterate over the raw logs and unpacked data for WhitelistedAddressRemoved events raised by the SuperCoin contract.
type SuperCoinWhitelistedAddressRemovedIterator struct {
	Event *SuperCoinWhitelistedAddressRemoved // Event containing the contract specifics and raw log

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
func (it *SuperCoinWhitelistedAddressRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuperCoinWhitelistedAddressRemoved)
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
		it.Event = new(SuperCoinWhitelistedAddressRemoved)
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
func (it *SuperCoinWhitelistedAddressRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuperCoinWhitelistedAddressRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuperCoinWhitelistedAddressRemoved represents a WhitelistedAddressRemoved event raised by the SuperCoin contract.
type SuperCoinWhitelistedAddressRemoved struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedAddressRemoved is a free log retrieval operation binding the contract event 0xf1abf01a1043b7c244d128e8595cf0c1d10743b022b03a02dffd8ca3bf729f5a.
//
// Solidity: event WhitelistedAddressRemoved(addr address)
func (_SuperCoin *SuperCoinFilterer) FilterWhitelistedAddressRemoved(opts *bind.FilterOpts) (*SuperCoinWhitelistedAddressRemovedIterator, error) {

	logs, sub, err := _SuperCoin.contract.FilterLogs(opts, "WhitelistedAddressRemoved")
	if err != nil {
		return nil, err
	}
	return &SuperCoinWhitelistedAddressRemovedIterator{contract: _SuperCoin.contract, event: "WhitelistedAddressRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistedAddressRemoved is a free log subscription operation binding the contract event 0xf1abf01a1043b7c244d128e8595cf0c1d10743b022b03a02dffd8ca3bf729f5a.
//
// Solidity: event WhitelistedAddressRemoved(addr address)
func (_SuperCoin *SuperCoinFilterer) WatchWhitelistedAddressRemoved(opts *bind.WatchOpts, sink chan<- *SuperCoinWhitelistedAddressRemoved) (event.Subscription, error) {

	logs, sub, err := _SuperCoin.contract.WatchLogs(opts, "WhitelistedAddressRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuperCoinWhitelistedAddressRemoved)
				if err := _SuperCoin.contract.UnpackLog(event, "WhitelistedAddressRemoved", log); err != nil {
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

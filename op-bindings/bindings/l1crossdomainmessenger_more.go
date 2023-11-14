// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L1CrossDomainMessengerStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)1012_storage\"},{\"astId\":1003,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_paused\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_bool\"},{\"astId\":1004,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)1011_storage\"},{\"astId\":1005,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"successfulMessages\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1006,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"xDomainMsgSender\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_address\"},{\"astId\":1007,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"msgNonce\",\"offset\":0,\"slot\":\"103\",\"type\":\"t_uint240\"},{\"astId\":1008,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"failedMessages\",\"offset\":0,\"slot\":\"104\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1009,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"105\",\"type\":\"t_array(t_uint256)1010_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)1010_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[45]\",\"numberOfBytes\":\"1440\"},\"t_array(t_uint256)1011_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)1012_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_uint240\":{\"encoding\":\"inplace\",\"label\":\"uint240\",\"numberOfBytes\":\"30\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L1CrossDomainMessengerStorageLayout = new(solc.StorageLayout)

var L1CrossDomainMessengerDeployedBin = "0x60806040526004361061015f5760003560e01c80636e296e45116100c0578063a4e7f8bd11610074578063b28ade2511610059578063b28ade251461039e578063d764ad0b146103be578063ecc70428146103d157600080fd5b8063a4e7f8bd1461033e578063b1b1b2091461036e57600080fd5b806383a74074116100a557806383a74074146102f35780638cbeeef2146102575780639fce812c1461030a57600080fd5b80636e296e45146102c95780638129fc1c146102de57600080fd5b80633f827a5a1161011757806354fd4d50116100fc57806354fd4d501461026d5780635644cfdf1461028f5780635c975abb146102a557600080fd5b80633f827a5a1461022f5780634c1d6a691461025757600080fd5b80630ff754ea116101485780630ff754ea146101ac5780632828d7e8146102055780633dbb202b1461021a57600080fd5b8063028f85f7146101645780630c56849814610197575b600080fd5b34801561017057600080fd5b50610179601081565b60405167ffffffffffffffff90911681526020015b60405180910390f35b3480156101a357600080fd5b50610179603f81565b3480156101b857600080fd5b506101e07f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161018e565b34801561021157600080fd5b50610179604081565b61022d610228366004611772565b61040f565b005b34801561023b57600080fd5b50610244600081565b60405161ffff909116815260200161018e565b34801561026357600080fd5b50610179619c4081565b34801561027957600080fd5b506102826105fd565b60405161018e9190611853565b34801561029b57600080fd5b5061017961138881565b3480156102b157600080fd5b5060335460ff165b604051901515815260200161018e565b3480156102d557600080fd5b506101e06106a0565b3480156102ea57600080fd5b5061022d61078c565b3480156102ff57600080fd5b5061017962030d4081565b34801561031657600080fd5b506101e07f000000000000000000000000000000000000000000000000000000000000000081565b34801561034a57600080fd5b506102b961035936600461186d565b60686020526000908152604090205460ff1681565b34801561037a57600080fd5b506102b961038936600461186d565b60656020526000908152604090205460ff1681565b3480156103aa57600080fd5b506101796103b9366004611886565b61091e565b61022d6103cc3660046118da565b61098c565b3480156103dd57600080fd5b506067547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1660405190815260200161018e565b6105277f000000000000000000000000000000000000000000000000000000000000000061043e85858561091e565b347fd764ad0b000000000000000000000000000000000000000000000000000000006104896067547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690565b338a34898c8c6040516024016104a597969594939291906119a9565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915261113f565b3373ffffffffffffffffffffffffffffffffffffffff85167fdd28cef75ff18fb538e43317144469f339702f973eace2bc808f2acc37db310e34868661058c6067547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1690565b8760405161059e959493929190611a08565b60405180910390a35050606780547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff0000000000000000000000000000000000000000000000000000000000009091161790555050565b60606106287f00000000000000000000000000000000000000000000000000000000000000006111f4565b6106517f00000000000000000000000000000000000000000000000000000000000000006111f4565b61067a7f00000000000000000000000000000000000000000000000000000000000000006111f4565b60405160200161068c93929190611a40565b604051602081830303815290604052905090565b60665460009073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff21530161076f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f43726f7373446f6d61696e4d657373656e6765723a2078446f6d61696e4d657360448201527f7361676553656e646572206973206e6f7420736574000000000000000000000060648201526084015b60405180910390fd5b5060665473ffffffffffffffffffffffffffffffffffffffff1690565b600054610100900460ff16158080156107ac5750600054600160ff909116105b806107c65750303b1580156107c6575060005460ff166001145b610852576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610766565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905580156108b057600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b6108b86112b2565b801561091b57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b6000611388619c4080603f61093a604063ffffffff8816611ae5565b6109449190611b15565b61094f601088611ae5565b61095c9062030d40611b63565b6109669190611b63565b6109709190611b63565b61097a9190611b63565b6109849190611b63565b949350505050565b60f087901c60018110610a47576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604760248201527f43726f7373446f6d61696e4d657373656e6765723a206f6e6c7920766572736960448201527f6f6e2030206d6573736167657320697320737570706f7274656420617420746860648201527f69732074696d6500000000000000000000000000000000000000000000000000608482015260a401610766565b6000610a8d898989898989898080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061137792505050565b9050610a9761139a565b15610acf57853414610aab57610aab611b8f565b60008181526068602052604090205460ff1615610aca57610aca611b8f565b610c21565b3415610b83576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605060248201527f43726f7373446f6d61696e4d657373656e6765723a2076616c7565206d75737460448201527f206265207a65726f20756e6c657373206d6573736167652069732066726f6d2060648201527f612073797374656d206164647265737300000000000000000000000000000000608482015260a401610766565b60008181526068602052604090205460ff16610c21576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520636160448201527f6e6e6f74206265207265706c61796564000000000000000000000000000000006064820152608401610766565b610c2a876114be565b15610cdd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f43726f7373446f6d61696e4d657373656e6765723a2063616e6e6f742073656e60448201527f64206d65737361676520746f20626c6f636b65642073797374656d206164647260648201527f6573730000000000000000000000000000000000000000000000000000000000608482015260a401610766565b60008181526065602052604090205460ff1615610d7c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520686160448201527f7320616c7265616479206265656e2072656c61796564000000000000000000006064820152608401610766565b610d9d85610d8e611388619c40611b63565b67ffffffffffffffff16611535565b1580610dc3575060665473ffffffffffffffffffffffffffffffffffffffff1661dead14155b15610edc5760008181526068602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555182917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff3201610ed5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d657373616765000000000000000000000000000000000000006064820152608401610766565b5050611115565b606680547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8a161790556000610f6d88619c405a610f309190611bbe565b8988888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061155392505050565b606680547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055905080156110045760008281526065602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c91a2611111565b60008281526068602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff3201611111576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d657373616765000000000000000000000000000000000000006064820152608401610766565b5050505b50505050505050565b905090565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b6040517fe9e05c4200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063e9e05c429084906111bc908890839089906000908990600401611bd5565b6000604051808303818588803b1580156111d557600080fd5b505af11580156111e9573d6000803e3d6000fd5b505050505050505050565b606060006112018361156d565b600101905060008167ffffffffffffffff81111561122157611221611c2d565b6040519080825280601f01601f19166020018201604052801561124b576020820181803683370190505b5090508181016020015b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a850494508461125557509392505050565b600054610100900460ff16611349576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610766565b606680547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055565b600061138787878787878761164f565b8051906020012090509695505050505050565b60003373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614801561111e57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16639bf62d826040518163ffffffff1660e01b8152600401602060405180830381865afa15801561147e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114a29190611c5c565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b600073ffffffffffffffffffffffffffffffffffffffff821630148061152f57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b92915050565b600080603f83619c4001026040850201603f5a021015949350505050565b600080600080845160208601878a8af19695505050505050565b6000807a184f03e93ff9f4daa797ed6e38ed64bf6a1f01000000000000000083106115b6577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef810000000083106115e2576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061160057662386f26fc10000830492506010015b6305f5e1008310611618576305f5e100830492506008015b612710831061162c57612710830492506004015b6064831061163e576064830492506002015b600a831061152f5760010192915050565b606086868686868660405160240161166c96959493929190611c79565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fd764ad0b0000000000000000000000000000000000000000000000000000000017905290509695505050505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461091b57600080fd5b60008083601f84011261172257600080fd5b50813567ffffffffffffffff81111561173a57600080fd5b60208301915083602082850101111561175257600080fd5b9250929050565b803563ffffffff8116811461176d57600080fd5b919050565b6000806000806060858703121561178857600080fd5b8435611793816116ee565b9350602085013567ffffffffffffffff8111156117af57600080fd5b6117bb87828801611710565b90945092506117ce905060408601611759565b905092959194509250565b60005b838110156117f45781810151838201526020016117dc565b83811115611803576000848401525b50505050565b600081518084526118218160208601602086016117d9565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006118666020830184611809565b9392505050565b60006020828403121561187f57600080fd5b5035919050565b60008060006040848603121561189b57600080fd5b833567ffffffffffffffff8111156118b257600080fd5b6118be86828701611710565b90945092506118d1905060208501611759565b90509250925092565b600080600080600080600060c0888a0312156118f557600080fd5b873596506020880135611907816116ee565b95506040880135611917816116ee565b9450606088013593506080880135925060a088013567ffffffffffffffff81111561194157600080fd5b61194d8a828b01611710565b989b979a50959850939692959293505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b878152600073ffffffffffffffffffffffffffffffffffffffff808916602084015280881660408401525085606083015263ffffffff8516608083015260c060a08301526119fb60c083018486611960565b9998505050505050505050565b858152608060208201526000611a22608083018688611960565b905083604083015263ffffffff831660608301529695505050505050565b60008451611a528184602089016117d9565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551611a8e816001850160208a016117d9565b60019201918201528351611aa98160028401602088016117d9565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff80831681851681830481118215151615611b0c57611b0c611ab6565b02949350505050565b600067ffffffffffffffff80841680611b57577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b92169190910492915050565b600067ffffffffffffffff808316818516808303821115611b8657611b86611ab6565b01949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b600082821015611bd057611bd0611ab6565b500390565b73ffffffffffffffffffffffffffffffffffffffff8616815284602082015267ffffffffffffffff84166040820152821515606082015260a060808201526000611c2260a0830184611809565b979650505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600060208284031215611c6e57600080fd5b8151611866816116ee565b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a0830152611cc460c0830184611809565b9897505050505050505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L1CrossDomainMessengerStorageLayoutJSON), L1CrossDomainMessengerStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1CrossDomainMessenger"] = L1CrossDomainMessengerStorageLayout
	deployedBytecodes["L1CrossDomainMessenger"] = L1CrossDomainMessengerDeployedBin
}

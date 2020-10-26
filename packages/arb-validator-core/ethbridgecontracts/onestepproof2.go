// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethbridgecontracts

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// OneStepProof2ABI is the input ABI used to generate the binding from.
const OneStepProof2ABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"inboxAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"messagesAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"logsAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"bproof\",\"type\":\"bytes\"}],\"name\":\"executeStep\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"gas\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[5]\",\"name\":\"fields\",\"type\":\"bytes32[5]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// OneStepProof2FuncSigs maps the 4-byte function signature to its string representation.
var OneStepProof2FuncSigs = map[string]string{
	"1041c884": "executeStep(bytes32,bytes32,bytes32,bytes,bytes)",
}

// OneStepProof2Bin is the compiled bytecode used for deploying new contracts.
var OneStepProof2Bin = "0x608060405234801561001057600080fd5b5061312d806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80631041c88414610030575b600080fd5b610105600480360360a081101561004657600080fd5b8135916020810135916040820135919081019060808101606082013564010000000081111561007457600080fd5b82018360208201111561008657600080fd5b803590602001918460018302840111640100000000831117156100a857600080fd5b9193909290916020810190356401000000008111156100c657600080fd5b8201836020820111156100d857600080fd5b803590602001918460018302840111640100000000831117156100fa57600080fd5b50909250905061014e565b60405167ffffffffffffffff83168152602081018260a080838360005b8381101561013a578181015183820152602001610122565b505050509050019250505060405180910390f35b6000610158612f14565b610160612f32565b6101d68a8a8a8a8a8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8e018190048102820181019092528c815292508c91508b90819084018382808284376000920191909152506101fb92505050565b90506101e1816106c1565b6101ea81610a29565b925092505097509795505050505050565b610203612f32565b60008360008151811061021257fe5b602001015160f81c60f81b60f81c905060008460018151811061023157fe5b602001015160f81c60f81b60f81c905060006002905060608360040160ff1660405190808252806020026020018201604052801561028957816020015b610276612fcb565b81526020019060019003908161026e5790505b50905060608360040160ff166040519080825280602002602001820160405280156102ce57816020015b6102bb612fcb565b8152602001906001900390816102b35790505b50905060005b8560ff1681101561030c576102e98985610a8c565b84518590849081106102f757fe5b602090810291909101015293506001016102d4565b5060005b8460ff16811015610348576103258985610a8c565b835184908490811061033357fe5b60209081029190910101529350600101610310565b50610351613008565b61035b8985610c4e565b9094509050610368612f32565b604051806101e0016040528083815260200161038384610cff565b81526020018e81526020018d81526020018c8152602001600067ffffffffffffffff1681526020016103b3610d74565b81526020016000801b815260200160405180604001604052808a60ff16815260200187815250815260200160405180604001604052808960ff1681526020018681525081526020018b878151811061040757fe5b602001015160f81c60f81b60f81c60ff16600114151581526020018b876001018151811061043157fe5b0160209081015160f81c825281018c90526002870160408201526060018a90528a519091506000908b908790811061046557fe5b602001015160f81c60f81b60f81c905060008b876001018151811061048657fe5b01602001516002979097019660f81c905060ff821615806104aa57508160ff166001145b6040518060400160405280600b81526020016a04241445f494d4d5f5459560ac1b815250906105575760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561051c578181015183820152602001610504565b50505050905090810190601f1680156105495780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50610560612fcb565b60ff831661057d57835151610576908390610dbb565b905061061d565b6000875111604051806040016040528060068152602001654e4f5f494d4d60d01b815250906105ed5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561051c578181015183820152602001610504565b5061061a828560000151600001518960018e0360ff168151811061060d57fe5b6020026020010151610e1f565b90505b61062681610ea5565b84515260005b838b0360ff1681101561066b5761066388828151811061064857fe5b60200260200101518660000151610fe990919063ffffffff16565b60010161062c565b5060005b8960ff168110156106ac576106a487828151811061068957fe5b6020026020010151866000015161100390919063ffffffff16565b60010161066f565b50929f9e505050505050505050505050505050565b60008060006130736106da85610160015160ff1661101d565b93509350935093506106ec858361110b565b156106fa5750505050610a26565b610100850151518411156107af57610718610713610d74565b610ea5565b610729866020015160200151610ea5565b146040518060400160405280600d81526020016c535441434b5f4d495353494e4760981b8152509061079c5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561051c578181015183820152602001610504565b506107a685611177565b50505050610a26565b6101208501515183111561084a576107c8610713610d74565b6107d9866020015160400151610ea5565b146040518060400160405280600b81526020016a4155585f4d495353494e4760a81b8152509061079c5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561051c578181015183820152602001610504565b600084118061085c5750846101400151155b801561086d57506101008501515184145b8061089557508461014001518015610883575083155b80156108955750610100850151516001145b6040518060400160405280600a815260200169535441434b5f4d414e5960b01b815250906109045760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561051c578181015183820152602001610504565b50610120850151516040805180820190915260088152674155585f4d414e5960c01b60208201529084146109795760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561051c578181015183820152602001610504565b50610987858263ffffffff16565b60005b610100860151518110156109d3576109cb8661010001516020015182815181106109b057fe5b60200260200101518760200151610fe990919063ffffffff16565b60010161098a565b5060005b61012086015151811015610a2057610a188661012001516020015182815181106109fd57fe5b6020026020010151876020015161100390919063ffffffff16565b6001016109d7565b50505050505b50565b6000610a33612f14565b8260a001516040518060a00160405280610a5086600001516111e0565b8152602001610a6286602001516111e0565b81526020018560400151815260200185606001518152602001856080015181525091509150915091565b6000610a96612fcb565b83518310610adc576040805162461bcd60e51b815260206004820152600e60248201526d1a5b9d985b1a59081bd9999cd95d60921b604482015290519081900360640190fd5b600080610ae986866112ba565b91509150610af56112e1565b60ff168160ff161415610b29576000610b0e87846112e6565b909350905082610b1d8261135a565b94509450505050610c47565b610b31611413565b60ff168160ff161415610b5357610b488683611418565b935093505050610c47565b610b5b6114ba565b60ff168160ff161415610b83576000610b7487846112e6565b909350905082610b1d826114bf565b610b8b61157c565b60ff168160ff161415610ba257610b488683611581565b610baa611615565b60ff168160ff1610158015610bcb5750610bc261161a565b60ff168160ff16105b15610c07576000610bda611615565b820390506060610beb82898661161f565b909450905083610bfa826116b8565b9550955050505050610c47565b6040805162461bcd60e51b815260206004820152601060248201526f696e76616c69642074797065636f646560801b604482015290519081900360640190fd5b9250929050565b6000610c58613008565b610c60613008565b6000610100820181905280610c7587876112e6565b9096509150610c848787611581565b60208501529550610c958787611581565b60408501529550610ca68787610a8c565b60608501529550610cb78787610a8c565b60808501529550610cc887876112e6565b60a08501529550610cd987876112e6565b9096509050610ce88787610a8c565b60e085015291835260c08301529590945092505050565b610d07613008565b60405180610120016040528083600001518152602001836020015181526020018360400151815260200183606001518152602001836080015181526020018360a0015181526020018360c0015181526020018360e00151815260200183610100015181525090505b919050565b610d7c612fcb565b60408051600080825260208201909252610db691610db0565b610d9d612fcb565b815260200190600190039081610d955790505b506116b8565b905090565b610dc3612fcb565b6040805160608101825260ff851681526020808201859052825160008082529181018452610e1693830191610e0e565b610dfb612fcb565b815260200190600190039081610df35790505b5090526117d1565b90505b92915050565b610e27612fcb565b604080516001808252818301909252606091816020015b610e46612fcb565b815260200190600190039081610e3e5790505090508281600081518110610e6957fe5b6020026020010181905250610e9a60405180606001604052808760ff168152602001868152602001838152506117d1565b9150505b9392505050565b6000610eaf6112e1565b60ff16826080015160ff161415610ed2578151610ecb9061183f565b9050610d6f565b610eda611413565b60ff16826080015160ff161415610ef857610ecb8260200151611863565b610f0061157c565b60ff16826080015160ff161415610f2257815160a0830151610ecb9190611960565b610f2a611615565b60ff16826080015160ff161415610f6357610f43612fcb565b610f5083604001516119b1565b9050610f5b81610ea5565b915050610d6f565b610f6b611b13565b60ff16826080015160ff161415610f8457508051610d6f565b610f8c6114ba565b60ff16826080015160ff161415610fa857506060810151610d6f565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b610ff7826020015182611b18565b82602001819052505050565b611011826040015182611b18565b82604001819052505050565b6000808061307360a085141561104157506001925060009150829050611b96611104565b60a185141561105f57506002925060009150600a9050611c0e611104565b60a285141561107d57506002925060009150600a9050611caa611104565b60a385141561109b57506002925060009150600a9050611d25611104565b60a48514156110b957506003925060009150600a9050611da0611104565b60a58514156110d757506003925060009150600a9050611e6c611104565b60a68514156110f557506003925060009150600a9050611f15611104565b5060009250829150819050611fbe5b9193509193565b60a0808301805167ffffffffffffffff90840181169091526020840151909101516000918316111561115657602083015160001960a09091015261114e83611177565b506001610e19565b50602082015160a001805167ffffffffffffffff8316900390526000610e19565b60408051600160f81b6020808301919091526000602183018190526022808401919091528351808403909101815260429092019092528051908201209082015160c0015114156111d3576111ce8160200151611fc7565b610a26565b6020015160c08101519052565b6000600282610100015114156111f857506000610d6f565b6001826101000151141561120e57506001610d6f565b8151602083015161121e90610ea5565b61122b8460400151610ea5565b6112388560600151610ea5565b6112458660800151610ea5565b8660a001518760c0015161125c8960e00151610ea5565b6040516020018089815260200188815260200187815260200186815260200185815260200184815260200183815260200182815260200198505050505050505050604051602081830303815290604052805190602001209050610d6f565b600080826001018484815181106112cd57fe5b016020015190925060f81c90509250929050565b600090565b600080828451101580156112fe575060208385510310155b61133b576040805162461bcd60e51b81526020600482015260096024820152681d1bdbc81cda1bdc9d60ba1b604482015290519081900360640190fd5b6020830161134f858563ffffffff611fd216565b915091509250929050565b611362612fcb565b6040805160c08101825283815281516060810183526000808252602082810182905284518281528082018652939490850193908301916113b8565b6113a5612fcb565b81526020019060019003908161139d5790505b509052815260408051600080825260208281019093529190920191906113f4565b6113e1612fcb565b8152602001906001900390816113d95790505b5081526000602082018190526040820152600160609091015292915050565b600190565b6000611422612fcb565b8260008061142e612fcb565b600061143a89866112ba565b909550935061144989866112ba565b9095509250600160ff8516141561146a576114648986610a8c565b90955091505b611474898661202b565b9095509050600160ff8516141561149f5784611491848385610e1f565b965096505050505050610c47565b846114aa8483610dbb565b9650965050505050509250929050565b600c90565b6114c7612fcb565b6040805160c0810182526000808252825160608101845281815260208181018390528451838152808201865293949085019391929083019190611520565b61150d612fcb565b8152602001906001900390816115055790505b5090528152604080516000808252602082810190935291909201919061155c565b611549612fcb565b8152602001906001900390816115415790505b50815260208101849052600c604082015260016060909101529050919050565b600290565b600061158b612fcb565b828451101580156115a0575060408385510310155b6115dc576040805162461bcd60e51b81526020600482015260086024820152671d1bc81cda1bdc9d60c21b604482015290519081900360640190fd5b6000806115e9868661202b565b90945091506115f886856112e6565b9094509050836116088383612042565b9350935050509250929050565b600390565b600d90565b60006060600083905060608660ff1660405190808252806020026020018201604052801561166757816020015b611654612fcb565b81526020019060019003908161164c5790505b50905060005b8760ff168160ff1610156116ab576116858784610a8c565b8351849060ff851690811061169657fe5b6020908102919091010152925060010161166d565b5090969095509350505050565b6116c0612fcb565b6116ca82516120fa565b61171b576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b83518110156117525783818151811061173557fe5b602002602001015160a00151820191508080600101915050611720565b506040805160c08101825260008082528251606081018452818152602081810183905284518381528082018652939490850193919290830191906117ac565b611799612fcb565b8152602001906001900390816117915790505b5090528152602081019490945260006040850152600360608501526080909301525090565b6117d9612fcb565b6040805160c0810182526000808252602080830186905283518281529081018452919283019190611820565b61180d612fcb565b8152602001906001900390816118055790505b5081526000602082015260016040820181905260609091015292915050565b60408051602080820193909352815180820384018152908201909152805191012090565b600060028260400151511061187457fe5b6040820151516118d957611886611413565b8251602080850151604080516001600160f81b031960f896871b8116828601529490951b909316602185015260228085019190915282518085039091018152604290930190915281519101209050610d6f565b6118e1611413565b826000015161190784604001516000815181106118fa57fe5b6020026020010151610ea5565b8460200151604051602001808560ff1660ff1660f81b81526001018460ff1660ff1660f81b8152600101838152602001828152602001945050505050604051602081830303815290604052805190602001209050919050565b600061196a611615565b8383604051602001808460ff1660ff1660f81b8152600101838152602001828152602001935050505060405160208183030381529060405280519060200120905092915050565b6119b9612fcb565b600882511115611a07576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015611a34578160200160208202803883390190505b508051909150600160005b82811015611a9757611a568682815181106118fa57fe5b848281518110611a6257fe5b602002602001018181525050858181518110611a7a57fe5b602002602001015160a00151820191508080600101915050611a3f565b506000835184604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015611adc578181015183820152602001611ac4565b5050505090500192505050604051602081830303815290604052805190602001209050611b098183612042565b9695505050505050565b606490565b611b20612fcb565b6040805160028082526060828101909352816020015b611b3e612fcb565b815260200190600190039081611b365790505090508281600081518110611b6157fe5b60200260200101819052508381600181518110611b7a57fe5b6020026020010181905250611b8e816119b1565b949350505050565b611b9e612fcb565b611bac826101000151612101565b9050611bb781612143565b611bca57611bc48261214e565b50610a26565b611c0a826101000151611c056000801b60405160200180828152602001915050604051602081830303815290604052805190602001206114bf565b61216d565b5050565b611c16612fcb565b611c24826101000151612101565b9050611c2e612fcb565b611c3c836101000151612101565b9050611c4781612143565b1580611c595750611c5782612197565b155b15611c6e57611c678361214e565b5050610a26565b6000611c9083606001518360000151611c8b876101c001516121a4565b61229a565b9050611ca4846101000151611c058361135a565b50505050565b611cb2612fcb565b611cc0826101000151612101565b9050611cca612fcb565b611cd8836101000151612101565b9050611ce381612143565b1580611cf55750611cf382612197565b155b15611d0357611c678361214e565b6000611c9083606001518360000151611d20876101c001516121a4565b6122bc565b611d2d612fcb565b611d3b826101000151612101565b9050611d45612fcb565b611d53836101000151612101565b9050611d5e81612143565b1580611d705750611d6e82612197565b155b15611d7e57611c678361214e565b6000611c9083606001518360000151611d9b876101c001516121a4565b61241b565b611da8612fcb565b611db6826101000151612101565b9050611dc0612fcb565b611dce836101000151612101565b9050611dd8612fcb565b611de6846101000151612101565b9050611df182612143565b1580611e035750611e0181612143565b155b80611e145750611e1283612197565b155b15611e2a57611e228461214e565b505050610a26565b6000611e51846060015184600001518460000151611e4c896101c001516121a4565b61254e565b9050611e65856101000151611c05836114bf565b5050505050565b611e74612fcb565b611e82826101000151612101565b9050611e8c612fcb565b611e9a836101000151612101565b9050611ea4612fcb565b611eb2846101000151612101565b9050611ebd82612143565b1580611ecf5750611ecd81612143565b155b80611ee05750611ede83612197565b155b15611eee57611e228461214e565b6000611e51846060015184600001518460000151611f10896101c001516121a4565b612597565b611f1d612fcb565b611f2b826101000151612101565b9050611f35612fcb565b611f43836101000151612101565b9050611f4d612fcb565b611f5b846101000151612101565b9050611f6682612143565b1580611f785750611f7681612143565b155b80611f895750611f8783612197565b155b15611f9757611e228461214e565b6000611e51846060015184600001518460000151611fb9896101c001516121a4565b6126df565b610a268161214e565b600161010090910152565b60008160200183511015612022576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b50016020015190565b6000806020830161134f858563ffffffff611fd216565b61204a612fcb565b6040805160c08101825284815281516060810183526000808252602082810182905284518281528082018652939490850193908301916120a0565b61208d612fcb565b8152602001906001900390816120855790505b509052815260408051600080825260208281019093529190920191906120dc565b6120c9612fcb565b8152602001906001900390816120c15790505b50815260006020820152600260408201526060019290925250919050565b6008101590565b612109612fcb565b612111612fcb565b826020015160018460000151038151811061212857fe5b60209081029190910101518351600019018452915050919050565b6080015160ff161590565b61215781611177565b6101008101516000908190526101209091015152565b80826020015183600001518151811061218257fe5b60209081029190910101525080516001019052565b6080015160ff16600c1490565b6121ac613075565b60606121eb83846000815181106121bf57fe5b602001015160f81c60f81b856001815181106121d757fe5b01602001516001600160f81b0319166127b0565b90506060612218848560018151811061220057fe5b602001015160f81c60f81b866002815181106121d757fe5b90506060612245858660028151811061222d57fe5b602001015160f81c60f81b876003815181106121d757fe5b90506060612272868760038151811061225a57fe5b602001015160f81c60f81b886004815181106121d757fe5b6040805160808101825295865260208601949094529284019190915250606082015292915050565b6000611b8e6122b2856020865b048560000151612843565b6020855b066129b1565b6040805160088082528183019092526000916060919060208201818038833901905050905060006122f6866020875b048660000151612843565b90506020808606600801106123c957600061231d876020885b046001018760400151612843565b905060005b6018601f88166008030181101561237157612343838260208a5b06016129b1565b60f81b84828151811061235257fe5b60200101906001600160f81b031916908160001a905350600101612322565b506018601f8716600803015b60088110156123c2576123948260208984016122b6565b60f81b8482815181106123a357fe5b60200101906001600160f81b031916908160001a90535060010161237d565b5050612412565b60005b6008811015612410576123e2828260208961233c565b60f81b8382815181106123f157fe5b60200101906001600160f81b031916908160001a9053506001016123cc565b505b611b09826129be565b60408051602080825281830190925260009160609190602082018180388339019050509050600061244e866020876122eb565b905060208086066020011061250757600061246b8760208861230f565b905060005b601f87166020038110156124b95761248b838260208a61233c565b60f81b84828151811061249a57fe5b60200101906001600160f81b031916908160001a905350600101612470565b50601f86166008035b60208110156123c2576124d98260208984016122b6565b60f81b8482815181106124e857fe5b60200101906001600160f81b031916908160001a9053506001016124c2565b60005b602081101561241057612520828260208961233c565b60f81b83828151811061252f57fe5b60200101906001600160f81b031916908160001a90535060010161250a565b60008061255d866020876122a7565b9050600061256f8260208806876129fb565b9050600061258b88602089048488600001518960200151612a3a565b98975050505050505050565b600060606125a484612aca565b905060006125b4876020886122eb565b90506020808706600801106126955760005b6018601f88166008030181101561260e57612604826020898401068584601801815181106125f057fe5b01602001516001600160f81b031916612b3c565b91506001016125c6565b50612628876020885b048387600001518860200151612a3a565b965060006126388860208961230f565b90506018601f8816600803015b6008811015612671576126678260208a8401068684601801815181106125f057fe5b9150600101612645565b5061268d88602089046001018388604001518960600151612a3a565b9750506126d4565b60005b60088110156126c4576126ba828260208a06018584601801815181106125f057fe5b9150600101612698565b506126d187602088612617565b96505b509495945050505050565b600060606126ec84612aca565b905060006126fc876020886122eb565b905060208087066020011061278d5760005b601f871660200381101561273d57612733828260208a5b06018584815181106125f057fe5b915060010161270e565b5061274a87602088612617565b9650600061275a8860208961230f565b9050601f87166020035b6020811015612671576127838260208a8401068684815181106125f057fe5b9150600101612764565b60005b60208110156126c4576127a6828260208a612725565b9150600101612790565b606060008360f81c8360f81c0360ff16905060008460f81c60ff1690506060826040519080825280602002602001820160405280156127f9578160200160208202803883390190505b50905060005b838110156128385761281688828501602002612b58565b60001b82828151811061282557fe5b60209081029190910101526001016127ff565b509695505050505050565b60008151600014156128ac57612859600061183f565b84146128a4576040805162461bcd60e51b815260206004820152601560248201527432bc3832b1ba32b21032b6b83a3c90313ab33332b960591b604482015290519081900360640190fd5b506000610e9e565b60006128cb836000815181106128be57fe5b602002602001015161183f565b905060015b8351811015612935578460011660011415612909576129028482815181106128f457fe5b602002602001015183612b90565b9150612929565b6129268285838151811061291957fe5b6020026020010151612b90565b91505b600194851c94016128d0565b50848114612982576040805162461bcd60e51b8152602060048201526015602482015274195e1c1958dd19590818dbdc9c9958dd081c9bdbdd605a1b604482015290519081900360640190fd5b8315612992575060009050610e9e565b8260008151811061299f57fe5b60200260200101519150509392505050565b601f036008021c60ff1690565b600080805b83518110156129f457600882901b91508381815181106129df57fe5b016020015160f81c91909117906001016129c3565b5092915050565b60006060612a0885612aca565b90508260f81b818581518110612a1a57fe5b60200101906001600160f81b031916908160001a905350610e9a816129be565b60008151600314612a7c5760405162461bcd60e51b81526004018080602001828103825260228152602001806130d76022913960400191505060405180910390fd5b611b098686868686600081518110612a9057fe5b602002602001015160001c87600181518110612aa857fe5b602002602001015188600281518110612abd57fe5b6020026020010151612bbc565b6040805160208082528183019092526060918391839160208201818038833901905050905060005b6020811015612b34578260f81b8282601f0381518110612b0e57fe5b60200101906001600160f81b031916908160001a90535060089290921c91600101612af2565b509392505050565b60006060612b4985612aca565b905082818581518110612a1a57fe5b600080805b6020811015612b3457600882901b91508481850181518110612b7b57fe5b016020015160f81c9190911790600101612b5d565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b600080612bc88761183f565b9050612bd5898988612843565b506060612be0612e4d565b905060018751036001901b8910612cb25787612c00578992505050612e42565b6000612c0b8a612eee565b88519091505b60018203811015612c3957612c2f8c84600184038151811061291957fe5b9b50600101612c11565b5060015b60018203811015612c9d578a60011660011415612c7b57612c74836001830381518110612c6657fe5b602002602001015185612b90565b9350612c91565b612c8e8484600184038151811061291957fe5b93505b60019a8b1c9a01612c3d565b50612ca88b84612b90565b9350505050612e42565b60015b8751811015612d325760008a600116600114612cd15783612ce6565b888281518110612cdd57fe5b60200260200101515b905060008b600116600114612d0e57898381518110612d0157fe5b6020026020010151612d10565b845b9050612d1c8282612b90565b60019c8d1c9c909550929092019150612cb59050565b508715612d4157509050612e42565b808681518110612d4d57fe5b602002602001015184141580612d61575085155b612db2576040805162461bcd60e51b815260206004820152601c60248201527f726967687420737562747265652063616e6e6f74206265207a65726f00000000604482015290519081900360640190fd5b60008615612dc957612dc48686612b90565b612dcb565b855b905080875b60018a5103811015612df657612dec8285838151811061291957fe5b9150600101612dd0565b50838114612e3c576040805162461bcd60e51b815260206004820152600e60248201526d0caf0e0cac6e8cac840dac2e8c6d60931b604482015290519081900360640190fd5b50925050505b979650505050505050565b60408051818152610820810182526060918291906020820161080080388339019050509050612e7c600061183f565b81600081518110612e8957fe5b602090810291909101015260015b6040811015612ee857612ec9826001830381518110612eb257fe5b602002602001015183600184038151811061291957fe5b828281518110612ed557fe5b6020908102919091010152600101612e97565b50905090565b600081612efd57506001610d6f565b612f0a600183901c612eee565b6001019050610d6f565b6040518060a001604052806005906020820280388339509192915050565b604051806101e00160405280612f46613008565b8152602001612f53613008565b81526000602082018190526040820181905260608201819052608082015260a001612f7c612fcb565b815260006020820152604001612f9061309d565b8152602001612f9d61309d565b8152602001600015158152602001600060ff1681526020016060815260200160008152602001606081525090565b6040518060c0016040528060008152602001612fe56130b7565b815260606020820181905260006040830181905290820181905260809091015290565b6040805161012081019091526000815260208101613024612fcb565b8152602001613031612fcb565b815260200161303e612fcb565b815260200161304b612fcb565b81526000602082018190526040820152606001613066612fcb565b8152602001600081525090565bfe5b6040518060800160405280606081526020016060815260200160608152602001606081525090565b604051806040016040528060008152602001606081525090565b604080516060808201835260008083526020830152918101919091529056fe6e6f726d616c697a6174696f6e2070726f6f66206861732077726f6e672073697a65a265627a7a7231582017ed119964a7d20c5c89fe6afaaa2cb0c623a8ab88071b378c7ee33c0f0ad05664736f6c63430005110032"

// DeployOneStepProof2 deploys a new Ethereum contract, binding an instance of OneStepProof2 to it.
func DeployOneStepProof2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProof2, error) {
	parsed, err := abi.JSON(strings.NewReader(OneStepProof2ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OneStepProof2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProof2{OneStepProof2Caller: OneStepProof2Caller{contract: contract}, OneStepProof2Transactor: OneStepProof2Transactor{contract: contract}, OneStepProof2Filterer: OneStepProof2Filterer{contract: contract}}, nil
}

// OneStepProof2 is an auto generated Go binding around an Ethereum contract.
type OneStepProof2 struct {
	OneStepProof2Caller     // Read-only binding to the contract
	OneStepProof2Transactor // Write-only binding to the contract
	OneStepProof2Filterer   // Log filterer for contract events
}

// OneStepProof2Caller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProof2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProof2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProof2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProof2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProof2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProof2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProof2Session struct {
	Contract     *OneStepProof2    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OneStepProof2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProof2CallerSession struct {
	Contract *OneStepProof2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// OneStepProof2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProof2TransactorSession struct {
	Contract     *OneStepProof2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// OneStepProof2Raw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProof2Raw struct {
	Contract *OneStepProof2 // Generic contract binding to access the raw methods on
}

// OneStepProof2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProof2CallerRaw struct {
	Contract *OneStepProof2Caller // Generic read-only contract binding to access the raw methods on
}

// OneStepProof2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProof2TransactorRaw struct {
	Contract *OneStepProof2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProof2 creates a new instance of OneStepProof2, bound to a specific deployed contract.
func NewOneStepProof2(address common.Address, backend bind.ContractBackend) (*OneStepProof2, error) {
	contract, err := bindOneStepProof2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProof2{OneStepProof2Caller: OneStepProof2Caller{contract: contract}, OneStepProof2Transactor: OneStepProof2Transactor{contract: contract}, OneStepProof2Filterer: OneStepProof2Filterer{contract: contract}}, nil
}

// NewOneStepProof2Caller creates a new read-only instance of OneStepProof2, bound to a specific deployed contract.
func NewOneStepProof2Caller(address common.Address, caller bind.ContractCaller) (*OneStepProof2Caller, error) {
	contract, err := bindOneStepProof2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProof2Caller{contract: contract}, nil
}

// NewOneStepProof2Transactor creates a new write-only instance of OneStepProof2, bound to a specific deployed contract.
func NewOneStepProof2Transactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProof2Transactor, error) {
	contract, err := bindOneStepProof2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProof2Transactor{contract: contract}, nil
}

// NewOneStepProof2Filterer creates a new log filterer instance of OneStepProof2, bound to a specific deployed contract.
func NewOneStepProof2Filterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProof2Filterer, error) {
	contract, err := bindOneStepProof2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProof2Filterer{contract: contract}, nil
}

// bindOneStepProof2 binds a generic wrapper to an already deployed contract.
func bindOneStepProof2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OneStepProof2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProof2 *OneStepProof2Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OneStepProof2.Contract.OneStepProof2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProof2 *OneStepProof2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProof2.Contract.OneStepProof2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProof2 *OneStepProof2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProof2.Contract.OneStepProof2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProof2 *OneStepProof2CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OneStepProof2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProof2 *OneStepProof2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProof2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProof2 *OneStepProof2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProof2.Contract.contract.Transact(opts, method, params...)
}

// ExecuteStep is a free data retrieval call binding the contract method 0x1041c884.
//
// Solidity: function executeStep(bytes32 inboxAcc, bytes32 messagesAcc, bytes32 logsAcc, bytes proof, bytes bproof) view returns(uint64 gas, bytes32[5] fields)
func (_OneStepProof2 *OneStepProof2Caller) ExecuteStep(opts *bind.CallOpts, inboxAcc [32]byte, messagesAcc [32]byte, logsAcc [32]byte, proof []byte, bproof []byte) (struct {
	Gas    uint64
	Fields [5][32]byte
}, error) {
	ret := new(struct {
		Gas    uint64
		Fields [5][32]byte
	})
	out := ret
	err := _OneStepProof2.contract.Call(opts, out, "executeStep", inboxAcc, messagesAcc, logsAcc, proof, bproof)
	return *ret, err
}

// ExecuteStep is a free data retrieval call binding the contract method 0x1041c884.
//
// Solidity: function executeStep(bytes32 inboxAcc, bytes32 messagesAcc, bytes32 logsAcc, bytes proof, bytes bproof) view returns(uint64 gas, bytes32[5] fields)
func (_OneStepProof2 *OneStepProof2Session) ExecuteStep(inboxAcc [32]byte, messagesAcc [32]byte, logsAcc [32]byte, proof []byte, bproof []byte) (struct {
	Gas    uint64
	Fields [5][32]byte
}, error) {
	return _OneStepProof2.Contract.ExecuteStep(&_OneStepProof2.CallOpts, inboxAcc, messagesAcc, logsAcc, proof, bproof)
}

// ExecuteStep is a free data retrieval call binding the contract method 0x1041c884.
//
// Solidity: function executeStep(bytes32 inboxAcc, bytes32 messagesAcc, bytes32 logsAcc, bytes proof, bytes bproof) view returns(uint64 gas, bytes32[5] fields)
func (_OneStepProof2 *OneStepProof2CallerSession) ExecuteStep(inboxAcc [32]byte, messagesAcc [32]byte, logsAcc [32]byte, proof []byte, bproof []byte) (struct {
	Gas    uint64
	Fields [5][32]byte
}, error) {
	return _OneStepProof2.Contract.ExecuteStep(&_OneStepProof2.CallOpts, inboxAcc, messagesAcc, logsAcc, proof, bproof)
}

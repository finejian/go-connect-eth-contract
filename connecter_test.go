package coin

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	userAddr0       = common.HexToAddress("0x0000000000000000000000000000000000000000")
	userAddr1       = common.HexToAddress("0x8987282dcad8feee4cab0b71ae01b126f9ae27f4")
	userAddr2       = common.HexToAddress("0x204b936b9329c8b1b780610630cbbbf193e2347f")
	userKeystore1   = `{"address":"8987282dcad8feee4cab0b71ae01b126f9ae27f4","crypto":{"cipher":"aes-128-ctr","ciphertext":"aac8164a46a7dc526950024a53f7e7ac0ae3e9ef408b3572ba8ebb90dfbc1851","cipherparams":{"iv":"82e0b442471f5c0917b252078d708196"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"2f1281a07ea3b7e4485ac03c8807bedc0fa033153ccad511d70a3b752b4caf5c"},"mac":"e2ac09d068be64d2ede5677d815dd78e83b73a59a37f7042975a6688c918fb3f"},"id":"9e476b7a-ec3c-4a41-977c-24b05526a1e8","version":3}`
	userPassphrase1 = ""
	userKeystore2   = `{"address":"204b936b9329c8b1b780610630cbbbf193e2347f","crypto":{"cipher":"aes-128-ctr","ciphertext":"e94b5a5d52e5619bd842d3de2b3c0e7cd1639a463e33446fcf595ed2e3ed6343","cipherparams":{"iv":"15ddb401f5d61599a99881165c04ff5f"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"3d27281f418b3509dcb452e07825ad3e2335f58e352a7c01192ccc0420e66663"},"mac":"8175b96635275741e609ad54fe4c05aed4e02de4210eff1f48a9fc404d7ab8bb"},"id":"efa188b6-992f-4883-bc3a-4888522a67ee","version":3}`
	userPassphrase2 = "123456"
)

func Test_query(t *testing.T) {
	Convey("Query logics", t, func() {
		c := NewConnecter()

		Println("Latest block nunmber is: ", c.BlockNumber())
		Println("Contract name is: ", c.ContractName())
		Println("User No.1 ETH balance is: ", c.BalanceOfEth(userAddr1))
		Println("User No.1 SC balance is: ", c.BalanceOfCoin(userAddr1))
		Println("SC total supplpy is: ", c.TotalSupply())

		So(0, ShouldEqual, -1)
	})
}

func Test_transfer(t *testing.T) {
	Convey("Transfer logics", t, func() {
		c := NewConnecter()

		Println("User No.2 SC balance is: ", c.BalanceOfCoin(userAddr2))
		fromAuth := AuthAccount(userKeystore1, userPassphrase1)
		Println("Transfer 1024 SC from user No.1 to user No.2 result: ", c.TransferCoin(fromAuth, userAddr2, big.NewInt(1024)))
		Println("User No.2 SC balance is: ", c.BalanceOfCoin(userAddr2))

		c.TransferLogs([]common.Address{userAddr1}, nil)

		So(0, ShouldEqual, -1)
	})
}

func Test_whitelist(t *testing.T) {
	Convey("Whitelist logics", t, func() {
		c := NewConnecter()

		owerAuth := AuthAccount(userKeystore1, userPassphrase1)

		Println("User No.2 on the white list: ", c.ExistsWhiteList(userAddr2))
		Println("Remove user No.2 from the white list: ", c.RemoveFromWhitelist(owerAuth, userAddr2))
		Println("User No.2 on the white list: ", c.ExistsWhiteList(userAddr2))
		Println("Add user No.2 to the white list: ", c.AddToWhitelist(owerAuth, userAddr2))
		Println("User No.2 on the white list: ", c.ExistsWhiteList(userAddr2))

		So(0, ShouldEqual, -1)
	})
}

func Test_mint(t *testing.T) {
	Convey("mint logics", t, func() {
		c := NewConnecter()

		owerAuth := AuthAccount(userKeystore1, userPassphrase1)

		Println("User No.2 on the white list: ", c.ExistsWhiteList(userAddr2))
		Println("User No.2 SC balance is: ", c.BalanceOfCoin(userAddr2))
		Println("Mint SC to user No.2: ", c.Mint(owerAuth, userAddr2, big.NewInt(1024)))
		Println("User No.2 SC balance is: ", c.BalanceOfCoin(userAddr2))

		c.MintLogs()

		So(0, ShouldEqual, -1)
	})
}

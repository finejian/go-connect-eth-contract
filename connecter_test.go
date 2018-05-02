package coin

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	userAddr0       = common.HexToAddress("0x0000000000000000000000000000000000000000")
	userAddr1       = common.HexToAddress("0x0a7ca39ed27d24626219f8a4147484f9a8537f51")
	userAddr2       = common.HexToAddress("0x6fb96014e851659ff031ca1413ad438203c8788f")
	userAddr3       = common.HexToAddress("0x86bdfc9d950e5eb927e44174c64bc8b66a2207a8")
	userKeystore1   = `{"address":"0a7ca39ed27d24626219f8a4147484f9a8537f51","crypto":{"cipher":"aes-128-ctr","ciphertext":"a995296dbca789c2e8334c2cc5761e1c78b6a9a27a877e74e7d6105da935eeb5","cipherparams":{"iv":"85a110c69409de030e89939687c373f2"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"f8557d90c835f8b1103577db50ce3deda043b6494e0c1f52135514684f193837"},"mac":"504346cd0dfdbb9b1a4505f63441b8a3a57200f39de7e2ab51706f850f25430f"},"id":"27b2b139-c4c3-4cb9-8733-514a077d9d50","version":3}`
	userPassphrase1 = ""
	userKeystore2   = `{"address":"6fb96014e851659ff031ca1413ad438203c8788f","crypto":{"cipher":"aes-128-ctr","ciphertext":"88f4941608756f2b61fc087d1a666d12545b5f2d1af199d642f9d8c43fbecf94","cipherparams":{"iv":"0a83f7f86463988bb4a81cecd0cc79c9"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"768c97d244f72748869f0c8310e474c2dee809e503ab4bd61526a3b4302d1978"},"mac":"1ce9be55283b69c1fd8b77619796dfae7331fe0d44868b72694bbd2f6b52f4f1"},"id":"84244989-8093-401f-9f84-7572fc66ccb5","version":3}`
	userPassphrase2 = "123456"
	userKeystore3   = `{"address":"86bdfc9d950e5eb927e44174c64bc8b66a2207a8","crypto":{"cipher":"aes-128-ctr","ciphertext":"93911eed55101caa8d5e477411aa242e00c9a40cb8f261209bed6e00dadccd09","cipherparams":{"iv":"1388ba0610ea45076c84913468f73847"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"720e4fa904fa74c32d1776583def02b8e5fc4d26d2c322743d1514bdc67d3605"},"mac":"caafcf3ab02ec39eb3204228be1d71833ee77fcf4832f9ff345532f60dc56737"},"id":"19d86e8b-1c2a-487f-99d8-20376e5b518f","version":3}`
	userPassphrase3 = "123456"
)

func Test_deploy(t *testing.T) {
	Convey("Deploy logics", t, func() {
		ownerAuth := AuthAccount(userKeystore1, userPassphrase1)
		c := NewConnecterWithDeploy(ownerAuth)

		Println("Contract address is: ", CoinAddr.String())
		Println("Contract name is: ", c.ContractName())

		So(0, ShouldEqual, -1)
	})
}

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

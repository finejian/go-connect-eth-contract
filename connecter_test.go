package coin

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	userAddr0       = common.HexToAddress("0x0000000000000000000000000000000000000000")
	userAddr1       = common.HexToAddress("0x6fb96014e851659ff031cA1413AD438203C8788F")
	userAddr2       = common.HexToAddress("0x86bDFc9D950e5Eb927E44174c64BC8B66a2207a8")
	userKeystore1   = `{"address":"6fb96014e851659ff031ca1413ad438203c8788f","crypto":{"cipher":"aes-128-ctr","ciphertext":"88f4941608756f2b61fc087d1a666d12545b5f2d1af199d642f9d8c43fbecf94","cipherparams":{"iv":"0a83f7f86463988bb4a81cecd0cc79c9"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"768c97d244f72748869f0c8310e474c2dee809e503ab4bd61526a3b4302d1978"},"mac":"1ce9be55283b69c1fd8b77619796dfae7331fe0d44868b72694bbd2f6b52f4f1"},"id":"84244989-8093-401f-9f84-7572fc66ccb5","version":3}`
	userPassphrase1 = "123456"
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

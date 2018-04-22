package coin

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	userAddr1 = common.HexToAddress("0x8987282dcad8feee4cab0b71ae01b126f9ae27f4")
)

func Test_query(t *testing.T) {
	Convey("query logics", t, func() {
		c := NewConnecter()

		Println("latest block nunmber is: ", c.BlockNumber())
		Println("contract name is: ", c.ContractName())
		Println("user one eth balance is: ", c.BalanceOfEth(userAddr1))
		Println("user one super coin balance is: ", c.BalanceOfCoin(userAddr1))
		Println("super coin total supplpy is: ", c.TotalSupply())

		So(0, ShouldEqual, -1)
	})
}

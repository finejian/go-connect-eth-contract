package coin

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	coinAddr = common.HexToAddress("0x5A4E05aCd772BAe3109e6C424907BE9F4e35b6Db")
	coinHash = coinAddr.Hash()
)

// Connecter SuperCoin连接者
type Connecter struct {
	ctx  context.Context
	conn *ethclient.Client
	coin *SuperCoin
}

// NewConnecter 生成一个SuperCoin连接者
func NewConnecter() *Connecter {
	conn, err := ethclient.Dial("ws://127.0.0.1:8546")
	if err != nil {
		panic(err)
	}
	coin, err := NewSuperCoin(coinAddr, conn)
	if err != nil {
		panic(err)
	}
	return &Connecter{
		ctx:  context.Background(),
		conn: conn,
		coin: coin,
	}
}

// BlockNumber 当前块高度
func (c *Connecter) BlockNumber() *big.Int {
	blockNumber, err := c.conn.BlockByNumber(c.ctx, nil)
	if err != nil {
		log.Fatal("Get block number error", err)
		return big.NewInt(0)
	}
	return blockNumber.Number()
}

// ContractName SuperCoin名称
func (c *Connecter) ContractName() string {
	name, err := c.coin.Name(nil)
	if err != nil {
		log.Fatal("Get contract name error", err)
		return ""
	}
	return name
}

// BalanceOfEth 以太币余额
func (c *Connecter) BalanceOfEth(addr common.Address) *big.Int {
	balance, err := c.conn.BalanceAt(c.ctx, addr, nil)
	if err != nil {
		log.Fatal("Get address eth balance error", err)
		return big.NewInt(0)
	}
	return balance
}

// BalanceOfCoin SuperCoin余额
func (c *Connecter) BalanceOfCoin(addr common.Address) *big.Int {
	balance, err := c.coin.BalanceOf(nil, addr)
	if err != nil {
		log.Fatal("Get address SuperCoin balance error", err)
		return big.NewInt(0)
	}
	return balance
}

// TotalSupply SuperCoin已铸币量
func (c *Connecter) TotalSupply() *big.Int {
	totalSupply, err := c.coin.TotalSupply(nil)
	if err != nil {
		log.Fatal("Get SuperCoin totalSuply error", err)
		return big.NewInt(0)
	}
	return totalSupply
}

// AuthAccount 解锁账户
func AuthAccount(key, passphrase string) *bind.TransactOpts {
	auth, err := bind.NewTransactor(strings.NewReader(key), passphrase)
	if err != nil {
		log.Fatalf("Auth account error: %v", err)
		return nil
	}
	log.Printf("Gas price is: %v, Gas limit is: %v", auth.GasPrice, auth.GasLimit)
	return auth
}

// TransferCoin SuperCoin转账
func (c *Connecter) TransferCoin(fromAuth *bind.TransactOpts, to common.Address, amount *big.Int) bool {
	tx, err := c.coin.Transfer(fromAuth, to, amount)
	if err != nil {
		log.Fatalf("Transfer SuperCoin from %s to %s amount %s error: %v", fromAuth.From.String(), to.String(), amount.String(), err)
		return false
	}
	// 等待执行
	receipt, err := bind.WaitMined(c.ctx, c.conn, tx)
	if err != nil {
		log.Fatalf("Wait Transfer Mined error: %v", err)
		return false
	}
	return receipt.Status == 1
}

// TransferLogs SuperCoin转账记录
func (c *Connecter) TransferLogs(froms []common.Address, tos []common.Address) {
	iter, err := c.coin.FilterTransfer(&bind.FilterOpts{}, froms, tos)
	defer iter.Close()
	if err != nil {
		panic(err)
	}
	for {
		if !iter.Next() {
			break
		}
		log.Printf("Transfer log: %s", iter.Event.String())
	}
}

func (s *SuperCoinTransfer) String() string {
	return fmt.Sprintf("From: %s, To: %s, Amount: %s, Log: %s", s.From.Hex(), s.To.Hex(), s.Value, s.Raw.String())
}

// ExistsWhiteList 地址是否在白名单
func (c *Connecter) ExistsWhiteList(addr common.Address) bool {
	exists, err := c.coin.Whitelist(&bind.CallOpts{}, addr)
	if err != nil {
		log.Fatalf("Query address on the white list error: %v", err)
		return false
	}
	return exists
}

// AddToWhitelist 地址是否在白名单
func (c *Connecter) AddToWhitelist(ownerAuth *bind.TransactOpts, addr common.Address) bool {
	tx, err := c.coin.AddToWhitelist(ownerAuth, addr)
	if err != nil {
		log.Fatalf("Add address to the white list error: %v", err)
		return false
	}
	// 等待执行
	receipt, err := bind.WaitMined(c.ctx, c.conn, tx)
	if err != nil {
		log.Fatalf("Wait AddToWhitelist Mined error: %v", err)
		return false
	}
	return receipt.Status == 1
}

// RemoveFromWhitelist 地址是否在白名单
func (c *Connecter) RemoveFromWhitelist(ownerAuth *bind.TransactOpts, addr common.Address) bool {
	tx, err := c.coin.RemoveFromWhitelist(ownerAuth, addr)
	if err != nil {
		log.Fatalf("Remove address from the white list error: %v", err)
		return false
	}
	// 等待执行
	receipt, err := bind.WaitMined(c.ctx, c.conn, tx)
	if err != nil {
		log.Fatalf("Wait RemoveFromWhitelist Mined error: %v", err)
		return false
	}
	return receipt.Status == 1
}

// Mint 给白名单地址铸币
func (c *Connecter) Mint(ownerAuth *bind.TransactOpts, addr common.Address, amount *big.Int) bool {
	tx, err := c.coin.Mint(ownerAuth, addr, amount)
	if err != nil {
		log.Fatalf("Mint to address SC error: %v", err)
		return false
	}
	// 等待执行
	receipt, err := bind.WaitMined(c.ctx, c.conn, tx)
	if err != nil {
		log.Fatalf("Wait Mint SC Mined error: %v", err)
		return false
	}
	return receipt.Status == 1
}

// MintLogs 铸币记录
func (c *Connecter) MintLogs() {
	iter, err := c.coin.FilterMint(&bind.FilterOpts{}, nil)
	defer iter.Close()
	if err != nil {
		panic(err)
	}
	for {
		if !iter.Next() {
			break
		}
		log.Printf("Mint log: %s", iter.Event.String())
	}
}

func (s *SuperCoinMint) String() string {
	return fmt.Sprintf("Address: %s, Amount: %s, Log: %s", s.To.String(), s.Amount.String(), s.Raw.String())
}

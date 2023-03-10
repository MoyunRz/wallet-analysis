package utils

import (
	"context"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
	"testing"
	"wallet-analysis/common/log"
)

func Test_api(t *testing.T) {
	cli, err := ethclient.Dial("http://124.71.12.16:9933")
	if err != nil {
		t.Fatal(err.Error())
	}

	block, err := cli.BlockByNumber(context.Background(), big.NewInt(1))
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(String(block))
	block, err = cli.BlockByHash(context.Background(), common.HexToHash("0x69a8794663704c3afd679b80ac486ca4f198780513ef4538884ec8ae8342f799"))
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(String(block))
	t.Log(cli.BlockNumber(context.Background()))
	receipt, err := cli.TransactionReceipt(context.TODO(), common.HexToHash("0xdade9d7f316fa709ade853f465d320fa58f81fbefdaf94f1da87e6265e2c715c"))
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(String(receipt))
}

func Test_acc(t *testing.T) {
	a := "0xa9059cbb0000000000000000000000005aa0729c9e76361c538df71d88d4b5b9fab9337600000000000000000000000000000000000000000000006054d4350ced240000"
	t.Log(a[34:74])
}

func Test_acli(t *testing.T) {
	rpc := NewRpcClient("https://node.waltonchain.pro")
	h, err := rpc.BlockNumber()
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(h)
	block, err := rpc.BlockByNumber(1)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(String(block))
	receipt, err := rpc.TransactionReceipt("0xdade9d7f316fa709ade853f465d320fa58f81fbefdaf94f1da87e6265e2c715c")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(String(receipt))
	d, err := hex.DecodeString(strings.TrimPrefix(receipt.Logs[0].Data, "0x"))
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(receipt.Logs[0].Data, big.NewInt(0).SetBytes(d).String())
}

func TestClient(t *testing.T) {
	cli, err := ethclient.Dial("ws://119.23.219.232:9944")
	if err != nil {
		log.Error("开始区块事件订阅失败，无法进行socket连接eth")
		return
	}
	defer cli.Close()

}

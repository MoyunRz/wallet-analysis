package utils;

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math/big"
	"strconv"
	"strings"
	"time"
)


func RemoveHex0x(hexStr string) string {
	if strings.HasPrefix(hexStr, "0x") {
		return hexStr[2:]
	}
	return hexStr
}

var hexDecMap = map[string]decimal.Decimal{
	"0": decimal.NewFromInt(0),
	"1": decimal.NewFromInt(1),
	"2": decimal.NewFromInt(2),
	"3": decimal.NewFromInt(3),
	"4": decimal.NewFromInt(4),
	"5": decimal.NewFromInt(5),
	"6": decimal.NewFromInt(6),
	"7": decimal.NewFromInt(7),
	"8": decimal.NewFromInt(8),
	"9": decimal.NewFromInt(9),
	"a": decimal.NewFromInt(10),
	"b": decimal.NewFromInt(11),
	"c": decimal.NewFromInt(12),
	"d": decimal.NewFromInt(13),
	"e": decimal.NewFromInt(14),
	"f": decimal.NewFromInt(15),
}

var hexDec = decimal.NewFromInt(16)

func HexToDec(hex string) *big.Int {
	if strings.HasPrefix(hex, "0x") {
		hex = hex[2:]
	}

	bigIntValue, ok := new(big.Int).SetString(hex, 16)
	if !ok {
		return big.NewInt(-1)
	}
	return bigIntValue
}

func DecToHex(dec int64) string {
	return "0x" + strconv.FormatInt(dec, 16)
}
func Timestamp(seconds int64) string {
	var timelayout = "2006-01-02 T 15:04:05.000" // 时间格式

	datatimestr := time.Unix(seconds, 0).Format(timelayout)

	return datatimestr

}

func Del0xToLower(address string) string {
	if strings.HasPrefix(address, "0x") {
		return strings.ToLower(strings.TrimPrefix(address, "0x"))
	}
	return strings.ToLower(address)
}

// ParseBigInt parse hex string value to big.Int
func ParseBigInt(value string) (*big.Int, error) {
	i := &big.Int{}
	_, err := fmt.Sscan(value, i)

	return i, err
}

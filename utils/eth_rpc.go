package utils

import (
	"net/url"
)

func EthGetBlockByHeight(height int) {
	values := url.Values{}
	values.Add("method", "eth_getBlockByNumber")
	values.Add("params", "[0x01]")
	values.Add("id", "90")
	values.Add("jsonrpc", "2.0")
	res, err := HttpPostForm("http://124.71.12.16:9933", values)
	if err != nil {
		println(err.Error)
	}
	println(string(res))
}

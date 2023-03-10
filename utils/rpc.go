package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
	"wallet-analysis/common/conf"
	"wallet-analysis/common/log"

	"golang.org/x/net/context"
	"google.golang.org/appengine/urlfetch"
)

type ErrorCode int

var ErrNullResult = errors.New("result is null")

type Error struct {
	// A Number that indicates the error type that occurred.
	Code ErrorCode `json:"code"` /* required */
	// A String providing a short description of the error.
	// The message SHOULD be limited to a concise single sentence.
	Message string `json:"message"` /* required */
	// A Primitive or Structured value that contains additional information about the error.
	Data interface{} `json:"data"` /* optional */
}

func (e *Error) Error() string {
	return e.Message
}

// Response RPC 响应返回数据结构
type Response struct {
	ID      string          `json:"id"`
	JSONRPC string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	Error   *Error          `json:"error"`
}

// RPC 请求参数数据结构
type request struct {
	ID      string        `json:"id"`
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

// RpcClient 包装的RPC-HTTP 客户端
type RpcClient struct {
	client      *http.Client
	url         string
	Debug       bool
	mutex       *sync.Mutex
	Credentials string //访问权限认证的 base58编码
}

// NewRpcClient 使用给定的 url 创建新的 rpc RpcClient
func NewRpcClient(url string, options ...func(rpc *RpcClient)) *RpcClient {
	rpc := &RpcClient{
		client:      http.DefaultClient,
		url:         url,
		mutex:       &sync.Mutex{},
		Credentials: "",
	}
	for _, option := range options {
		option(rpc)
	}

	return rpc
}

// URL 获取RPC服务URL
func (rpc *RpcClient) URL() string {
	return rpc.url
}

func (rpc *RpcClient) Urlfetch(ctx context.Context, seconds ...int) {

	if len(seconds) > 0 {
		ctx, _ = context.WithDeadline(
			ctx,
			time.Now().Add(time.Duration(1000000000*seconds[0])*time.Second),
		)
	}

	rpc.client = urlfetch.Client(ctx)
}

// CallNoAuth 没有权限认证的RPC请求。
func (rpc *RpcClient) CallNoAuth(method string, target interface{}, params ...interface{}) error {
	result, err := rpc.call(method, params...)
	if err != nil {
		return err
	}

	if target == nil {
		return nil
	}

	return json.Unmarshal(result, target)
}

// CallWithRPC 需要权限认证的RPC请求。
func (rpc *RpcClient) CallWithRPC(method string, target interface{}, params ...interface{}) error {

	result, err := rpc.call(method, params...)
	if err != nil {
		//log.Info("CallWithAuth", method, err.Error(), rpc.url)
		return err
	}

	if target == nil {
		return nil
	}

	return json.Unmarshal(result, target)
}

// Call returns raw response of method call
func (rpc *RpcClient) call(method string, params ...interface{}) (json.RawMessage, error) {
	if params == nil {
		params = []interface{}{}
	}
	request := request{
		ID:      fmt.Sprintf("%d", conf.Cfg.ChainId),
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", rpc.url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	//log.Infof("rpc.client.Do %v \n", req)
	response, err := rpc.client.Do(req)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		log.Info(err)
		return nil, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Info(err)
		return nil, fmt.Errorf("ReadAll err: %v", err)
	}

	//rpc.log.Println(fmt.Sprintf("%s\nResponse: %s\n", method, data))
	resp := new(Response)
	if err := json.Unmarshal(data, resp); err != nil {
		log.Info(err.Error())
		return nil, fmt.Errorf("resp: %v , err: %v", resp, err)
	}

	if resp.Error != nil {
		log.Info(resp.Error.Error())
		return nil, resp.Error
	}

	return resp.Result, nil

}

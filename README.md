# wallet-analysis

该项目为eth的扫链程序，主要作用是为app钱包提供查询服务和分析归类区块数据

## 主要运用包
```
github.com/ethereum/go-ethereum
```

## 扫链方式

- 通过rpc请求进行一个一个区块的获取和解析交易
- 通过订阅的方式监听合约内部的事件进行处理


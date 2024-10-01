# 现货交易文档

https://developers.binance.com/docs/zh-CN/binance-spot-api-docs/README

# 合约交易文档

https://developers.binance.com/docs/zh-CN/derivatives/Introduction

# 现货测试网络

https://testnet.binance.vision/

# 合约测试网络

https://testnet.binancefuture.com/en/futures/BTCUSDT

# 如何生成 Ed25519 API keys?

1.生成私钥

```shell
openssl genpkey -algorithm ed25519 -out test-prv-key.pem
```

2.用私钥生成公钥

```shell
openssl pkey -pubout -in test-prv-key.pem -out test-pub-key.pem
```

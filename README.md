# coin-go

## 中文名字: 币狗

coin-go 是一个用 Go 语言编写的开源库，旨在封装各大加密货币交易所的 API，简化交易所 API 的调用。目前，coin-go
支持币安（Binance），并计划扩展到更多交易所。

## 特性

- **币安 API 封装**：提供对币安交易所的便捷接口，包括市场数据获取、账户管理、下单等功能。
- **轻量且高效**：基于 Go 语言的高性能实现，适合开发需要高并发的应用程序。
- **可扩展性**：为未来支持更多交易所做好准备，具备良好的扩展性设计。
- **简单易用**：提供简单的调用方式，方便开发者快速集成到自己的项目中。

## 安装

要在项目中使用 coin-go，请确保已安装 Go 环境，并按照以下步骤操作：

使用 `go get` 命令获取项目：

```bash
go get github.com/sleep-go/coin-go
```

演示代码

```go
package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/sleep-go/coin-go/binance"
	"github.com/sleep-go/coin-go/binance/consts"
	"github.com/sleep-go/coin-go/binance/consts/enums"
	"github.com/sleep-go/coin-go/binance/spot/endpoints/general"
	"github.com/sleep-go/coin-go/binance/spot/endpoints/market"
)

func TestNewExchangeInfo(t *testing.T) {
	client := binance.NewClient(
		"vmPUZE6mv9SD5VNHk4HlWFsOr6aKE2zvsw0MuIgwCIPy6utIco14y7Ju91duEh8A",
		"NhqPtmdSJYdKjVHjA7PZj4Mge3R5YNiP1e3UZjInClVN65XAbvqqM6A7H5fATj0j",
		consts.REST_API,
	)
	client.Debug = true
	response, err := general.NewExchangeInfo(client, []string{"ETHUSDT"}, nil).Call(context.Background())
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Println(response)
}

func TestAggTrades(t *testing.T) {
	res, err := market.NewAggTrades(client, "BTCUSDT", enums.Limit10).
		SetStartTime(time.Now().UnixMilli() - 60*60*24*30*365*5).
		SetEndTime(time.Now().UnixMilli()).
		SetFromId(3031206).
		Call(context.Background())
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	for _, r := range res {
		fmt.Println(r)
	}
}

```

# 目前支持的交易所

- **币安**：[Binance API 文档](https://developers.binance.com/docs/zh-CN)

| 功能                              | 完成度 |
|---------------------------------|-----|
| 现货交易(REST接口 通用接口)               | 完成  |
| 现货交易(REST接口 行情接口)               | 完成  |
| 现货交易(REST接口 账户接口)               | 完成  |
| 现货交易(Web Socket 行情接口)           | 完成  |
| 现货交易(WebSocket 账户接口)            | 完成  |
| 现货交易(Binance 的公共 WebSocket API) | 未完成 |

# 支持的功能

* 获取市场数据（如行情、深度、K线等）
* 账户信息和资产余额查询
* 下单、撤单等交易操作

# 未来计划

* 支持更多交易所，如火币、OKEx、Kraken 等。
* 添加更多高级功能，如自动交易机器人、数据分析工具。

# 贡献

我们欢迎社区贡献，任何形式的反馈或建议都非常宝贵。请查看 贡献指南 以了解如何提交贡献。

# 许可证

本项目采用 MIT 许可证。

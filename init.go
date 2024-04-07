package pay

import (
	"github.com/jimu-server/logger"
	"github.com/jimu-server/pay/control"
	"github.com/jimu-server/web"
)

var logs = logger.Logger

func init() {
	ali := web.Engine.Group("/ali")
	ali.POST("/pay", control.CreateOrderPlay) // 创建支付订单
	ali.POST("/notify", control.AliPayNotify) // 异步回调通知
}

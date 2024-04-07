package control

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/xlog"
	"github.com/jimu-server/common/resp"
	"github.com/jimu-server/logger"
	"github.com/jimu-server/pay/ali"
	"github.com/jimu-server/util/uuidutils/uuid"
	"github.com/jimu-server/web"
	"net/http"
)

var logs = logger.Logger

func CreateOrderPlay(c *gin.Context) {
	var err error
	var args *PayArgs
	web.BindJSON(c, &args)
	var url string
	if url, err = ali.CreateOrderPlay("测试支付", uuid.String(), args.Value); err != nil {
		logs.Error(err.Error())
		c.JSON(500, resp.Error(err, resp.Msg("下单失败")))
		return
	}
	c.JSON(200, resp.Success(url, resp.Msg("下单成功")))
}

func AliPayNotify(c *gin.Context) {
	notifyReq, err := alipay.ParseNotifyToBodyMap(c.Request) // c.Request 是 gin 框架的写法
	if err != nil {
		xlog.Error(err)
		return
	}
	if status := notifyReq.Get("trade_status"); status == "TRADE_SUCCESS" {
		logs.Info("支付成功")
	}
	// 支付宝异步通知验签（公钥模式）
	ok, err := alipay.VerifySign(string(ali.AlipayPublicKey), notifyReq)
	if ok {
		logs.Info("支付宝异步通知验签成功")
	}
	c.String(http.StatusOK, "%s", "success")
}

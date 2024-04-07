package ali

import (
	"context"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/jimu-server/config"
	"github.com/jimu-server/logger"
	"os"
)

var AliPayClient *alipay.Client
var logs = logger.Logger
var privateKey = "MIIEpQIBAAKCAQEAh2rPZ/o8LJuXsm5+cpV3NFpwpDzsccWsY98ztYe1CyPqvORc4/c0PLQrWUDblwao5eO9v43yiSixtqeifThEree5Dg1vCXv5+IRUhjroWkVENk4zOeqtKkY7mJJqbsBV8ngWjJUqQNXFOrkGcP+zBitzcI2qdu438zHudZcduESx4at7UmZQT38fo1YbA9dHiq1SH6wZuSsQ6ZiATpqxkkHwWauQzq42oLCh0TwaKtXannqIY4fJW+2dJ07AS18R1Bdckcvy9Hp8Bb64a+hfwqvh8K17vZ+UlNh89pRJegjskh8s3O+ItVs4zN/HvDn+aF3t+8Dl8Hip723TaWnWcwIDAQABAoIBAGuoKLfbIte74wACBBkIZrqCZCbOIJPauVC09CEPgIkYxtfhHVBHCYpxGj1c6LbKnqAVTJbrPLR6W76AyxeOElvHa0GWwH3jyDkgyynjzzFk+/PIe8VBouWG7H5BeMgPtMSNYC/nrq3YxXIIAuz2+9ad2KrUlX7zqBt5yMBHoWMKXaxkgVD+Uxy6chPefSSs2Exun3gVbHPALelL32CgxwRfDWH+vxMiGvX3ClbSH4t0ymklR+I8gFy8zqWJIvfxQagVE8Aw6mswBJ0pgF0MrK+Eqp2Z/OjzLRk/n7dabK891otbFPQRQcUiE9nPqrEwPnuneWZJinxeLPYU0KhFHnECgYEA3IcQQ0tb5FnU6bL5QqIdX+m+nrd45u3EoCPxPpN3NCGdqw9uu+eGRUTt6/rOh3c0xbY4ZZBiRuEQkpEekVhtAWYKgSDz46XYadYjcy6/Pi05pq/FcoiveR/hm51BWO+z7ag3w0q4Hv2YIdcymNH+dIyxY4IXgBRht/8iOkX6+k0CgYEAnTMNYQ6jfpYSOWgAi9Wf29mKYPI71+tsP/btJuu7Vi+RWGXrV4hAh3k2NEBR97HPQkI0jNvzE0DUebALNlLdV10Q1oBgg0iRWL85zMxIt1Chu+LUqigYAQ9PxPvXF1407oQ/KjYtSEFg4AIhKFiunOSs9VL1dsWw4Gf/Yxra878CgYEArDlW/BcoZse5PeImRGGzKyqzUeTbqQ3b12z6hPJJ0A7IwcVFp74C4KoaXYb8MFNqhVXv6XG/LreqZ6yqALzcNJFqdrozfoAQ6WbxPI0vkfFj6sevHemdkKzTVuKTHE/nZx1On1gFPz6xxwv3Wd32KTOPfbUlgUNppXa9VmE8xqkCgYEAgVPH9QqTDYivc5URpflZLAkb5EhFXY0soK8oSjX0CKLXw88NwBSzagEZzAECrnIVnmBTVXN61mMmqvpfLxEbUk2Zla9GN5vTIB0qk0eZp17pwGaMKXUU0oJjUR8TxQDXMUeRo8uwk1peNczqtIPJHACyHp5inZVkwCovHLyyV4cCgYEAyJN6+dGtBaJqmpMrD65zQatmUfOkZ6ttBpocNTmbgG0hQLDN8I/W8KoxNVgr6YloAmsH/k9lJGqBw/NQi9/b0Fv2ZZbXHT8bFdkDqkGQ7KwjlIL0yNgIsKrGzAGsgNL4oVymnpYsOYu3BsFIWqHaW2VjzIsNfDswFdo2dM00B3M="
var AppPublicCert, AlipayPublicKey, AlipayRootKey []byte

func init() {
	AliPayInit()
}
func AliPayInit() {
	// appid
	AppId := config.Evn.App.Pay.AliPay.AppId
	appPublicCertPath := config.Evn.App.Pay.AliPay.AppPublicCert
	alipayPublicKeyPath := config.Evn.App.Pay.AliPay.AliPayPublicKey
	alipayRootCertPath := config.Evn.App.Pay.AliPay.AliPayRootKey
	var err error

	// 应用程序公钥
	if AppPublicCert, err = os.ReadFile(appPublicCertPath); err != nil {
		logs.Error(err.Error())
		return
	}
	if AlipayPublicKey, err = os.ReadFile(alipayPublicKeyPath); err != nil {
		logs.Error(err.Error())
		return
	}

	if AlipayRootKey, err = os.ReadFile(alipayRootCertPath); err != nil {
		logs.Error(err.Error())
		return
	}

	// 初始化支付宝客户端
	// appid：应用ID
	// privateKey：应用私钥，支持PKCS1和PKCS8
	// isProd：是否是正式环境，沙箱环境请选择新版沙箱应用。
	if AliPayClient, err = alipay.NewClient(AppId, privateKey, false); err != nil {
		logs.Error(err.Error())
		return
	}
	AliPayClient.DebugSwitch = gopay.DebugOn
	// 设置时区，不设置或出错均为默认服务器时间
	AliPayClient.SetLocation(alipay.LocationShanghai)
	// 设置字符编码，不设置默认 utf-8
	AliPayClient.SetCharset(alipay.UTF8)
	// 设置签名类型，不设置默认 RSA2
	AliPayClient.SetSignType(alipay.RSA2)

	// 前端设置返回url
	AliPayClient.SetReturnUrl("http://localhost:5173/#/app/dev")
	// 设置异步通知URL 支付完成以后的回调接口
	AliPayClient.SetNotifyUrl("http://jimuos.v7.idcfengye.com/ali/notify")

	AliPayClient.SetAppAuthToken("")

	// 传入 alipayCertPublicKey_RSA2.crt 内容
	AliPayClient.AutoVerifySign(AlipayPublicKey)
	err = AliPayClient.SetCertSnByContent(AppPublicCert, AlipayRootKey, AlipayPublicKey)
}

func CreateOrderPlay(title, orderId, amount string) (url string, err error) {
	params := map[string]any{
		"subject":         title,
		"out_trade_no":    orderId,
		"total_amount":    amount,
		"timeout_express": "6m",
		"product_code":    "FAST_INSTANT_TRADE_PAY",
		"qr_pay_mode":     "3",
	}
	url, err = AliPayClient.TradePagePay(context.Background(), params)
	return
}

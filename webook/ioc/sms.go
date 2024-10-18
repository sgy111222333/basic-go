package ioc

import (
	"github.com/sgy111222333/basic-go/webook/internal/service/sms"
	"github.com/sgy111222333/basic-go/webook/internal/service/sms/tencent"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	TencentSMS "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111" //引入sms
	"os"
)

func InitSMSService() sms.Service {
	//return nil
	return initTencentSMSService()
}

func initTencentSMSService() sms.Service {
	secretId, ok := os.LookupEnv("SMS_SECRET_ID")
	if !ok {
		panic("没找到腾讯 SMS 的 secret id")
	}
	secretKey, ok := os.LookupEnv("SMS_SECRET_KEY")
	if !ok {
		panic("没找到腾讯 SMS 的 secret key")
	}
	c, err := TencentSMS.NewClient(
		common.NewCredential(secretId, secretKey),
		"ap-beijing",
		profile.NewClientProfile(),
	)
	if err != nil {
		panic(err)
	}
	return tencent.NewService(c, "1400866577", "一型码农广广公众号")
}

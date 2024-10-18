package test

import (
	"encoding/json"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"os"
	"testing"
)

func TestTencentSMS(t *testing.T) {
	println("test")
	credential := common.NewCredential(
		os.Getenv("SMS_SECRET_ID"),
		os.Getenv("SMS_SECRET_KEY"),
	)

	cpf := profile.NewClientProfile()
	// 实例化要请求产品(以sms为例)的client对象
	client, _ := sms.NewClient(credential, "ap-beijing", cpf)
	request := sms.NewSendSmsRequest()
	/* 模板参数: 模板参数的个数需要与 TemplateId 对应模板的变量个数保持一致，若无模板参数，则设置为空*/
	request.TemplateParamSet = common.StringPtrs([]string{"472464", "10"})
	/* 短信应用ID: 短信SdkAppId在 [短信控制台] 添加应用后生成的实际SdkAppId，示例如1400006666 */
	request.SmsSdkAppId = common.StringPtr("1400866577")
	/* 短信签名内容: 使用 UTF-8 编码，必须填写已审核通过的签名 */
	request.SignName = common.StringPtr("一型码农广广公众号")
	/* 模板 ID: 必须填写已审核通过的模板 ID */
	request.TemplateId = common.StringPtr("1986336")

	request.PhoneNumberSet = common.StringPtrs([]string{"+8615021580082"})
	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, err := client.SendSms(request)
	// 处理异常
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		panic(err)
	}
	b, _ := json.Marshal(response.Response)
	// 打印返回的json字符串
	fmt.Printf("%s", b)
}

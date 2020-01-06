package main

import (
	"fmt"

	snssms "github.com/atEaE/golang-aws-sdk-sample/internal/aws/sns-sms"
	"github.com/atEaE/golang-aws-sdk-sample/internal/conf"
)

func main() {
	appConf, err := conf.GetSnsConf()
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}

	client, err := snssms.GetClient(appConf.GlobalConf.AccessKeyId, appConf.GlobalConf.SecretKeyId, appConf.GlobalConf.Region)
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}

	msgIn := snssms.CreateInputMessage("Test Message", appConf.PhoneNum)

	result, err := client.Publish(msgIn)
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}

	fmt.Printf("Result: %s", result.String())
}

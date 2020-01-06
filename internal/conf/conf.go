package conf

import (
	"gopkg.in/ini.v1"
)

type SnsAppConf struct {
	GlobalConf AwsGlobalConf
	PhoneNum   string
}

type AwsGlobalConf struct {
	AccessKeyId string
	SecretKeyId string
	Region      string
}

func GetSnsConf() (*SnsAppConf, error) {
	conf, err := ini.Load("app.conf")
	if err != nil {
		return nil, err
	}

	appConf := &SnsAppConf{
		GlobalConf: AwsGlobalConf{
			AccessKeyId: conf.Section("global").Key("accessKeyId").String(),
			SecretKeyId: conf.Section("global").Key("secretKeyId").String(),
			Region:      conf.Section("global").Key("region").String(),
		},
		PhoneNum: conf.Section("sns").Key("phoneNum").String(),
	}
	return appConf, nil
}

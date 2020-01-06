package snssms

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

// AWS_ACCESS_KEY : acsId
// AWS_SECRET_ACCESS_KEY : secId
// REGION : reg
func GetClient(acsId string, secId string, reg string) (*sns.SNS, error) {
	creds := credentials.NewStaticCredentials(acsId, secId, "")
	sess, err := session.NewSession(&aws.Config{
		Credentials: creds,
		Region:      aws.String(reg),
	})

	if err != nil {
		return nil, err
	}

	return sns.New(sess), nil
}

func CreateInputMessage(msg string, phoneNum string) *sns.PublishInput {
	pin := &sns.PublishInput{}
	pin.SetMessage(msg)
	pin.SetPhoneNumber(phoneNum)
	return pin
}

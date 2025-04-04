package aws_s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/matiniiuu/mcommon/pkg/email"
	"github.com/matiniiuu/mcommon/pkg/mconfig"
)

type (
	AwsSES struct {
		client *ses.Client
	}
)

func New(cfg *mconfig.AwsSES) (email.Email, error) {
	awsConfig, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(cfg.SESRegion),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.SESAccessKey, cfg.SESSecretKey, "")),
	)
	if err != nil {
		return nil, err
	}

	return &AwsSES{client: ses.NewFromConfig(awsConfig)}, nil
}

func (a *AwsSES) SendHtml(from string, to []string, subject, messageBody string) error {
	_, err := a.client.SendEmail(context.TODO(), &ses.SendEmailInput{
		Source:      aws.String(from),
		Destination: &types.Destination{ToAddresses: to},
		Message: &types.Message{
			Subject: &types.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(subject),
			},
			Body: &types.Body{
				Html: &types.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(messageBody),
				},
			}},
	})
	if err != nil {
		return err
	}

	return nil
}

func (a *AwsSES) SendText(from string, to []string, subject, messageBody string) error {
	_, err := a.client.SendEmail(context.TODO(), &ses.SendEmailInput{
		Source:      aws.String(from),
		Destination: &types.Destination{ToAddresses: to},
		Message: &types.Message{
			Subject: &types.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(subject),
			},
			Body: &types.Body{
				Text: &types.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(messageBody),
				},
			}},
	})
	if err != nil {
		return err
	}

	return nil
}

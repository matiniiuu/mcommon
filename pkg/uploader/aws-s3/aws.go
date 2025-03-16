package aws_s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/matiniiuu/mcommon/pkg/mconfig"
	"github.com/matiniiuu/mcommon/pkg/uploader"
)

type (
	AwsS3 struct {
		baseUrl string
		bucket  string
		awsS3   *s3.Client
	}
)

func New(cfg *mconfig.AwsS3) (uploader.Uploader, error) {
	awsConfig, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(cfg.S3Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.S3AccessKey, cfg.S3SecretKey, "")),
	)
	if err != nil {
		return nil, err
	}

	return &AwsS3{
		awsS3:   s3.NewFromConfig(awsConfig),
		bucket:  cfg.S3Bucket,
		baseUrl: cfg.S3BaseUrl,
	}, nil
}

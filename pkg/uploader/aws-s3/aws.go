package cloudinary

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/matiniiuu/mcommon/pkg/mconfig"
	"github.com/matiniiuu/mcommon/pkg/uploader"
)

type (
	AwsS3 struct {
		bucket string
		awsS3  *s3.Client
	}
)

func New(cfg mconfig.AwsS3) uploader.Uploader {
	awsConfig, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(cfg.S3Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.S3AccessKey, cfg.S3SecretKey, "")),
	)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &AwsS3{awsS3: s3.NewFromConfig(awsConfig), bucket: cfg.S3Bucket}
}

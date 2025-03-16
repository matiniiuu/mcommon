package aws_s3

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/matiniiuu/mcommon/pkg/uploader"
)

func (a *AwsS3) Upload(file multipart.File, folder, filename string) (string, error) {
	mimetype, err := uploader.GetMimetype(file)
	if err != nil {
		return "", err
	}
	fileKey := fmt.Sprintf("%s/%s", folder, filename)
	_, err = a.awsS3.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:             aws.String(a.bucket),
		Key:                aws.String(fileKey),
		Body:               file,
		ContentType:        aws.String(mimetype),
		ACL:                types.ObjectCannedACLPublicRead,
		ContentDisposition: aws.String("inline"),
	})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s", a.baseUrl, fileKey), nil
}

func (a *AwsS3) UploadBase24(base64File, folder, filename string) (string, error) {
	decode, err := base64.StdEncoding.DecodeString(base64File[strings.IndexByte(base64File, ',')+1:])
	if err != nil {
		return "", err
	}
	fileKey := fmt.Sprintf("%s/%s", folder, filename)
	_, err = a.awsS3.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:             aws.String(a.bucket),
		Key:                aws.String(fmt.Sprintf("%s/%s", folder, filename)),
		Body:               bytes.NewReader(decode),
		ContentType:        aws.String(base64File[strings.IndexByte(base64File, ':')+1 : strings.IndexByte(base64File, ';')]),
		ACL:                types.ObjectCannedACLPublicRead,
		ContentDisposition: aws.String("inline"),
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", a.baseUrl, fileKey), nil
}

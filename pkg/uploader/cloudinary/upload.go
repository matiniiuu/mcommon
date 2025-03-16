package cloudinary

import (
	"mime/multipart"
)

func (c *Cloudinary) Upload(file multipart.File, folder, filename string) (path string, err error) {
	// newFilename := uuid.New().String()
	// resp, err := c.cld.Upload.Upload(
	// 	context.Background(),
	// 	file,
	// 	uploader.UploadParams{Folder: folder, PublicID: newFilename})
	// if err != nil {
	// 	return "", derrors.InternalError()
	// }
	// filePath := resp.SecureURL
	return "", nil
	// return filePath, nil
}

func (c *Cloudinary) UploadBase24(base64File, folder, filename string) (string, error) {
	// decode, err := base64.StdEncoding.DecodeString(base64File[strings.IndexByte(base64File, ',')+1:])
	// if err != nil {
	// 	return "", "", err
	// }
	// sss := bytes.NewReader(decode)
	// // awsSession := initAWSConnection()

	// uploadParams := &s3.PutObjectInput{
	// 	Bucket: aws.String("AWS_S3_BUCKET_NAME"),
	// 	Key:    aws.String(objectKey),
	// 	Body:   bytes.NewReader(decode),
	// }

	// _, err = awsSession.PutObject(uploadParams)

	return "", nil
}

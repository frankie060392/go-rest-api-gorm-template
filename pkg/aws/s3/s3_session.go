package s3

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func InitAWSConnection(id string, key string, token string, region string) *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(id, key, token),
	})
	if err != nil {
		exitErrorf("Can not init session")
	}
	return sess
}

func S3UploadBase64(awsSession *session.Session, bucketName string, base64File string, objectKey string) error {
	decode, err := base64.StdEncoding.DecodeString(base64File)

	if err != nil {
		return err
	}

	uploadInputs := &s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   bytes.NewReader(decode),
	}

	uploader := s3manager.NewUploader(awsSession)

	_, err = uploader.Upload(uploadInputs)

	return err
}

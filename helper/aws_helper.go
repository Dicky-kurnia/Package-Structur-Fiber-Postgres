package helper

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadFileToAWS(file *multipart.FileHeader) (string, error) {
	ImgFile, err := file.Open()
	if err != nil {
		return "", err
	}
	ImgFileExt := filepath.Ext(file.Filename)
	ImgFilename := uuid.NewString()

	err = UploadFileAWS(os.Getenv("AWS_BUCKET"), fmt.Sprintf("%s%s", ImgFilename, ImgFileExt), ImgFile)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%s", ImgFilename, ImgFileExt), nil
}

func UploadFileAWS(bucket, key string, body io.Reader) error {
	config := &aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_KEYID"), os.Getenv("AWS_SECRET"), ""),
	}

	sess := session.Must(session.NewSession(config))

	uploader := s3manager.NewUploader(sess)

	input := &s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   body,
	}

	_, err := uploader.UploadWithContext(context.Background(), input)

	return err
}

func GetUrlFileAWS(bucket, key string) (url string, err error) {
	if key == "" {
		return
	}
	config := &aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_KEYID"), os.Getenv("AWS_SECRET"), ""),
	}
	sess := session.Must(session.NewSession(config))

	// fmt.Println()
	// fmt.Println("mime :", "image/"+GetExtFile(key))
	input := &s3.GetObjectInput{
		Bucket:              aws.String(bucket),
		Key:                 aws.String(key),
		ResponseContentType: aws.String("image/" + filepath.Ext(key)),
	}

	svc := s3.New(sess)

	req, _ := svc.GetObjectRequest(input)

	url, err = req.Presign(15 * time.Minute)
	return
}

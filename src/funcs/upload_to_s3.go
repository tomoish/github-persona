package funcs

import (
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/tomoish/readme/src/aws" // `NewSession` を含むパッケージへのパスを調整
)

func UploadToS3(filePath, bucketName string) (string, error) {
	sess := aws.NewSession() // AWSセッションを作成
	uploader := s3manager.NewUploader(sess)

	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filepath.Base(filePath)),
		Body:   file,
	})
	if err != nil {
		return "", err
	}
	return result.Location, nil // URLを返す
}

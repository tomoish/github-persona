package aws

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/joho/godotenv"
    "log"
    "os"
)

func init() {
    // .envファイルから環境変数を読み込む
    if err := godotenv.Load("../funcs/.env"); err != nil {
        log.Fatal("Error loading .env file")
    }
}

func NewSession() *session.Session {
    sess, err := session.NewSession(&aws.Config{
        Region:      aws.String(os.Getenv("AWS_REGION")),
        Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
    })
    if err != nil {
        panic(err)
    }
    return sess
}

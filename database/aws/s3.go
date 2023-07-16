package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func InitS3() *s3.S3 {
	ACCESS_KEY := os.Getenv("ACCESS_KEY")
	SECRET_KEY := os.Getenv("SECRET_KEY")
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-southeast-1"),
		Credentials: credentials.NewStaticCredentials(ACCESS_KEY, SECRET_KEY, ""),
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	svc := s3.New(sess)
	return svc
}

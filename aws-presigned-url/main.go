package main
//
// import (
// 	"fmt"
// 	"log"
// 	"time"
//
// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/aws/session"
// 	"github.com/aws/aws-sdk-go/service/s3"
// 	"github.com/aws/aws-sdk-go/service/s3/s3manager"
// )
//
// type Config struct {
// 	Region    string
// 	Bucket    string
// 	AccessKey string
// 	SecretKey string
// }
//
// func main() {
// 	cfg, err := loadConfig()
// 	if err != nil {
// 		panic(fmt.Errorf("loading config err %w", err))
// 	}
//
// 	sess, err := session.NewSession(&aws.Config{
// 		Region: aws.String(cfg.Region)},
// 	)
// 	// s := s3.New(sess)
// 	log.Println("The URL is", urlStr)
// }
//
// func loadConfig() (Config, error) {
// 	return Config{
// 		Region:    "",
// 		Bucket:    "",
// 		AccessKey: "",
// 		SecretKey: "",
// 	}, nil
// }

// package main

// import (
// 	"log"

// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/aws/credentials"
// 	"github.com/aws/aws-sdk-go/aws/session"
// 	"github.com/aws/aws-sdk-go/service/s3"
// )

// func main() {
// 	// ctx := context.Background()
// 	endpoint := "https://s3-sgn09.fptcloud.com"
// 	accessKeyID := "122b6c03649fc4b36723"
// 	secretAccessKey := "FJm7hzvP63vaW4d0FOBeqMGA5HIDZSluzKieO4cx"
// 	region := "us-east-1"
// 	// useSSL := true

// 	Config := &aws.Config{
// 		Credentials:      credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
// 		Region:           aws.String(region),
// 		Endpoint:         aws.String(endpoint),
// 		DisableSSL:       aws.Bool(true),
// 		S3ForcePathStyle: aws.Bool(true),
// 	}
// 	sess := session.New(Config)

// 	cli := s3.New(sess)
// 	// bucket, err := cli.ListBuckets(&s3.ListBucketsInput{})
// 	buc, err := cli.GetBucketLocation(&s3.GetBucketLocationInput{Bucket: aws.String("test-backup-bucket")})
// 	if err != nil {
// 		log.Printf("new AWS session failed: %v", err)
// 	}
// 	log.Printf("%s", buc.String())
// 	// buckets := bucket.Buckets
// 	// for _, i := range buckets {
// 	// 	log.Printf("bucket: %s \n", i.Name)
// 	// }

// }

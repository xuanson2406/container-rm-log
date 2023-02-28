package main

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	ctx := context.Background()
	endpoint := "s3-sgn09.fptcloud.com"
	accessKeyID := "122b6c03649fc4b36723"
	secretAccessKey := "FJm7hzvP63vaW4d0FOBeqMGA5HIDZSluzKieO4cx"
	useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
	// log.Printf("S3: %#v\n", minioClient) // minioClient is now set up
	// log.Printf("S3 region: %s", s3utils.GetRegionFromURL(*minioClient.EndpointURL()))
	// location, _ := minioClient.GetBucketLocation(ctx, "test-backup-bucket")
	// log.Printf("S3: %s\n", location)
	// buckets, _ := minioClient.ListBuckets(ctx)
	// for _, i := range buckets {
	// 	log.Printf("bucket: %s \n", i.Name)
	// }
	// bool, err := minioClient.BucketExists(ctx, "stg-awx-util")
	// if err == nil {
	// 	log.Printf("%v", err)
	// }
	// if bool {
	// 	log.Printf("True")
	// } else {
	// 	log.Printf("False")
	// }
	// err = minioClient.MakeBucket(ctx, "test-backup-bucket", minio.MakeBucketOptions{
	// 	Region:        "us-east-1",
	// 	ObjectLocking: false,
	// })
	// if err != nil {
	// 	log.Printf("Error when creating backup bucket: %v", err)
	// }
	// storage, _ := minioClient.EndpointURL().Query().Del()
	// log.Printf("policy: %s", storage)
	// err = minioClient.RemoveBucket(ctx, "test-backup-bucket")
	// if err != nil {
	// 	log.Printf("Error when creating backup bucket: %v", err)
	// }

	// Open the file to be uploaded
	// file, err := os.Open("/home/sondx/Downloads/iqbvjeh3-kubeconfig")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer file.Close()

	// // Get file info and size
	// fileInfo, err := file.Stat()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fileSize := fileInfo.Size()
	// minioClient.PutObject(ctx, "test-backup-bucket", "test-kubeconfig", file, fileSize, minio.PutObjectOptions{})
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	err = minioClient.FGetObject(ctx, "test-backup-bucket", "test-kubeconfig", "/home/sondx/Downloads/test-download", minio.GetObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Object uploaded successfully")
}

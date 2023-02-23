package main

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/minio/minio-go/v7/pkg/s3utils"
)

func main() {
	ctx := context.Background()
	endpoint := "s3-sgn09.fptcloud.com"
	accessKeyID := "7b96df8a6d81cca01be8"
	secretAccessKey := "Z7BEnkbg22L353IjTDhXB+66C54ixUyypofoViiA"
	useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("S3: %#v\n", *minioClient) // minioClient is now set up
	log.Printf("S3 region: %s", s3utils.GetRegionFromURL(*minioClient.EndpointURL()))
	location, _ := minioClient.GetBucketLocation(ctx, "m-fke-prd-backup")
	log.Printf("S3: %s\n", location)
	buckets, _ := minioClient.ListBuckets(ctx)
	for _, i := range buckets {
		log.Printf("bucket: %s \n", i.Name)
	}
}

package main

import (
	"context"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"k8s.io/klog/v2"
)

func main() {
	ctx := context.Background()
	useSSL := true
	// clusterName := "sondx12-yditlfwp"
	// endpoint := "s3-sgn09.fptcloud.com"
	// accessKeyID := "122b6c03649fc4b36723"
	// secretAccessKey := "FJm7hzvP63vaW4d0FOBeqMGA5HIDZSluzKieO4cx"
	// bucketName := "m-fke-stg-cluster-audit-log"
	endpoint := os.Getenv("ENDPOINT")
	accessKeyID := os.Getenv("ACCESS_KEY")
	secretAccessKey := os.Getenv("SECRET_KEY")
	bucketName := os.Getenv("BUCKET_NAME")
	clusterName := os.Getenv("CLUSTER_NAME")
	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	fileName := clusterName + ".log"
	err = os.Remove("/etc/kubernetes/audit-log/" + fileName) // remove a log file
	if err != nil {
		klog.Info(err)
	}

	objectCh := minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{
		Prefix:    clusterName,
		Recursive: true,
	})

	// Collect objects to be removed
	var objectsToDelete []minio.ObjectInfo
	for object := range objectCh {
		klog.Infof("Ojb name: %s\n", object.Key)
		if object.Err != nil {
			log.Fatalln(object.Err)
		}
		objectsToDelete = append(objectsToDelete, object)
	}
	for _, object := range objectsToDelete {
		err := minioClient.RemoveObject(context.Background(), bucketName, object.Key, minio.RemoveObjectOptions{})
		if err != nil {
			log.Fatalln(err)
		}
	}
	klog.Infof("Log file of cluster %s is removed", clusterName)
	return
}

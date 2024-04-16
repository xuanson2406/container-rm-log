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
	endpoint := os.Getenv("ENDPOINT")
	accessKeyID := os.Getenv("ACCESS_KEY")
	secretAccessKey := os.Getenv("SECRET_KEY")
	bucketName := os.Getenv("BUCKET_NAME")
	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
		Region: "us-east-1",
	})
	if err != nil {
		log.Fatalln(err)
	}
	clusterName := os.Getenv("CLUSTER_NAME")
	fileName := clusterName + ".log"
	err = os.Remove("/etc/kubernetes/audit-log/" + fileName) // remove a log file
	if err != nil {
		klog.Info(err)
		// return
	}
	klog.Infof("bucket: %s, prefix: %s", bucketName, clusterName)
	objectCh := minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{
		// Prefix:    clusterName,
		// Recursive: true,
	})
	klog.Infof("%d", len(objectCh))
	if len(objectCh) != 0 {
		// Collect objects to be removed
		var objectsToDelete []minio.ObjectInfo
		for object := range objectCh {
			klog.Infof("Ojb name: %s\n", object.Key)
			if object.Err != nil {
				log.Fatalln(object.Err)
			}
			objectsToDelete = append(objectsToDelete, object)
		}
		klog.Infof("Remove folder log of cluster %s", clusterName)
		for _, object := range objectsToDelete {
			err := minioClient.RemoveObject(context.Background(), bucketName, object.Key, minio.RemoveObjectOptions{})
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
	klog.Infof("Log file of cluster %s is removed", clusterName)
	return
}

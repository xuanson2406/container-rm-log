package main

import (
	"context"
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
	bucketName := "824032ee-7014-42e1-a27b-d23bc20fd08f"
	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
		Region: "us-east-1",
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
	// minioClient.PutObject(ctx, "824032ee-7014-42e1-a27b-d23bc20fd08f", "test-kubeconfig/", nil, 0, minio.PutObjectOptions{})
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// Mở file và đọc dữ liệu vào một đối tượng kiểu io.Reader
	// file, err := os.Open("/home/sondx/Downloads/100MB.bin")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// // Đọc dữ liệu từ file vào một byte array
	// fileInfo, _ := file.Stat()
	// fileSize := fileInfo.Size()
	// buffer := make([]byte, fileSize)
	// _, err = file.Read(buffer)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(fileSize)
	// Đưa dữ liệu từ byte array vào một đối tượng kiểu bytes.Buffer
	// reader := bytes.NewReader(buffer)

	// err = minioClient.FGetObject(ctx, "test-backup-bucket", "test-kubeconfig", "/home/sondx/Downloads/test-download", minio.GetObjectOptions{})
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// minioClient.IN
	// n, err := minioClient.PutObject(ctx, "test-backup-bucket", "file-test", reader, fileSize, minio.PutObjectOptions{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Object uploaded successfully %d byte.", n)

	object := minioClient.ListObjects(ctx, bucketName, minio.ListObjectsOptions{})
	log.Printf("List object in bucket %s: \n", bucketName)
	for k := range object {
		log.Printf("Name: %s; CreatedOn: %s", k.Key, k.LastModified.Format("2006-01-02 15:04:05"))
	}
	// minio.ListObjectsOptions
	// arr := []string{"abc", "dasdad"}
	// str := strings.Join(arr, "/")
	// fmt.Printf(str)

}

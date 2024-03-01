package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	ctx := context.Background()
	endpoint := "s3-hfx03.fptcloud.com"
	accessKeyID := "VQNDG4HVO3GGV6P85YL2"
	secretAccessKey := "uMhpUKfhC3SmhpN0tTJwLoE2YgjGcJ4yDaheVDvS"
	useSSL := true
	// bucketName := "824032ee-7014-42e1-a27b-d23bc20fd08f"
	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
		Region: "",
	})
	if err != nil {
		log.Fatalln(err)
	}
	time.Sleep(5 * time.Second)
	// Load in-cluster Kubernetes configuration
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	// Create a new Kubernetes client using the provided config
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	pods, err := clientset.CoreV1().Pods("kube-system").List(context.TODO(), metav1.ListOptions{LabelSelector: "app=kube-bench"})
	if err != nil {
		panic(err.Error())
	}
	if pods != nil {
		for _, pod := range pods.Items {
			podLogOptions := &corev1.PodLogOptions{
				Container: "kube-bench",
			}
			podLogs, err := clientset.CoreV1().Pods("kube-system").GetLogs(pod.Name, podLogOptions).Stream(context.TODO())
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error opening stream to pod logs: %v\n", err)
				os.Exit(1)
			}
			defer podLogs.Close()

			// Open a file to write the logs
			file, err := os.Create("/var/log/kube-bench/benchmark-clusterid.log")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error creating log file: %v\n", err)
				os.Exit(1)
			}
			defer file.Close()

			// Copy pod logs to the file
			_, err = io.Copy(file, podLogs)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error writing pod logs to file: %v\n", err)
				os.Exit(1)
			}

			fmt.Println("Logs written to file successfully.")
		}
	}
	// var log_file string
	// files, err := os.ReadDir("/var/log/containers")
	// if err != nil {
	// 	fmt.Println("Error reading directory:", err)
	// 	return
	// }

	// // Print the list of files
	// fmt.Println("Files in directory:")
	// for _, file := range files {
	// 	if strings.Contains(file.Name(), "kube-bench-node-fptcloud") && !strings.Contains(file.Name(), "container-pushlog") {
	// 		fmt.Println(file.Name())
	// 		log_file = file.Name()
	// 	}
	// }
	// Open the file to be uploaded
	file, err := os.Open("/var/log/kube-bench/benchmark-clusterid.log")
	if err != nil {
		log.Fatalln("Can not open file: ", err)
	}
	defer file.Close()

	// Get file info and size
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
	}
	// // Đọc dữ liệu từ file vào một byte array
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	_, err = file.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewReader(buffer)

	// minioClient.IN
	_, err = minioClient.PutObject(ctx, "1c7896b7-15ea-424d-9bda-89624019ded5", "file-test", reader, fileSize, minio.PutObjectOptions{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Object uploaded successfully")
}

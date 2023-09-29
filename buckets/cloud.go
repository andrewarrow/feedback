package buckets

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/storage"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go/aws"
	"google.golang.org/api/option"
)

func StoreInAws(data []byte, filename string) {

	bucketName := os.Getenv("PUBLIC_STORAGE_BUCKET")

	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion("us-west-2"),
	)
	fmt.Println(err)

	client := s3.NewFromConfig(cfg)

	dataReader := bytes.NewReader(data)

	putObjectInput := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
		Body:   dataReader,
	}
	putObjectInput.ACL = types.ObjectCannedACLPublicRead

	_, err = client.PutObject(context.Background(), putObjectInput)
	fmt.Println(err)
}

func StoreInGoogle(data []byte, filename string) {
	bucket := os.Getenv("PUBLIC_STORAGE_BUCKET")
	keyPath := os.Getenv("KEY_PATH")

	gcsClient, err := storage.NewClient(context.Background(),
		option.WithCredentialsFile(keyPath))
	fmt.Println(err, bucket, keyPath, len(data), filename)

	w := gcsClient.Bucket(bucket).Object(filename).NewWriter(context.Background())
	w.ContentType = "application/octet-stream"
	_, err = w.Write(data)
	fmt.Println("write", err)
	w.Close()
}

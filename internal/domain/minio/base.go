package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Minio struct {
	Host      string        `json:"host"`
	AccessKey string        `json:"access_key"`
	SecretKey string        `json:"secret_key"`
	Client    *minio.Client `json:"minio"`
}

func InitMinio(host, accessKey, secretKey string, secure bool) (*Minio, error) {
	var _minio Minio

	client, err := minio.New(host, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: secure,
	})
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	bucketNames := []string{"artifacts"}
	for _, bucketName := range bucketNames {
		err = client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			exists, errBucketExists := client.BucketExists(ctx, bucketName)
			if errBucketExists != nil && !exists {
				return nil, err
			}
		}
	}

	_minio.Client = client
	_minio.Host = host
	_minio.AccessKey = accessKey
	_minio.SecretKey = secretKey

	return &_minio, nil
}

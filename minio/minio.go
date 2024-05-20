package minio

import (
	"context"
	"io"
	"net/url"
	"time"

	"cloudstorage"
	minioV7 "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type minio struct {
	client *minioV7.Client
}

type Config struct {
	Endpoint        string
	AccessKey       string
	SecretAccessKey string
	UseSSL          bool
}

type Option func(mn *minioV7.Options) error

func New(cfg Config, opt ...Option) (cloudstorage.CloudStorage, error) {
	minioOpt := &minioV7.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretAccessKey, ""),
		Secure: cfg.UseSSL,
	}

	for _, fn := range opt {
		err := fn(minioOpt)
		if err != nil {
			return nil, err
		}
	}

	minioClient, err := minioV7.New(cfg.Endpoint, minioOpt)
	if err != nil {
		return nil, err
	}
	return &minio{
		client: minioClient,
	}, err
}

func (m *minio) UploadPresign(ctx context.Context, bucketName string, expires time.Duration, file cloudstorage.UploadOption) (*url.URL, error) {
	info, err := m.Upload(ctx, bucketName, file)
	if err != nil {
		return nil, err
	}

	presign, err := m.Presign(ctx, info.Bucket, info.Key, expires)
	if err != nil {
		return nil, err
	}

	return presign, nil
}

func (m *minio) Download(ctx context.Context, bucketName, objectName string) (io.Reader, error) {
	object, err := m.client.GetObject(ctx, bucketName, objectName, minioV7.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	return object, nil
}

func (m *minio) Upload(ctx context.Context, bucketName string, file cloudstorage.UploadOption) (cloudstorage.FileInfo, error) {
	info, err := m.client.PutObject(ctx, bucketName, file.ObjectName, file.File, file.Size, minioV7.PutObjectOptions{ContentType: file.ContentType})
	if err != nil {
		return cloudstorage.FileInfo{}, err
	}

	return cloudstorage.FileInfo{
		Bucket:       info.Bucket,
		Key:          info.Key,
		Size:         info.Size,
		LastModified: info.LastModified,
		Location:     info.Location,
		VersionID:    info.VersionID,
		Expiration:   info.Expiration,
	}, nil
}

func (m *minio) Presign(ctx context.Context, bucketName, objectName string, expires time.Duration) (*url.URL, error) {
	presignedURL, err := m.client.PresignedGetObject(ctx, bucketName, objectName, expires, nil)
	if err != nil {
		return nil, err
	}

	return presignedURL, nil
}

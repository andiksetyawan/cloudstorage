//go:generate go run github.com/vektra/mockery/v2@v2.43.0 --all
package cloudstorage

import (
	"context"
	"io"
	"net/url"
	"time"
)

type FileInfo struct {
	Bucket       string
	Key          string
	Size         int64
	LastModified time.Time
	Location     string
	VersionID    string
	Expiration   time.Time
}

type UploadOption struct {
	File        io.Reader
	ObjectName  string
	ContentType string
	Size        int64
}

type CloudStorage interface {
	Download(ctx context.Context, bucketName, objectName string) (io.Reader, error)
	Upload(ctx context.Context, bucketName string, file UploadOption) (FileInfo, error)
	Presign(ctx context.Context, bucketName, objectName string, expires time.Duration) (*url.URL, error)
	UploadPresign(ctx context.Context, bucketName string, expires time.Duration, file UploadOption) (*url.URL, error)
}

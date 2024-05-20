package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"cloudstorage"
	"cloudstorage/minio"
)

type r2config struct {
	R2Endpoint        string
	R2AccessKey       string
	R2SecretAccessKey string
	R2UseSSL          bool
}

func main() {
	r2cfg := r2config{
		R2Endpoint:        "akdjkajdaiwruwieuweiwwekaaas.r2.cloudflarestorage.com",
		R2AccessKey:       "aasdfakjdfadweienasdkjnkjasd",
		R2SecretAccessKey: "ppqowkjnaksdjfkajndkfajdskfadjfnakjdnfjkabwueybaoueioaniniualnicsd",
		R2UseSSL:          false,
	}

	config := minio.Config{
		Endpoint:        r2cfg.R2Endpoint,
		AccessKey:       r2cfg.R2AccessKey,
		SecretAccessKey: r2cfg.R2SecretAccessKey,
		UseSSL:          r2cfg.R2UseSSL,
	}

	minioClient, err := minio.New(config)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	file, err := os.Open("user1.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stat, _ := file.Stat()
	uploadOption := cloudstorage.UploadOption{
		File:        file,
		ObjectName:  "/user/123123-user1.jpg",
		ContentType: "image/jpeg",
		Size:        stat.Size(),
	}
	presign, _ := minioClient.UploadPresign(ctx, "bucket-a", 1*time.Hour, uploadOption)
	fmt.Printf("presigned url: %s", presign.String())
}

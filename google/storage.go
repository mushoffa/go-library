package google

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"time"

	"cloud.google.com/go/storage"
)

type CloudStorageService interface {
	GetInstance() *storage.Client
	CreateBucket(string) error
	UploadFile(multipart.File, string, string) error
	DownloadFile() error
}

type storage struct {
	client 		*storage.Client
	id 			string
	bucket 		string
}

func NewCloudStorageClient(projectID, bucketName string) (CloudStorageService, error) {
	// Creates Google Cloud Storage client agent
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	defer client.Close()

	return &storage{client, projectID, bucketName}, nil
}

func (g *storage) GetInstance() *storage.Client {
	return g.client
}

func (g *storage) CreateBucket(bucketName string) error {
	ctx := context.Background()

	// Creates the new bucket.
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	// Creates a Bucket instance.
	bucket := g.client.Bucket(bucketName)
	if err := bucket.Create(ctx, g.id, nil); err != nil {
		// log.Fatalf("Failed to create bucket: %v", err)
		return fmt.Errorf("Failed to create bucket: %v", err)
	}

	return nil
}

func (g *storage) UploadFile(file multipart.File, folderName, fileName string) error {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	filePath := folderName + "/" + fileName

	wc := g.client.Bucket(g.bucket).Object(filePath).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
        return fmt.Errorf("io.Copy: %v", err)
    }

    if err := wc.Close(); err != nil {
        return fmt.Errorf("Writer.Close: %v", err)
    }

    // fmt.Fprintf(w, "Blob %v uploaded.\n", filePath)   

	return nil
}

func (g *storage) DownloadFile() error {
	return nil
}
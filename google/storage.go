package google

import (
	"context"
	"fmt"
	"io"
	// "io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

type CloudStorageService interface {
	GetInstance() *storage.Client
	CreateBucket(string) error
	ListBucket() ([]string, error)
	UploadFile(multipart.File, string, string, string) error
	DownloadFile(string, string) ([]byte, error)
	CloseClient()
}

type gcstorage struct {
	client 		*storage.Client
	id 			string
}

func NewCloudStorageClient(projectID string) (CloudStorageService, error) {
	// Creates Google Cloud Storage client agent
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return &gcstorage{client, projectID}, nil
}

// GetInstance ...
func (g *gcstorage) GetInstance() *storage.Client {
	return g.client
}

// CreateBucket ...
func (g *gcstorage) CreateBucket(bucketName string) error {
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

// ListBucket ...
func (g *gcstorage) ListBucket() ([]string, error) {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	var buckets []string
	it := g.client.Buckets(ctx, g.id)
	for {
		battrs, err := it.Next()
        if err == iterator.Done {
            break
        }

        if err != nil {
        	return nil, err
        }

        buckets = append(buckets, battrs.Name)
	}

	return buckets, nil
}


// UploadFile ...
func (g *gcstorage) UploadFile(file multipart.File, bucketName, folderName, fileName string) error {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	filePath := folderName + "/" + fileName

	wc := g.client.Bucket(bucketName).Object(filePath).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
        return fmt.Errorf("io.Copy: %v", err)
    }

    if err := wc.Close(); err != nil {
        return fmt.Errorf("Writer.Close: %v", err)
    }

    // fmt.Fprintf(w, "Blob %v uploaded.\n", filePath)   

	return nil
}

// DownloadFile ...
func (g *gcstorage) DownloadFile(bucketName, fileName string) ([]byte, error) {
	ctx := context.Background()	

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
    defer cancel()

    // Save image to file system
    f, err := os.Create("./ktp2.jpeg")
    if err != nil {
    	return nil, err
    }

    rc, err := g.client.Bucket(bucketName).Object(fileName).NewReader(ctx)
	if err != nil {
		return nil, err
	}

	defer rc.Close()

	if _, err := io.Copy(f, rc); err != nil {
		return nil, err
	}

	if err := f.Close(); err != nil {
		return nil, err
	}

	// dataFile, err := ioutil.ReadAll(rc)
	// if err != nil {
	// 	return nil, err
	// }

	// return dataFile, nil
	return nil, nil
}

func (g *gcstorage) CloseClient() {
	g.client.Close()
}
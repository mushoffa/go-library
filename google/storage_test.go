package google

import (
	"log"
	"testing"

	"github.com/mushoffa/go-library/google"
)

func TestNewCloudStorageClient_Success(t *testing.T) {
	gcstorage, err := google.NewCloudStorageClient("test")
	if err != nil {
		t.Errorf("Error creating cloud storage client: %v", err)
	}

	if gcstorage == nil {
		t.Errorf("Invalid cloud storage client: %v", gcstorage)
	}
}

func TestListBucket_Success(t *testing.T) {
	gcstorage, err := google.NewCloudStorageClient("YOUR_PROJECT_ID")
	if err != nil {
		t.Errorf("Error creating cloud storage client: %v", err)
	}

	buckets, err := gcstorage.ListBucket()
	if err != nil {
		t.Errorf("Error getting buckets: %v", err)
	} else {
		log.Println(buckets)
	}
}
package google

import (
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
package google

import (
	// "bytes"
	"context"
	"fmt"
	// "image"
	// "image/jpeg"
	"log"
	// "os"
	"testing"
	"time"

	"github.com/mushoffa/go-library/google"
	"google.golang.org/api/iterator"
)

const (
	// Change these parameters according to your project environment
	GOOGLE_CLOUD_STORAGE_PROJECT_ID  string = "YOUR_PROJECT_ID"
	GOOGLE_CLOUD_STORAGE_BUCKET_NAME string = "YOUR_BUCKET_NAME"
)

func TestNewCloudStorageClient_Success(t *testing.T) {
	gcstorage, err := google.NewCloudStorageClient(GOOGLE_CLOUD_STORAGE_PROJECT_ID)
	if err != nil {
		t.Errorf("Error creating cloud storage client: %v", err)
	}

	if gcstorage == nil {
		t.Errorf("Invalid cloud storage client: %v", gcstorage)
	}
}

func TestCreateBucket_Success(t *testing.T) {
	gcstorage, err := google.NewCloudStorageClient(GOOGLE_CLOUD_STORAGE_PROJECT_ID)
	if err != nil {
		t.Errorf("Error creating cloud storage client: %v", err)
	}

	if err := gcstorage.CreateBucket("hijra-customer-corporate"); err != nil {
		t.Errorf("Error creating bucket: %v", err)
	}
}

func TestListBuckets_Success(t *testing.T) {
	gcstorage, err := google.NewCloudStorageClient(GOOGLE_CLOUD_STORAGE_PROJECT_ID)
	if err != nil {
		t.Errorf("Error creating cloud storage client: %v", err)
	}

	buckets, err := gcstorage.ListBucket()
	if err != nil {
		t.Errorf("Error getting buckets: %v", err)
	} else {
		for _, bucket := range buckets {
			log.Println(bucket)
		}
	}
}

func TestListFiles_Success(t *testing.T) {
	gcstorage, err := google.NewCloudStorageClient(GOOGLE_CLOUD_STORAGE_PROJECT_ID)
	if err != nil {
		t.Errorf("Error creating cloud storage client: %v", err)
	}

	ctx := context.Background()

	// Creates the new bucket.
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	it := gcstorage.GetInstance().Bucket(GOOGLE_CLOUD_STORAGE_BUCKET_NAME).Objects(ctx, nil)
	for {
		attrs, err := it.Next()
        if err == iterator.Done {
            break
        }

        if err != nil {
        	t.Errorf("%v", err)
        }

        fmt.Println(attrs.Name)
	}
}

func TestDownloadFile_Success(t *testing.T) {
	gcstorage, err := google.NewCloudStorageClient(GOOGLE_CLOUD_STORAGE_PROJECT_ID)
	if err != nil {
		t.Errorf("Error creating cloud storage client: %v", err)
	}

	_, err = gcstorage.DownloadFile(GOOGLE_CLOUD_STORAGE_BUCKET_NAME, "test-2022_03_22_16:23:52/ktp1.jpeg")
	if err != nil {
		t.Errorf("Error downloading file: %v", err)
	} else {
		// fmt.Println(string(file))

		// img, _, err := image.Decode(bytes.NewReader(file))
		// if err != nil {
		// 	t.Errorf("Error decoding image: %v", err)
		// }

		// save, _ := os.Create("./ktp1.jpeg")
		// defer save.Close()

		// var opts jpeg.Options
		// opts.Quality = 1

		// if err := jpeg.Encode(save, img, &opts); err != nil {
		// 	t.Errorf("Error encoding image to JPEG format: %v", err)
		// }		
	}
}
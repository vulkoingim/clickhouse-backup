package gcs

import (
	"cloud.google.com/go/storage"
)

type GCS struct {
	client         *storage.Client
	bucketName     string
	bucketHandle   *storage.BucketHandle
	maxParallelism int
}

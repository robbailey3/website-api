package storage

import (
	"cloud.google.com/go/storage"
	"context"
	"io"
)

type Client interface {
	Upload(ctx context.Context, name string, file io.Reader) error
	GetFile(ctx context.Context, name string) (io.Reader, error)
}

type client struct {
	bucket *storage.BucketHandle
}

func NewClient(bucketName string) (Client, error) {
	storageClient, err := storage.NewClient(context.Background())

	if err != nil {
		return nil, err
	}

	bucket := storageClient.Bucket(bucketName)

	c := &client{
		bucket,
	}

	return c, nil
}

func (c *client) Upload(ctx context.Context, name string, file io.Reader) error {
	wc := c.bucket.Object(name).NewWriter(ctx)

	_, err := io.Copy(wc, file)

	if err != nil {
		return err
	}

	if err := wc.Close(); err != nil {
		return err
	}

	return nil
}

func (c *client) GetFile(ctx context.Context, name string) (io.Reader, error) {
	rc, err := c.bucket.Object(name).NewReader(ctx)

	if err != nil {
		return nil, err
	}

	defer func(rc *storage.Reader) {
		err := rc.Close()
		if err != nil {

		}
	}(rc)

	return rc, nil
}

package image

import (
	vision "cloud.google.com/go/vision/apiv1"
	"context"
	pb "google.golang.org/genproto/googleapis/cloud/vision/v1"
	"io"
)

type VisionAiClient interface {
	DetectLabels(ctx context.Context, file io.Reader, maxResults int) ([]*Label, error)
	DetectFaces(ctx context.Context, file io.Reader, maxResults int) ([]*pb.FaceAnnotation, error)
	DetectImageProperties(ctx context.Context, file io.Reader) (*pb.ImageProperties, error)
}

type visionAiClient struct {
	c *vision.ImageAnnotatorClient
}

func NewVisionClient() (VisionAiClient, error) {
	client, err := vision.NewImageAnnotatorClient(context.Background())

	if err != nil {
		return nil, err
	}

	return &visionAiClient{
		c: client,
	}, nil
}

func (v *visionAiClient) DetectLabels(ctx context.Context, file io.Reader, maxResults int) ([]*Label, error) {
	image, err := vision.NewImageFromReader(file)
	if err != nil {
		return nil, err
	}
	result, err := v.c.DetectLabels(ctx, image, nil, maxResults)

	if err != nil {
		return nil, err
	}

	var labels []*Label

	for _, r := range result {
		labels = append(labels, &Label{Score: r.Score, Description: r.Description})
	}

	return labels, nil
}

func (v *visionAiClient) DetectFaces(ctx context.Context, file io.Reader, maxResults int) ([]*pb.FaceAnnotation, error) {
	image, err := vision.NewImageFromReader(file)
	if err != nil {
		return nil, err
	}
	faces, err := v.c.DetectFaces(ctx, image, nil, maxResults)

	if err != nil {
		return nil, err
	}

	return faces, nil
}

func (v *visionAiClient) DetectImageProperties(ctx context.Context, file io.Reader) (*pb.ImageProperties, error) {
	image, err := vision.NewImageFromReader(file)
	if err != nil {
		return nil, err
	}
	faces, err := v.c.DetectImageProperties(ctx, image, nil)

	if err != nil {
		return nil, err
	}

	return faces, nil
}

package detector

import (
	"context"
	"thesis_back/internal/transport/grpc/detector/proto"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type detectorClient struct {
	client proto.ObjectDetectorClient
	logger *zap.Logger
}

type IDetectorClient interface {
	Detect(imageID string, imageBytes []byte) (*proto.DetectResponse, error)
}

func NewDetectorClient(address string, logger *zap.Logger) (IDetectorClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logger.Error("grpc conn error", zap.Error(err))
		return nil, err
	}

	client := proto.NewObjectDetectorClient(conn)
	return &detectorClient{
		client: client,
		logger: logger,
	}, nil
}

func (d *detectorClient) Detect(imageID string, imageBytes []byte) (*proto.DetectResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	req := &proto.DetectRequest{
		ImageId:   imageID,
		ImageData: imageBytes,
	}

	resp, err := d.client.Detect(ctx, req)
	if err != nil {
		d.logger.Error("grpc detect error", zap.Error(err))
		return nil, err
	}

	return resp, nil
}

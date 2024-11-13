package server

import (
	"context"
	"fmt"
	pb "happenedapi/gen/protos/v1"
	"happenedapi/gen/protos/v1/happenedv1connect"
	"log"
	"log/slog"
	"time"

	"connectrpc.com/connect"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	storage_go "github.com/supabase-community/storage-go"
)

// Ensure interface satisfaction
var _ happenedv1connect.HappenedServiceHandler = (*HappenedServer)(nil)

const (
	HappenedBucketName = "happened-bucket"
)

type HappenedServer struct {
	s3Client *s3.Client
	storageClient *storage_go.Client
}

func New(s3Client *s3.Client, storageClient *storage_go.Client) *HappenedServer {
	return &HappenedServer{
		s3Client: s3Client,
		storageClient: storageClient,
	}
}

func (s *HappenedServer) GetUploadImageURL(
	ctx context.Context,
	req *connect.Request[pb.GetUploadImageURLRequest]) (*connect.Response[pb.GetUploadImageURLResponse], error) {
	imageKey := req.Msg.ImageKey
	slog.Info("generating presigned", slog.String("imageKey", imageKey))
	presignClient := s3.NewPresignClient(s.s3Client)

	presignedPutRequest, err := presignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket:  aws.String(HappenedBucketName),
		Key:     aws.String(imageKey),
		ContentType: aws.String("image/*"),
		Expires: aws.Time(time.Now().Add(time.Minute * 5)),
	})
	if err != nil {
		return nil, err
	}

	


	// resp, err := s.storageClient.CreateSignedUploadUrl(HappenedBucketName, imageKey)
	// if err != nil {
	// 	return nil, err
	// }
	
	log.Println("presigned url", presignedPutRequest.URL)
	response := connect.NewResponse(&pb.GetUploadImageURLResponse{
		UploadUrl: presignedPutRequest.URL,
	})

	return response, nil
}

func (s *HappenedServer) CreateEvent(ctx context.Context, c *connect.Request[pb.CreateEventRequest]) (*connect.Response[pb.CreateEventResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (s *HappenedServer) Greet(
	ctx context.Context,
	req *connect.Request[pb.GreetRequest]) (*connect.Response[pb.GreetResponse], error) {
	log.Println("Request headers", req.Header())

	res := connect.NewResponse(&pb.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}

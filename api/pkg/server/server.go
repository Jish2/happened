package server

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	pb "happenedapi/gen/protos/v1"
	"happenedapi/gen/protos/v1/happenedv1connect"
	"log"
	"time"
)

// Ensure interface satisfaction
var _ happenedv1connect.HappenedServiceHandler = (*HappenedServer)(nil)

const (
	HappenedBucketName = "happened-bucket"
)

type HappenedServer struct {
	s3Client *s3.Client
}

func (s *HappenedServer) GetUploadImageURL(
	ctx context.Context,
	req *connect.Request[pb.GetUploadImageURLRequest]) (*connect.Response[pb.GetUploadImageURLResponse], error) {
	//TODO implement me

	imageKey := uuid.New().String()
	presignClient := s3.NewPresignClient(s.s3Client)
	presignedPutRequest, err := presignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket:  aws.String(HappenedBucketName),
		Key:     aws.String(imageKey),
		Expires: aws.Time(time.Now().Add(time.Minute * 5)),
	})
	if err != nil {
		return nil, err
	}

	log.Println("presignedPutRequest", presignedPutRequest)
	response := connect.NewResponse(&pb.GetUploadImageURLResponse{
		PresignedUrl: presignedPutRequest.URL,
	})

	return response, nil
}

func New(s3Client *s3.Client) *HappenedServer {
	return &HappenedServer{
		s3Client: s3Client,
	}
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

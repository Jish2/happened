package server

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	happenedv1 "happenedapi/gen/protos/v1"
	"happenedapi/gen/protos/v1/happenedv1connect"
	"log"
)

// Ensure interface satisfaction
var _ happenedv1connect.HappenedServiceHandler = (*HappenedServer)(nil)

type HappenedServer struct {
	s3Client *s3.Client
}

func New(s3Client *s3.Client) *HappenedServer {
	return &HappenedServer{
		s3Client: s3Client,
	}
}

func (s *HappenedServer) UploadImage(ctx context.Context, req *connect.Request[happenedv1.UploadImageRequest]) (*connect.Response[happenedv1.UploadImageResponse], error) {

	//TODO implement me
	panic("implement me")
}

func (s *HappenedServer) CreateEvent(ctx context.Context, c *connect.Request[happenedv1.CreateEventRequest]) (*connect.Response[happenedv1.CreateEventResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (s *HappenedServer) Greet(
	ctx context.Context,
	req *connect.Request[happenedv1.GreetRequest]) (*connect.Response[happenedv1.GreetResponse], error) {
	log.Println("Request headers", req.Header())

	res := connect.NewResponse(&happenedv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}

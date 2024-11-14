package server

import (
	"context"
	"errors"
	"fmt"
	happenedv1 "happenedapi/gen/protos/v1"
	"happenedapi/gen/protos/v1/happenedv1connect"
	"log"

	"connectrpc.com/connect"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Ensure interface satisfaction
var _ happenedv1connect.HappenedServiceHandler = (*HappenedServer)(nil)

type HappenedServer struct {
	s3Client *s3.Client
}

func (s *HappenedServer) GetUploadImageURL(
	ctx context.Context,
	req *connect.Request[happenedv1.GetUploadImageURLRequest]) (*connect.Response[happenedv1.GetUploadImageURLResponse], error) {
	//TODO implement me
	panic("implement me")
}

func New(s3Client *s3.Client) *HappenedServer {
	return &HappenedServer{
		s3Client: s3Client,
	}
}

func (s *HappenedServer) CreateEvent(ctx context.Context, c *connect.Request[happenedv1.CreateEventRequest]) (*connect.Response[happenedv1.CreateEventResponse], error) {
	//TODO implement me
	panic("implement me")
}

func (s *HappenedServer) Greet(
	ctx context.Context,
	req *connect.Request[happenedv1.GreetRequest]) (*connect.Response[happenedv1.GreetResponse], error) {
	log.Println("Request headers", req.Header())

	if req.Msg.Name == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("name is required"))
	}

	res := connect.NewResponse(&happenedv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}

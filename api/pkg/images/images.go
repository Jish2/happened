package images

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

const (
	HappenedBucketName = "happened-bucket"
)

type Service struct {
	s3Client *s3.Client
}

func NewService(s3Client *s3.Client) *Service {
	return &Service{
		s3Client: s3Client,
	}
}

type UploadURLResult struct {
	Method        string
	URL           string
	SignedHeaders http.Header
}

func (s *Service) CreateUploadURL(
	ctx context.Context,
	imageKey string,
) (*UploadURLResult, error) {

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
	result := &UploadURLResult{
		Method:        presignedPutRequest.Method,
		URL:           presignedPutRequest.URL,
		SignedHeaders: presignedPutRequest.SignedHeader,
	}

	return result, nil
}

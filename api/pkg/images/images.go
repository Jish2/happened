package images

import (
	"context"
	"log"
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
	Method       string
	URL          string
	SignedHeader map[string]string
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

	headers := make(map[string]string, len(presignedPutRequest.SignedHeader))
	for key, values := range presignedPutRequest.SignedHeader {
		if len(values) > 0 {
			headers[key] = values[0]
		} else {
			// should record this as a metric
		}
	}

	log.Println("presignedPutRequest", presignedPutRequest)
	result := &UploadURLResult{
		Method:       presignedPutRequest.Method,
		URL:          presignedPutRequest.URL,
		SignedHeader: headers,
	}

	return result, nil
}

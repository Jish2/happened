package images

import (
	"context"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

const (
	HappenedBucketName = "happened-bucket"
)

type Service struct {
	s3PresignClient S3PresignClient
}

//go:generate mockery --name=S3PresignClient
type S3PresignClient interface {
	PresignPutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.PresignOptions)) (*v4.PresignedHTTPRequest, error)
}

func NewService(s3PresignClient S3PresignClient) *Service {
	return &Service{
		s3PresignClient: s3PresignClient,
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

	presignedPutRequest, err := s.s3PresignClient.PresignPutObject(ctx, &s3.PutObjectInput{
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
			log.Println(values)
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

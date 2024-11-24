package test

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/danielgtaylor/huma/v2/humatest"
	"github.com/stretchr/testify/require"
	"happenedapi/pkg/images"
	"happenedapi/pkg/server"
	"io"
	"testing"
)

func MakeTestServer(t *testing.T) humatest.TestAPI {
	ctx := context.Background()
	_, api := humatest.New(t)
	cfg, err := config.LoadDefaultConfig(ctx)
	require.NoError(t, err)
	s3Client := s3.NewFromConfig(cfg)
	imageService := images.NewService(s3Client)
	server.RegisterAPI(api, nil, imageService)

	return api
}

func DecodeAs[T any](body io.Reader, t *testing.T) T {
	var data T
	err := json.NewDecoder(body).Decode(&data)
	require.NoError(t, err)
	return data
}

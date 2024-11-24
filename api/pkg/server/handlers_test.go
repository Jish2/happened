package server

import (
	"context"
	"encoding/json"
	"fmt"
	"happenedapi/pkg/images"
	"net/http"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/danielgtaylor/huma/v2/humatest"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_CreateUploadURLHandler(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc string
		// params is the input sent when creating a presigned upload url
		params map[string]any
		// uploadFilePath is the path of the file to upload to the presigned URL
		uploadFilePath string

		expectedStatusCode int
		expectedErr        bool
		expectedMethod     string
	}{
		{
			desc: "happy path: image key and valid image",
			params: map[string]any{
				"image_key": fmt.Sprintf("%s.jpeg", uuid.New().String()),
			},
			expectedStatusCode: http.StatusOK,
			expectedErr:        false,
			expectedMethod:     http.MethodPut,
		},
		// {
		// 	desc: "file is not a valid image",
		// },
		// {
		// 	desc: "file is a png",
		// },
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			ctx := context.Background()
			_, api := humatest.New(t)
			cfg, err := config.LoadDefaultConfig(ctx)
			require.NoError(t, err)
			s3Client := s3.NewFromConfig(cfg)
			imageService := images.NewService(s3Client)

			RegisterAPI(api, nil, imageService)

			resp := api.Get("/create-upload-url", tc.params)
			assert.Equal(t, tc.expectedStatusCode, resp.Code)

			if !tc.expectedErr {
				var body CreateUploadURLBody
				err = json.NewDecoder(resp.Body).Decode(&body)
				assert.NoError(t, err)

				assert.Equal(t, tc.expectedMethod, body.Method)
				assert.NotEmpty(t, body.UploadURL)
			}
		})
	}
}

package server_test

import (
	"fmt"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/danielgtaylor/huma/v2/humatest"
	"github.com/stretchr/testify/mock"
	"happenedapi/pkg/images"
	"happenedapi/pkg/images/mocks"
	"happenedapi/pkg/server"
	"happenedapi/pkg/test"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_CreateUploadURLHandler(t *testing.T) {
	t.Parallel()

	var (
		successfulPresignResponse = &v4.PresignedHTTPRequest{
			// example presign url
			URL:    "https://happened-bucket.s3.us-west-1.amazonaws.com/635b6887-665d-41fe-b471-3775f0eb4630.jpeg?X-Amz-Algorithm=AWS4-HMAC-SHA256\u0026X-Amz-Credential=AKIAZMUGYJXGC2SYP6BY%2F20241128%2Fus-west-1%2Fs3%2Faws4_request\u0026X-Amz-Date=20241128T225728Z\u0026X-Amz-Expires=900\u0026X-Amz-SignedHeaders=expires%3Bhost\u0026x-id=PutObject\u0026X-Amz-Signature=7d02a328833aeb4785d3a13627f78c8eb29b913f9463ebd6e90f0e3647e0e467",
			Method: http.MethodPut,
			SignedHeader: http.Header{
				"Expires": {"Thu, 28 Nov 2024 23:02:28 GMT"},
				"Host":    {"happened-bucket.s3.us-west-1.amazonaws.com"},
			},
		}
	)

	tests := []struct {
		desc string
		// params is the input sent when creating a pre-signed upload url
		params map[string]any
		// uploadFilePath is the path of the file to upload to the pre-signed URL
		uploadFilePath    string
		s3PresignResponse *v4.PresignedHTTPRequest
		s3PresignError    awserr.Error

		expectedStatusCode int
		expectedErr        bool
		expectedMethod     string
	}{
		{
			desc: "happy path: image key and valid image",
			params: map[string]any{
				"image_key": fmt.Sprintf("%s.jpeg", uuid.New().String()),
			},
			s3PresignResponse: successfulPresignResponse,

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
			mockS3PresignClient := mocks.NewS3PresignClient(t)

			mockS3PresignClient.
				On("PresignPutObject", mock.Anything, mock.Anything, mock.Anything).
				Return(tc.s3PresignResponse, tc.s3PresignError)

			api := server.New(test.MakeLocalDB(t), images.NewService(mockS3PresignClient))
			testapi := humatest.Wrap(t, api)

			resp := testapi.Get("/create-upload-url", tc.params)
			assert.Equal(t, tc.expectedStatusCode, resp.Code)

			//require.True(t, false, "stop")
			if !tc.expectedErr {
				body := test.DecodeAs[server.CreateUploadURLBody](resp.Body, t)
				assert.Equal(t, tc.expectedMethod, body.Method)
				assert.NotEmpty(t, body.UploadURL)
				assert.NotNil(t, body.SignedHeaders)
			}
		})
	}
}

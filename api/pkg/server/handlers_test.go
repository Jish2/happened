package server_test

import (
	"fmt"
	"happenedapi/pkg/server"
	test "happenedapi/pkg/tests"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_CreateUploadURLHandler(t *testing.T) {
	t.Parallel()

	tests := []struct {
		desc string
		// params is the input sent when creating a pre-signed upload url
		params map[string]any
		// uploadFilePath is the path of the file to upload to the pre-signed URL
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
			api := test.MakeTestServer(t)

			resp := api.Get("/create-upload-url", tc.params)
			assert.Equal(t, tc.expectedStatusCode, resp.Code)

			if !tc.expectedErr {
				body := test.DecodeAs[server.CreateUploadURLBody](resp.Body, t)
				assert.Equal(t, tc.expectedMethod, body.Method)
				assert.NotEmpty(t, body.UploadURL)
			}
		})
	}
}

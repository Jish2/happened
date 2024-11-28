package server

import (
	"context"
	"fmt"
	"happenedapi/pkg/images"
	"log/slog"
)

type GreetingInput struct {
	Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
}

type HumaHandler[I, O any] func(ctx context.Context, input *I) (*O, error)

// GreetingOutput represents the greeting operation response.
type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

func GreetHandler() HumaHandler[GreetingInput, GreetingOutput] {
	return func(ctx context.Context, input *GreetingInput) (*GreetingOutput, error) {
		resp := &GreetingOutput{}
		resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
		return resp, nil
	}
}

func protectedGreetHandler() HumaHandler[GreetingInput, GreetingOutput] {
	return func(ctx context.Context, input *GreetingInput) (*GreetingOutput, error) {
		resp := &GreetingOutput{}
		resp.Body.Message = fmt.Sprintf("Hello protected, %s!", input.Name)
		return resp, nil
	}
}

type CreateUploadURLRequest struct {
	Body struct {
		ImageKey string `json:"image_key" example:"my-image.jpeg"`
	}
}

type CreateUploadURLBody struct {
	Method        string            `json:"method"`
	UploadURL     string            `json:"upload_url"`
	SignedHeaders map[string]string `json:"signed_headers"`
}

type CreateUploadURLResponse struct {
	Body CreateUploadURLBody
}

func CreateUploadURLHandler(imageService *images.Service) HumaHandler[CreateUploadURLRequest, CreateUploadURLResponse] {
	return func(ctx context.Context, request *CreateUploadURLRequest) (*CreateUploadURLResponse, error) {
		slog.Info("CreateUploadURLHandler", slog.Any("request", request))
		result, err := imageService.CreateUploadURL(ctx, request.Body.ImageKey)
		if err != nil {
			return nil, err
		}
		resp := &CreateUploadURLResponse{
			Body: CreateUploadURLBody{
				UploadURL:     result.URL,
				Method:        result.Method,
				SignedHeaders: result.SignedHeader,
			},
		}
		return resp, nil
	}
}

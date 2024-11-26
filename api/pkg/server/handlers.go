package server

import (
	"context"
	"fmt"
)

type HumaHandler[I, O any] func(ctx context.Context, input *I) (*O, error)

// GreetingOutput represents the greeting operation response.
type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

type GreetingInput struct {
	Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
}

func greetHandler() HumaHandler[GreetingInput, GreetingOutput] {
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

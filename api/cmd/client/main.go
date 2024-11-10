package main

import (
	"context"
	"log"
	"net/http"

	greetv1 "happenedapi/gen/greet/v1"
	"happenedapi/gen/greet/v1/greetv1connect"

	"connectrpc.com/connect"
)

func main() {
	client := greetv1connect.NewGreetServiceClient(
		http.DefaultClient,
		"http://localhost:8080",
	)

	res, err := client.Greet(
		context.Background(),
		connect.NewRequest(&greetv1.GreetRequest{Name: "Jane"}),
	)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res.Msg.Greeting)
}

package main

import (
	"context"
	happenedv1 "happenedapi/gen/protos/v1"
	"happenedapi/gen/protos/v1/happenedv1connect"
	"log"
	"net/http"

	"connectrpc.com/connect"
)

func main() {

	client := happenedv1connect.NewHappenedServiceClient(http.DefaultClient,
		"http://localhost:8080",
	)

	res, err := client.Greet(
		context.Background(),
		connect.NewRequest(&happenedv1.GreetRequest{Name: "Jane"}),
	)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res.Msg.Greeting)
}

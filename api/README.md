
Bootstrap Go server using ConnectRPC quickstart guide.
Example code for server and client.

Start server using make watch

Using HTTP
```
curl \
    --header "Content-Type: application/json" \
    --data '{"name": "Jane"}' \
    http://localhost:8080/happened_service.v1.HappenedService/Greet
``` 
Using gRPC
```
grpcurl \
    -protoset <(buf build -o -) -plaintext \
    -d '{"name": "Jane"}' \
    localhost:8080 greet_service.v1.HappenedService/Greet
```

FROM golang:1.22.2 AS builder
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api ./cmd/api

FROM alpine:3.20.3

WORKDIR /app

COPY --from=builder /app/api .

EXPOSE 8080
CMD ["./api"]
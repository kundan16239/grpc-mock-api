FROM golang:1.24.6-alpine3.22 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main ./cmd/grpc_server/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 50051
CMD ["./main"]
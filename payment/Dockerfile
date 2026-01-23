# Build stage
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum* ./
COPY payment/go.mod payment/go.sum* ./payment/

RUN cd payment && go mod download

COPY . .

RUN cd payment && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/payment/main .

EXPOSE 50052

CMD ["./main"]

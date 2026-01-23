# Build stage
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum* ./
COPY shipping/go.mod shipping/go.sum* ./shipping/

RUN cd shipping && go mod download

COPY . .

RUN cd shipping && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/shipping/main .

EXPOSE 50053

CMD ["./main"]

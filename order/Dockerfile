# Build stage
FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum* ./
COPY order/go.mod order/go.sum* ./order/

RUN cd order && go mod download

COPY . .

RUN cd order && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/order/main .

EXPOSE 50051

CMD ["./main"]

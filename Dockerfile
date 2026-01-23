# syntax=docker/dockerfile:1

ARG GO_VERSION=1.24
ARG SERVICE=order
ARG SERVICE_PORT=50051

FROM golang:${GO_VERSION} AS builder

ARG SERVICE
WORKDIR /app

# Bring proto definitions into the build context before resolving modules
COPY microservices-proto ./microservices-proto

# Prime module cache with the target service's go.mod/go.sum
COPY microservices/${SERVICE}/go.mod microservices/${SERVICE}/go.sum ./microservices/${SERVICE}/

WORKDIR /app/microservices/${SERVICE}
RUN go mod download

# Copy the full service source after dependencies are cached
COPY microservices/${SERVICE} .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/${SERVICE} ./cmd

FROM alpine:3.20

ARG SERVICE
ARG SERVICE_PORT
RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /app/bin/${SERVICE} /app/main

EXPOSE ${SERVICE_PORT}
CMD ["/app/main"]

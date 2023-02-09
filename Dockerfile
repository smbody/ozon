# Base environment (alias: base)
FROM golang:1.20 AS base

# Development backend environment
FROM base as dev
WORKDIR /root
RUN apt-get update && apt-get install -y fswatch
RUN go install github.com/go-delve/delve/cmd/dlv@latest
WORKDIR /ozon

FROM base as go-builder
COPY . /ozon
WORKDIR /ozon
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o ozon-server

#Production backend environment
FROM scratch as backend
COPY --from=go-builder /ozon/config/config.yml /ozon/config/config.yml
COPY --from=go-builder /ozon/ozon-server /ozon/ozon-server
WORKDIR /ozon

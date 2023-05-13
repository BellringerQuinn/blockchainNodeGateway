# build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go mod verify
COPY . .
RUN go build -o nodeGateway -v ./cmd


# run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/nodeGateway .

HEALTHCHECK  --interval=5m --timeout=3s \
    CMD wget --no-verbose --tries=1 --spider http://localhost/ || exit 1
EXPOSE 8080
CMD /app/nodeGateway

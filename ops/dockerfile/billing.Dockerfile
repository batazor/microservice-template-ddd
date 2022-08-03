FROM golang:1.19-alpine as builder

ARG CI_COMMIT_TAG

WORKDIR /go/src/microservice-template-ddd
COPY . .

# Load dependencies
RUN go mod vendor

# Build project
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
  go build \
  -a \
  -mod vendor \
  -ldflags "-X main.CI_COMMIT_TAG=$CI_COMMIT_TAG" \
  -installsuffix cgo \
  -trimpath \
  -o app ./cmd/billing

FROM alpine:latest

# 50051: gRPC server
EXPOSE 50051

# Install dependencies
RUN \
    apk update && \
    apk add curl && \
    rm -rf /var/cache/apk/*

RUN addgroup -S billing && adduser -S -g billing billing
USER billing

WORKDIR /app/
COPY --from=builder /go/src/microservice-template-ddd/app .
CMD ["./app"]

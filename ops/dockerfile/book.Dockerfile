FROM golang:1.15-alpine as builder

ARG CI_COMMIT_TAG

WORKDIR /go/src/robovoice-template
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
  -o app ./cmd/book

FROM alpine:latest

# 50051: gRPC server
EXPOSE 50051

# Install dependencies
RUN \
    apk update && \
    apk add curl && \
    rm -rf /var/cache/apk/*

RUN addgroup -S book && adduser -S -g book book
USER book

WORKDIR /app/
COPY --from=builder /go/src/robovoice-template/app .
CMD ["./app"]

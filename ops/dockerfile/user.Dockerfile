FROM golang:1.17-alpine as builder

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
  -o app ./cmd/user

FROM alpine:latest

# 50051: gRPC server
EXPOSE 50051

# Install dependencies
RUN \
    apk update && \
    apk add curl && \
    rm -rf /var/cache/apk/*

RUN addgroup -S user && adduser -S -g user user
USER user

WORKDIR /app/
COPY --from=builder /go/src/robovoice-template/app .
CMD ["./app"]

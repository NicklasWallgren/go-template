# syntax=docker/dockerfile:1.4
FROM golang:1.18-alpine as builder

RUN apk --no-cache add git ca-certificates
ENV CGO_ENABLED=0

WORKDIR /app

# Download all necessary dependencies in a separate layer
COPY go.mod go.sum ./
RUN --mount=type=cache,mode=0777,target=/go/pkg/mod go mod download

# Add the source code
COPY ./ .

RUN go build -o ./build/go-template .

FROM alpine:3.9

WORKDIR /app

COPY --from=builder --chmod=0755 /app/build/go-template /app/go-template
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/.env.docker /app/.env

# This container exposes port 8080
EXPOSE 8080

ENTRYPOINT ["./go-template", "start"]
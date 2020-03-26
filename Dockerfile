# Builder stage
FROM golang:1.14.0-alpine AS builder

WORKDIR /app/src

COPY go.mod go.sum .
RUN go mod download

COPY . .
RUN go build -o /app/bin/ .


# Deploy stage
FROM alpine:latest

RUN apk update
RUN apk add --no-cache bash
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/bin/vizix .

EXPOSE 9999

ENTRYPOINT ["/app/bin/vizix"]
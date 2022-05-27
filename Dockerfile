FROM golang:1.18-alpine AS builder
RUN apk --update add ca-certificates

WORKDIR /app
ENV LANG en_US.UTF-8

COPY go.mod go.sum ./
RUN go mod download
RUN go get github.com/golang/mock/mockgen

COPY . .

COPY main.go .
RUN CGO_ENABLED=0 go build -o /build/confetti

FROM alpine:latest

RUN apk add --no-cache --upgrade bash

COPY --from=builder /build/ /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY migrations /migrations
COPY scripts/run.sh /app/run.sh

VOLUME /etc/confetti

EXPOSE 8000

CMD ["sh", "/app/run.sh"]

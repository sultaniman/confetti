FROM golang:1.16-alpine AS builder
RUN apk --update add ca-certificates

WORKDIR /app
ENV LANG en_US.UTF-8

COPY go.mod go.sum ./
RUN go mod download
RUN go get github.com/golang/mock/mockgen

COPY . .

COPY main.go .
RUN CGO_ENABLED=0 go build -o /build/confetti

FROM scratch
COPY --from=builder /build/ /

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY migrations /migrations

VOLUME /etc/confetti

EXPOSE 8000

CMD ["./scripts/run.sh"]

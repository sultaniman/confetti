FROM golang:1.16-alpine AS builder
RUN apk --update add ca-certificates

WORKDIR /app
ENV LANG en_US.UTF-8

COPY go.mod go.sum ./
RUN go mod download
RUN go get github.com/golang/mock/mockgen

COPY . .
RUN go generate ./...

COPY main.go .
RUN CGO_ENABLED=0 go build -o /build/getout

FROM scratch
COPY --from=builder /build/ /

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY migrations /migrations

VOLUME /etc/getout

EXPOSE 8000

ENTRYPOINT ["/getout"]
CMD ["serve"]
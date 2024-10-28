FROM golang:1.22.5 AS builder

ENV GOPATH=/
ENV GOTOOLCHAIN=local
WORKDIR /app

COPY ./ ./
RUN apt-get update && \ 
    go mod download && \ 
    go build -o server cmd/app/main.go

FROM alpine:latest

COPY --from=builder /app/server /app/server

CMD ["/app/server"]

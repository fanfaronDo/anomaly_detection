FROM golang:1.22.5 AS builder

ENV GOPATH=/
ENV GOTOOLCHAIN=local
WORKDIR /app

COPY . .

RUN go mod download && \ 
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/client /app/cmd/client/main.go


FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/client .
RUN chmod +x client

CMD ["./client"]
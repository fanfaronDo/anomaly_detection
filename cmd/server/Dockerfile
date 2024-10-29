FROM golang:1.22.5 AS builder

ENV GOPATH=/
ENV GOTOOLCHAIN=local
WORKDIR /app

COPY . .

RUN go mod download && \ 
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/server /app/cmd/server/main.go


FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/server .
RUN chmod +x server

CMD ["/app/server"]

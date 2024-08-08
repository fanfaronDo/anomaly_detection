FROM ubuntu:latest

WORKDIR /app

COPY ./server .

CMD ["./server"]



FROM golang:1.18 AS builder
WORKDIR /application
COPY application ./
RUN go build main.go
ENTRYPOINT ["./main"]
EXPOSE 8080
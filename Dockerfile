FROM golang:1.18 AS builder
WORKDIR /application
COPY application ./
EXPOSE 8080
RUN go mod download
# CMD [ "go", "run", "main.go" ]
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder application ./
ENTRYPOINT [ "./app" ]
FROM golang:1.14.4-alpine

WORKDIR /go/src/

COPY . .

RUN go clean --modcache
RUN GOOS=linux go build server.go

EXPOSE 9000
ENTRYPOINT ["./server"]
FROM golang:alpine

ADD ./src /go/src/app
WORKDIR /go/src/app
EXPOSE 3000
CMD ["go", "run", "server.go"]
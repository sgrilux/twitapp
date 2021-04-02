FROM golang:1.16.2

LABEL maintainer="sgrilux@gmail.com"

WORKDIR $GOPATH/src/github.com/sgrilux/twitapp

COPY . .
RUN go mod download
RUN go build -o bin/twitapp cmd/main.go

ENTRYPOINT ["bin/twitapp"]
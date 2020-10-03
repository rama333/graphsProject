FROM golang:latest

WORKDIR /app

COPY ./ /app


RUN apt-get update && apt-get install -y libaio1 wget unzip
RUN go get github.com/githubnemo/CompileDaemon
RUN go get -u github.com/gin-gonic/gin
RUN go get github.com/googollee/go-socket.io
RUN go get github.com/gomodule/redigo/redis


WORKDIR /app/cmd/graphs/

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main
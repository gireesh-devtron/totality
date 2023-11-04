FROM golang:1.20 AS build-env

RUN apt update
RUN apt install git gcc musl-dev make -y

WORKDIR /go/src/github.com/gireesh-devtron/totality/server
ADD . /go/src/github.com/gireesh-devtron/totality/server

FROM ubuntu
RUN apt update
RUN apt install ca-certificates -y
RUN apt clean autoclean
RUN apt autoremove -y && rm -rf /var/lib/apt/lists/*
COPY --from=build-env  /go/src/github.com/gireesh-devtron/totality/server/server .

CMD ["./server"]
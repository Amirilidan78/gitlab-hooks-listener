FROM golang:1.18.3-bullseye

WORKDIR /var/www/html/api

ENV GO111MODULE=on
ENV GOROOT "/usr/local/go"
ENV GOPATH "/root/go"
ENV PATH "$PATH:$GOROOT/bin"
ENV PATH "$PATH:$GOPATH/bin"

COPY ./go.mod ./go.mod

COPY ./go.sum ./go.sum

RUN go mod download

RUN go mod tidy

COPY . .

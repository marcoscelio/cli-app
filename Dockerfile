FROM golang:alpine
RUN apk add build-base
#RUN mkdir -p /go/src
RUN mkdir /app
COPY ./ /app
WORKDIR /app
CMD go test cli_test.go cli_create_test.go
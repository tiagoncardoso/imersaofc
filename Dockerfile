FROM golang:1.17-alpine

WORKDIR /go/src

CMD ["tail", "-f", "/dev/null"]
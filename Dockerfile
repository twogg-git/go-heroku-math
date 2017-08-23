FROM golang:alpine

ADD /main.go /go/src/gomath/

RUN go install gomath

ENTRYPOINT /go/bin/gomath

EXPOSE 8080

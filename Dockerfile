FROM golang

ADD /main.go /go/src/gomath/

WORKDIR /go/src/gomath/

RUN go get github.com/gonum/stat

RUN go install gomath

ENTRYPOINT /go/bin/gomath

EXPOSE 8080

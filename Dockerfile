FROM golang

ADD /char.go /go/src/gomath/

WORKDIR /go/src/gomath/

RUN go get -v .

RUN go install gomath

ENTRYPOINT /go/bin/gomath

EXPOSE 8080

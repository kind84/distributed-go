FROM golang:1.12

ADD . /go/src/hello

RUN go install hello

ENV NAME Bob

ENTRYPOINT /go/bin/hello

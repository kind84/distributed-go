FROM golang:1.12.3

ADD . /go/src/github.com/kind84/distributed-go/chapter5/goophr/librarian

WORKDIR /go/src/github.com/kind84/distributed-go/chapter5/goophr/librarian

RUN go install github.com/kind84/distributed-go/chapter5/goophr/librarian

ENTRYPOINT [ "/go/bin/librarian" ]

EXPOSE 9000
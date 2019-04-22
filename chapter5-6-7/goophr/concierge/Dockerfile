FROM golang:1.12.3

ADD . /go/src/github.com/kind84/distributed-go/chapter5/goophr/concierge

WORKDIR /go/src/github.com/kind84/distributed-go/chapter5/goophr/concierge

RUN go install github.com/kind84/distributed-go/chapter5/goophr/concierge

ENTRYPOINT [ "/go/bin/concierge" ]

EXPOSE 9000
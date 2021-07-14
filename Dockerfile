FROM golang:1.16

COPY . /go/src/app
WORKDIR /go/src/app
RUN go get app
RUN go install
ENTRYPOINT ["/go/bin/app"]

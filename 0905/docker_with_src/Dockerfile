FROM golang:alpine

WORKDIR /go/src/hello
COPY src/main.go .
RUN go install

CMD /go/bin/hello

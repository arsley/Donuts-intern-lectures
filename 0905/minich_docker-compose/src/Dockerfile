FROM golang:alpine

RUN apk add git

WORKDIR /go/src/server
RUN go get -u github.com/golang/dep/...
COPY . .
RUN dep ensure
RUN go install

EXPOSE 8080

CMD /go/bin/server

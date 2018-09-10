# FROM golang AS build
FROM alpine

WORKDIR /root/
# RUN go get -u github.com/golang/dep/cmd/dep
ADD . .
# RUN dep ensure
# RUN CGO_ENABLED=0 GOOS=linux go install
CMD ["./chat"]

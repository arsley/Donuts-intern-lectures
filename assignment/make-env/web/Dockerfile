# FROM golang AS build

# WORKDIR /go/src/web
# RUN go get -u github.com/golang/dep/cmd/dep
# ADD . .
# RUN dep ensure
# RUN CGO_ENABLED=0 GOOS=linux go install

FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /root/
# ADD ./template ./template
ADD . .
# COPY --from=build /go/bin/web .
CMD ["./web"]

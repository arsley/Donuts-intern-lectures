# Usage

```
docker build -t reqsrv/sample .
docker run --rm -d -p 8080:8080 reqsrv/sample
# curl localhost:8080/status/500
#=> Internal server error
```

# Usage

```
docker build -t rtmp/sample .
docker run -p 80:80 -p 1935:1935 --rm -d rtmp/sample
# access http://localhost/hls/localhost.m3u8
```

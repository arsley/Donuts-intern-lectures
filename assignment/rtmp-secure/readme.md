# Usage

```
# makehash
echo $(($(date +%s) + 3600))
echo -n 'sec:live/fuga`put calculated time here`' | openssl md5 -binary | openssl base64 | tr +/ -_ | tr -d = | pbcopy
docker-compose up -d

# OBS setting
rtmp://localhost/live
fuga?t=`time`&s=`hash`
```

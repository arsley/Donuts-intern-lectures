# Usage

```
docker run --rm -v `pwd`/conf.d:/etc/nginx/conf.d -v `pwd`/html:/usr/share/nginx/html -d -p 80:80 nginx

# check
curl -H "Host: web1" http://127.0.0.1/
curl -H "Host: web2" http://127.0.0.1/
```

## Learn

`docker run` with `-v` option needs absolute path

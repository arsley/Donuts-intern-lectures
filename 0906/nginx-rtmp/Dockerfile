FROM tiangolo/nginx-rtmp

RUN mkdir -p /tmp/hls
RUN chmod 777 /tmp/hls
COPY nginx.conf /etc/nginx/nginx.conf
EXPOSE 80

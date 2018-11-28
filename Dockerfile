FROM alpine:latest

MAINTAINER "wsl <wu12490@gmail.com>"

ADD ./build/server /usr/bin/

ENV EtcdEndPoints="http://192.168.0.59:2379"
ENV AuthAccount="admin"
ENV AuthPassword="admin"

EXPOSE 8080

CMD ["server"]
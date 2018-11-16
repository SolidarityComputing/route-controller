FROM alpine:latest

MAINTAINER "wsl <wu12490@gmail.com>"

ADD ./build/server /usr/bin/

ENV EtcdEndPoints="http://sys-core_sys-etcd:2379"

EXPOSE 8080

CMD ["server"]

## run etcd

docker run -d -p 2379:2379 quay.io/coreos/etcd \
etcd \
--listen-client-urls=http://0.0.0.0:2379 \
--advertise-client-urls=http://0.0.0.0:2379 \
--data-dir=/var/lib/etcd


## build

```
go build -o server main.go

build using docker golang

docker run -it -v /Users/wsl/Projects/go/src:/go/src golang:1.10.3-alpine3.8 sh

cd src/github.com/kfcoding-container-api/ && go build -o server main.go && exit
```

## run in docker alpine:latest

```
docker run -it -v /Users/wsl/Projects/go/src/github.com/kfcoding-container-api:/home -p 8080:8080 alpine sh
export EtcdEndPoints="http://192.168.200.179:2379"
cd /home && ./server
```

## build docker image
```
docker build -t server:v1 .
docker run -it server:v1
```
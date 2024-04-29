# Simple NSQ Auth service
The simple authentication service for [NSQ](https://github.com/nsqio/nsq)

## Build

### Binary
```shell
go build -o nsq-auth ./cmd/nsq-auth/...

# or
go build -ldflags "-s -w -X 'main.Version=MY_VERSION'" -o nsq-auth ./cmd/nsq-auth/...
```

### Docker
```shell
# build local docker image
docker buildx bake image-local

# build and push multiple architecture image
APP_VERSION=0.0.0 docker buildx bake image-all --push
```

## Configure

### `nsqauth.yaml` config file
```yaml
# IP address and port for the HTTP server binding (optional, default: 0.0.0.0:4181)
bind_addr: 0.0.0.0
bind_port: 4181

# Identities list
identities:
  - identity: my-app-message-sender # identity name
    secret: twlEgK525guP7ByWiSZPMkok2OHTEJLN # identity secret key
    authorizations: # access lists
      - topic: ^test$
        channels:
          - .*
        permissions:
          - publish
          - subscribe
```

### Docker-Compose example
```yaml
version: '3'
services:
  nsqlookupd:
    image: nsqio/nsq
    command: /nsqlookupd
    ports:
      - "4160:4160"
      - "4161:4161"
  nsqd:
    image: nsqio/nsq
    command: /nsqd --lookupd-tcp-address=nsqlookupd:4160 -broadcast-address 10.0.0.2 -auth-http-address nsq-auth:4181
    depends_on:
      - nsqlookupd
    ports:
      - "4150:4150"
      - "4151:4151"
  nsqadmin:
    image: nsqio/nsq
    command: /nsqadmin --lookupd-http-address=nsqlookupd:4161
    depends_on:
      - nsqlookupd
    ports:
      - "4171:4171"

  nsq-auth:
    image: ghcr.io/yousysadmin/nsq-auth:0.0.1
    restart: always
    volumes:
      - ./nsqauth.yaml:/nsqauth.yaml
```

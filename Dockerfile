ARG TARGETOS
ARG TARGETARCH
ARG GOLANG_VERSION="1.22"
ARG APP_VERSION="unversioned"

FROM  golang:${GOLANG_VERSION} AS golang

FROM golang as build
ARG TARGETOS
ARG TARGETARCH
ARG APP_VERSION

WORKDIR /build
COPY go.mod go.sum ./
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags "-s -w -X 'main.Version=v${APP_VERSION}'" -o /nsq-auth ./cmd/nsq-auth/...

FROM scratch
COPY --from=build /nsq-auth /
ENTRYPOINT ["/nsq-auth"]

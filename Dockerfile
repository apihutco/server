FROM northes/golang:1.20-alpine-gcc AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN export GOPROXY=https://proxy.golang.com.cn,direct && go build -o ./bin/apihut .

FROM northes/alpine:gcc

WORKDIR /app

COPY --from=builder /build/bin/apihut .
COPY ./conf/config.sample.yaml ./conf/config.yaml
#COPY ./deploy/setup ./deploy/setup

ENTRYPOINT ["/app/apihut"]
CMD ["-f","./conf/config.yaml"]
EXPOSE 8282

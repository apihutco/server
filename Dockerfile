FROM golang:1.18-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories && \
    apk --no-cache add build-base

WORKDIR /build

COPY . .

RUN export GOPROXY=https://proxy.golang.com.cn,direct && go build -o ./bin/apihut .

FROM alpine

WORKDIR /app

RUN apk --no-cache add tzdata ca-certificates libc6-compat libgcc libstdc++

COPY --from=builder /build/bin/apihut .
COPY ./conf/config.sample.yaml ./conf/config.yaml

ENTRYPOINT ["/app/apihut"]
CMD ["-f","./conf/config.yaml"]
EXPOSE 8090

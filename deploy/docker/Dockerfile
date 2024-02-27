FROM golang:1.20

WORKDIR /app/go-zero-check
COPY . .

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
#    GOOS=darwin \
#    GOARCH=arm64 \
    GOPROXY="https://goproxy.cn,direct" \
    TZ=Asia/Shanghai \
    APP_ENV=docker

RUN go build -o main

EXPOSE 8080
ENTRYPOINT ["./main"]
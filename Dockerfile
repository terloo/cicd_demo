FROM golang:1.17 as builder
ENV GOPROXY=https://goproxy.cn,direct
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server

FROM alpine
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /app/server /server
ENTRYPOINT ["/server"]
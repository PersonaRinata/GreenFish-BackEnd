FROM alpine:latest

WORKDIR /app

COPY  ./user .
COPY  ./config.yaml ./server/service/aigc/config.yaml

# 暴露端口
EXPOSE 10006

CMD ["./aigc"]
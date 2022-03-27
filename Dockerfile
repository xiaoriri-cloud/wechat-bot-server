# 打包依赖阶段使用golang作为基础镜像
FROM golang:1.17-alpine as builder

# 启用go module
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .

# 指定OS等，并go build
RUN  go build .


# 运行阶段指定scratch作为基础镜像
FROM alpine

WORKDIR /app

# 将上一个阶段publish文件夹下的所有文件复制进来
COPY --from=builder /app .

# 指定运行时环境变量
#ENV GIN_MODE=release \
#    PORT=80

EXPOSE 8080

ENTRYPOINT ["./wechat-bot-server"]
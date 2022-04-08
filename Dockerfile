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
FROM alpine:3.14

RUN echo -e  "http://mirrors.aliyun.com/alpine/v3.14/main\nhttp://mirrors.aliyun.com/alpine/v3.14/community" >  /etc/apk/repositories \
&& apk update && apk add tzdata \
&& cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
&& echo "Shanghai/Asia" > /etc/timezone \
&& apk del tzdata

WORKDIR /app

# 将上一个阶段publish文件夹下的所有文件复制进来
COPY --from=builder /app/wechat-bot-server .

# 指定运行时环境变量
#ENV GIN_MODE=release \
#    PORT=80

EXPOSE 8080

ENTRYPOINT ["./wechat-bot-server"]
# 使用官方 Go 镜像作为基础镜像
FROM golang:1.22.5-alpine

# 设置工作目录
WORKDIR /app

# 将当前目录下的所有文件复制到容器的 /app 目录
COPY . .

# 下载并安装依赖项
RUN go mod tidy

# 编译 Go 应用
RUN go build -o main .

# 公开容器的端口（假设你的服务在 8080 端口运行）
EXPOSE 8080

# 启动应用
CMD ["./main"]

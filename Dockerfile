# 使用 Golang 作为构建镜像
FROM golang:1.23 AS builder

# 设置工作目录
WORKDIR /usr/src/gochat

# 设置 Go 模块代理
ENV GOPROXY=https://goproxy.io,direct

# 复制 go.mod 和 go.sum 文件，并下载依赖
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# 复制源码到构建镜像
COPY . .

# 编译项目
RUN go build -v -o /opt/gochat/server

# 使用 Busybox 作为运行时镜像
FROM busybox

# 复制构建好的可执行文件到最终镜像
COPY --from=builder /opt/gochat/server /opt/gochat/server

# 复制配置文件
COPY ../../config.toml /opt/gochat/config.toml

# 设置工作目录
WORKDIR /opt/gochat

# 设置容器启动时执行的命令
CMD ["./server"]

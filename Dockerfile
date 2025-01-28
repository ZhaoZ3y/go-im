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

# 编译项目（指定 main.go 路径）
RUN go build -v -o /opt/gochat/server ./cmd

FROM busybox

COPY --from=builder /opt/gochat/server /opt/gochat/server

COPY ./config.yaml /opt/gochat/config.yaml

WORKDIR /opt/gochat

CMD ["./server"]
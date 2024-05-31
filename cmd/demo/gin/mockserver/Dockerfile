# 打包
FROM golang:latest AS builder

# 启用go module
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY *.go ./

# Build
RUN go mod init mockserver
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/mockserver

# 运行阶段指定scratch作为基础镜像
FROM scratch

WORKDIR /app

# 拷贝程序
COPY --from=builder /app/mockserver .

EXPOSE 8080

# Run
CMD ["./mockserver"]
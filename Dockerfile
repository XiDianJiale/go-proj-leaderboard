# -----------------------------
# Stage 1: Build (Go builder)
# -----------------------------
FROM golang:1.25-alpine AS builder

WORKDIR /app

# 让 Go module 更快
ENV GOPROXY=https://goproxy.cn,direct

# 拷贝 go.mod / go.sum
COPY go.mod go.sum ./
RUN go mod download

# 拷贝全部代码
COPY . .

# 构建二进制
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/main.go


# -----------------------------
# Stage 2: Run (small runtime)
# -----------------------------
FROM alpine:3.20

WORKDIR /app

# 拷贝可执行文件
COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]

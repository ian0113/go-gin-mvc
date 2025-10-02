# 使用官方 Golang 映像作為建置階段
FROM golang:1.21-alpine AS builder

# 設定工作目錄
WORKDIR /app

# 拷貝 go.mod 和 go.sum 並先行下載依賴
COPY go.mod go.sum ./
RUN go mod download

# 拷貝其他程式碼
COPY . .

# 編譯應用程式（可調整檔名）
RUN go build -o go-gin-mvc .

# 使用更輕量的映像作為最終執行環境
FROM alpine:latest

# 安裝 ca-certificates 以支援 HTTPS 請求
RUN apk --no-cache add ca-certificates

# 設定工作目錄
WORKDIR /root/

# 從建置階段拷貝編譯好的應用程式
COPY --from=builder /app/go-gin-mvc .

# 曝露應用程式使用的 Port（可根據你專案實際設定）
EXPOSE 8080

# 設定啟動指令
CMD ["./go-gin-mvc"]

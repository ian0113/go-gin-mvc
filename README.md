# go-gin-mvc

🚀 A clean and modular MVC-style project using [Gin](https://github.com/gin-gonic/gin) framework in Go.

本專案採用分層架構設計（MVC + Service/Repository Layer），讓專案結構更清晰、易於維護與擴充。

---

## 📁 專案資料夾結構

```
.
├── config/          設定檔
├── controllers/     控制器層，處理 HTTP 請求與回應
├── infra/           基礎設施（Config,DB,Redis,Logger,Queue...等等）
├── middlewares/     中介層功能（驗證、日誌 等）
├── models/          資料模型定義（struct 與 DB 映射）
├── repositories/    資料操作邏輯（資料庫 CRUD）
├── routes/          路由註冊與綁定控制器
├── services/        業務邏輯處理
├── utils/           工具函式（如加密、轉換等）
├── main.go          專案啟動點
├── go.mod           Go 模組定義
└── go.sum           相依套件資訊
```

---

## 🧠 專案執行順序

```
main.go
├── 載入設定（config）
│   └── 讀取環境變數或初始化參數（如 App,DB,Redis 設定）
│
├── 初始化基礎設施（infra）
│   └── 建立DB/Redis/Logger等資源
│
├── 載入中間件（middlewares）
│   └── 如 CORS、Logger、驗證、錯誤處理等
│
├── 註冊路由（routes）
│   └── 綁定 URL 路徑與對應的 controller 函式
│
├── 啟動 Gin 引擎
    └── 開始接收 HTTP 請求
```

---

## 🔁 邏輯流程

```
     ┌──────────────┐                       ┌──────────────┐
     │ HTTP request │                       │ HTTP response│
     └─────┬────────┘                       └─────┬────────┘
           ↓         請求進入系統                  ↑         回傳給 client
     ┌─────┴────────┐                             │
     │   routes     │                             │
     └─────┬────────┘                             │
           ↓         路由匹配與請求分派             ↑        ??? 可能可以驗證輸出結果 ???
     ┌ - - ┴ - - -  ┐                       ┌ - - ┴ - - -  ┐
     │ middlewares  │                       │ middlewares  │
     └ - - ┬ - - -  ┘                       └ - - ┬ - - -  ┘
           ↓         中介層處理(認證、日誌等)       ↑        封裝結果
     ┌─────┴────────┐                       ┌─────┴────────┐
     │ controllers  │                       │ controllers  │
     └─────┬────────┘                       └─────┬────────┘
           ↓         處理請求、呼叫服務層           ↑        資料回傳
     ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
     ··· 處理業務邏輯與資料存取 ··· 處理業務邏輯與資料存取  ···
     ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
           ↓         呼叫服務層
     ┌─────┴────────┐
     │  services    │
     └─────┬────────┘
           ↓         業務邏輯處理
     ┌─────┴────────┐
     │ repositories │
     └─────┬────────┘
           ↓         存取資料來源
      ┌────┴─────────┐
      ↓              ↓
┌─────┴────────┐    非關聯式資料庫 (Redis)
│   Models     │
└─────┬────────┘
      ↓
     關聯式資料庫 (Postgres/MySQL)
```

---

## 🐳 Docker 啟動方式

專案已整合 Dockerfile 與 docker-compose.yml，可一鍵啟動 Go Gin API 與相關服務（如 Postgres、Redis）。

### 使用 Docker Compose

執行

```
docker-compose up --build
```

### 手動執行（不用 compose, 但你得自己建 DB/Redis, 並修改 config）

建立映像

```
docker build -t go-gin-mvc .
```

執行 container

```
docker run -p 8080:8080 go-gin-mvc
```

啟動完成後，你可以在瀏覽器或 Postman 訪問：

```
http://localhost:8080
```


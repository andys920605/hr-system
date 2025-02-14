# HR System API Documentation

## 簡介

本專案為 **HR 系統**，基於 **Golang (Gin)** 開發，提供 RESTful API 來管理員工資訊。題目需求包括下列：

- **Gin**：輕量級、高效能的 Web 框架
- **MySQL**：作為資料庫，並使用 **GORM** 進行 ORM 操作
- **Redis**：作為快取層，提高查詢效能
- **GORM Migration & SEED Data**：初始化資料庫
- **Unit Test**：確保核心邏輯的正確性
- **Makefile**：統一管理建置與部署流程
- **Docker Compose**：一鍵啟動整個服務

---

## 架構介紹
[菱形對稱架構](https://segmentfault.com/a/1190000040533813/ "游標顯示")
菱形對稱架構介紹
菱形對稱架構是一種結合 六邊形架構 (Hexagonal Architecture)、Clean Architecture 與 領域驅動設計 (DDD) 的架構模式，旨在讓系統的內部結構保持清晰，並提高可維護性與可擴展性。

架構概念
這個架構的核心思想是：

1. 將業務邏輯獨立於基礎設施與框架，確保核心業務邏輯不被技術細節耦合。
2. 輸入與輸出對稱 (Input & Output Symmetry)，也就是說，無論是外部 API、資料庫、訊息隊列或是 UI 介面，它們與業務邏輯的交互方式都是一致的，保持架構的統一性與一致性。
3. 以 DDD 的方式組織代碼，使用 聚合 (Aggregate)、領域服務 (Domain Service)、應用層 (Application Layer) 來確保業務邏輯的可讀性與擴展性。

適用場景
高複雜度的業務邏輯，如金融、電商、廣告系統
需要長期維護的企業級應用
需要高度解耦與可擴展的架構設計
總結
菱形對稱架構的目標是 透過清晰的分層來降低耦合，讓系統更容易維護與擴展。
它保留了 六邊形架構的開放性、Clean Architecture 的清晰分層，並結合 DDD 的業務驅動，讓整體架構更符合現代後端開發的需求。


## 快速上手

### 使用 Docker Compose 啟動服務

```bash
docker-compose up -d
```

## API 介面說明
### 創建員工
- Endpoint：http://127.0.0.1:8080/srv/employees
- HTTP Method：POST
- Request Body 範例：
```bash
{
  "name": "John Doe",
  "email": "johndoe3@example.com",
  "phone": "+886912345678",
  "address": "123 Main St, Taipei, Taiwan",
  "level": "L2",
  "position": "engineer"
}
```
Response：
狀態碼：200

### 查詢在職員工
- Endpoint：http://127.0.0.1:8080/srv/employees/{id}
- HTTP Method：GET
- Response Body 範例：
```bash
{
  "code": 0,
  "msg": "ok",
  "data": {
    "id": "1123669401745228800",
    "name": "Alice",
    "email": "alice@example.com",
    "phone": "123456789",
    "address": "Taipei",
    "level": "L3",
    "position": "engineer"
  }
}
```
# Go Backend Fabric

> Scalable backend building blocks in Go — woven into a powerful infrastructure fabric.

A modular and scalable backend infrastructure toolkit built in Go, providing reusable components like rate limiting, caching, and load balancing for high-performance systems.

---

## 🚀 Features

- ⚡ Rate Limiter (middleware-based)
- 🧠 In-memory Cache (extensible interface)
- 🏗 Clean Architecture (domain, infrastructure, delivery)
- 🔌 Pluggable components (easy to extend with Redis, etc.)
- 🧵 Concurrency-safe implementations using Go routines & sync primitives

---

## 📂 Project Structure

```
go-backend-fabric/
├── cmd/server/                # Application entry point
├── internal/
│   ├── domain/               # Core interfaces & business logic
│   ├── infrastructure/       # Implementations (cache, rate limiter)
│   ├── delivery/             # HTTP layer (middlewares, handlers)
│
├── go.mod
└── README.md
```

---

## 🧩 Components

### Rate Limiter
- Middleware-based request limiting
- Token bucket / custom logic (extendable)
- Designed to integrate easily with HTTP servers

### Cache
- Generic cache interface
- In-memory implementation (`memorystore`)
- Easily extendable to Redis or distributed caches

---

## ⚙️ Getting Started

### 1. Clone the repo
```bash
git clone https://github.com/ShubhamKharde45/go-backend-fabric.git
cd go-backend-fabric
```

### 2. Run the server
```bash
go run cmd/server/main.go
```

---

## 🛠️ Usage Example

```go
// Example: Using rate limiter middleware
router.Use(RateLimiterMiddleware)

// Example: Using cache
cache.Set("key", "value")
val := cache.Get("key")
```

---

## 🧠 Design Philosophy

This project is built with the idea of a **“backend fabric”** —  
a set of core infrastructure components woven together to form a scalable system foundation.

- Focus on **modularity**
- Follow **clean architecture principles**
- Keep components **independent and reusable**

---

## 📈 Future Improvements

- Redis-based distributed cache
- Distributed rate limiter
- Load balancer implementation
- Metrics & monitoring
- Config-driven architecture

---

## 🤝 Contributing

Contributions are welcome! Feel free to open issues or submit PRs.

---

## 📜 License

This project is licensed under the MIT License.

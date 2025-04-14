# 🎭 Backdrop

**Backdrop** is a concurrent, asynchronous task orchestration service built with Go. It allows users to upload files, trigger background processing, track progress, and optionally cancel jobs in-flight — all with real-time feedback and robust status tracking.

---

## 📌 Features

- Asynchronous file processing (non-blocking)
- Unique upload URLs per request
- Task cancellation support (server-side invalidation)
- Polling endpoint for task status
- Clean context-aware Goroutine management
- JWT-based user authentication
- Task persistence via PostgreSQL

---

## 🧰 Tech Stack

| Component | Purpose |
|----------|---------|
| **Go** | Core backend logic (goroutines, context, channel-based workers) |
| **PostgreSQL** | Task metadata, user data, and login tokens |
| **Redis** | Fast pub/sub and task state caching |
| **Polling** | Task status updates to frontend |
| **Docker** | Containerization for local or cloud deployment |

---

## 📁 Project Structure

```
backdrop/
├── api
│   └── user
│   └── task
├── cmd
│   └── server
├── go.mod
├── go.sum
├── internal
│   ├── middleware
│   ├── router
│   └── util
├── migration.sql
├── pkg
│   ├── constants
│   ├── crypto
│   ├── database
│   ├── formatter
│   ├── text
│   └── validator
├── README.md
└── TODO.md
```

---

## 🚀 Getting Started

### Prerequisites
- Go 1.20+
- PostgreSQL running locally

### Clone the project
```bash
git clone https://github.com/rranand/backdrop.git
cd backdrop
```

### Initialize Go module
```bash
go mod tidy
```

### Setup `.env` file
```env
PORT=8080
DB_URL=postgres://user:pass@localhost:5432/backdrop?sslmode=disable
JWT_SECRET=your-secret-key
```

---

## 📬 API Overview

| Method | Endpoint            | Description |
|--------|---------------------|-------------|
| POST   | `/auth/login`       | Login and receive token |
| POST   | `/task/request`     | Request a file upload session |
| POST   | `/task/upload/:id`  | Upload file to task |
| POST   | `/task/cancel/:id`  | Cancel ongoing task |
| GET    | `/task/status/:id`  | Poll task status |
| GET    | `/task/result/:id`  | Fetch result (if completed) |

---

## 📈 Future Enhancements

- [ ] WebSocket-based live task updates
- [ ] Retry policy for failed tasks
- [ ] Admin dashboard (task queues, worker health)
- [ ] Multi-user role support
- [ ] Metrics + alerting (Prometheus/Grafana)

---

## 🧠 Learnings and Concepts Covered

- Goroutines, context cancelation, worker pools
- Safe concurrent access to shared resources
- Token-based authentication
- Structuring scalable Go services
- Database schema design for async systems

---

## 📄 License

MIT © 2025 Rohit Anand
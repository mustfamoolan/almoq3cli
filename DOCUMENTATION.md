# 🛡️ almoq3 Framework - Documentation
**Version:** `1.0.3` | **Language:** `Go (Golang)` | **Ecosystem:** `NPM / GitHub`

## 🌟 Introduction
**almoq3** is a high-performance, enterprise-grade Go framework designed for developers who love the elegant Developer Experience (DX) of Laravel but demand the raw performance and concurrency of Go. It provides a complete suite of tools to scaffold, develop, and deploy modern web applications.

---

## 🚀 Key Features

### 1. Robust Scaffolding (CLI Engine)
The `almoq3` CLI is the heart of the framework, allowing you to generate components in seconds:
- **Controllers**: `make:controller [Name]`
- **Services (Business Logic)**: `make:service [Name]`
- **Form Requests (Validation)**: `make:request [Name]`
- **Middleware**: `make:middleware [Name]`
- **Database Migrations**: `make:migration [Name]`
- **Seeders**: `make:seeder [Name]`
- **Background Jobs**: `make:job [Name]`
- **Automated Tests**: `make:test [Name]`

### 2. Enterprise Security & Auth
- **JWT Authentication**: Built-in stateless auth system using `golang-jwt/v5`.
- **Bcrypt Hashing**: Secure password management out-of-the-box.
- **Auth Scaffolding**: `make:auth` generates Login, Register, Middleware, and User models instantly.

### 3. Database & Persistence
- **ORM**: Powered by **GORM** for seamless database interactions.
- **Engines**: Supports MySQL, PostgreSQL, and SQLite.
- **Seeding Engine**: Laravel-like seeding system with `db:seed`.

### 4. High-Performance Infrastructure
- **Web Engine**: Powered by **Fiber (v2)**, the fastest Go web framework.
- **Caching**: Native **Redis** integration for high-speed caching.
- **Logging**: Advanced logging using **Uber Zap** with automatic **File Rotation** (Lumberjack).

### 5. DevOps & Deployment
- **Docker Ready**: Every project includes a multi-stage `Dockerfile` and `docker-compose.yml`.
- **NPM Distribution**: Installable globally via `npm install -g almoq3-cli`.
- **Self-Update**: Built-in version checker and self-updating mechanism via GitHub.

---

## 📂 Project Architecture
```text
├── app/
│   ├── controllers/    # Request handlers
│   ├── models/         # Database entities
│   ├── services/       # Business logic layer
│   ├── middleware/     # HTTP filters
│   ├── requests/       # Validation logic
│   └── jobs/           # Background tasks
├── bootstrap/          # Framework core initializers (DB, Cache, Logger)
├── config/             # Application configuration
├── database/
│   ├── migrations/     # Schema versioning
│   └── seeders/        # Dummy data generators
├── resources/
│   └── views/          # HTML templates (Welcome page)
├── routes/             # API and Web route definitions
├── storage/            # Logs and uploads
└── main.go             # Entry point
```

---

## 🛠️ Developer Commands
- `almoq3 new [Project]`: Create a fresh masterpiece.
- `almoq3 run`: Start the development server (Cross-platform).
- `almoq3 key:generate`: Set the `APP_KEY` for security.
- `almoq3 migrate`: Sync your database schema.
- `almoq3 self-update`: Stay up-to-date with the latest framework features.

---

## 🎨 Design Philosophy
almoq3 follows a **Clean Architecture** approach while maintaining a **Minimalist Aesthetic**. The premium welcome page and the intuitive CLI interactions are designed to inspire developers from the first command.

---

## 📜 License
Crafted with ❤️ in **Iraq**. Released under the **MIT License**.

---
**Official Website:** [https://almoq3.com](https://almoq3.com)  
**Documentation:** [https://documentation.almoq3.com](https://documentation.almoq3.com)

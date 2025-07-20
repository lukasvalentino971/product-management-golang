# 🔒 JWT Auth Product Management API

A high-performance RESTful API built with Golang, Gin, and GORM, featuring robust JWT authentication and comprehensive product management capabilities.

![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?logo=go)
![Gin](https://img.shields.io/badge/Gin-1.9+-000000?logo=go)
![JWT](https://img.shields.io/badge/JWT-Auth-000000?logo=json-web-tokens)

## 🚀 Features

### 🔐 Authentication
- User registration with email/password validation
- Secure login with JWT token generation
- Password hashing using bcrypt
- Role-based access control (ready for extension)

### 🛡️ Security
- JWT authentication middleware
- Rate limiting on sensitive endpoints (login protection)
- Secure password storage
- Configurable token expiration

### 📦 Product Management
- Full CRUD operations for products
- Paginated product listings
- Search and filter capabilities
- Data validation and sanitization

### ⚙️ Infrastructure
- Clean architecture with separation of concerns
- Database auto-migrations
- Environment configuration
- Modular and testable design

## 📦 Tech Stack

| Component       | Technology |
|-----------------|------------|
| Framework       | Gin        |
| ORM             | GORM       |
| Authentication  | JWT        |
| Database        | MySQL |
| Rate Limiting   | ulule/limiter |
| Testing         | Testify    |

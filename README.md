# 🛒 ecom — REST API for E-commerce in Go

`ecom` is a lightweight RESTful API written in Go that provides backend functionality for a simple e-commerce platform. It supports user registration, authentication using JWT, product listing, and cart checkout.

---

## 🚀 Features

- 🔐 User Registration & Login (JWT-based)
- 📦 Product browsing
- 🛒 Cart checkout & order creation
- ✅ Middleware for protected routes
- 🧰 Environment-based configuration
- 🗃️ Database migrations (raw SQL)

---

## 🧑‍💻 Tech Stack

- **Language**: Go (Golang)
- **Routing**: [gorilla/mux](https://github.com/gorilla/mux)
- **Database**: MySQL
- **Auth**: JWT (`golang-jwt/jwt/v5`)
- **HTTP Server**: net/http
- **Migrations**: raw `.sql` files (with custom runner)

---

## 📁 Project Structure

```text
ecom/
├── bin/                         # Output binaries
├── cmd/
│   ├── api/                     # Server entry point (api.go)
│   └── migrate/migrations/     # SQL migration files
├── config/
│   └── env.go                  # Loads environment variables
├── db/
│   └── db.go                   # MySQL connection setup
├── service/
│   ├── auth/                   # JWT & middleware
│   ├── cart/                   # Checkout logic
│   ├── product/                # Product logic
│   └── user/                   # Registration/login
├── types/
│   └── types.go                # Structs & interfaces
├── utils/
│   └── utils.go                # JSON handling, token extraction
├── .env                        # Environment configuration
├── Makefile                    # Dev commands (run, migrate)
└── go.mod                      # Module file

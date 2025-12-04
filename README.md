# Go Fiber REST API - Todo App

🛠️ Tech Stack
Language: Go (Golang)
Web Framework: Fiber v2 (Express-inspired, zero memory allocation)
ORM: GORM
Database: SQLite (Embedded)
Validation: Go Playground Validator


## 🚀 Key Features

* **Layered Architecture:** Clear separation of `DAL` (Data Access Layer), `Services`, and `Controllers`.
* **RESTful Endpoints:** Full CRUD operations for managing Todo items.
* **Database ORM:** Uses **GORM** for database interactions with **SQLite** (easily swappable for PostgreSQL/MySQL).
* **Input Validation:** Robust request validation using `go-playground/validator`.
* **Modern Go Patterns:** Utilizes structs, pointers, and interfaces effectively.

## 📂 Project Structure

This project follows a modular structure to maintain scalability:

```go-rest
├── dal/          # Data Access Layer (Database queries & GORM operations)
├── database/     # Database connection and configuration
├── services/     # Business Logic Layer (Handles core functionality)
├── types/        # Data Transfer Objects (DTOs) and Request Models
├── main.go       # Application entry point and route definitions
└── go.mod        # Dependency management 

 ## ⚡ Getting Started

Prerequisites
Go 1.19 or higher installed.

##  Installation
Clone the repository:
Install dependencies:go mod tidy
Run the application: go run main.go

🔮 Future Improvements
To further enhance this project, the following features are planned:

[ ] Migration to PostgreSQL for production environments.

[ ] Implementation of Unit Tests using Go's testing package.

[ ] Dockerization (Dockerfile & Docker Compose).

[ ] Swagger/OpenAPI documentation.
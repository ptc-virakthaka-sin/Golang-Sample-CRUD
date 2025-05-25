# Golang CRUD Service

This is a sample CRUD service built with Golang, demonstrating a layered architecture that separates concerns into handler, service, repository, and model layers. The project uses **TiDB** as the relational database and **Redis** for caching and other in-memory operations.

## 🚀 Project Purpose

Build a sample CRUD service API with Golang to showcase clean code structure and common best practices in web API development.

## 🛠️ Tech Stack

- **[Golang](https://golang.org)** — Core programming language
- **[Fiber](https://gofiber.io)** — Fast HTTP router
- **[GORM](https://gorm.io)** — ORM for database operations
- **[TiDB](https://www.pingcap.com)** — Distributed SQL database
- **[Redis](https://redis.io)** — In-memory data store
- **[Validator](https://github.com/go-playground/validator)** — Input validation
- **[Copier](https://github.com/jinzhu/copier)** — Object copying for DTOs

## 📁 Project Structure

```
├── cmd/
│   └── app/
│       └── main.go                # Application entry point
├── pkg/                           # Shared package configure
│   └── bcrypt/
|       └── bcrypt.go              # Bcrypt hashing                 
│   └── redis/                     
|       └── redis.go               # Redis configuration
│   └── db/                        
|       └── db.go                  # DB configuration
├── api/
│   └── student/
│       ├── route.go               # Student-related endpoints
│       └── handler.go             # Logic for handling student API requests
├── internal/
│   ├── service/                   # Business logic, called from handler
│   ├── repository/                # Data access layer (calls GORM & queries DB)
│   ├── model/                     # Database models/entities
│   └── dto/                       # DTOs for transforming request/response data
```

## 📌 Key Concepts

- **Layered Architecture**: Divides the project into clear layers for maintainability and scalability.
- **DTO Pattern**: Uses `copier` to map between internal models and response/request structures.
- **Validation**: Ensures request payloads are validated with `go-playground/validator`.
- **ORM**: Leverages GORM to interact with TiDB in a concise and type-safe way.

## 🧪 Running the Project

1. **Configure Environment**
   Create a `.env` file. find in a `.env.example` file for environment:
   - JWT Secret key
   - TiDB connection
   - Redis connection

2. **Run the App**
   ```bash
   go run cmd/app/main.go
   ```

## 🔧 Future Enhancements

- Add unit tests and integration tests
- Implement authentication and authorization
- Dockerize the application
- Add Swagger/OpenAPI documentation

## 📄 License

This project is open-source and available under the [MIT License](LICENSE).

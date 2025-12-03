# Go CRUD API (Backend)

A clean and professional implementation of a CRUD API using **Go**, **Gin**, and **GORM**. This project demonstrates how to build a scalable REST API with proper project structure, validation, database integration, and error handling.

---

## 🚀 Features

* Full **CRUD (Create, Read, Update, Delete)** operations
* Clean project layout using **repository, service, handler** layers
* Database integration with **GORM**
* JSON-based request/response formatting
* Environment-based configuration loading (`.env`)

---

## 📂 Project Structure

```
.
├── cmd/server/
│       └── main.go
├── internal/
|   ├── config/
│   ├── handler/
│   ├── service/
│   ├── repository/
│   ├── model/
|   ├── router/
│   ├── domain/
|   └── dto/
├── go.mod
└── go.sum
```

**Explanation:**

* `cmd/` — Application entry point
* `internal/` — Core modules your application uses internally
* `config/` — Load config from env file locally
* `handler/` — HTTP handlers (Gin)
* `service/` — Business logic
* `repository/` — Database operations using GORM
* `model/` — GORM models
* `domain/` — Domain objects for business logic
* `dto/` — Data Transfer Object for request / response

---

## 🛠️ Installation

### 1. Clone the repository

```sh
git clone https://github.com/yourname/go-crud.git
cd go-crud
```

### 2. Install dependencies

```sh
go mod tidy
```

### 3. Configure `.env`

Create a `.env` file:

```ini
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=yourpassword
DB_NAME=cruddb
JWT_SECRET=your_secret_here
```

---

## 🗄️ Database Setup

Make sure Docker is setup or PostgreSQL is running.

Create the database:

```bash
docker-compose up -d
```

GORM will auto-migrate the required tables.

---

## ▶️ Run the Application

```sh
go run ./cmd/server
```

Server will start at:

```
http://localhost:8080
```

---

## 📌 API Endpoints

### 👉 Create User

**POST** `/users`

```json
{
  "name": "Alex",
  "email": "alex@gmail.com",
  "password": "123456"
}
```

### 👉 Get All Users

**GET** `/users`

### 👉 Get User by ID

**GET** `/users/:id`

### 👉 Update User

**PUT** `/users/:id`

```json
{
  "name": "Updated Name",
  "email": "new@gmail.com"
}
```

### 👉 Delete User

**DELETE** `/users/:id`


---

## 🔐 Authentication ()

A version of this project includes JWT authentication using:

```go
jwt.MapClaims{"user_id": id, "user_email": email "exp": expirationTime}
```

### 👉 login

**POST** `/login`

You can get the JWT token when you login.

```json
{
  "email": "example@gmail.com",
  "password": "example"
}
```

### 👉 signup

**POST** `/signup`

Before login via above api, you have to sign up with below API.

```json
{
  "email": "example@gmail.com",
  "password": "example"
}
```

---

## 🧱 Technologies Used

* Go (Golang)
* Gin framework
* GORM ORM
* PostgreSQL
* GoDotEnv

---

## 🤝 Contributing

Pull requests are welcome! For major changes, open an issue first to discuss what you’d like to change.

---

## 📜 License

This project is licensed under the MIT License.

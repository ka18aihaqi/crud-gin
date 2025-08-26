# 👥 CRUD Gin - User Management

A simple **CRUD (Create, Read, Update, Delete)** REST API built with [Gin]((https://gin-gonic.com/)) for managing users.  
The project demonstrates the basics of backend development with Golang and Gin by implementing a `users` table.

---

## 🚀 Features
- ✨ Create a new user
- 📋 Get all users
- 🔍 Get a user by ID
- 📝 Update user data
- ❌ Delete a user
- 🔐 User model includes:
  - `email`
  - `username`
  - `name`
  - `password`

---

## 🛠️ Tech Stack
- [Go]((https://go.dev/))
- [Gin](https://gin-gonic.com/)
- [Gorm](https://gorm.io/) *(ORM for database)*
- MySQL

---

## ⚙️ Installation & Usage
1. Clone this repository
   ```bash
   git clone https://github.com/ka18aihaqi/crud-express.js.git
   cd crud-gin
2. Install dependencies
   ```bash
   go mod tidy
4. Run the server
   ```bash
   go run main.go
6. The server will start at:
   ```bash
   http://localhost:8080

---

## 📌 API Endpoints

| Method | Endpoint         | Description       |
|--------|------------------|-------------------|
| GET    | `/api/users`     | Get all users     |
| GET    | `/api/users/:id` | Get user by ID    |
| POST   | `/api/users`     | Create a new user |
| PUT    | `/api/users/:id` | Update user by ID |
| DELETE | `/api/users/:id` | Delete user by ID |

---

## 📂 User Schema

Example of a user object:
```json
{
  "id": 1,
  "email": "john.doe@example.com",
  "username": "johndoe",
  "name": "John Doe",
  "password": "hashedpassword123"
}
```

🔥 Example Request (POST)
```json
curl -X POST http://localhost:3000/api/users \
   -H "Content-Type: application/json" \
   -d '{
     "email": "jane.doe@example.com",
     "username": "janedoe",
     "name": "Jane Doe",
     "password": "mypassword123"
   }'
```

## 🎯 Purpose
This project was created as a learning exercise to practice building a REST API with Gin, focusing on user management with basic CRUD operations.

## 📄 License
This project is not licensed under anyone.
Feel free to use, modify, and distribute it as needed.

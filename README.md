# Rekeningku API

Rekeningku API is a Go (Golang) based backend service for managing bank account data with features such as authentication with JWT, profile management and account management.

## Main Features

- User registration and login.
- Account management (CRUD).
- JWT token authentication.
- Pagination for account data retrieval.

## Technologies Used

- **Golang** - Main programming language.
- **Gin Framework** - HTTP framework for API development.
- **GORM** - ORM for database interaction.
- **JWT** - For token-based authentication.
- **MySQL** - Database used.

## Installation

### Prerequisites

Make sure you have installed:

- [Go](https://go.dev/) (minimum version 1.19).
- [MySQL](https://www.mysql.com/).
- [Postman](https://www.postman.com/) to test the API (optional).

### Installation Steps

1. Clone this repository:

```bash
git clone https://github.com/username/rekeningku-api.git
cd rekeningku-api
```

2. Create a file `/internal/configs/mysql.go` in the project root by replacing your dsn:

```env
dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
```

3. Install dependencies:

```bash
go mod tidy
```

4. Run the server:

```bash
go run main.go
```

The server will run on `http://localhost:9001`.

## API Testing

Use the [Postman API Collection](https://documenter.getpostman.com/view/7362955/2sAYBbf9eZ) to test the API endpoints. Click the link to see the documentation and import the collection into your Postman.

### Main Endpoints

#### Authentication
- `POST /api/v1/auth/register` - Register a user.
- `POST /api/v1/auth/login` - Login a user.

#### Users
- `GET /api/v1/users` - Get a user profile.
- `PATCH /api/v1/users` - Update a user profile.

#### Accounts
- `GET /api/v1/accounts` - Get a list of accounts with pagination.
- `GET /api/v1/accounts/:accountId` - Get account details.
- `POST /api/v1/accounts` - Create a new account.
- `PATCH /api/v1/accounts/:accountId` - Update account data.
- `DELETE /api/v1/accounts/:accountId` - Delete an account.

## Collaboration

If you want to collaborate feel free to contact me!

## License

This project is licensed under the [MIT License](LICENSE).

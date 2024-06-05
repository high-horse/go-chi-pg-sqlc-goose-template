# Go-Chi-PG-SQLC-Goose Template

This is a template project for building a RESTful API server in Go using the following technologies:
- [Chi](https://github.com/go-chi/chi): Lightweight and fast HTTP router for Go.
- [pg](https://github.com/go-pg/pg): PostgreSQL client and ORM for Go.
- [sqlc](https://github.com/kyleconroy/sqlc): Generate type-safe Go from SQL.
- [goose](https://github.com/pressly/goose): Database migrations for Go.

## Getting Started

### Prerequisites

- [Go](https://golang.org/) installed on your machine.
- PostgreSQL installed and running on your local machine.

### Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/high-horse/go-chi-pg-sqlc-goose-template.git
   ```
  
2. Navigate to the project directory:
   ```bash
   cd go-chi-pg-sqlc-goose-template
   ```
   
3. Install dependencies:
   ```bash
   go mod tidy
   ```
   
4. Configure the PostgreSQL connection details in config/config.go.
   ```bash
   goose -dir migrations postgres "YOUR_DATABASE_URL" up
   ```
   
5. Run database migrations:
   ```bash
   goose -dir migrations postgres "YOUR_DATABASE_URL" up
   ```
   
6. Build and run the server:
   ```bash
   go run cmd/server/main.go
   ```
   
### Usage

The server should now be running locally on http://localhost:8080. You can use tools like Postman or curl to interact with the API endpoints.

### Project Structure
go-chi-pg-sqlc-goose-template/
├── cmd/
│ └── main.go # Entrypoint for the server
├── models/
│ └── shared.go # Configuration settings
├── sql/
│ ├── database			# SQLC Query output methods 
│ ├── queries			# SQLC queries files
│ └── schema			# SQL schema files
├── sqlc.yaml 			# SQLC configuration file
├── helper/
│ └──passwords.go		# Password encryption and decryption 
├── handler/ 			# HTTP request handlers
│ └── util.go
├── middleware/ 		# HTTP middleware
│ └── logger.go 		# Request logger middleware
├── go.mod
├── go.sum
└── README.md

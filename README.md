# golang-bootstrap-project

### Technologies and frameworks used in this template

- Gin for routing
- pgx for database driver (postgresql)
- go-migrate for handling migration
- zap for logging 

### How to configure application

First you need to create a .env file at the root path (similar to Node.js), then you need to insert the following variables.

```
DATABASE_URL=postgresql://postgres:docker@localhost:5432/app_db?sslmode=disable  
ENV=development
HTTP_PORT=:3000
```

### Useful commands

#### Application related commands

- go run main.go server

#### Migration related commads

Creating a migration file (should follow the go-migrate guidelines)
- go run main.go migration create initial migration -m /Users/appname/Projects/Golang/golang-bootstrap-project/db/migrations/

Running up migration 
- go run main.go migration down -m /Users/appname/Projects/Golang/golang-bootstrap-project/db/migrations/
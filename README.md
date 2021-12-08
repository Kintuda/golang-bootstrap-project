# golang-bootstrap-project

### TODO

- Create unit test and the required scripts to run
- Add oauth2 module
- Add service and repository layers
- Add dockerfile and docker-compose files
- Add pipeline for build and lint

### Technologies and frameworks used in this template

- Gin for routing
- pgx for database driver (postgresql)
- go-migrate for handling migrations
- zap for logging 
- go-env for configuring environment variables
- cobra for cli

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

#### Migration related commands

Script already generate a new revision number 

How to create a migration file (should follow the go-migrate guidelines for migration name)
- go run main.go migration create initial migration -m /Users/appname/Projects/Golang/golang-bootstrap-project/db/migrations/

Running migration up 
- go run main.go migration down -m /Users/appname/Projects/Golang/golang-bootstrap-project/db/migrations/
# Test tast for NIX course

## Build and Run (Locally)
### Prerequisites
- go 1.13

Create .env file in root directory and add following values:
```
DB_USER=root
DB_PASSWORD=1212
DB_HOST=localhost
DB_PORT=3306
DB_NAME=parser_db

```

Use `migrate -path migrations -database "mysql://root:1212@tcp(localhost:5436)/restapi_dev" up` to up migration
Use `go run cmd/app/main.go` to run app
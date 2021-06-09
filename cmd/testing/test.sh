#!/bin/sh
go run ./cmd/migrations/migrate.go down
go run ./cmd/migrations/migrate.go up
go test ./internal/handlers/ktp_test.go
go run ./cmd/migrations/migrate.go down

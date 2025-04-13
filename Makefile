new_migration:
	migrate create -ext sql -dir ./pkg/database/migrations -seq users

migrate:
	migrate -path ./pkg/database/migrations -database "sqlite3://pkg/database/database.db" -verbose up

migratedown:
	migrate -path ./pkg/database/migrations -database "sqlite3://pkg/database/database.db" -verbose down

test run:
	go test ./test/... -v -count=1

dev:
	go run ./cmd/api-gateway

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative user_service_grpc/user_service.proto
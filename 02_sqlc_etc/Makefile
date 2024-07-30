postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mysecretpassword -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb --username=root --owner=root simple_bank

migrateup:
	migrate -path ./db/migration -database "postgresql://root:mysecretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path ./db/migration -database "postgresql://root:mysecretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path ./db/migration -database "postgresql://root:mysecretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path ./db/migration -database "postgresql://root:mysecretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

server:
	 go run main.go

mock:
	mockgen -package mockdb  -destination db/mock/store.go izumarth_grpc/db/sqlc Store 

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server mock migrateup1 migratedown1 test
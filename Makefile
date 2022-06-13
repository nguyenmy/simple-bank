postgres:
	docker run --name my-postgres -p 5400:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest
rm-postgres:
	docker stop my-postgres
	docker rm my-postgres
createdb:
	docker exec -it my-postgres createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it my-postgres dropdb simple_bank
connectdb:
	docker exec -it my-postgres psql -U root 
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5400/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5400/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
.PHONY: rm-postgres postgres createdb dropdb connectdb migrate-up migrate-down sqlc test
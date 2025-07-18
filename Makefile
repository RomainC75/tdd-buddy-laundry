migrateup:
	migrate --path db/migration --database "postgresql://name:pass@localhost:5432/securitest?sslmode=disable" --verbose up

migratedown:
	migrate --path db/migration --database "postgresql://name:pass@localhost:5432/securitest?sslmode=disable" --verbose down

sqlc:
	cd adapters/secondary/repositories/sqlc && sqlc generate && cd -

test:
	go test -v -cover ./...

test-it:
	go test -v -tags=integration ./... 

gotestsum-it:
	gotestsum --watch -- -v -tags=integration ./... 

.PHONY: migrateup migratedown sqlc test
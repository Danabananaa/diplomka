# go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
# go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
run:
	go run ./cmd/main.go

migup:
	migrate -path ./db/migrations/ -database sqlite3://db/data.db up

migdown:
	migrate -path ./db/migrations/ -database sqlite3://db/data.db down

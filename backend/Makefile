run:
	go run ./cmd/main.go

migup:
	go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	migrate -path ./db/migrations/ -database sqlite3://db/data.db up

migdown:
	go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	migrate -path ./db/migrations/ -database sqlite3://db/data.db down


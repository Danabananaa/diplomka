# go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
# go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
run:
	go run ./cmd/main.go

migup:
	migrate -path ./db/migrations/ -database sqlite3://db/data.db up

migdown:
	migrate -path ./db/migrations/ -database sqlite3://db/data.db down

# curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
# mv migrate.linux-amd64 $GOPATH/bin/migrate
install-deps:
	go install github.com/golang-migrate/migrate@latest

migrate-up:
	migrate -database ${POSTGRESQL_URL} -path db/migrations up

migrate-down:
	migrate -database ${POSTGRESQL_URL} -path db/migrations down
install-migrate:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

install-staticcheck:
	go install honnef.co/go/tools/cmd/staticcheck@latest

install-golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.1
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.50.1
	wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.50.1

install-deps:
	make install-migrate
	make install-staticcheck
	make install-golangci-lint

migrate-up:
	migrate -database ${POSTGRESQL_URL} -path db/migrations up

migrate-down:
	migrate -database ${POSTGRESQL_URL} -path db/migrations down

lint:
	go vet ./...
	staticcheck ./...
	golangci-lint run ./...

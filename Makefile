install-migrate:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

install-staticcheck:
	go install honnef.co/go/tools/cmd/staticcheck@latest

install-golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.1
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.50.1
	wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.50.1

install-present:
	go get golang.org/x/net
	go get golang.org/x/tools
	go install golang.org/x/tools/cmd/present

install-deps:
	make install-migrate
	make install-staticcheck
	make install-golangci-lint
	make install-present

migrate-up:
	migrate -database ${POSTGRESQL_URL} -path db/migrations up

migrate-down:
	migrate -database ${POSTGRESQL_URL} -path db/migrations down

lint:
	go vet ./...
	staticcheck ./...
	golangci-lint run ./...

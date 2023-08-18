build:
	docker-compose build --no-cache
up:
	docker-compose up
down:
	docker-compose down -v
run:
	./app
tidy:
	go mod tidy

download:
	go mod download

lint:
	go install golang.org/x/tools/cmd/goimports@latest
	$(go env GOPATH)/bin/goimports -w .
	gofmt -s -w .

build_app:
	go build -v .

test:
	go test -v ./...

# see:
#     docker exec  -it go-tesma-api-app-1 cat

# db_login:
#   docker exec -it go-tesma-api-db-1 bin/bash
# mysql -u user -p
# USE dbname;


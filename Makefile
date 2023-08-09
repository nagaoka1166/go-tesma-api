build:
	docker build --no-cache -t go-tesma-api .
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


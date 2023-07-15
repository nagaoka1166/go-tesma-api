build:
    docker build --no-cache -t go-tesma-api .

up:
	docker-compose up -d

run:
    ./app

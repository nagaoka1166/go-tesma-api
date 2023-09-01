# Dockerfile
FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . /app

CMD ["go", "run", "main.go"]
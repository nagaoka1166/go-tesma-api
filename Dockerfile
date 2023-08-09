# Dockerfile
FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . /app

RUN go build -o main ./main.go

CMD ["go", "run", "main.go"]

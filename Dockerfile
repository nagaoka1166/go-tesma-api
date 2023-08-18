# Dockerfile
FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# COPY wait-for-it.sh /wait-for-it.sh
# RUN chmod +x /wait-for-it.sh

# COPY entrypoint.sh /entrypoint.sh
# RUN chmod +x /entrypoint.sh
# ENTRYPOINT ["/entrypoint.sh"]

COPY . /app

CMD ["go", "run", "main.go"]
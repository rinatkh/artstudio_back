FROM golang:1.19.3-alpine3.16

WORKDIR /app

COPY . .

CMD ["go", "run", "/app/cmd/api/main.go"]

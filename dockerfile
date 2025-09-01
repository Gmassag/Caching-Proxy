FROM golang:1.25-alpine3.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o caching-proxy cmd/main.go

EXPOSE 3000
ENTRYPOINT ["./caching-proxy"]
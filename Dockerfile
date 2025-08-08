FROM golang:1.23-alpine

WORKDIR /app

# Copiar go.mod e go.sum primeiro
COPY go.mod go.sum ./
RUN go mod download

# Copiar código fonte
COPY . .

# Build
RUN go build -o initdb ./cmd/initdb
RUN go build -o ordersystem ./cmd/ordersystem

# Executar sem sleep (health checks cuidam disso)
CMD ["sh", "-c", "./initdb && ./ordersystem"]
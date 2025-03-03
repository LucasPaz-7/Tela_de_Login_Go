# Estágio de build
FROM golang:1.24-alpine AS builder

# Instalando dependências de build
RUN apk add --no-cache gcc musl-dev

# Definindo o diretório de trabalho
WORKDIR /app

# Copiando os arquivos de dependência
COPY go.mod go.sum ./

# Baixando dependências
RUN go mod download

# Copiando o código fonte
COPY . .

# Compilando o aplicativo
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Estágio final
FROM alpine:latest

# Instalando certificados SSL/TLS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copiando o binário compilado do estágio anterior
COPY --from=builder /app/main .

# Expondo a porta que a aplicação utiliza
EXPOSE 8080

# Comando para executar a aplicação
CMD ["./main"] 
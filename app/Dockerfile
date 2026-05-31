FROM golang:1.23.12-alpine3.22 AS build
WORKDIR /build

# Copiando e baixando os modulos GO para o container
COPY go.mod go.sum ./
RUN go mod download

# Copiando a aplicação para o container e construindo o binário
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /build/http-server-projeto-korp

FROM alpine:3.22
WORKDIR /app

# Criando um usuário não-root para rodar a aplicação com segurança
RUN addgroup -S appuser && adduser -S -G appuser -H -s /sbin/nologin appuser

# Definindo proprietario do diretório de trabalho para o usuário criado e copiando o binário da aplicação para o container
COPY --from=build --chown=appuser:appuser /build/http-server-projeto-korp /app/http-server-projeto-korp
USER appuser
ENTRYPOINT ["/app/http-server-projeto-korp"]
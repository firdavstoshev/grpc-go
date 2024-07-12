# gRPC Account Management Service

Это пример простого gRPC сервиса для управления аккаунтами и балансами. Сервис поддерживает получение аккаунта, изменение имени аккаунта, изменение баланса аккаунта, создание и удаление аккаунта. Также предоставляется CLI для взаимодействия с сервисом.

## Структура проекта

- `account/account.proto`: Описание gRPC сервиса и сообщений.
- `server/server.go`: Реализация серверной части.
- `client/client.go`: Реализация клиентской части (CLI).

## Требования

- Go 1.16 или выше
- `protoc` (Protocol Buffers Compiler)
- Плагины для Go:
    - `protoc-gen-go`
    - `protoc-gen-go-grpc`

## Установка

### Установка protoc

#### macOS

```sh
brew install protobuf
```

### Команда генерации
```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative account.proto
```
## Запуск

### Запуск сервера

```
go run server/server.go
```

### Использование CLI

#### Создание аккаунта
```
go run client/client.go --action=create --name="Firdavs"
```

#### Получение аккаунта
```
go run client/client.go --action=get --id="1"
```

#### Изменение имени аккаунта
```
go run client/client.go --action=update-name --id="1" --name="Tekhronsho"
```

#### Изменение баланса аккаунта
```
go run client/client.go --action=update-balance --id="1" --balance=100.50
```

#### Удаление аккаунта
```
go run client/client.go --action=delete --id="1"
```

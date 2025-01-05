```go
go get github.com/budka-tech/gserv
```

# GServ

GServ - это легковесная обертка для gRPC сервера на Go, разработанная для упрощения процесса создания и управления gRPC сервисами.

## Особенности

- Простая инициализация и настройка gRPC сервера
- Поддержка пользовательских опций сервера
- Встроенное логирование с использованием пакета logit-go
- Graceful shutdown


## Использование
```go
import (
    "context"
    "github.com/budka-tech/gserv"
    "github.com/budka-tech/iport"
    "github.com/budka-tech/logit-go"
)

func main() {
    logger := logit.New()
    params := gserv.Params{
        Host:   "localhost",
        Port:   8080,
        Logger: logger,
    }

    server := gserv.NewGServ(params)

    ctx := context.Background()
    err := server.Init(ctx, 
        func(s *grpc.Server) {
            // Регистрация ваших gRPC сервисов
        },
    )

    if err != nil {
        logger.Error(ctx, err)
        return
    }

    // Ваш код для ожидания сигнала завершения

    server.Dispose(ctx)
}
```
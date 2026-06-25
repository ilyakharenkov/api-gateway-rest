# api-gateway-rest
```
Структура сервиса:
├── cmd/                        # Точки входа (исполняемые файлы)
│   └── main.go                 # Сборка зависимостей и запуск (очень тонкий слой)
│
├── internal/                   # Закрытый код (не доступен для импорта извне)
│   ├── usecases/               # Сценарии (оркестрация) — UseCases / Services
│   ├── domain/                 # Бизнес-сущности, правила (чистая логика)
│   ├── ports/                  # Интерфейсы для внешних слоев (driven ports)
│   └── adapters/               # Реализации интерфейсов (БД, Kafka, HTTP клиенты)
│       ├── repository/         # Работа с БД (sqlc, gorm)
│       ├── kafka/              # Продюсеры/консюмеры
│       └── rest/               # HTTP-клиенты к внешним API
│
├── pkg/                        # Публичные библиотеки (можно шарить между сервисами)
│   ├── logger/                 # Обертка над логгером
│   ├── errors/                 # Кастомные ошибки
│   └── utils/                  # Утилиты (слайсы, строки)
│
├── api/                        # Контракты (Protobuf/OpenAPI)
│   └── proto/                  # .proto файлы
│
├── configs/                    # Конфиги (.env, yaml, docker-compose)
├── deployments/                # Kubernetes/Docker манифесты
├── migrations/                 # SQL-миграции
├── go.mod
└── Makefile
```
# Gateway Service Documentation

[1. Инструкции по запуску](#Запуск-сервиса)\
[2. Дополнительные полезные команды](#Дополнительные-полезные-команды)\
[3. Архитектура](#Архитектура)

Searcher - сервис, который оперирует API эксплорер-сервиса [TonAPI](https://tonapi.io/) накладывая допольнительную бизнес-логику на получаемые данные 

## Запуск сервиса

Чтобы запустить сервис требуется описать .env файл и положить его в корень проекта

```
SEARCHER_TON_BASE_URL - адрес tonApi
SEARCHER_HOST - host сервиса
SEARCHER_PORT - порт сервиса
SEARCHER_SWAGGER_PATH - полный путь до сваггер документации (нахождится в api/swagger-ui)
```

Пример готового env файла есть в корне проекта и называется "env"

Затем следуйте данной инструкции для запуска

```
go mod download
go mod verify

make run-searcher
```

## Дополнительные полезные команды

```
/* Генерация сваггер-документации */
make swagger-regen
/* Запуск тестов */
make test
/* Запуск линтера */
make lint
```

## Архитектура

**internal/ton_client**

Сервис, который создает и отправляет запросы в предоставленную API от TonAPI

**internal/usecases**

Основаня бизнес-логика (пока просто в роли обертки над TonAPI) 

**internal/http_server**

HTTP-сервер сервиса

**internal/searcher**

Структура для сборки всех компонентов сервиса и последующего запуска сервиса
# Income Flow REST Service

## О проекте
Go REST API сервис для управления товарами на базе фреймворка Fiber v3. Предоставляет HTTP endpoints для выполнения CRUD операций с товарами.

## Технологии
- **Go 1.24.5** - основной язык
- **Fiber v3** - веб-фреймворк
- **JSON-iterator** - быстрая JSON сериализация
- **Logrus** - логирование
- **FastHTTP** - HTTP сервер

## Структура проекта
- `cmd/main.go` - точка входа в приложение
- `cmd/server.go` - настройка HTTP сервера
- `handler/handler.go` - роутинг API endpoints
- `handler/income_flow.go` - бизнес-логика (CRUD операции)
- `model/good.go` - структура товара
- `model/error.go` - структура ошибки
- `go.mod` - зависимости Go

## API Endpoints
- `POST /good/create` - создание товара
- `POST /good/get` - получение товара по ID
- `POST /good/get_all` - получение всех товаров
- `POST /good/update` - обновление товара
- `POST /good/delete` - удаление товара

## Модель данных
```go
type Good struct {
    ID          int     `json:"id"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Count       int     `json:"count"`
    Weight      float64 `json:"weight"`
}
```

## Запуск
```
go run cmd/main.go
```
Сервис запустится на порту 8080.

## Начальные данные
При запуске сервис содержит два предустановленных товара:
- iPhone 16 Pro (ID: 1, 1 шт, 0.2 кг)
- Dyson V12 (ID: 2, 2 шт, 3.0 кг)

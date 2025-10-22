## URL Shortener (Go + PostgreSQL + Docker)

Простой **сокращатель ссылок**, написанный на **Go** для закрепления навыков работы с:
- HTTP-сервером (`net/http`)
- обработчиками (handlers)
- PostgreSQL (через `database/sql`)
- Docker (контейнеризация приложения и БД)
- Graceful shutdown и контекстами (`context`)

---

## Функционал

- **POST /shorten** — принимает длинный URL и возвращает короткий код  
- **GET /{code}** — перенаправляет на исходный URL  
- **GET /links** — выводит список всех сохранённых ссылок  

---

## Технологии

- Go 1.22+
- PostgreSQL
- Docker / Docker Compose
  


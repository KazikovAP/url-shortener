[![Go](https://img.shields.io/badge/-Go-464646?style=flat-square&logo=Go)](https://go.dev/)
[![PostgreSQL](https://img.shields.io/badge/-PostgreSQL-464646?style=flat-square&logo=PostgreSQL)](https://www.postgresql.org/)
[![docker](https://img.shields.io/badge/-Docker-464646?style=flat-square&logo=docker)](https://www.docker.com/)

# url-shortener
# Сервис, предоставляющий API по созданию сокращённых ссылок

---
## Технологии
* Go 1.21.1
* PostgreSQL
* REST API
* Docker
* Postman

---
## Взаимодейастиве с сервисом

### Экспорировать путь до конфига
`export CONFIG_PATH="<path>\url-shortener\config\local.yaml"` 

### Запуск приложения в хранилище in-memory
`go run cmd/url-shortener/main.go --in_memory`

*Пример POST запроса на адрес* `http://localhost:8000/url`:
**Request:**
```JSON
{
    "url": "https://dzen.ru/news/story/story_one"
}
```
**Response:**
```JSON
{
    "status": "OK",
    "alias": "VRNYvumJmZ"
}
```
*Пример GET запроса на адрес* `http://localhost:8000/VRNYvumJmZ`:
**Response:**
```JSON
{
    "status": "OK",
    "url": "https://dzen.ru/news/story/story_one"
}
```
*Повторный POST запрос на адрес* `http://localhost:8000/url`:
**Request:**
```JSON
{
    "url": "https://dzen.ru/news/story/story_one"
}
```
**Response:**
```JSON
{
    "status": "Error",
    "error": "url already exists"
}
```

### Запуск докер контейнера с Postgres
`docker compose -p url-shortener -f ./build/docker-compose.yml up -d`

### Запуск приложения с Postgres
`go run cmd/url-shortener/main.go`

*Пример POST запроса на адрес* `http://localhost:8000/url`:
**Request:**
```JSON
{
    "url": "url": "https://www.youtube.com/videos/nature"
}
```
**Response:**
```JSON
{
    "status": "OK",
    "alias": "2f62KXpQms"
}
```
*Пример GET запроса на адрес* `http://localhost:8000/2f62KXpQms`:
**Response:**
```JSON
{
    "status": "OK",
    "url": "https://www.youtube.com/videos/nature"
}
```
*Повторный POST запрос на адрес* `http://localhost:8000/url`:
**Request:**
```JSON
{
    "url": "https://www.youtube.com/videos/nature"
}
```
**Response:**
```JSON
{
    "status": "Error",
    "error": "url already exists"
}
```

### Остановка и удаление докер контейнер с Postgres
`docker compose -p url-shortener -f ./build/docker-compose.yml down`

---
## Разработал:
[Aleksey Kazikov](https://github.com/KazikovAP)

---
## Лицензия:
[MIT](https://opensource.org/licenses/MIT)

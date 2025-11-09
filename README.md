# Erzi — Интернет-магазин воды

**Erzi** — это полноценный интернет-магазин для продажи воды и оборудования. Пользователи могут регистрироваться, просматривать товары, добавлять их в корзину и оформлять заказы. Администраторы могут создавать, изменять и удалять продукты.

Проект реализует логику интернет-магазина с хранением данных в PostgreSQL, авторизацией через JWT и бизнес-логикой для корзины и продуктов.

---

## Основные возможности

- Регистрация и вход пользователей
- CRUD для продуктов (для админов)
- Работа с корзиной пользователя
- Аутентификация JWT (user/admin)
- Валидация данных через `ozzo-validation`
- Хеширование паролей через `bcrypt`
- Логирование запросов
- Конфигурация через TOML (`config.toml`)
- Запуск через Docker Compose

---

## Архитектура проекта

.
├── cmd/
│ └── apiserver/ # HTTP-сервер
├── internal/
│ ├── handler/ # REST API
│ ├── service/ # бизнес-логика
│ ├── repository/ # доступ к PostgreSQL
│ ├── store/ # инициализация БД
│ └── config/ # загрузка TOML-конфига
├── db/
│ ├── migrations/ # миграции базы
│ └── insert/ # сиды
├── docker/
│ ├── Dockerfile # сборка API
│ └── Dockerfile.postgres
├── docker-compose.yml
├── config.toml
└── go.mod

---

## Требования

- Go 1.24+
- Docker + Docker Compose
- PostgreSQL

---

## Установка и запуск

1. Клонируем проект:

```bash
git clone <repo-url>
cd erzi-shop
Поднимаем инфраструктуру:
docker compose up -d db
Запуск сервера:
go run cmd/apiserver/main.go -config config.toml
Проверка работы сервера:
curl http://localhost:8080/products
Конфигурация (config.toml)
bind_addr = ":8080"
log_level = "debug"

[store]
database_url = "postgres://appuser:secret@localhost:5432/erziapp?sslmode=disable"
API
Пользователи
Метод	URL	Описание
POST	/user/register	Регистрация
POST	/user/login	Вход (JWT)
Продукты
Метод	URL	Описание
GET	/products	Список всех продуктов
GET	/products/:id	Получение продукта по ID
POST	/products/create	Создание продукта (admin)
PUT	/products/:id	Обновление продукта (admin)
DELETE	/products/:id	Удаление продукта (admin)
Корзина
Метод	URL	Описание
POST	/cart/items/:product_id/add	Добавить товар в корзину
GET	/cart/items	Получить корзину
PUT	/cart/items/:id/increment	Увеличить количество
PUT	/cart/items/:id/decrement	Уменьшить количество
DELETE	/cart/items/:id	Удалить товар
PUT	/cart/status	Пометить корзину как "deleted"
Модель данных
Users
id UUID PRIMARY KEY
username TEXT
email TEXT UNIQUE
password TEXT
role TEXT DEFAULT 'user'
created_at TIMESTAMP
Products
id SERIAL PRIMARY KEY
title TEXT
description TEXT
price NUMERIC
image_url TEXT
category TEXT
quantity INT
created_at TIMESTAMP
Carts
id SERIAL PRIMARY KEY
user_id UUID REFERENCES users(id)
status TEXT DEFAULT 'active'
created_at TIMESTAMP
Cart_items
id SERIAL PRIMARY KEY
cart_id INT REFERENCES carts(id)
product_id INT REFERENCES products(id)
quantity INT DEFAULT 1
created_at TIMESTAMP
Docker Compose
version: '3.8'
services:
  db:
    build:
      context: .
      dockerfile: docker/Dockerfile.postgres
    environment:
      POSTGRES_USER: appuser
      POSTGRES_PASSWORD: secret
      POSTGRES_DB: erziapp
    ports:
      - "5432:5432"
    volumes:
      - pgdata:/var/lib/postgresql/data

  api:
    build: .
    depends_on: [db]
    environment:
      DATABASE_URL: postgres://appuser:secret@db:5432/erziapp?sslmode=disable
    ports:
      - "8080:8080"

volumes:
  pgdata:
Тестирование
go test ./...
Возможное развитие
Refresh-токены и улучшенная авторизация JWT
Кэширование корзины через Redis
Метрики Prometheus
OpenAPI документация
Веб-панель администратора
Полезные ссылки
Gin
PostgreSQL
bcrypt
JWT (lestrrat-go/jwx)
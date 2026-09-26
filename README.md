# Hotel_Project

[![Go](https://img.shields.io/badge/Go-1.23-00ADD8?style=flat-square&logo=go&logoColor=white)]()
[![Gin](https://img.shields.io/badge/Gin-Framework-00ADD8?style=flat-square)]()
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=flat-square&logo=postgresql&logoColor=white)]()
[![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat-square&logo=docker&logoColor=white)]()

Go programlama dili ile geliştirilmiş, staj projem kapsamında hazırlanan bir otel yönetim REST API'si.

## Proje Hakkında

Uygulama; müşteri, oda, oda tipi, rezervasyon, ödeme, kullanıcı ve admin kaynaklarını yöneten bir REST API sunar. [Gin](https://github.com/gin-gonic/gin) web framework'ü ve [GORM](https://gorm.io/) ORM'i ile PostgreSQL üzerinde çalışır; JWT tabanlı kimlik doğrulama içerir.

## Kullanılan Teknolojiler

- Go, Gin
- GORM + PostgreSQL
- JWT (golang-jwt)
- Docker / Docker Compose

## API Uç Noktaları

Uygulama `/api` altında şu route grupları ile çalışır:

- `/api/customers`
- `/api/roomTypes`
- `/api/rooms`
- `/api/payments`
- `/api/users`
- `/api/reservations`
- `/api/admins`

## Kurulum ve Çalıştırma

### Docker ile

```bash
git clone https://github.com/mrvbyrm/Hotel_Project.git
cd Hotel_Project/stajprojesi
docker-compose up --build
```

### Manuel

```bash
cd Hotel_Project/stajprojesi
go mod download
go run main.go
```

Uygulama varsayılan olarak `8081` portunda çalışır (ortam değişkeni `PORT` ile değiştirilebilir). Veritabanı bağlantı bilgileri `DB_HOST`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, `DB_PORT` ortam değişkenleriyle yapılandırılır; örnek değerler için `.env.example` dosyasına bakın.

## İletişim

Merve — [GitHub](https://github.com/mrvbyrm)

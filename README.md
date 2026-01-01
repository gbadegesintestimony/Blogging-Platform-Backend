# Blogging Platform Backend

A simple blogging platform backend built with Go, Gin, and GORM (Postgres).

## Features

- User registration & login (JWT)
- CRUD for posts
- Categories and comments

## Requirements

- Go 1.24+
- PostgreSQL

## Setup

1. Copy `.env.example` to `.env` and fill in values:

   PORT=8080
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=your_password_here
   DB_NAME=blog_db
   DB_SSLMODE=disable
   JWT_SECRET=your_jwt_secret_here
   JWT_EXPIRATION_HOURS=24

2. Install dependencies and run:

```bash
go mod download
go run ./cmd/
```

The server will run on the port specified in `.env`.

## API

- POST `/api/auth/register` — register new user
- POST `/api/auth/login` — login and receive JWT
- GET `/api/posts` — list posts
- GET `/api/posts/:id` — get a post
- GET `/api/categories` — list categories
- Protected routes (require `Authorization: Bearer <token>`):
  - POST `/api/posts` — create post
  - PUT `/api/posts/:id` — update post
  - DELETE `/api/posts/:id` — delete post
  - POST `/api/categories` — create category
  - POST `/api/posts/:postid/comments` — add comment to post

## Notes

- `.env` is ignored by `.gitignore`. Use `.env.example` as a template.
- The project uses GORM AutoMigrate in development; manual migrations are recommended for production.

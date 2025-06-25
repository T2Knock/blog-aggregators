# Blog Aggregators

A Go-based service that aggregates blog posts from RSS feeds.

## Prerequisites

- Go 1.20+ installed
- Git
- PostgreSQL database

## Installation

```bash
go install github.com/T2Knock/blog-aggregators
```

## Configuration

Configure environment variables via `~/.gatorconfig` or update `internal/config/config.go`:

Example configuration:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": "joe"
}
```

## Database Migrations

Schema files are in `sql/schema`. Apply migrations using [goose](https://github.com/pressly/goose) or manually via `psql` :

```bash
goose up
```

```bash
psql "$DATABASE_URL" -f sql/schema/20250618134959_001_users.sql
psql "$DATABASE_URL" -f sql/schema/20250624040904_001_feeds.sql
# ...repeat for other .sql files
```

## Build

```bash
go build
```

## License

MIT

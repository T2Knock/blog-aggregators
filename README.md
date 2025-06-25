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

## Project Structure

```sh

.
├── internal
│   ├── config
│   │   ├── config.go
│   │   ├── read.go
│   │   └── set_user.go
│   └── database
│       ├── db.go
│       ├── feed_follows.sql.go
│       ├── feeds.sql.go
│       ├── models.go
│       ├── posts.sql.go
│       └── users.sql.go
├── sql
│   ├── queries
│   │   ├── feed_follows.sql
│   │   ├── feeds.sql
│   │   ├── posts.sql
│   │   └── users.sql
│   └── schema
│       ├── 20250618134959_001_users.sql
│       ├── 20250624040904_001_feeds.sql
│       ├── 20250624065950_create_feed_follows.sql
│       ├── 20250624111901_add_last_fetched_at.sql
│       └── 20250624125645_create_posts.sql
├── AGENTS.md
├── commands.go
├── go.mod
├── go.sum
├── handler_agg.go
├── handler_feed.go
├── handler_feed_follow.go
├── handler_post.go
├── handler_user.go
├── main.go
├── middleware.go
├── README.md
├── rss_feed.go
└── sqlc.yaml
```

## License

MIT

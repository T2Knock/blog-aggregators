# Blog Aggregators

A simple Go service to aggregate blog posts from RSS feeds.

## Prerequisites

- Go 1.20+
- Git
- PostgreSQL

## Installation

```bash
go install github.com/T2Knock/blog-aggregators
```

## Configuration

Set environment variables in `~/.gatorconfig` or edit `internal/config/config.go`:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": "joe"
}
```

## Database

Apply migrations in `sql/schema` using goose or psql:

```bash
goose up
```

or

```bash
psql "$DATABASE_URL" -f sql/schema/20250618134959_001_users.sql
# repeat for other .sql files
```

## Build

```bash
go build
```

## Commands

- `login <name>`: Switch current user.
- `register <name>`: Register and switch user.
- `reset`: Delete all users.
- `users`: List users.
- `agg <duration>`: Aggregate feeds every duration (e.g. 10s, 1m).
- `addfeed <name> <url>`: Add and follow feed.
- `feeds`: List feeds.
- `follow <url>`: Follow feed.
- `unfollow <url>`: Unfollow feed.
- `following`: List followed feeds.
- `browse [limit]`: Browse posts, default 2.

🤖 Generated with [opencode](https://opencode.ai)

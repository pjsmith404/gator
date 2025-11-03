# Gator

A CLI tool for aggregating RSS feeds

# Requirements

- Postgres v17
- Go

# Installing

```go install github.com/pjsmith404/gator```

# Getting Started

1. Create the file `~/.gatorconfig.json` in your homedir, per the below example. Adjust the DB URL as appropriate.

```json
{
    "db_url":"postgres://postgres:@localhost:5432/gator?sslmode=disable"
}
```
2. Register a user using `gator register <username>`
3. Add a feed using `gator addfeed <name> <url>`
4. Start collecting posts by running `gator agg <sleep duration>`
5. Browse recent posts using `gator browse`

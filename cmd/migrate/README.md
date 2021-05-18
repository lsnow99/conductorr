# Migrate

Internal utility for development - allows developers to use [go-migrate](https://github.com/golang-migrate/migrate) without installing the binary.

This was necessary since the default cli tool that they provide does not come shipped with `sqlite3` support, as it requires CGo (like all current SQLite drivers for Go)

## Usage

To run the utility, use `go run ./cmd/migrate` and supply a command

Available commands:
- `help` - Show a help menu
- `create` - Create a new migration
- `force` - Force update the database version
# Migrate

Internal utility for development - allows developers to use [go-migrate](https://github.com/golang-migrate/migrate) without installing the binary.

(The history here is that the original `go-migrate` binary did not have built-in support for `sqlite3` as it requires cgo. So, I made this quick-and-dirty utility for dealing with migrations that somewhat wraps the `go-migrate` binary but adds `sqlite3` support. Now, both `go-migrate` and Conductorr have transitioned to supporting the `sqlite` driver so this utility is no longer needed if you have the `go-migrate` binary. However, I opted to leave it here for convenience as it should still work fine)

## Usage

To run the utility, use `go run ./cmd/migrate` and supply a command

Available commands:
- `help` - Show a help menu
- `create` - Create a new migration
- `force` - Force update the database version

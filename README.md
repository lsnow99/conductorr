# Conductorr

## Usage

Conductorr targets both SQLite and PostgreSQL as databases. Currently, PostgreSQL support is untested and may be unstable. For the time being it is recommended to use SQLite (no extra action is required, simply do not set the `PG_` environment variables).

Conductorr has a public [Docker image](https://hub.docker.com/r/logansnow/conductorr-pre), but it can also be run as a regular executable. To build the executable for your system, simply clone the repository and run the `build.sh` script to generate cross-platform binaries in the `bin` folder.

### Environment Variables

- `JWT_SECRET` Secret string that JWTs will be signed with. (**required**)
- `OMDB_API_KEY` API Key for [The Open Movie Database](https://omdbapi.com/). (**required** for searching new content)
- `CONDUCTORR_DEBUG` Run Conductorr in debug mode if this environment variable is set to anything other than the empty string (optional, default is empty string)
- `JWT_EXP_DAYS` Number of days you want a session to stay valid for (optional, default 7)
- `RESET_USER` If you forgot your password, set this environment variable to a non empty string and Conductorr will prompt you to set up a new user.
- `PG_USER` PostgreSQL user. If set, Conductorr will not use SQLite and will instead use PostgreSQL (optional - required for postgres support, default is empty string)
- `PG_PASS` PostgreSQL user password (optional - required for postgres support, default is empty string)
- `PG_NAME` PostgreSQL database name (optional - required for postgres support, default is empty string)
- `PG_PORT` PostgreSQL database port (optional - required for postgres support, default is empty string)
- `PG_SSL` If set to anything other than the empty string, Conductorr will force SSL for PostgreSQL connections (optional)
- `DB_PATH` When using SQLite, this is the path for the database file. Do not append any URL parameters to the end.

## Developing

To develop on Conductorr, first clone the repository.

[Visual Studio Code](https://code.visualstudio.com/) is recommended for development, and some instructions are specific to VS Code.

### Backend Development

Conductorr backend requires [go](https://golang.org/) to be installed. For development, any stable Go version should do fine. However, to build the production binaries Conductorr requires Go version ^1.16.0. This is because the production binary uses the [embed](https://golang.org/pkg/embed/) package which was introduced in Go v1.16.0.

Within VS Code, select "Launch Backend" from the debug menu. You can set breakpoints and debug as usual within VS Code. To run the backend without VS Code, use the command `CONDUCTORR_DEBUG=true go run ./cmd/conductorr`

### WebAssembly CSL Module

Conductorr's profile editor page uses a WebAssembly module to do client-side validation and parsing of CSL scripts. Before developing on the frontend, you must generate this module using `./build_csl.sh` (For Windows, you can either run this in WSL or just run the equivalent commands in cmd/PowerShell). This script assumes you have [Brotli](https://github.com/google/brotli) installed and available in your PATH.

> NOTE: If you skip this step, it is likely you will run into an issue where `wasm_exec.js` is not found when testing locally in the frontend, or when building for production.

### Frontend Development

Conductorr is built on Vue 3 using Vite and [TailwindCSS](https://tailwindcss.com/).

- `cd web`
- `yarn install`
- `yarn dev`

This will start the web frontend and should be available at http://localhost:3000

Modifying files within `web/src` and saving them will cause the website to hot-reload in your browser, aiding in development.

Conductorr uses Tailwind for all styling. [Oruga](https://oruga.io) components are used and styled via Tailwind.

### Database Migrations

Conductorr uses [go-migrate](https://github.com/golang-migrate/migrate) to manage database migrations. The source for migrations are located in the `migrations` folder. When building for release, the migration files are bundled into Conductorr's binary executable via the [embed](https://golang.org/pkg/embed/) package. Alternatively, the directory for Conductorr to look for migration files can be specified via the `MIGRATIONS_PATH` environment variable.

#### Creating new Migrations

Conductorr comes with a customized cli for go-migrate. Simply run `go run ./cmd/migrate create {migration name}` within Conductorr's root folder to create a new migration. This will create two SQL files in the migrations folder. One that contains the SQL to spin up the migration, and one to spin it down. Alternatively, you could make these files manually by following the naming pattern in the directory (increment the version number for the new up and down migrations).

In order to write effective migrations, abide by the following conventions:
- Always create both an up and down migration. The down migration should undo any changes made to the database schema by the up migration (ie. if the up migration creates a table or adds a column, the down migration should drop it). Where possible, data should be preserved. Down migrations are not intended ever to be run by end users, and are primarily good practice and useful for development. If a table needs to be removed in a future version, a proper up migration should handle it (with a corresponding down migration to add it back).
- Always wrap migration files in transactions. This means the first line should be `BEGIN;` and the last line should be `COMMIT;` in any migration file.
- Since Conductorr attempts to target and support both SQLite and PostgreSQL, try to keep all SQL in migration files cross-compatible between PostgreSQL and SQLite. Where not possible, default to the SQLite syntax (ie. use `blob` instead of `bytea`).  Conductorr runs some basic find/replace operations to convert SQLite SQL to PostgreSQL-compatible SQL before applying PostgreSQL migrations. To check if your code will convert correctly, look in the file `./database/conversion.go`.

#### Dealing with a Failed Migration

When developing, sometimes you might write a migration and run Conductorr to test it, and it fails. `go-migrate` will still increase the version of the database and mark it dirty (When starting Conductorr again it will fail due to a dirty database). In general the only thing you will need to do to fix a dirty database is to run this command: `go run ./cmd/migrate force 0 -path migrations/ -database sqlite3://conductorr.db` (replace 0 with one less than the version the database just got bumped to - in other words, if Conductorr outputs "Dirty database version 12. Fix and force version" after a failed up migration, then run `go run ./cmd/migrate force 11 -path migrations/ -database sqlite3://conductorr.db`). Note that these instructions are for when Conductorr tries to perform the up migration. If for some reason you see this error during a down migration, then you'll want to increase the version number by one instead of decreasing it.

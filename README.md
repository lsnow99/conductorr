<p align="center">
  <img src="./frontend/public/logo-rect.svg" alt="Conductorr logo image" width="200px" /> 
</p>

# Conductorr

Conductorr is a PVR for BitTorrent and Usenet that aims to strike a balance between lightweight, powerful, and easy to use. It supports both TV Shows and Movies, and allows you to fine tune exactly which releases to prefer via the custom scripting language CSL.

## Overview

### Major Features Include:
- Search for movies and TV shows to add them to your library, along with descriptions, posters, and ratings.
- Monitor new releases for pisodes, seasons, full shows, and movies for automatic download
- Ability to add media servers (plex, jellyfin - WIP) which will be notified of new content to scan automatically.
- Sync availability status from media servers for visibility in your Conductorr library
- Accessible and pleasant UI
- Highly-advanced customization for release filtering and priortiziation via CSL scripts

### Status of Support for External Services:
- [NZBGet](https://github.com/nzbget/nzbget) - **YES**
- [Transmission](https://github.com/transmission/transmission) - **YES**
- [qBittorrent](https://github.com/qbittorrent/qBittorrent) - **NO/PLANNED**
- [rTorrent/ruTorrent](https://github.com/rakshasa/rtorrent) - **NO/PLANNED**
- [Deluge](https://github.com/deluge-torrent/deluge) - **NO/PLANNED**
- [Plex](https://plex.tv) - **YES**
- [JellyFin](https://github.com/deluge-torrent/deluge) - **NO/PLANNED**
- [Emby](https://emby.media/) - **NO/PLANNED**
- [Kodi](https://github.com/xbmc/xbmc) - **NO/PLANNED**

If the service you use is not yet supported or planned, please feel free to request it in a Feature Request. Pull requests are also welcome. New downloaders simply need to implement the interface specified in `integration/downloaders.go` and pass testing.

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

## FAQ

### Why is Conductorr Needed and How Does it Differ from Other PVRs?
Conductorr actually began as an entirely [different project]() that glued together Radarr, Sonarr, Filebot, and Plex to process downloads, specifically in Kubernetes where a library is mounted over NFS. After growing frustrated of managing this array of different services, I thought it would be a fun project to build a replacement to this pipeline that handled everything in one elegant service. Other more mature PVRs like Sonarr and Radarr are great choices for most users, however, Conductorr differs in a few key ways that may make it more desirable for certain setups. For one, Conductorr supports both movies and tv series all in one interface. Conductorr also behaves nicely in containerized environments. For power users, CSL scripts can be customized very precisely to ensure the correct releases are always downloaded.

## Developing

To develop for Conductorr, first clone the repository.

[Visual Studio Code](https://code.visualstudio.com/) is recommended for development, and some instructions are specific to VS Code. However, you can simply ignore the specific instructions and use your own editor.

### Backend Development

Conductorr backend requires [go](https://golang.org/) to be installed. For development, any stable Go version should be fine. However, to build the production binaries Conductorr requires Go version ^1.16.0. This is because the production binary uses the [embed](https://golang.org/pkg/embed/) package which was introduced in Go v1.16.0.

Within VS Code, select "Launch Backend" from the debug menu. You can set breakpoints and debug as usual within VS Code. To run the backend without VS Code, use the command `CONDUCTORR_DEBUG=true go run ./cmd/conductorr`.

### WebAssembly CSL Module

Conductorr's profile editor page uses a WebAssembly module to do client-side validation and evaluationms of CSL scripts. Before developing on the frontend, you must generate this module using `./build.sh csl` (For Windows, you can either run this in WSL or just run the equivalent commands in cmd/PowerShell). This script assumes you have [Brotli](https://github.com/google/brotli) installed and available in your PATH.

> NOTE: If you skip this step, it is likely you will run into an issue where `wasm_exec.js` is not found when serving the frontend.

### Shared Library

The Shared Library is a Vue 3 project that exports components which are intended to be shared accross both the documentation site and the web application.

To develop the Shared Library with hot-reloading:

- `pnpm install`
- `cd shared`
- `pnpm dev`

To build the Shared Library (so that the `frontend` and `docusite` projects can use it):

- `pnpm install`
- `cd shared`
- `pnpm build`

> NOTE: This build step must be performed before developing in the `frontend` or `docusite` projects, and any time the Shared Library changes.

### Frontend Development

> NOTE: Requires the [WebAssembly CSL Module](#webassembly-csl-module) and [Shared Library](#shared-library) to be built first. This only needs to happen the first time after you clone the repository or any time code is modified in the CSL module or the Shared Library.

Conductorr is built on Vue 3 using Vite and [TailwindCSS](https://tailwindcss.com/).

- `pnpm install`
- `cd frontend`
- `pnpm dev`

This will start the frontend app and should be available at http://localhost:3000/

Modifying files within `frontend/src` and saving them will cause the website to hot-reload in your browser, aiding in development.

Conductorr uses Tailwind for all styling. [Oruga](https://oruga.io/) components are used and styled via Tailwind.

### Database Migrations

Conductorr uses [go-migrate](https://github.com/golang-migrate/migrate) to manage database migrations. The source for migrations are located in the `migrations` folder. When building for release, the migration files are bundled into Conductorr's binary executable via the [embed](https://golang.org/pkg/embed/) package. Alternatively, the directory for Conductorr to look for migration files can be specified via the `MIGRATIONS_PATH` environment variable.

#### Creating new Migrations

Conductorr comes with a customized cli for go-migrate. Simply run `go run ./cmd/migrate create {migration name}` within Conductorr's root folder to create a new migration. This will create two SQL files in the migrations folder. One that contains the SQL to spin up the migration, and one to spin it down. Alternatively, you could make these files manually by following the naming pattern in the directory (increment the version number for the new up and down migrations).

In order to write effective migrations, abide by the following conventions:
- Always create both an up and down migration. The down migration should undo any changes made to the database schema by the up migration (ie. if the up migration creates a table or adds a column, the down migration should drop it). Where possible, data should be preserved. Down migrations are not intended ever to be run by end users, and are primarily good practice and useful for development. If a table needs to be removed in a future version, a proper up migration should handle it (with a corresponding down migration to add it back).
- Always wrap migration files in transactions. This means the first line should be `BEGIN;` and the last line should be `COMMIT;` in any migration file.
- Since Conductorr attempts to target and support both SQLite and PostgreSQL, try to keep all SQL in migration files cross-compatible between PostgreSQL and SQLite. Where not possible, default to the SQLite syntax (ie. use `blob` instead of `bytea`).  Conductorr runs some basic find/replace operations to convert SQLite SQL to PostgreSQL-compatible SQL before applying PostgreSQL migrations. To check if your code will convert correctly, look in the file `./database/conversion.go`.

#### Dealing with a Failed Migration

When developing, sometimes you might write a migration and run Conductorr to test it, and it fails. `go-migrate` will still increase the version of the database and mark it dirty (When starting Conductorr again it will fail due to a dirty database). In general the only thing you will need to do to fix a dirty database is to run this command: `go run ./cmd/migrate -database sqlite://conductorr.db force 0` (replace 0 with one less than the version the database just got bumped to - in other words, if Conductorr outputs "Dirty database version 12. Fix and force version" after a failed up migration, then run `go run ./cmd/migrate -database sqlite://conductorr.db force 11 `). Note that these instructions are for when Conductorr tries to perform the up migration. If for some reason you see this error during a down migration, then you'll want to increase the version number by one instead of decreasing it.

#### Rolling back a Successful Migration

If the migration succeeded, but you wish to undo it later on, you can run the following command to roll it back: `go run ./cmd/migrate -database sqlite://conductorr.db down 1`. Replace `1` with the number of migrations you would like to roll back.

## Open Source

### Attributions
Like most software, Conductorr is built upon the backs of many open source projects.
The direct dependencies of Conductorr are listed here:
- [Oruga](https://oruga.io/) - CSS agnostic Vue components
- [Font Awesome](https://fontawesome.com/) - Free, regular, brands, and core icon packages
- [Luxon](https://moment.github.io/luxon/) - Successor to Moment.js; a utility to deal with dates and times in JavaScript
- [Vue Prism Editor](https://github.com/koca/vue-prism-editor) - Vue component for code editing
- [prism.js](https://prismjs.com) - Code editor syntax highlighting
- [split.js](https://split.js.org) - Utility to split html views
- [Vue](https://v3.vuejs.org/) - Frontend framework
- [Vite](https://vitejs.dev/) - Frontend tooling with HMR and building via Rollup
- [mitt](https://github.com/developit/mitt) - Lightweight JavaScript event emitter/pubsub
- [PostCSS](https://postcss.org/) - Transform CSS with JavaScript
- [TailwindCSS](https://tailwindcss.com) - CSS framework
- [semver for golang](https://github.com/blang/semver) - Semantic Versioning library written in golang
- [jwt-go](https://github.com/golang-jwt/jwt) - JWT package for Go
- [go-ethereum](https://github.com/ethereum/go-ethereum) Golang implementation of the Ethereum protocol - used for it's implementation of rpc
- [migrate](https://github.com/golang-migrate/migrate) - Database migrations for Go
- [gorilla/mux](https://github.com/gorilla/mux) - HTTP request multiplexer for Go
- [TransmissionRPC](https://github.com/hekmon/transmissionrpc) - Golang bindings to Transmission (bittorrent) RPC interface
- [pgx](https://github.com/jackc/pgx) - Pure Go driver for PostgreSQL
- [sqlite](https://gitlab.com/cznic/sqlite) - Package sqlite is a CGo-free port of SQLite/SQLite3 v3.37.0
- [archiver](https://github.com/mholt/archiver) - Pure Go utility to interface with common archive formats
- [go-newznab](https://github.com/mrobinsn/go-newznab) - Newznab/Torznab bindings for Go

For a full list of packages used and their versions/licenses, please refer to the [go.mod](./go.mod) and package.json files

### License
- [GNU GPL v3](https://www.gnu.org/licenses/gpl-3.0.html)

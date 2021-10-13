# Configuration

Conductorr uses the following environment variables for configuration

| Environment Variable  | Description  |
| --------------------- | ------------ |
| TMDB_API_KEY          | API key for searching media and fetching media metadata from [The Movie Database](https://themoviedb.org)  |
| CONDUCTORR_DEBUG      | Run Conductorr in debug mode if this environment variable is set to anything |
| JWT_EXP_DAYS          | Number of days a user should stay logged in for without logging in |
| RESET_ADMIN_USER      | If you forgot your password, set this environment variable to true and Conductorr will prompt you to set up a new user next time it starts |
| PG_USER               | PostgreSQL user. If set, Conductorr will not use SQLite and will instead use PostgreSQL (optional - required for postgres support) |
| PG_PASS               | PostgreSQL user password (optional - required for postgres support)
| PG_DATABASE           | PostgreSQL database name (optional - required for postgres support) |
| PG_PORT               | PostgreSQL server port (optional - default is 5432) |
| PG_SSL                | If set, Conductorr will force SSL for PostgreSQL connections (optional) |
| DB_PATH               | When using SQLite, this is the path for the database file. Do not append any URL parameters to the end |

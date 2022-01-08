# durn 2.0

## Local Setup

1. Install `golang` and `PostgreSQL`.
2. Start a `psql` server on your system
    - Create a database with some name, eg `CREATE DATABASE durn;`
3. Set the required env-variables
4. `go run main.go`
    - Automatically installs dependencies

## Environment variables:

Name | Description | Default
--- | --- | ---
`ADDR` | Address for the service to run on | `localhost:8080`
`LOGIN_API_KEY` | Login API key for KTH authentication. Value is ignored if `SKIP_AUTH` is true | ---
`DB_HOST` | Host URL for the database | `localhost`
`DB_PORT` | Host port for the database | `5432`
`DB_USER` | Username for the `psql` server | ---
`DB_PASSWORD` | Password for the `psql` user | ---
`DB_NAME` | Name of the `psql` database | ---
`SKIP_AUTH` | Skips auth steps if set to `true`. Only for local development, should never be set in production | `false`

Variables can be set in an `.env` file in root, and will then automatically be loaded into the environment upon `go run`.

Will be added later:
- PLS_URL
- LOGIN_URL
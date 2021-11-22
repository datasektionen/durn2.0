# durn 2.0

## Setup

1. install and start a psql server on your system
    - create a database with some name, eg `CREATE DATABASE durn;`
    - create tables with `\i database/models.sql`
2. Set required env-variables
3. `go run main.go`

## Required environment variables:

Currently:
- `DB_PORT` - defaulted to `5432` if not set 
- `DB_USER` - psql username
- `DB_PASSWORD` - psql user password
- `DB_NAME` - name of the databse in psql
- `ADDR` - address that the server will be ran on, set to something like "`:8000`" for local development

Will be added later:
- LOGIN_API_KEY
- PLS_URL
- LOGIN_URL
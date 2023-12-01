# database migration

Using [golang-migrate](https://github.com/golang-migrate/migrate) to scaffold a local postgres database

## Prerequisites

## How to
1. Create a new `snippetbox` database in local postgres server using `root` user and default `template1` database :
```sh
psql -h localhost -d template1 -c "CREATE DATABASE snippetbox;"
```

2. Create a new user `snippetbox` with password:
```sh
psql -h localhost -d template1 -c "CREATE ROLE snippetbox WITH LOGIN PASSWORD 'password';"
```

3. We can now login into our `snippetbox` db as `snippetbox` user:
```sh
psql -h localhost -d snippetbox -U snippetbox
```


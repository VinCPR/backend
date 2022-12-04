# VinCPR backend server

## Prerequisites
* Go 1.18
* Docker
* golang-migrate
* sqlc

## Quickstart

### Create a new file `app.env` and copy from `app.env-example`

### Start postgres container
```shell
$ make postgres
```

### Create database
```shell
$ make createdb
```

### Run database migration up
```shell
$ make migrateup
```

### Run test
```shell
$ make test
```

### Run server
```shell
$ make server
```

## Optional: Generating code and docs

### Generate SQL CRUD with sqlc
```shell
$ make sqlc
```

### Generate swagger docs
```shell
$ make gen-swagger
```

# 🚀 Confetti

[Password cards](https://github.com/sultaniman/pwc) as a service backend.

Note: `keys/key.pem` is development key so can be public.

## WORK IN PROGRESS

## OpenAPI spec

It is under `docs`

## Build

To build project you need to use

```sh
$ go build -o confetti
```

## Configuration

All configuration is done via environment variables and the following variables are expected to be defined

```dotenv
CO_DB_URI=postgresql://postgres:postgres@localhost:5432/confetti?sslmode=disable
CO_MIGRATIONS=file:///<PATH_TO>/confetti/migrations
CO_PRIVATE_KEY=file:///<PATH_TO>/confetti/keys/key.pem
CO_KEY_LOADER=local
CO_REFRESH_TOKEN_TTL=4320h
CO_ACCESS_TOKEN_TTL=1h
```

## Generate key

```sh
$ openssl genrsa -out key.pem 2048
```

## To generate swagger

```sh
$ go install github.com/golang/mock/mockgen@v1.6.0
```

## Migrate

First create database then run migrations to create tables

```sh
$ ./confetti migrate
```

## Run the server

To run the server the following command should be used once it starts
it is ready to accept requests at `http://localhost:4000`

```sh
$ ./confetti serve
```

# 🚀 Confetti API

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

## Migrate

To create database tables

```sh
$ ./confetti migrate
```

## Run the server

To run the server the following command should be used once it starts
it is ready to accept requests at `http://localhost:4000`

```sh
$ ./confetti serve
```

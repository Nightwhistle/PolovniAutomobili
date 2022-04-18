# MeridianBet Betradar API
This microservice serves data fetched from Betradar from mongodb that's being populated by [**Populate stream application**](https://gitlab.com/meridianbet/meridianbet-stream-populate) through mem-cache [**Ristretto**](https://github.com/dgraph-io/ristretto). Application also serves as proxy server for frontend to communicate with Betradar since Betradar only whitelists single IP address

## Installation

- Clone the repository

- Install go language by following official tutorial: [link](https://go.dev/doc/install)

- Copy `.env.example` to `.env` and edit it to adjust your environment

- Install dependencies with `go mod tidy`

- Start application with `go run main.go`


## Info
- Aplication starts server on port defined in `.env`
- Both [**Populate**](https://gitlab.com/meridianbet/meridianbet-stream-populate) and [**API**](https://gitlab.com/meridianbet/meridianbet-stream-api) stream applications use same mongo database. If you want you can use `docker-compose up` to run mongodb inside docker container on port `27017`. You can use `mongodb-express` application that runs inside docker container and runs on [**http://localhost:8081**](http://localhost:8081)



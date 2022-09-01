# Merchant Reporting RESTful API

## Package
The project uses the following Go packages:

* Routing: [echo](https://echo.labstack.com/)
* Database access: [database/sql](https://pkg.go.dev/database/sql")
* Database migration: [golang-migrate](https://github.com/golang-migrate/migrate)
* Data validation: [go-playground/validator](https://github.com/go-playground/validator)

## Project Layout

Merchant Reporting RESTful API uses the following project layout:
 
```
.
├── config
│   └── db
├── controller
│   ├── auth
│   └── transaction
├── database
│   ├── dump
│   ├── migrations
│   └── seeds
├── docs
│   └── postman_collection.json
├── model
├── schema
├── service
│   ├── auth
│   ├── merchant
│   ├── outlet
│   └── transaction
└── utils
```

## Getting Started

If this is your first time encountering Go, please follow [the instructions](https://golang.org/doc/install) to
install Go on your computer. The kit requires **Go 1.18 or above**.

[Docker](https://www.docker.com/get-started) is also needed if you want to try the kit without setting up your
own database server. The kit requires **Docker 17.05 or higher** for the multi-stage build support.

[migrator-tools](https://github.com/golang-migrate/migrate/releases) download from [golang-migrate](https://github.com/golang-migrate/migrate/releases) in release page

After installing Go, Docker, and Migrator tools run the following commands to start experiencing this starter kit:

```shell
# download the starter kit
git clone https://github.com/fianalfian/merchant-report.git

cd merchant-report

# start a merchant-reporting and MySQL database server in a Docker container
make deploy

# Usually you should run this command each time after you pull new code from the code repo. 
make migrate

# seed the database with some test data
make majoo-seed

```

At this time, you have a RESTful API server running at `http://localhost:8000`. It provides the following endpoints:

* `POST /api/login`: login to get token
* `GET /api/transactions/merchant/:id`: returns a list transaction merchant
* `GET /api/transactions/outlet/:id`: returns a list transaction Outlet

Try the URL `http://localhost:8000/` in a browser, and you should see something like `"OK v1.0.0"` displayed.

You can test Merchant Reporting RESTful API with preview file `docs/Merchant Reporting API.postman_collection.json` on postman
```
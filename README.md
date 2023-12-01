# chi-orders-api
This project is an example API service build in Go's Chi framework. It will serve as a learning opportunity for me to learn Chi and Go for building micro services.

The application will be backed by a local Postgres database, it will not be deployed. I may make a containerised version in the future.


## Running the application
To run the application you will need to have a Postgres database running locally. The application will create the required tables on startup.

```bash
createdb chi-orders-db
```

The server runs on port 3000 by default. It can be started with the following commands:

```bash
go run .

OR

env SERVER_PORT=<port> go run .
```

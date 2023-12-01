# chi-orders-api
This project is an example API service build in Go's Chi framework. It will serve as a learning opportunity for me to learn Chi and Go for building micro services.

The application will be backed by a local Postgres database, it will not be deployed.

I have included dockerfile and docker-compose files to be able to run the application in a container, but it is not required. Do *NOT* use this for deployment as it is not secure (although it doesn't actually work with important data).


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

## Running the application in a container

To run the application in a container, you will need to have docker and docker-compose installed. The application will handle the database on startup.

```bash
docker-compose build
docker-compose up -d
```

Shutdown the container with: `docker-compose down`

You can connect to the container through port 80 (default http port).


## Sending requests to the API

The API has the following methods:

- GET /orders
- POST /orders
- GET /orders/{id}
- PUT /orders/{id}
- DELETE /orders/{id}

The API will communicate via JSON and will return JSON responses. Here are some example requests you can make via `curl`:

In the scripts folder, there is a python script which will let you send some example data to the API. You can run it with the following command (Python 3.8+ required):

```bash
# docker
python scripts/make-example-orders.py

# local
env SERVER_PORT=<port> python scripts/make-example-orders.py
```

### GET /orders

```bash
# docker
curl -X GET http://localhost/orders

# local
curl -X GET http://localhost:<port>/orders
```


### GET /orders/{id}

```bash
# docker
curl -X GET http://localhost/orders/1

# local
curl -X GET http://localhost:<port>/orders/1
```

### PUT /orders/{id}

You can send PUT requests to update the status of an order. An order must be shipped before it can be marked as complete. The status can be one of the following:

- shipped
- completed

```bash
# docker
# update order 1 to shipped
curl -X PUT -d '{"status": "shipped"}' http://localhost/orders/1
# update order 1 to completed
curl -X PUT -d '{"status": "completed"}' http://localhost/orders/1

# local
# update order 1 to shipped
curl -X PUT -d '{"status": "shipped"}' http://localhost:<port>/orders/1
# update order 1 to completed
curl -X PUT -d '{"status": "completed"}' http://localhost:<port>/orders/1
```




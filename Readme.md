RE Pack Size Calculator
======================

Calculate the number of packs needed to fill a given order of x items.
It also includes a UI to manage the pack sizes and request the shipment calculation for an order.

The API endpoints are as follows:

- GET /api/packs - Get the pack sizes
- POST /api/packs - Set the pack sizes
- GET /api/order/fulfill - Get the shipment for an order
- GET / - Return the UI

Install
-------

Build the application with the following command:

```bash
make build
```

Usage
-----

Start the server to run the application with port 8085 and default pack sizes:

```bash
make start
```

Send a GET request to the following endpoint to calculate the shipment directly via the API:

```bash
curl -X GET "http://localhost:8085/api/order/fulfill?order=4"
```

Response:

```json
{
  "shipment": {
    "250": 1
  }
}
```

You can also access the UI by navigating to the following URL:

```bash
http://localhost:8085/
# or use port 8086 if starting from docker
http://localhost:8086/
```

Docker Container
----------------

The application can also be run in a docker container.
Use the following command to do so:

```bash
make docker-up
```

Please note that the docker-compose.yml file sets to port to 8086 instead of the default 8085.

When done destroy the container with:

```bash
make docker-down
```

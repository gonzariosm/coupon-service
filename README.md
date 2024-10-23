# Coupon Service

This project is a simple Coupon Service API built with Golang and Gin, designed to handle operations related to coupon creation, application, and retrieval. The service runs in a Docker container and is configured to be easily deployable using Docker Compose.

## Features

- **Create Coupons**: Create discount coupons with a specific code, discount, and minimum basket value.
- **Apply Coupons**: Apply a valid coupon to a basket if the basket value meets the coupon's minimum requirements.
- **Get Coupons**: Retrieve existing coupons by their codes.

## Requirements

- Docker
- Docker Compose

## How to Start the Application with Docker Compose

To build and run the application using Docker Compose, follow these steps:

1. Clone this repository:

```bash
git clone <repository_url>
cd coupon-service
```

2. Build and start the application using Docker Compose:

```bash
docker-compose up --build
```

This will:

- Build the application image.
- Start the `coupon-service` container and expose the API on port `8080`.

3. To stop the application, use:

```bash
docker-compose down
```

## Running the Application without Tests

In the `docker-compose.yml`, the application is configured to skip running tests by default. You can control this behavior using the `RUN_TESTS` build argument in the `Dockerfile`.

To ensure tests are skipped, the following argument is used:

```yaml
services:
  coupon-service:
    build:
      args:
        RUN_TESTS: "false"
```

To run the application with tests, set the `RUN_TESTS` argument to `"true"` when building the image.

## Example CURL Requests for Local Testing

Once the application is running, you can interact with the API using the following example curl requests.

### 1. Create a Coupon

This request creates a new coupon.

```bash
curl --location --request POST 'http://localhost:8080/api/create' \
--header 'Content-Type: application/json' \
--data-raw '{
    "discount": 15,
    "code": "DISCOUNT15",
    "min_basket_value": 50
}'
```

### 2. Apply a Coupon

This request applies a coupon to a basket.

```bash
curl --location --request POST 'http://localhost:8080/api/apply' \
--header 'Content-Type: application/json' \
--data-raw '{
    "code": "DISCOUNT15",
    "basket": {
        "value": 100
    }
}'
```

### 3. Get Coupons

This request retrieves details of one or more coupons by their codes.

```bash
curl --location --request GET 'http://localhost:8080/api/coupons' \
--header 'Content-Type: application/json' \
--data-raw '{
    "codes": ["DISCOUNT15", "DISCOUNT10"]
}'
```

## Environment Variables

The service supports the following environment variables:

- `API_PORT`: The port on which the API runs (default: 8080).
- `CORES`: Number of cores required to run (default: 1).

> `CORES` environment variable can trigger a Warning, but the application will run in any case.

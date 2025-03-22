# be-product-service
## Features

The API lets you manage products easily. Here’s what you can do:

- Add New Product

    Create a new product by sending all the necessary details (like category, supplier, unit, name, description, base price, and stock) to the `/products` endpoint.

    **Example**
    ```bash
    curl -X POST http://localhost:8080/products \
    -H "Content-Type: application/json" \
    -d '{
        "category_id": "00000000-0000-0000-0000-000000000001",
        "supplier_id": "00000000-0000-0000-0000-000000000011",
        "unit_id": "00000000-0000-0000-0000-000000000021",
        "name": "Kangkung Potong 1",
        "description": "Kangkung Potong segar",
        "base_price": 3000,
        "stock": 100
    }'
    ```

- List All Products

    Retrieve a list of all available products. You can also filter by adding query parameters.

    **Example**
    ```bash
    curl -X GET http://localhost:8080/products
    ```

- Get Product By ID

    Fetch detailed information about a specific product using its unique ID.

    **Example**
    ```bash
    curl -X GET http://localhost:8080/products/66100efd-e17c-470e-aa8c-5fba02949266
    ```

- Search Products by Name

    Find products by providing a search string for the product name.

    **Example**
    ```bash
    curl -X GET "http://localhost:8080/products?product_name=kangkung"
    ```

- Search & Filter Products

    Combine search and filter by product name and category type.

    **Example**
    ```bash
    curl -X GET "http://localhost:8080/products?product_name=kangkung&category_type=Sayuran"
    ```

- Sort Products

    Retrieve a sorted list of products by specifying a sort field and direction.

    **Example 1**
    ```bash
    curl -X GET "http://localhost:8080/products?sort=name&directive=asc"
    ```

    **Example 2**
    ```bash
    curl -X GET "http://localhost:8080/products?sort=base_price&directive=desc"
    ```

## Requirements

To run this project you need to have the following installed:

1. [Go](https://golang.org/doc/install) version 1.23.5
2. [GNU Make](https://www.gnu.org/software/make/) version 3.8
3. [Docker](https://docs.docker.com/get-docker/) version 20
4. [Docker Compose](https://docs.docker.com/compose/install/) version 2

## Initiate The Project

1. Clone the repository:
    ```bash
    git clone https://github.com/gunawanpras/be-product-service.git
    cd be-product-service
    ```

2. To build and start the project, execute
    ```
    make all
    ```

3. The service will start on port `8080`.

## Makefile
- all: Runs the full setup by cleaning up, initializing the project, migrating and seeding the database, and running both unit and e2e tests.
- clean: Stops and removes Docker containers and volumes for a fresh start.
- init: Prepares the project by tidying Go dependencies, building and starting Docker containers, and waiting for the database.
- db_migrate: Resets and applies database migrations to update the schema.
- db_seed: Resets and applies seed data to initialize the database with default values.
- unit_test: Executes unit tests for the product repository.
- e2e_test: Runs end-to-end tests to check overall system functionality.

## Documentations
[https://documenter.getpostman.com/view/819887/2sAYk8uhgG](https://documenter.getpostman.com/view/819887/2sAYk8uhgG)

## Performance and Testing

**Cache**: This project uses Redis to store frequently accessed data, reducing the load on the database and improving overall performance. Both SQL queries and responses are cached to minimize latency.

**Unit Tests**: The project includes a set of unit tests to ensure the handlers, services and repositories are functioning properly.

**End-to-End (E2E) Tests**: The project also includes E2E tests that simulate real-world scenarios.
# be-product-service

## Features
**Cache**: This project uses Redis to store frequently accessed data, reducing the load on the database and improving overall performance. Both SQL queries and responses are cached to minimize latency.

**Unit Tests**: The project includes a set of unit tests to ensure the handlers, services and repositories are functioning properly. This helps catch bugs and regressions early in the development cycle.

**End-to-End (E2E) Tests**: The project also includes E2E tests that simulate real-world scenarios, verifying that the entire application works seamlessly from start to finish.

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

## Documentations
[https://documenter.getpostman.com/view/819887/2sAYk8uhgG](https://documenter.getpostman.com/view/819887/2sAYk8uhgG)
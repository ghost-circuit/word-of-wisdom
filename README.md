# Word-of-wisdom Server

## Project Overview

This project contains a TCP server that implements Proof of Work (PoW) to protect against DDoS attacks, serving quotes from a “Word of Wisdom” collection. The project also includes a client that connects to the server, solves the PoW challenge, and receives quotes.

## Choosing a PoW Algorithm

We will use HashCash as the PoW algorithm. It is simple and widely used in scenarios where the server needs to limit client requests through computational effort.

This approach is computationally challenging for the client, making it harder to flood the server with requests, thus mitigating DDoS attacks.

## Simple and fast project running using docker-compose
   ```bash
   make docker-up
   ```

**Important: the `.env` file redeclare base env variables from `configs` directory.**

**Important: you don't need to configure `.env` file if you run application through `docker-compose`**

### Makefile Overview

The provided Makefile simplifies the process of building, running, testing, and linting the project, as well as building Docker images for both the server and the client.

### Prerequisites

- **Go**: Ensure that you have Go installed. You can install it from [here](https://golang.org/doc/install).
- **Docker**: Make sure Docker is installed to build and run the Docker containers. Install Docker from [here](https://docs.docker.com/get-docker/).
- `protoc`: The project uses Protocol Buffers for defining the gRPC service. Install the `protoc` compiler from [here](https://grpc.io/docs/protoc-installation/).
- **install tools**: 
  ```bash
  make install
  ```

### Available Make Commands

The Makefile supports various commands to manage your development workflow.

### Migrations
1) pre requirements:
    ```bash
    make install
    ```
2) create migration:
    ```bash
    make migration-create MIGRATION_NAME=<migration_name>
    ```
3) Open `.env` file and set environments.

### Linting and Testing

1. **Run Lint Checks**:
   Runs `golangci-lint` to check the Go code for style, bugs, and best practices.
   ```bash
   make lint
   ```

2. **Run Tests**:
   Runs all the tests in the project using `go test`.
   ```bash
   make test
   ```

### Docker Commands

1. **Start Docker Containers for Server**:
   ```bash
   make docker-up
   ```

2. **Stop Docker Containers for Client**:
   ```bash
   make docker-stop
   ```

3. **Remove Docker Containers for Client**:
   ```bash
   make docker-down
   ```

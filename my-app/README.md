# My App - API Backend and Client

## Project Structure

```bash
my-app
├── cmd
│   └── client
│       └── client.go      # API Client to test different HTTP methods and endpoints
├── compose.yaml           # Docker Compose file to run the backend
├── Dockerfile             # Dockerfile for building the Go application
├── go.mod                 # Go module dependencies
├── go.sum                 # Checksums for Go modules
├── README.md              # Project documentation (This file)
└── server.go              # Go application (API server)
```

## Introduction

This project demonstrates a simple API with various endpoints built using Go and the Gin framework. It consists of a backend server (`server.go`) that simulates operations on a set of devices and an API client (`cmd/client/client.go`) that can call these APIs using different HTTP methods.

The backend provides basic CRUD operations and simulates random latency and errors in certain operations. The client repeatedly makes requests to the backend API for testing purposes.

## Backend Server Overview

The backend server is built using the Gin framework and provides the following API endpoints:

- **GET /devices**: Fetches a list of devices.
- **POST /devices**: Creates a new device (dummy response).
- **PUT /devices/:id**: Upgrades a specific device.
- **DELETE /devices/:id**: Deletes a device (simulated failure).
- **POST /login**: Simulates a login request (always returns an internal server error).

### Simulated Device Data

The server initializes with two devices:
- **Device 1**: ID: 1, MAC: `E9-CF-45-FD-18-B3`
- **Device 2**: ID: 2, MAC: `CD-6A-9B-70-BF-EA`

### Randomized Delays and Errors

The backend introduces randomness to simulate real-world network latency and failures:
- **Normal Sleep**: Functions like `getDevices`, `createDevices`, and `upgradeDevice` introduce random delays using the `sleep()` function.
- **Error Simulation**: Certain endpoints (like `deleteDevice` and `login`) simulate errors and return failure responses.

### Server Endpoints

| **Method** | **Endpoint**             | **Description**                                  | **Example Response**                   |
|------------|--------------------------|--------------------------------------------------|----------------------------------------|
| `GET`      | `/devices`               | Fetch the list of devices.                       | `[{id:1, mac:"E9-CF-45-FD-18-B3"}, ...]` |
| `POST`     | `/devices`               | Create a new device.                             | `{"message": "Created!"}`              |
| `PUT`      | `/devices/:id`           | Upgrade a device with the given ID.              | `{"message": "Upgrade started..."}`    |
| `DELETE`   | `/devices/:id`           | Delete a device (always fails).                  | `{"message": "failed to delete."}`     |
| `POST`     | `/login`                 | Simulate a login (always returns error).         | `{"message": "Internal error!"}`       |

### Simulating API Calls

You can run the API client (`client.go`) to test these endpoints using various HTTP methods.

---

## API Client Overview

The API client, located in the `cmd/client/client.go` file, makes multiple API calls to different backend endpoints using various HTTP methods such as `GET`, `POST`, `PUT`, and `DELETE`.

The client continuously sends requests to the following endpoints:
- `http://panchanandevops.com/devices` (GET, POST, PUT)
- `http://panchanandevops.com/login` (POST)
- `http://panchanandevops.com/devices/:id` (PUT, DELETE)

The `req()` function handles the HTTP request generation and logs the response.

---

## How to Run

### Prerequisites

- Go (v1.20.5 or higher)
- Docker
- Docker Compose

### Running the Backend Locally

You can run the backend server using Docker Compose. The `compose.yaml` file is already set up to build the application and run it in a container.

1. **Clone the repository**:
   ```bash
   git clone https://github.com/panchanandevops/my-app
   cd my-app
   ```

2. **Build and run the backend server**:
   ```bash
   docker-compose up --build
   ```

   The backend server will be available on `http://localhost:8000`.

3. **Access the APIs**:  
   You can test the API endpoints using `curl` or any API client like Postman.

   Example:
   ```bash
   curl http://localhost:8000/devices
   ```

### Running the API Client

The API client is designed to simulate continuous API requests. You can run it with the following steps:

1. **Navigate to the client directory**:
   ```bash
   cd cmd/client
   ```

2. **Run the client**:
   ```bash
   go run client.go
   ```

The client will start sending repeated HTTP requests to the backend API.

---

## Dockerfile

The `Dockerfile` is a multi-stage build that compiles the Go application and uses a lightweight `distroless` image for production.

```dockerfile
FROM golang:1.20.5-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify

COPY server.go ./
RUN go build -o /my-app

FROM gcr.io/distroless/base-debian12
COPY --from=build /my-app /my-app

ENTRYPOINT ["/my-app"]
```

### Building and Running the Docker Image

1. **Build the Docker image**:
   ```bash
   docker build -t panchanandevops/sample-backend-go-app:v1.0.0 .
   ```

2. **Run the Docker container**:
   ```bash
   docker run -p 8000:8080 panchanandevops/sample-backend-go-app:v1.0.0
   ```

---

## Docker Compose

The `compose.yaml` file is set up to run the backend API in a Docker container. It maps port `8000` on your local machine to port `8080` in the container.

```yaml
services:
  backend:
    container_name: sample-backend-monitoring
    image: panchanandevops/sample-backend-go-app:v1.0.0
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - 8000:8080
```

To use Docker Compose, simply run:

```bash
docker-compose up --build
```

---



This project provides a basic API server with a variety of endpoints and an API client to simulate traffic. The backend is highly modular and can be extended with more endpoints as needed.

Feel free to clone this project, customize it, and experiment with Docker, Go, and API development!

# My App - Golang API and Argo CD Deployment



## Project Overview

The project includes a backend API built with Go, designed to handle basic CRUD operations on a set of "device" resources with simulated latency and error responses. The system is automated for building, testing, and deploying with GitHub Actions for CI/CD and Argo CD for declarative Kubernetes deployment.


## API Server Overview

The API backend is built using the [Gin framework](https://github.com/gin-gonic/gin) in Go and exposes the following endpoints for managing devices:

| **Method** | **Endpoint**             | **Description**                                |
|------------|--------------------------|------------------------------------------------|
| `GET`      | `/devices`               | Retrieve list of devices.                      |
| `POST`     | `/devices`               | Create a new device.                           |
| `PUT`      | `/devices/:id`           | Upgrade a specific device.                     |
| `DELETE`   | `/devices/:id`           | Attempt to delete a device (fails intentionally). |
| `POST`     | `/login`                 | Simulate a login request (always returns error).|

### Simulated Device Data and Behavior

- **Initial Devices**: Two devices are preloaded into the system:
  - Device 1 - ID: 1, MAC: `E9-CF-45-FD-18-B3`
  - Device 2 - ID: 2, MAC: `CD-6A-9B-70-BF-EA`
- **Randomized Behavior**: Delays and simulated failures are introduced to mimic network latency and possible API failures.

---

## Development and Deployment Workflows

### GitHub Actions - Build and Push Docker Image

The `build-push-docker-image.yaml` workflow automates the CI/CD pipeline:
- **Triggers**: Activates on Git tag pushes (e.g., `v1.0.0`) and directory changes.
- **Steps**:
  - **Build and Push**: Builds the Go API Docker image and pushes it to Docker Hub.
  - **Update Deployment YAML**: Modifies the `1-deployment.yaml` file with the new image tag.
  - **Pull Request**: Creates a PR with the updated image tag for review.

### Argo CD - GitOps Deployment

The `k8s-manifesto.yaml` file defines an Argo CD `Application` to deploy the backend API on a Kubernetes cluster:
- **Automated Sync**: Ensures that the Kubernetes cluster state aligns with the Git repository, handling updates, and cleaning up unused resources.
- **Configuration**:
  - **Source**: Uses the `Deploy/` folder from the repository.
  - **Sync Policy**: Automatic sync with self-healing and pruning enabled to match Git-defined state.

---

## Setup Instructions

### Prerequisites

- **Docker Hub Account**: For image storage, accessible via GitHub Secrets.
- **Kubernetes Cluster**: Required for deploying the backend API with Argo CD.
- **Argo CD**: Installed on the cluster, with `argocd` namespace.

### Configuration and Secrets

1. **GitHub Secrets**: Add the following secrets for GitHub Actions:
   - `DOCKER_USERNAME`: Your Docker Hub username.
   - `DOCKER_PASSWORD`: Your Docker Hub password.
   - `PAT_TOKEN`: GitHub PAT for creating pull requests.
2. **Directory Structure**: Ensure the repository follows the structure shown above for workflow consistency.

---

## Usage

1. **Build and Push the Docker Image**:
   - Push a new Git tag (e.g., `v1.0.0`) to trigger the GitHub Action:
     ```bash
     git tag v1.0.0
     git push origin v1.0.0
     ```
   - The workflow will build and push the image, update the Kubernetes YAML, and create a PR.

2. **Deploy with Argo CD**:
   - Clone the repository and apply the Argo CD configuration:
     ```bash
     git clone https://github.com/panchanandevops/Golang-ObservabilityStack.git
     kubectl apply -f path/to/argo-application.yaml
     ```

---



## Observability

This application’s observability is powered by **Grafana**, **Prometheus**, and **Loki** for monitoring and analysis.

- **Grafana**: Provides visualizations and dashboards for real-time monitoring of key metrics, with alerts for performance anomalies.
- **Prometheus**: Collects and stores time-series metrics from the application and infrastructure, enabling alerting for critical issues.
- **Loki**: Aggregates and retains application logs for efficient debugging and pattern analysis.


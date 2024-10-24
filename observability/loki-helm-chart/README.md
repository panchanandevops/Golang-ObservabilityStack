
# Loki Helm Chart

This repository contains the Helm chart for deploying Loki, a log aggregation system designed for efficiently storing and querying logs from Kubernetes.



## Charts

- **Loki:** The main chart for deploying Loki, either in single-binary mode or as a scalable distributed system.
- **Grafana Agent Operator:** Used for managing Grafana agents and collecting logs.
- **Minio:** A dependency for providing object storage when using Loki in certain environments.

## Templates

Templates are defined for various components of Loki including:

- `admin-api`
- `backend`
- `compactor`
- `distributor`
- `ingester`
- `gateway`
- `index-gateway`

Each component has Kubernetes resources like deployments, services, and HPA configurations.

## Values Files

There are different `values.yaml` files provided for various deployment scenarios:

- **distributed-values.yaml:** For deploying Loki in a distributed mode.
- **simple-scalable-values.yaml:** For scalable but simple deployments.
- **single-binary-values.yaml:** For deploying Loki as a single binary, useful in small environments.

## Usage

1. Clone the repository.
2. Choose the appropriate `values.yaml` file based on your needs (e.g., `distributed-values.yaml`).
3. Install the Helm chart using:

   ```bash
   helm install loki ./loki -f distributed-values.yaml
   ```

4. Monitor the deployment and ensure all components are running smoothly.

## Contributions

Contributions are welcome! Feel free to open issues or submit pull requests to improve this chart.


This `README.md` should give users an overview of the repository and how to use the Helm chart to deploy Loki.
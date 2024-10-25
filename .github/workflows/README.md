

# Build and Push Golang API Docker Image

This GitHub Action automates the process of building, tagging, and pushing a Docker image for your Golang API whenever a new tag is pushed or changes are made in the specified directory (`my-app/`). Additionally, it updates the deployment YAML file with the new image tag and creates a pull request with these updates.

## Workflow Overview

### Workflow Name:
`build-push-docker-image.yaml`

### Workflow Triggers:
- **On Tag Push:** The workflow triggers on pushing a Git tag that matches the pattern `v*.*.*` (e.g., `v1.0.0`, `v2.1.3`).
- **On Path Changes:** It also triggers when changes are made to files within the `my-app/` directory or the workflow file itself.

### Jobs:
This workflow consists of the following key job:

1. **build-backend:** Runs the following steps:
   - **Checkout Code:** Clones the repository to the GitHub runner.
   - **Set up Docker Buildx:** Prepares Docker Buildx for building multi-platform images.
   - **Log in to Docker Hub:** Authenticates to Docker Hub using the provided secrets.
   - **Extract Docker Metadata:** Extracts tags and labels for the Docker image based on the Git tag.
   - **Build and Push Docker Image:** Builds the Docker image for the Golang API from the provided context (`./my-app`) and pushes it to Docker Hub with cache optimization.
   - **Update Deployment YAML:** Updates the image tag in the Kubernetes deployment YAML file (`Deploy/1-deployment.yaml`) with the new Docker image tag.
   - **Commit and Push Changes:** Creates a new branch, commits the updated deployment YAML, and pushes the changes.
   - **Create Pull Request:** Automatically creates a pull request with the updated image tag for the deployment YAML.

---

## Setup Instructions

### 1. Prerequisites
- Ensure Docker Hub credentials are stored as GitHub Secrets:
  - `DOCKER_USERNAME`
  - `DOCKER_PASSWORD`
- A Personal Access Token (PAT) for creating pull requests is stored as `PAT_TOKEN`.
- The repository must contain a Kubernetes deployment YAML file at `Deploy/1-deployment.yaml` which includes the Docker image tag.

### 2. Directory Structure
Your repository should follow this structure:
```bash
.
├── .github/
│   └── workflows/
│       └── build-push-docker-image.yaml
├── my-app/
│   └── Dockerfile
└── Deploy/
    └── 1-deployment.yaml
```

### 3. Secrets Configuration
Navigate to your repository's settings and add the following secrets:
- **DOCKER_USERNAME:** Your Docker Hub username.
- **DOCKER_PASSWORD:** Your Docker Hub password.
- **PAT_TOKEN:** A personal access token to create a pull request with the changes.

---

## Detailed Workflow Steps

### 1. **Checkout Code**
The action checks out the repository to the GitHub runner to ensure all necessary files are available for building the Docker image.
```yaml
- name: Checkout code
  uses: actions/checkout@v4
```

### 2. **Set up Docker Buildx**
Docker Buildx is set up for building and pushing the Docker image with enhanced features like multi-platform support.
```yaml
- name: Set up Docker Buildx
  uses: docker/setup-buildx-action@v3
```

### 3. **Log in to Docker Hub**
Docker login credentials are passed from GitHub Secrets to authenticate and push the Docker image.
```yaml
- name: Log in to Docker Hub
  uses: docker/login-action@v3
  with:
    username: ${{ secrets.DOCKER_USERNAME }}
    password: ${{ secrets.DOCKER_PASSWORD }}
```

### 4. **Extract Docker Metadata**
The workflow dynamically extracts metadata for the Docker image such as tags and labels based on the Git tag.
```yaml
- name: Extract Docker metadata
  id: meta
  uses: docker/metadata-action@v5
  with:
    images: panchanandevops/sample-backend-go-app
    flavor: latest=false
    tags: type=raw,value={{tag}}
```

### 5. **Build and Push Golang API Docker Image**
This step builds the Docker image for the Golang API, pushes it to Docker Hub, and caches layers for faster subsequent builds.
```yaml
- name: Build and push golang API Docker image
  uses: docker/build-push-action@v6
  with:
    context: ./my-app
    file: ./my-app/Dockerfile
    push: true
    tags: ${{ steps.meta.outputs.tags }}
    labels: ${{ steps.meta.outputs.labels }}
    cache-from: type=registry,ref=panchanandevops/sample-backend-go-app:cache
    cache-to: type=registry,ref=panchanandevops/sample-backend-go-app:cache,mode=max
```

### 6. **Update Deployment YAML**
Once the Docker image is built and pushed, the Kubernetes deployment YAML file is updated with the new Docker image tag.
```yaml
- name: Update deployment image in YAML
  run: |
    sed -i "s|image: panchanandevops/sample-backend-go-app:.*|image: ${{ steps.meta.outputs.tags }}|g" Deploy/1-deployment.yaml
```

### 7. **Create Pull Request**
A pull request is automatically created to update the deployment with the new Docker image tag.
```yaml
- name: Create Pull Request
  uses: peter-evans/create-pull-request@v6
  with:
    token: ${{ secrets.PAT_TOKEN }}
    branch: update-deployment-image-${{ github.sha }}
    base: main
    title: "Update golang API image tag to ${{ steps.meta.outputs.tags }}"
    body: "This PR updates the golang API image tag to ${{ steps.meta.outputs.tags }}."
```

---

## Example Tag Push Trigger
To trigger this workflow, push a new Git tag in the format `v*.*.*` (e.g., `v1.0.0`):
```bash
git tag v1.0.0
git push origin v1.0.0
```

This will trigger the workflow, build the Docker image, push it to Docker Hub, update the Kubernetes deployment, and create a pull request.






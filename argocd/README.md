
# Argo CD Application Configuration

This k8s-manifesto.yaml file contains the configuration for deploying a Kubernetes application using Argo CD. Argo CD is a declarative GitOps continuous delivery tool for Kubernetes. This configuration file deploys resources from a GitHub repository and manages them within a Kubernetes cluster.

## Overview

The configuration defines an Argo CD `Application` resource to automate the deployment of resources within a Kubernetes cluster from a Git repository. 

## Configuration Details

### Application YAML Breakdown

```yaml
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: my-app
  namespace: argocd
  finalizers:
    - resources-finalizer.argocd.argoproj.io
spec:
  project: default
  source:
    repoURL: https://github.com/panchanandevops/Golang-ObservabilityStack.git
    targetRevision: HEAD
    path: Deploy/
  destination:
    server: https://kubernetes.default.svc
  syncPolicy:
    automated:
      prune: true
      selfHeal: true
      allowEmpty: false
    syncOptions:
      - Validate=true
      - CreateNamespace=false
      - PrunePropagationPolicy=foreground
      - PruneLast=true
```

### Key Sections Explained

- **Metadata**
  - `name`: Name of the Argo CD Application (`my-app`).
  - `namespace`: Namespace for Argo CD resources (`argocd`).
  - `finalizers`: Ensures resource cleanup upon deletion.

- **spec**
  - `project`: Sets the project context, using the default Argo CD project.
  
#### Source Section

- **repoURL**: Points to the Git repository hosting the application manifests.
- **targetRevision**: Specifies the branch or tag (`HEAD` points to the latest commit on the default branch).
- **path**: Directory within the repository where deployment manifests are stored (`Deploy/`).

#### Destination Section

- **server**: The Kubernetes API server address to apply the configurations to. Here, `https://kubernetes.default.svc` uses the in-cluster Kubernetes API.

#### Sync Policy

Defines the rules for syncing resources from Git to the Kubernetes cluster:

- **automated**:
  - `prune`: Deletes resources from the cluster that are not defined in Git.
  - `selfHeal`: Automatically repairs resources to match the Git state if changes occur in the cluster.
  - `allowEmpty`: Disallows syncing empty applications.
  
- **syncOptions**:
  - `Validate=true`: Validates Kubernetes manifests during sync.
  - `CreateNamespace=false`: Disables automatic namespace creation during sync.
  - `PrunePropagationPolicy=foreground`: Sets the pruning policy to remove resources with dependencies first.
  - `PruneLast=true`: Performs pruning at the end of sync to ensure other operations complete.

## Usage

To use this configuration:

1. **Clone the Repository**  
   Clone the repository where this YAML file is stored.

   ```bash
   git clone https://github.com/panchanandevops/Golang-ObservabilityStack.git
   cd Golang-ObservabilityStack
   ```

2. **Apply the Application Configuration**  
   Apply this configuration to the Argo CD namespace by running:

   ```bash
   kubectl apply -f path/to/argo-application.yaml
   ```

3. **Monitor Application Sync**  
   Use the Argo CD web interface or CLI to monitor and manage the application's sync status.

## Notes

- This configuration assumes that the `argocd` namespace is present in the cluster where Argo CD is installed.
- Modify `repoURL`, `path`, and other parameters if your repository structure or URL changes.

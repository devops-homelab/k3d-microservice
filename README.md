# k3d-microservice

A tiny Go HTTP service with no external dependencies.

## Endpoints

- `GET /` returns `{"message":"Hello from k3d-microservice!"}`.
- `GET /health` returns `{"status":"ok"}`.

The service listens on port `8080`.

## Run locally

Requires Go 1.23 or newer.

```sh
go run .
```

In another terminal:

```sh
curl http://localhost:8080/
curl http://localhost:8080/health
```

Run tests with `go test ./...`.

## Run with Docker

```sh
docker build -t k3d-microservice .
docker run --rm -p 8080:8080 k3d-microservice
```

The Docker image uses a multi-stage build and runs as a non-root user.

## Publish to Docker Hub using GitHub Actions

1. Create a Docker Hub repository named `k3d-microservice` under your username.
2. Create a Docker Hub access token with **Read & Write** permission.
3. In your GitHub repository, go to **Settings → Secrets and variables → Actions**
   and add these repository secrets:

   | Secret | Value |
   | --- | --- |
   | `DOCKERHUB_USERNAME` | Your Docker Hub username |
   | `DOCKERHUB_TOKEN` | Your Docker Hub access token (not your password) |

4. Push this project to the `main` branch of your GitHub repository.
   The workflow tests the service, builds the image, and pushes this tag:
   - `YOUR_USERNAME/k3d-microservice:<full-commit-sha>`

You can also run it manually from **Actions → Build and push to Docker Hub →
Run workflow**. To use another Docker Hub repository name, change
`k3d-microservice` in the workflow's `tags` entry.

The workflow uses the self-hosted `k3d-gha-runners` scale set. Enable
`containerMode.type: dind` in its ARC Helm values so Docker Buildx can connect
to a Docker daemon. After changing the values, sync the runner scale set in
Argo CD and wait for replacement runner pods before rerunning the workflow.

Buildx uses `driver-opts: network=host` so its BuildKit container shares the
runner pod's network instead of using Docker's nested bridge. This avoids
registry TLS timeouts caused by the nested bridge's larger MTU in k3d.
Here, "host" means the Docker-in-Docker pod network, not the Kubernetes node.

Run the published image:

```sh
docker run --rm -p 8080:8080 YOUR_USERNAME/k3d-microservice:COMMIT_SHA
```

Replace `COMMIT_SHA` with the full commit SHA from the successful workflow run.
The image targets the runner's native architecture. When running it on a different
architecture, use Docker's `--platform` option with appropriate emulation support.
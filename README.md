# Tharun Kumaran - Go Web Application

A Go-based web application deployed using a complete **DevOps and GitOps workflow** with Docker, Kubernetes, Helm, GitHub Actions, Argo CD, and AWS EKS.

The project demonstrates how a simple Go web application can be containerized, tested, packaged, deployed, and continuously delivered to Kubernetes using modern DevOps practices.

---

## 🚀 Project Overview

This project implements an end-to-end CI/CD and GitOps pipeline:

```text
Developer
    │
    │ git push
    ▼
 GitHub
    │
    ▼
GitHub Actions
    │
    ├── Run Go Tests
    ├── Build Docker Image
    ├── Push Image to Docker Hub
    └── Update Helm Image Tag
    │
    ▼
 Git Repository
    │
    ▼
  Argo CD
    │
    ▼
   Helm
    │
    ▼
 AWS EKS
    │
    ├── Deployment
    │      └── Go Application Pod
    │
    ├── Service
    │
    └── Ingress
            │
            ▼
     AWS Application
     Load Balancer
            │
            ▼
        Internet
```

---

## 🛠️ Technologies Used

| Technology                        | Purpose                          |
| --------------------------------- | -------------------------------- |
| **Go**                            | Application development          |
| **Go Testing**                    | Application testing              |
| **Docker**                        | Containerization                 |
| **Docker Hub**                    | Container image registry         |
| **Kubernetes**                    | Container orchestration          |
| **Helm**                          | Kubernetes application packaging |
| **GitHub Actions**                | CI automation                    |
| **Argo CD**                       | GitOps-based continuous delivery |
| **AWS EKS**                       | Managed Kubernetes cluster       |
| **AWS Load Balancer Controller**  | Creates and manages AWS ALB      |
| **AWS Application Load Balancer** | External application access      |

---

## 📁 Project Structure

```text
tharun-go-web-app/
│
├── .github/
│   └── workflows/
│       └── cicd.yaml
│
├── helm/
│   └── go-web-app-chart/
│       ├── Chart.yaml
│       ├── values.yaml
│       └── templates/
│           ├── deployment.yaml
│           ├── service.yaml
│           └── ingress.yaml
│
├── k8s/
│   └── manifests/
│       ├── deployment.yaml
│       ├── service.yaml
│       └── ingress.yaml
│
├── static/
│   └── CSS and static assets
│
├── Dockerfile
├── go.mod
├── main.go
├── main_test.go
├── README.md
└── README-DevOps.md
```

---

## 🌐 Application

The application is written in Go and runs on port **8080**.

### Application Routes

| Route       | Description           |
| ----------- | --------------------- |
| `/`         | Home page             |
| `/about`    | About page            |
| `/projects` | Projects page         |
| `/contact`  | Contact page          |
| `/health`   | Health check endpoint |

The `/health` endpoint can be used to verify that the application is running correctly.

---

# 🐳 Docker

The application is containerized using a multi-stage Docker build.

### Dockerfile

The first stage uses Go 1.22 Alpine to build the application, while the final image uses Alpine Linux to keep the runtime image smaller.

```dockerfile
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY main.go main_test.go ./

RUN go build -o web-app main.go

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/web-app .
COPY static ./static

EXPOSE 8080

CMD ["./web-app"]
```

### Build the Image

```bash
docker build -t tharunm11/tharun-go-web-app:v1 .
```

### Run Locally

```bash
docker run -p 8080:8080 tharunm11/tharun-go-web-app:v1
```

Application:

```text
http://localhost:8080
```

---

# ☸️ Kubernetes

The project contains Kubernetes manifests for:

* Deployment
* Service
* Ingress

## Deployment

The Deployment manages the Go application Pod. The application container listens on port `8080`.

```yaml
kind: Deployment
```

The current configuration uses:

```yaml
replicas: 1
```

## Service

A Kubernetes `ClusterIP` Service exposes the application internally.

```text
Service Port: 80
Target Port: 8080
```

This allows Kubernetes traffic to reach the Go application running inside the Pod.

## Ingress

The application uses Kubernetes Ingress with the AWS Load Balancer Controller.

The Ingress is configured as an **internet-facing AWS Application Load Balancer** using IP targets.

Traffic flow:

```text
Internet
   ↓
AWS Application Load Balancer
   ↓
Kubernetes Ingress
   ↓
ClusterIP Service
   ↓
Go Application Pod
```

---

# 📦 Helm

The Kubernetes deployment is also packaged using Helm.

```text
helm/
└── go-web-app-chart/
    ├── Chart.yaml
    ├── values.yaml
    └── templates/
        ├── deployment.yaml
        ├── service.yaml
        └── ingress.yaml
```

The Helm chart contains configurable values such as:

```yaml
replicaCount: 1

image:
  repository: tharunm11/tharun-go-web-app
  pullPolicy: IfNotPresent
  tag: "<git-commit-sha>"
```

The image tag is updated automatically by the CI pipeline.

### Helm Installation

```bash
helm install tharun-go-web-app ./helm/go-web-app-chart
```

### Upgrade

```bash
helm upgrade tharun-go-web-app ./helm/go-web-app-chart
```

---

# 🔄 CI/CD Pipeline

GitHub Actions is used for Continuous Integration.

The pipeline runs when changes are pushed to or pull requests are created against the `main` branch.

### CI Flow

```text
Git Push
   ↓
GitHub Actions
   ↓
Run Go Tests
   ↓
Build Docker Image
   ↓
Push Image to Docker Hub
   ↓
Update Helm Image Tag
   ↓
Git Repository
```

### Docker Image Tagging

Docker images are tagged using the Git commit SHA.

Example:

```text
tharunm11/tharun-go-web-app:<commit-sha>
```

Using the commit SHA provides an immutable reference to the exact application version.

---

# 🔁 GitOps with Argo CD

Argo CD is used for Continuous Delivery.

Instead of GitHub Actions directly deploying to Kubernetes, the desired Kubernetes state is stored in Git.

```text
GitHub
   │
   │ Desired State
   ▼
Argo CD
   │
   │ Sync
   ▼
Helm
   │
   ▼
AWS EKS
```

Argo CD continuously monitors the Git repository and synchronizes the Kubernetes environment with the desired state defined in Git.

### Benefits

* Git acts as the source of truth
* Declarative deployments
* Automatic synchronization
* Easier rollback
* Improved deployment visibility
* Separation of CI and CD responsibilities

---

# ☁️ AWS Deployment

The application runs on **Amazon EKS**.

The deployment architecture consists of:

```text
AWS EKS
│
├── Worker Node
│    └── Go Application Pod
│
├── Worker Node
│
├── Kubernetes Service
│
└── Ingress
      │
      ▼
AWS Load Balancer Controller
      │
      ▼
AWS Application Load Balancer
```

The AWS Load Balancer Controller watches the Kubernetes Ingress resource and manages the corresponding AWS Application Load Balancer.

---

# 🔐 Security

Sensitive credentials are not stored directly inside the GitHub Actions workflow.

Docker Hub authentication is handled using GitHub repository secrets.

Example secrets:

```text
DOCKERHUB_USERNAME
DOCKERHUB_TOKEN
```

The workflow accesses them through GitHub Actions secrets.

---

# 🧪 Testing

The Go application contains automated tests in:

```text
main_test.go
```

Tests can be executed locally using:

```bash
go test ./...
```

The same tests are executed as part of the GitHub Actions pipeline before the Docker image is built.

---

# 🚀 Deployment Workflow

A typical application update follows this process:

### 1. Make a code change

```bash
git add .
git commit -m "Update application"
```

### 2. Push to GitHub

```bash
git push origin main
```

### 3. GitHub Actions runs

The pipeline:

```text
Run Tests
    ↓
Build Docker Image
    ↓
Push Docker Image
    ↓
Update Helm Image Tag
```

### 4. Argo CD detects the Git change

Argo CD detects the updated Helm configuration.

### 5. Argo CD synchronizes Kubernetes

```text
Git
 ↓
Argo CD
 ↓
Helm
 ↓
Kubernetes
```

### 6. Kubernetes starts the new version

The Deployment updates the application Pod with the new Docker image.

### 7. Application is available

```text
Internet
   ↓
AWS ALB
   ↓
Ingress
   ↓
Service
   ↓
Go Application
```

---

# 📊 DevOps Features

This project demonstrates:

* ✅ Go application development
* ✅ Automated Go testing
* ✅ Docker containerization
* ✅ Multi-stage Docker builds
* ✅ Docker Hub image publishing
* ✅ Kubernetes Deployment
* ✅ Kubernetes Service
* ✅ Kubernetes Ingress
* ✅ Helm packaging
* ✅ GitHub Actions CI
* ✅ Git SHA-based Docker image tagging
* ✅ GitOps workflow
* ✅ Argo CD continuous delivery
* ✅ AWS EKS deployment
* ✅ AWS Application Load Balancer
* ✅ AWS Load Balancer Controller
* ✅ Kubernetes-based application deployment

---

# 🎯 Project Objective

The main objective of this project is to demonstrate an end-to-end DevOps workflow for a Go web application.

The project takes the application from:

```text
Source Code
    ↓
Testing
    ↓
Docker Image
    ↓
Container Registry
    ↓
Kubernetes
    ↓
Helm
    ↓
GitOps
    ↓
AWS EKS
    ↓
AWS ALB
    ↓
Live Application
```

This demonstrates how development, containerization, CI/CD, Kubernetes, cloud infrastructure, and GitOps can be integrated into a single deployment workflow.

---

## 👨‍💻 Author

**Tharun Kumaran**

DevOps & Cloud Enthusiast

GitHub:
https://github.com/tharunkumaran11

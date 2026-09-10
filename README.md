# Tharun Kumaran - Go Web Application

A Go-based web application implemented with a complete **DevOps and GitOps workflow** using Docker, Kubernetes, Helm, GitHub Actions, Argo CD, and AWS EKS.

The project demonstrates how a simple Go application can be tested, containerized, pushed to a container registry, and continuously deployed to Kubernetes using GitOps practices.

---

## 🛠️ Technologies Used

- **Go** – Application development
- **Docker** – Containerization
- **Docker Hub** – Container registry
- **Kubernetes** – Container orchestration
- **Helm** – Kubernetes package management
- **GitHub Actions** – CI pipeline
- **Argo CD** – GitOps-based CD
- **AWS EKS** – Kubernetes platform
- **AWS Load Balancer Controller** – AWS ALB integration

---

## 📁 Project Structure

```text
tharun-go-web-app/
│
├── .github/workflows/
│   └── cicd.yaml
│
├── helm/go-web-app-chart/
│   ├── Chart.yaml
│   ├── values.yaml
│   └── templates/
│       ├── deployment.yaml
│       ├── service.yaml
│       └── ingress.yaml
│
├── k8s/manifests/
│   ├── deployment.yaml
│   ├── service.yaml
│   └── ingress.yaml
│
├── static/
├── Dockerfile
├── go.mod
├── main.go
└── main_test.go

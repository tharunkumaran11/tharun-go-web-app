# Tharun Kumaran - Go Web Application

This is a Go-based web application that I used to build and implement an end-to-end DevOps and GitOps workflow.

The project covers application development, testing, Docker containerization, Kubernetes deployment, Helm, GitHub Actions, Argo CD, and AWS EKS.

## Project Overview

The main idea of this project was to take a Go web application and implement the complete deployment process using DevOps tools and practices.

The application is first tested using Go testing. GitHub Actions is then used to build the Docker image and push it to Docker Hub. The Helm configuration is updated with the new Docker image tag.

Argo CD monitors the GitHub repository and deploys the updated application to Kubernetes running on AWS EKS.

For external access, I used the AWS Load Balancer Controller to create and manage an AWS Application Load Balancer.

The overall workflow is:

Developer → GitHub → GitHub Actions → Docker Hub → Helm → Argo CD → AWS EKS → AWS ALB → Application

## Technologies Used

- Go
- Docker
- Docker Hub
- Kubernetes
- Helm
- GitHub Actions
- Argo CD
- AWS EKS
- AWS Load Balancer Controller
- AWS Application Load Balancer

## Application

The application is written in Go and runs on port 8080.

It contains the following pages:

- Home
- About
- Projects
- Contact

It also has a `/health` endpoint for checking the application status.

## Docker

I containerized the Go application using a multi-stage Dockerfile.

The Docker image is built and pushed to Docker Hub. The image is tagged using the Git commit SHA so that every application version has a unique image tag.

## Kubernetes

The application is deployed to Kubernetes using a Deployment, Service and Ingress.

The Deployment manages the application Pods.

The Service is configured as a ClusterIP service and connects the Kubernetes network to the application running on port 8080.

The Ingress is used to expose the application through the AWS Application Load Balancer.

The traffic flow is:

Internet → AWS ALB → Ingress → Service → Application Pod

## Helm

I used Helm to package and manage the Kubernetes deployment.

The Helm chart contains the Kubernetes configuration and allows values such as the Docker image, image tag and replica count to be configured easily.

The Docker image tag is updated automatically by the CI pipeline.

## CI/CD

GitHub Actions is used for the CI part of the project.

Whenever changes are pushed to the main branch, the workflow:

1. Runs the Go tests
2. Builds the Docker image
3. Pushes the image to Docker Hub
4. Updates the image tag in the Helm values
5. Pushes the updated Helm configuration back to GitHub

## GitOps with Argo CD

Argo CD is used for the Continuous Delivery part.

Instead of GitHub Actions directly deploying to Kubernetes, the deployment configuration is stored in Git.

Argo CD monitors the repository and synchronizes the changes to the Kubernetes cluster.

This follows the GitOps approach where Git is used as the source of truth for the deployment configuration.

## AWS Deployment

The application is deployed on an AWS EKS cluster.

The AWS Load Balancer Controller is used with Kubernetes Ingress to provision an internet-facing AWS Application Load Balancer.

The final application flow is:

Internet → AWS ALB → Kubernetes Ingress → Service → Go Application

## Testing

The project contains Go tests which can be run using:

`go test ./...`

The same tests are also executed automatically through GitHub Actions before the Docker image is built.

## Project Objective

The objective of this project was to understand and implement a complete DevOps workflow for a Go application.

It covers:

- Application development
- Testing
- Docker containerization
- Docker image management
- Kubernetes deployment
- Helm
- CI using GitHub Actions
- GitOps using Argo CD
- Deployment on AWS EKS
- AWS Application Load Balancer

## Author

Tharun Kumaran

GitHub: https://github.com/tharunkumaran11/tharun-go-web-app

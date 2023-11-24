# Sample Microservice Application

![Kubernetes](https://img.shields.io/badge/kubernetes-%23326ce5.svg?style=for-the-badge&logo=kubernetes&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)


![TraefikLabs](https://img.shields.io/badge/Traefik%20Labs-yellow)



A basic restful microservice API

## Services

| Service         | Port | Purpose                                 |
| --------------- | ---- | --------------------------------------- |
| Product Service | 8000 | Manages Product                         |
| Reviews Service | 8090 | Manages Reviews                         |
| Ratings Service | 8080 | Manages Ratings                         |
| Load Generator  | n/a  | Generates load / unlimited http request |


## Running on Docker

![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)

We can run it thru `docker-compose` via:

```bash
# Build it first,
docker-compose build

# Then, run it:
docker-compose up
```


## Running on Kubernetes

To run it on kubernetes, you must have the following prerequisites before you can proceed.

### Pre-requisites

- Helm
- Kubernetes Cluster or Minikube
- kubectl


### Traefik Setup (Optional)

First, We need an edge proxy that will proxied our http request going to our microservice. 


```bash
# Add Traefik Helm Repository

helm repo add traefik https://helm.traefik.io/traefik
helm repo update


# Dry Run
helm upgrade --install traefik traefik/traefik \
    -n traefik \
    --create-namespace \
    --version 25.0.0  \
    --set additionalArguments='{"--log.level=DEBUG", "--accesslog=true",  "--accesslog.format=json"}' \
    --set ingressRoute.dashboard.enabled=false --dry-run

# Deployment
helm upgrade --install traefik traefik/traefik \
    -n traefik \
    --create-namespace \
    --version 25.0.0  \
    --set additionalArguments='{"--log.level=DEBUG", "--accesslog=true",  "--accesslog.format=json"}' \
    --set ingressRoute.dashboard.enabled=false
```

### MySQL Database Server Setup

| Version |
|---------|
| 8.0.35  |


Our microservices requires MySQL as it's backend database. To deploy mysql server:

```bash
# Add bitnami helm chart repository
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update

# Dry-Run
helm upgrade --install mysql-server bitnami/mysql \
    -n mysql-server \
    --create-namespace \
    --version 9.14.4 \
    --set auth.username="admin" \
    --set auth.password="admin" \
    --set auth.rootpassword="admin" \
    --set auth.database="microservice" \
    --set primary.persistence.enabled="true" \
    --dry-run

## Deployment 

helm upgrade --install mysql-server bitnami/mysql \
    -n mysql-server \
    --create-namespace \
    --version 9.14.4 \
    --set auth.username="admin" \
    --set auth.password="admin" \
    --set auth.rootpassword="admin" \
    --set auth.database="microservice" \
    --set primary.persistence.enabled="true"
```

### Microservice Deployment

Now, After deploying all our dependencies. We can start deploying our actual workload. There's 2 options available here.

- Normal Deployment (Without Traefik)
- Deployment with Traefik Proxy

### Deployment with Traefik Proxy

```bash
kubectl apply -f k8s/deployment-with-traefik.yaml
```

### Deployment without Traefik Proxy

```bash
kubectl apply -f k8s/deployment.yaml
```
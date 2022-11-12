# Food-detail-integrator-be

## Description
This little microservice is used to integrate food details from [openFoodFacts](https://world.openfoodfacts.org/). It is used by [food-track-be](https://github.com/nico-iaco/food-track-be) and [grocery-be](https://github.com/nico-iaco/grocery-be).

## Features

- [x] Integrate food details from openFoodFacts
- [x] Calculate kcal per quantity consumed

## Requirements

- [Redis](https://redis.io/) (optional)

## Installation

### Cluster installation

To install this app in a cluster, first create grocery namespace, then modify the kustomization.yaml file in /k8s/overlays/qa
changing the property to match your configuration and run the following command:

```bash
kubectl apply -k k8s/overlays/qa
```

### Local installation

You can run this app locally with docker. To do so, run the following command:

```bash
docker run -p 8080:8080 ghcr.io/nico-iaco/food-detail-integrator-be:latest -e {ALL_ENV_VARIABLES}
```

## Environment variables

| Name          | Description                                       | Default value |
|---------------|---------------------------------------------------|---------------|
| PORT          | Port on which the app will listen                 | 8080          |
| GIN_MODE      | Release type of app                               | debug         |
| REDIS_URL     | Redis host                                        |               |
| REDIS_ENABLED | Flag to enable redis cache                        | false         |
| IS_SANDBOX    | Flag for choosing open food facts api environment | false         |


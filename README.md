# Irmag Innovations Repository

This repository contains various services and tools developed by Irmag Innovations. It is structured into three primary branches, each serving a specific purpose within the project.

## Branches Overview

### 1. `master`
- **Contents**: Docker configuration files, Makefile.
- **Purpose**: This branch serves as the central configuration hub, containing Docker-related setup and build automation scripts.
- **Status**: Active.

### 2. `golang_client`
- **Contents**: Web page parser written in Go, utilizing goroutines for concurrent operations.
- **Purpose**: This branch is responsible for the Go-based web scraper client. It processes and parses data from web pages, leveraging goroutine to optimize performance.
- **Database**: MongoDB.
- **Kafka Integration**: This service communicates with the `product_service` via Kafka.

### 3. `java_client`
- **Contents**: Four microservices - `apigateway`, `auth_service`, `eureka_server`, `product_service`.
- **Purpose**: This branch handles the backend services for the project.
  - `apigateway`: Manages API requests.
  - `auth_service`: Handles authentication using PostgreSQL as the database.
  - `eureka_server`: Service registry for the microservices architecture.
  - `product_service`: Manages product data and interacts with MongoDB.
- **Kafka Integration**: `product_service` communicates with `golang_client` via Kafka.

## Docker and Docker Compose
All services across the branches are containerized using Docker and are orchestrated with Docker Compose. This ensures seamless deployment and scalability.


# ☁️ Cloud Assessment Tool

A production-style Go backend for managing cloud resources — VM lifecycle, user access, cost recommendations, audit logging, caching, event processing, and typed internal service communication.

Built to demonstrate real-world backend engineering patterns: layered architecture, authentication/authorization, caching strategy, asynchronous processing, gRPC contracts, and observability — not just CRUD.

![Go](https://img.shields.io/badge/Go-1.2x-00ADD8?logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-Compose-2496ED?logo=docker&logoColor=white)
![MySQL](https://img.shields.io/badge/MySQL-8-4479A1?logo=mysql&logoColor=white)
![Redis](https://img.shields.io/badge/Redis-7-DC382D?logo=redis&logoColor=white)
![gRPC](https://img.shields.io/badge/gRPC-enabled-244c5a?logo=grpc&logoColor=white)
![License](https://img.shields.io/badge/license-MIT-green)

---

## 📌 Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Architecture](#architecture)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [API Reference](#api-reference)
- [Caching Strategy](#caching-strategy)
- [Event Processing (Kafka)](#event-processing-kafka)
- [gRPC & Protocol Buffers](#grpc--protocol-buffers)
- [Testing](#testing)
- [Observability](#observability)
- [Production Readiness](#production-readiness)
- [Author](#author)

---

## Overview

Cloud environments need consistent management of resource lifecycle, cost, access, and operational state. This project models that problem end-to-end with a Go backend that handles:

- VM provisioning, start/stop/terminate lifecycle
- Role-based access control and JWT authentication
- Cost-optimization recommendations
- Audit logging for compliance
- Redis-backed caching with invalidation on state change
- Kafka-based async event processing
- gRPC for typed, low-latency internal service communication
- Prometheus + Grafana for metrics and health monitoring

## Features

| Area | Capability |
|---|---|
| **Auth** | JWT-based authentication, role-based authorization (RBAC) |
| **VM Management** | List, start, and terminate VMs; get cost-optimization recommendations |
| **Admin** | Enable/disable user accounts, list all users |
| **Cloud Sync** | Trigger synchronization of cloud resource state |
| **Audit** | Queryable audit log for administrative and lifecycle actions |
| **Caching** | Redis cache-aside pattern with automatic invalidation |
| **Async Events** | Kafka producer/consumer for decoupled event processing |
| **Internal RPC** | gRPC services with Protocol Buffer contracts |
| **Docs** | Auto-generated Swagger/OpenAPI UI |
| **Monitoring** | `/health` and `/metrics` endpoints, Prometheus + Grafana dashboards |

## Tech Stack

| Category | Technology |
|---|---|
| Language | Go |
| HTTP Framework | Gin |
| ORM | GORM |
| Database | MySQL 8 |
| Cache | Redis 7 |
| Authentication | JWT |
| RPC | gRPC / Protocol Buffers |
| Messaging | Apache Kafka |
| API Docs | Swagger / OpenAPI (Swaggo) |
| Containerization | Docker, Docker Compose |
| Orchestration | Kubernetes (Minikube) |
| Monitoring | Prometheus + Grafana |

## Architecture

```text
Client
  │
  ▼
Gin Router → JWT Middleware → RBAC → Handler
  │
  ▼
Service Layer ──► Redis (cache) ──► Kafka (events)
  │
  ▼
Repository → GORM → MySQL

Internal service-to-service calls use gRPC + Protocol Buffers
instead of REST, for typed contracts and lower latency.
```

## Project Structure

```text
cloud-assessment-tool/
├── cmd/                # application entrypoint
├── config/             # environment & app configuration
├── handlers/           # HTTP request handlers
├── middleware/         # JWT, RBAC, request processing
├── models/             # GORM models
├── repository/         # data access layer
├── services/           # business logic
├── routes/             # route definitions
├── grpc/
│   ├── proto/          # .proto contracts
│   ├── generated/      # generated Go stubs
│   └── services/       # gRPC service implementations
├── docs/                # Swagger/OpenAPI output
├── tests/
├── Dockerfile
├── docker-compose.yml
├── docker-compose.kafka.yml
└── docker-compose.monitoring.yml
```

## Getting Started

### Prerequisites
- Go 1.2x+
- Docker & Docker Compose

### Run locally

```bash
git clone https://github.com/<your-username>/cloud-assessment-tool.git
cd cloud-assessment-tool
docker compose up --build
```

### Verify it's running

| Resource | URL |
|---|---|
| API | http://localhost:8080 |
| Swagger UI | http://localhost:8080/swagger/index.html |
| Health check | http://localhost:8080/health |
| Metrics | http://localhost:8080/metrics |

## API Reference

| Method | Endpoint | Purpose | Auth |
|---|---|---|---|
| POST | `/register` | Register user | Public |
| POST | `/login` | Login, receive JWT | Public |
| GET | `/api/profile` | Current user profile | JWT |
| GET | `/api/vms` | List VMs | JWT |
| PUT | `/api/vms/{id}/start` | Start VM | JWT |
| PUT | `/api/vms/{id}/terminate` | Terminate VM | JWT |
| GET | `/api/vms/recommendations` | Cost recommendations | JWT |
| POST | `/api/cloud/sync` | Sync cloud resources | JWT |
| GET | `/api/audit-logs` | View audit logs | JWT + Role |
| GET | `/api/admin/users` | List users | JWT + Role |
| PUT | `/api/admin/users/{id}/disable` | Disable user | JWT + Role |
| PUT | `/api/admin/users/{id}/enable` | Enable user | JWT + Role |
| GET | `/health` | Health check | Public |
| GET | `/metrics` | Prometheus metrics | Public |

## Caching Strategy

Redis implements a cache-aside pattern for VM queries:

```text
GET /api/vms → check Redis → HIT: return cached data
                            → MISS: query MySQL → populate Redis → return
```

VM lifecycle operations (start/terminate) invalidate the relevant cache keys to prevent stale reads.

## Event Processing (Kafka)

Kafka decouples state-changing operations from downstream side effects (audit logging, cloud sync, notifications):

```text
Service → Kafka Producer → Topic → Kafka Consumer → Event Handler
```

## gRPC & Protocol Buffers

Internal services communicate over gRPC using `.proto`-defined contracts, giving compile-time type safety and generated client/server code instead of hand-rolled JSON serialization.

```protobuf
service VMService {
    rpc GetVM(GetVMRequest) returns (VMResponse);
}
```

Supports unary, server-streaming, client-streaming, and bidirectional streaming patterns.

## Testing

- **Unit tests** — services, business logic, repositories, validation
- **Handler tests** — status codes, JSON responses, auth/authz behavior
- **Integration tests** — full request path through MySQL and Redis
- **JWT test matrix** — valid, missing, malformed, expired, invalid signature, insufficient role

## Observability

```text
App → /health, /metrics → Prometheus → Grafana
```

Tracks API traffic, latency, error rates, and infrastructure health.

## Production Readiness

Beyond the core service, this project also includes:

- ✅ CI/CD pipeline with automated image publishing
- ✅ Kubernetes deployment manifests
- ✅ Automated DB migrations

## Author

**Paresh**
Backend Developer — Go · REST · gRPC · Docker · Distributed Systems

[GitHub](https://github.com/pareshlaptop23-source/cloud-assessment-tool-go) · [LinkedIn](https://www.linkedin.com/in/paresh-patil-044796204) · [Email](mailto:pareshspatil23@gmail.com)

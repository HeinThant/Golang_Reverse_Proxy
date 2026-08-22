# Go Reverse Proxy

A practical reverse proxy built with Go to demonstrate HTTP request forwarding, backend routing, and basic service abstraction.

This project supports my backend/platform engineering portfolio by showing the network-facing layer that often sits between clients and internal services.

## What It Demonstrates

- Reverse proxy behavior using Go standard library packages
- Forwarding client requests to an upstream backend service
- Separating client-facing endpoints from backend implementation details
- Local testing with a simple backend server
- Foundation concepts used by API gateways, load balancers, and service proxies

## Architecture

```text
Client
  |
  v
Reverse Proxy
  |
  v
Backend Service
```

The client communicates with the proxy, while the proxy forwards traffic to the backend service and returns the backend response.

## Tech Stack

- Go
- `net/http`
- Reverse proxy concepts
- Local service testing

## How To Run

Start the backend service:

```bash
go run backend.go
```

Start the reverse proxy in another terminal:

```bash
go run reverse_proxy_server.go
```

Send a request to the proxy endpoint according to the configured port in the source code.

## Production Considerations

For a production-grade reverse proxy, I would add:

- Config-driven upstream targets
- Structured logging with correlation/request IDs
- Request and response timeout controls
- Authentication and authorization at the proxy boundary
- TLS handling
- Health checks and failover behavior
- Metrics for latency, status codes, and upstream errors

## Why This Project Matters

Reverse proxying is a core concept behind API gateways, ingress controllers, load balancers, and secure backend integrations. This project keeps the implementation small so the routing behavior is easy to inspect.
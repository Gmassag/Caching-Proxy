# Caching Proxy

A **lightweight HTTP caching proxy** in Go with Redis backend, inspired by [roadmap.sh Caching Server](https://roadmap.sh/projects/caching-server).

**Features**

* HTTP Proxy: forwards requests to origin servers
* Redis Caching: in-memory cache (default 60s)
* `X-Cache` headers: `HIT` / `MISS`
* Preserves original response headers
* Docker-ready with logs for cache hits/misses and errors

---

## Quick Start

**Docker Compose (recommended)**

```bash
git clone https://github.com/Gmassag/Caching-Proxy.git
cd Caching-Proxy
docker compose up --build
```

Test:

```bash
curl -i http://localhost:3000/products/1  # MISS
curl -i http://localhost:3000/products/1  # HIT
```

**Manual Setup**

```bash
docker run -d -p 6379:6379 redis:8.0.3-alpine
export REDIS_URL=redis://localhost:6379
export PORT=3000
export ORIGIN=http://dummyjson.com
go run cmd/main.go
```

---

## Configuration

| Variable   | Description             | Default | Required |
| ---------- | ----------------------- | ------- | -------- |
| ORIGIN     | Upstream server URL     | -       | ✅        |
| REDIS\_URL | Redis connection string | -       | ✅        |
| PORT       | HTTP server port        | 3000    | ❌        |

---

## How It Works

1. Client sends request → Proxy checks Redis cache
2. **HIT** → returns cached response (`X-Cache: HIT`)
3. **MISS** → fetches from origin, caches response (`X-Cache: MISS`)

**Logs** show cache hits/misses and errors for easy monitoring.

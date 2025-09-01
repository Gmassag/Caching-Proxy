Caching Proxy
A lightweight HTTP caching proxy server built in Go with Redis backend, inspired by the roadmap.sh Caching Server project.
Overview
This caching proxy sits between clients and origin servers, forwarding requests and caching responses to improve performance and reduce backend load. When the same request is made again, it returns the cached response instead of hitting the origin server.
Features

✅ HTTP Proxy: Forwards requests to configurable origin servers
✅ Redis Caching: Uses Redis as a fast, in-memory cache backend
✅ Cache Headers: Adds X-Cache: HIT/MISS headers to responses
✅ Automatic Expiration: Cached responses expire after 60 seconds
✅ Header Preservation: Maintains original response headers
✅ Docker Support: Fully containerized deployment
✅ Detailed Logging: Cache hit/miss logging for monitoring

Quick Start
Using Docker Compose (Recommended)

Clone the repository:
bashgit clone https://github.com/Gmassag/Caching-Proxy.git
cd Caching-Proxy

Start the services:
bashdocker compose up --build

Test the proxy:
bash# First request (MISS) - fetches from origin
curl -i http://localhost:3000/products/1

# Second request (HIT) - returns from cache
curl -i http://localhost:3000/products/1


Manual Setup

Start Redis:
bashdocker run -d -p 6379:6379 redis:8.0.3-alpine

Set environment variables:
bashexport REDIS_URL=redis://localhost:6379
export PORT=3000
export ORIGIN=http://dummyjson.com

Run the application:
bashgo run cmd/main.go


Configuration
The proxy is configured via environment variables:
VariableDescriptionDefaultRequiredORIGINThe upstream server URL to proxy requests to-✅REDIS_URLRedis connection string-✅PORTHTTP server port3000❌
How It Works

Request Received: Client makes a request to the proxy
Cache Check: Proxy checks if response exists in Redis cache
Cache Hit: If cached, returns the stored response with X-Cache: HIT header
Cache Miss: If not cached:

Forwards request to origin server
Caches the response in Redis (60s TTL)
Returns response with X-Cache: MISS header



Example Usage
bash# Start the proxy pointing to DummyJSON API
docker compose up

# Test different endpoints
curl http://localhost:3000/products/1      # Returns product data
curl http://localhost:3000/users/1         # Returns user data
curl http://localhost:3000/posts           # Returns posts
Project Structure
.
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   └── redis.go             # Redis client setup
├── docker-compose.yml       # Docker services configuration
├── Dockerfile              # Application container
├── go.mod                  # Go module dependencies
└── go.sum                  # Go module checksums
Tech Stack

Go 1.25: Core application language
Redis 8.0: High-performance caching layer
Docker: Containerized deployment
Alpine Linux: Lightweight container images

Monitoring
The application provides detailed logging:
2025/09/01 12:55:53 Server avviato su :3000, origin: http://dummyjson.com
2025/09/01 12:56:15 CACHE MISS /products/1 - recupero da http://dummyjson.com/products/1
2025/09/01 12:56:20 CACHE HIT /products/1
Inspiration
This project is based on the Caching Server project idea from roadmap.sh:
https://roadmap.sh/projects/caching-server
# Simple HTTP Servers

Three simple HTTP servers that display "hello" and read from environment variables.

## Test build with Docker

All Dockerfiles use multi-stage builds with test stages. Tests are run during the build process and must pass before the final image is created.

**Go:**
```bash
cd go-server
docker build -t go-server .
docker run -p 8080:8080 -e MESSAGE="docker go" go-server
```

**Python:**
```bash
cd python-server
docker build -t python-server .
docker run -p 8080:8080 -e MESSAGE="docker python" python-server
```

**Node.js:**
```bash
cd nodejs-server
docker build -t node-server .
```
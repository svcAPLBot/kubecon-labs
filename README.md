# Simple HTTP Servers

Three simple HTTP servers that display "hello" and read from environment variables.

## Lab

Code, Build, Deploy, Expose

1. Create your feature branch
```
git fetch && git reset --hard origin/main
git checkout main -b <your-branch>
```

2. Test build locally (read: Test build with Docker)
3. Push your branch to remote origin
4. Open D1 bookmark in web browser
5. Log in to the `lab1` team in App Platform 
6. Create container image with your own tag
7. Create workload that uses that image
8. Create service that assigns a public URL to your workload
9. Click the service URL link to see the result (it may take few minutes until URL is available)

## Test build with Docker

All Dockerfiles use multi-stage builds with test stages. Tests are run during the build process and must pass before the final image is created.

**Go:**
```bash
cd go-server
docker build -t go-server .
```

**Python:**
```bash
cd python-server
docker build -t python-server .
```

**Node.js:**
```bash
cd nodejs-server
docker build -t node-server .
```
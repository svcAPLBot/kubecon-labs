# Simple HTTP Servers

Four simple HTTP servers that display "hello" and read from environment variables.

## Running the Servers


### Go
```bash
cd go-server
go run src/main.go
```

Run with custom environment variables:
```bash
PORT=8081 MESSAGE="from go" go run src/main.go
```

### Python
```bash
cd python-server
python3 src/server.py
```

Run with custom environment variables:
```bash
PORT=8081 MESSAGE="from python" python3 src/server.py
```

### Node.js
```bash
cd nodejs-server
node src/server.js
```

Run with custom environment variables:
```bash
PORT=8081 MESSAGE="from nodejs" node src/server.js
```

## Environment Variables

- **PORT**: Server port (default: 8080)
- **MESSAGE**: Custom message to display after "hello" (default: "world")

## Testing

### Unit Tests

**Go:**
```bash
cd go-server
go test -v
```

**Python:**
```bash
cd python-server
python3 -m pytest src/test_server.py -v
# or
python3 -m unittest src.test_server -v
```

**Node.js:**
```bash
cd nodejs-server
node src/test.js
```

### Manual Testing

After starting a server, test it in another terminal:
```bash
curl http://localhost:8080
```

Or with a custom port:
```bash
curl http://localhost:8081
```

## Docker

All Dockerfiles use multi-stage builds with test stages. Tests are run during the build process and must pass before the final image is created.

**Rust:**
```bash
cd rust-server
docker build -t rust-server .
docker run -p 8080:8080 -e MESSAGE="docker rust" rust-server
```
Stages: `builder` (compiles and runs tests) → final runtime image

**Go:**
```bash
cd go-server
docker build -t go-server .
docker run -p 8080:8080 -e MESSAGE="docker go" go-server
```
Stages: `builder` (compiles) → `tester` (runs tests) → final runtime image

**Python:**
```bash
cd python-server
docker build -t python-server .
docker run -p 8080:8080 -e MESSAGE="docker python" python-server
```
Stages: `tester` (runs tests) → final runtime image

**Node.js:**
```bash
cd nodejs-server
docker build -t node-server .
docker run -p 8080:8080 -e MESSAGE="docker nodejs" node-server
```
Stages: `tester` (runs tests) → final runtime image

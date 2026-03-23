FROM golang:1.21-alpine as builder
WORKDIR /app
COPY src/ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-server main.go

FROM golang:1.21-alpine as tester
WORKDIR /app
COPY src/ .
RUN go test -v

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/go-server /usr/local/bin/
EXPOSE 8080
CMD ["go-server"]

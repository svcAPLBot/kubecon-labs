FROM rust:latest as builder
WORKDIR /app
COPY Cargo.toml .
COPY src src/
COPY tests tests/
RUN cargo build --release
RUN cargo test --release

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
COPY --from=builder /app/target/release/rust-server /usr/local/bin/
EXPOSE 8080
CMD ["rust-server"]

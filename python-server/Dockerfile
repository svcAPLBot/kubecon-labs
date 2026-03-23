FROM python:3.11-slim
WORKDIR /app
COPY src/ .
RUN python3 -m unittest test_server -v

FROM python:3.11-slim
WORKDIR /app
COPY src/server.py .
EXPOSE 8080
CMD ["python3", "server.py"]

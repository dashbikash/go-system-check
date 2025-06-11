# Use the official Golang image as the build environment
FROM golang:1.24.4-alpine3.22 AS builder
WORKDIR /app

COPY ipc ./ipc

CMD ["sh", "-c", "while true; do sleep 3600; done"]
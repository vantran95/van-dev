# Build stage
FROM golang:1.23-alpine AS builder

# Set working directory for the build stage
WORKDIR /app

# Install necessary build dependencies
RUN apk add --no-cache git ca-certificates tzdata && \
    update-ca-certificates

# Copy go mod files first for better layer caching
COPY go.mod ./
RUN go mod download && go mod verify

# Copy the rest of the application code
COPY . .

# Expose the port that the application listens on
EXPOSE 8000

# Health check to verify the application is running properly
# This helps container orchestration systems know if the container is healthy
HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1


RUN GO111MODULE=on go install golang.org/x/tools/cmd/goimports@v0.25.0

RUN GO111MODULE=on go install github.com/volatiletech/sqlboiler/v4@v4.16.2 && \
    GO111MODULE=on go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@v4.16.2 \


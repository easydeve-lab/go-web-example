# Use the Go 1.23.3 base image
FROM golang:1.23.3 AS builder

# Set the current working directory in the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download and install the dependencies (this will cache dependencies if they don't change)
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go binary
RUN go build -o main .

# Use Ubuntu 22.04 as the base image to get a newer GLIBC version
FROM ubuntu:22.04

# Install bash and necessary dependencies (including GLIBC 2.34 or newer)
RUN apt-get update && apt-get install -y \
    bash \
    libc6 \
    ca-certificates \
    net-tools \
    && rm -rf /var/lib/apt/lists/*

# Set the current working directory
WORKDIR /app

# Copy the built Go binary from the builder stage
COPY --from=builder /app/main /app/

# Copy assets and templates to the container
COPY assets/ /app/assets/
COPY templates/ /app/templates/

# Expose the port the app runs on
EXPOSE 8080

# Set the entrypoint to run the Go application with bash
CMD ["bash", "-c", "./main"]

# ---------- Builder Stage ----------
# Use Golang as base image
FROM golang:1.22-alpine AS builder

# Install Git (required for Go modules)
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy
RUN go mod download

# Copy the project files
COPY . .

# Build the application
RUN go build -o go-todo ./cmd/main.go


# ---------- Final Stage ----------
# Use a minimal base image
FROM alpine:3.18

# Set working directory
WORKDIR /app

# Copy the compiled binary
COPY --from=builder /app/go-todo .

# 🔥 Copy wait-for.sh from your source repo to image
COPY wait-for.sh /wait-for.sh
RUN chmod +x /wait-for.sh

# Expose the application port
EXPOSE 9090

# Run the application
CMD ["./go-todo"]

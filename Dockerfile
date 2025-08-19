# Multi-stage build for Go job-listing app
# Build stage
FROM golang:1.21-alpine AS builder

# Install ca-certificates for HTTPS requests
RUN apk add --no-cache ca-certificates git

# Set working directory
WORKDIR /app

# Copy go mod files first (better caching)
COPY go.mod ./
RUN go mod download

# Copy source code
COPY *.go ./

# Build the application
# CGO_ENABLED=0 for static binary
# GOOS=linux for Linux target
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Production stage
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Create non-root user for security
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

# Set working directory
WORKDIR /app

# Copy binary from builder stage
COPY --from=builder /app/main .

# Copy static files and templates
COPY static/ ./static/
COPY *.html ./
COPY data.json ./

# Change ownership to non-root user
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

# Expose port
EXPOSE 8000

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8000/ || exit 1

# Run the application
CMD ["./main"]

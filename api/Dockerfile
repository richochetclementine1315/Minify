# ============================================
# Stage 1: Builder - Compile the Go application
# ============================================
# Use official Golang Alpine image as the base for building
FROM golang:alpine as builder

# Create the app directory inside the builder container
RUN mkdir /app

# Copy all source files from current directory to /build in the container
ADD . /build/

# Set the working directory to /build where source code is located
WORKDIR /build

# Compile the Go application into a binary named 'main'
RUN go build -o main .

# ============================================
# Stage 2: Runtime - Create minimal production image
# ============================================
# Use minimal Alpine Linux image for the final container (reduces image size)
FROM alpine

# Create a non-root user 'appuser' for security
# -S: Create a system user
# -D: Don't assign a password
# -H: Don't create a home directory
# -h /app: Set home directory to /app
RUN adduser -S -D -H -h /app appuser

# Switch to the non-privileged user to run the application
USER appuser

# Copy all files from current directory to /app (includes .env and other config files)
COPY . /app

# Copy the compiled binary from the builder stage to /app
COPY --from=builder /build/main /app/

# Set working directory to /app where the application will run
WORKDIR /app

# Expose port 3000 for the API server (matches APP_PORT in .env)
EXPOSE 3000

# Command to run the application when container starts
CMD ["./main"]
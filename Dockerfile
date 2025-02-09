# Stage 1(builder stage/environment): Build the application
FROM golang:1.23 AS builder

# Set environment variables
ENV GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0

# Set the working directory inside the container
WORKDIR /app

# Copy Go modules manifests,
# Copy new files or directories from and adds them to the filesystem of the container at the path.
COPY go.mod go.sum ./

# Download deendencies
RUN go mod tidy

# Copy the source code
COPY . .

# Build the Go binary
RUN go build -o main .

#Stage 2(runtime stage/environment): Create a lightweight image for production
FROM alpine:latest

# Set the working directory in the final image
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Copy the .env file if needed (optional; Docker Compose can set env vars too)
COPY .env .

# Expose the port your API listens on
EXPOSE 8080

# Command to run the binary
CMD ["./main"]

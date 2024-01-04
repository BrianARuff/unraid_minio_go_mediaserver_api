# First stage: Build the executable.
FROM golang:1.21.5 AS builder
WORKDIR /build

# Copy the go.mod and go.sum files first to leverage Docker cache
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the application source code from the local "src" directory to the "src" directory in the container
COPY src/ src/
COPY main.go .

# Build the application from the main.go file
RUN CGO_ENABLED=0 go build -ldflags='-w -s' -o media_app_one main.go

# Second stage: Create the final image with only the binary.
FROM alpine:latest  
WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /build/media_app_one .

# Copy the necessary certificate and key files for runtime
COPY .certs/ .certs/

# Copy the image directory for S3 uploads
# This ensures the image is present in the Docker container at the expected path
COPY images/ images/

# Run the application on container startup
CMD ["./media_app_one"]

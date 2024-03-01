#############      builder                                  #############
FROM golang:1.19.4 AS builder

WORKDIR /app


# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go dependencies
RUN go mod download

COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o upload_to_s3 .

# Start a new stage from scratch
FROM alpine:latest

# Install ca-certificates to support HTTPS connections
RUN apk --no-cache add ca-certificates

# Copy the binary from the first stage
COPY --from=builder /app/upload_to_s3 /upload_to_s3

# Define the entry point for the container
ENTRYPOINT ["/upload_to_s3"]   
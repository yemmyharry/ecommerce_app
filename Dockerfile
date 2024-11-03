# Dockerfile
# Use the official Golang image as a base
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to the working directory
COPY go.mod go.sum ./

# Download and cache dependencies
RUN go mod download

# Copy the rest of the application code to the working directory
COPY . .

# Set environment variables for the Go application
ENV GO111MODULE=on

# Build the Go application
RUN go build -o ecommerce-app .

# Expose the port that the app will run on
EXPOSE 8080

# Run the executable
CMD ["./ecommerce-app"]
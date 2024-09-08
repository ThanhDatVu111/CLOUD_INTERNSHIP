# Use the official Golang image as the base image
FROM golang:1.22.5-alpine

# Set the working directory inside the container
WORKDIR /app

# Set environment variables for the application
ENV GO111MODULE=on

# Copy only the necessary files
COPY api .
RUN go mod download

# Build the Go application, specifically the server.go file
RUN go build -o server ./cmd/server.go

# Expose the port the application will run on
EXPOSE 3000

# Run the executable
CMD ["./server"]

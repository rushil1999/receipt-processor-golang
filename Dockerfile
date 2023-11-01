# Use an official Go runtime as a parent image
FROM golang:1.21.1

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Build the Go application
RUN go build -o main

# Expose the API port
EXPOSE 8080

# Run the Go application
CMD ["./cmd/main"]

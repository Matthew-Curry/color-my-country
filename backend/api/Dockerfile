# Use the official Go image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod .
COPY go.sum .


# Download dependencies
RUN go mod download

# Copy the local package files to the container's working directory
COPY . .

#Set environment variable indicating docker environment
ENV DOCKERIZED=True

# Build the Go application
RUN go build -o main .

# Expose the port on which the API will run
EXPOSE 8080

# Command to run the application
CMD ["./main"]

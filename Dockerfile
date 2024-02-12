# Start from the official Go image to build our executable.
FROM golang:1.21

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go mod download

# Build the Go app
RUN go build -o main ./cmd

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["./main"]

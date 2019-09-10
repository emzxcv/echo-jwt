
# Start from the golang alpine base image
FROM golang:1.13.0-alpine3.10

# Add Maintainer Info
LABEL maintainer="Emily Ha <emilyha17@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 1323 to the outside world
EXPOSE 1323

# Command to run the executable
CMD ["./main"]

# Start from the latest golang base image
FROM golang:alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o /main

# execution stage
FROM alpine:latest

WORKDIR /app/

COPY --from=builder /main .

CMD ["./main", "worker"]

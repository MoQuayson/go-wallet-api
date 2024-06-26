# syntax=docker/dockerfile:1.4
FROM --platform=$BUILDPLATFORM golang:1.21-alpine AS builder

# Move to working directory /app
WORKDIR /app

# Copy the code into the container
COPY . .

# Build the application's binary
RUN go build -o main main.go

# stage 2: run stage
FROM alpine:latest

# Move to working directory /app
WORKDIR /app

# Copy the code into the container from builder
COPY --from=builder /app/main .
COPY --from=builder /app/.env ./.env
COPY --from=builder /app/.env ./backend

# Command to run the application when starting the container
CMD ["/app/main"]
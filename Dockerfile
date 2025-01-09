# Use the official Go image as the base for building
FROM golang:1.23 as builder

# Set the working directory
WORKDIR /app

COPY . .

# Build the pingcli binary (ensure static build)
RUN CGO_ENABLED=0 go mod tidy && CGO_ENABLED=0 go build -o /pingcli

# Debugging step: Check if the binary exists
RUN ls -al /pingcli

# Use Alpine as the runtime image
FROM alpine:latest

# Copy the binary from the builder stage
COPY --from=builder /pingcli /usr/local/bin/pingcli

# Set permissions for the binary
RUN chmod +x /usr/local/bin/pingcli

# Debugging step: Verify binary is in place
RUN ls -al /usr/local/bin/pingcli

# Set the entry point
ENTRYPOINT ["pingcli"]

# Allow subcommands
CMD []
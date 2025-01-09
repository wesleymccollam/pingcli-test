# Start from an empty scratch image
FROM alpine:latest

# Copy the statically compiled Go binary into the container
COPY pingcli /pingcli

# Set the binary as the entry point
ENTRYPOINT ["/pingcli"]

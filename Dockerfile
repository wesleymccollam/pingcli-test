# Use Alpine as the runtime image
FROM alpine:latest

# Copy the binary from the builder stage
COPY  pingcli /usr/local/bin/pingcli

# Set permissions for the binary
RUN chmod +x /usr/local/bin/pingcli

# Debugging step: Verify binary is in place
RUN ls -al /usr/local/bin/pingcli

# Set the entry point
ENTRYPOINT ["pingcli"]

# Allow subcommands
CMD []
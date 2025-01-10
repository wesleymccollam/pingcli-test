# Use Alpine as the runtime image
FROM alpine:latest

# Install glibc
RUN apk update && apk add --no-cache gcompat libstdc++

RUN mkdir -p /usr/local/bin/pingcli

# Copy in repo
COPY ./ /usr/local/bin/pingcli

# Set permissions for the binary
RUN chmod +x /usr/local/bin/pingcli

# Debugging step: Verify binary is in place
RUN ls -al /usr/local/bin/pingcli

# Set the entry point
ENTRYPOINT ["pingcli"]

# Allow subcommands
CMD []
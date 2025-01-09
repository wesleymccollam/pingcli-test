# Use the official Go image as the base for building
FROM scratch

# Copy content into working directory
COPY ./ /pingcli

# Set the entry point
ENTRYPOINT ["/pingcli"]

# Allow subcommands
CMD []
# Use Alpine as the runtime image
FROM scratch

# Set the entry point
ENTRYPOINT ["pingcli"]

# Allow subcommands
CMD []

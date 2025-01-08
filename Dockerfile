# Use Alpine as the runtime image
FROM scratch

# Set the entry point
ENTRYPOINT ["pingcli"]
COPY pingcli /

# Allow subcommands
CMD []

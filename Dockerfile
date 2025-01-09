FROM scratch

COPY pingcli /usr/bin/pingcli

# Set the entry point
ENTRYPOINT ["pingcli"]

# Allow subcommands
CMD []
#!/usr/bin/env bash

set -e

go build -o wft

printf "Created executable binary: wft\n"

sudo ln -sf "$(pwd)/wft" /usr/local/bin/wft

printf "Created symlink in /usr/local/bin\n"


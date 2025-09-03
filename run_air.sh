#!/bin/bash
# run_air.sh - run Air with a specific target binary
# Usage: ./run_air.sh webserver
#        ./run_air.sh cli

if [ -z "$1" ]; then
  echo "Error: No target specified."
  echo "Usage: $0 <target>"
  exit 1
fi

export TARGET="$1"

echo "Running Air for target: $TARGET"
air

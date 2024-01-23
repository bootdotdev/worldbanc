#!/bin/sh

# This script is used to print nice warning messages
# branded in worldbanc style

echo "============================================"
echo "=========== WORLDBANC WARNING =============="
echo "============================================"
echo "$WARN_MESSAGE"
echo "============================================"
echo "From: $WARN_FROM_NAME"
echo "============================================"

if [ -z "$WARN_MESSAGE" ]; then
    echo "WARN_MESSAGE is not set. Exiting with error."
    exit 69
fi

if [ -z "$WARN_FROM_NAME" ]; then
    echo "WARN_FROM_NAME is not set. Exiting with error."
    exit 69
fi

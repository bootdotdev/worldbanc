#!/usr/bin/env bash

# Check if a file path is provided
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <path_to_file>"
    exit 1
fi

# Read the file line by line
while IFS= read -r line; do

    # Extract the date from the line
    date="${line##*,}"

    # Extract the year from the date
    year="${date%%-*}"

    # Check if the year is before 2000
    if [[ $year -lt 2000 ]]; then
        # Print to stderr
        echo "$line" >&2
    else
        # Print to stdout
        echo "$line"
    fi
# skip the first line (header)
done < <(tail -n +2 "$1")

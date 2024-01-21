#!/bin/bash

# Check if a file path is provided
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <path_to_file>"
    exit 1
fi

# Read the file line by line
while IFS= read -r line; do
    # Skip the header line
    if [[ $line == amount,from_user_id,to_user_id,from_name,to_name,created_at ]]; then
        continue
    fi

    # Extract the date from the line
    date=$(echo "$line" | cut -d ',' -f 6)

    # Extract the year from the date
    year=$(echo "$date" | cut -d '-' -f 1)

    # Check if the year is before 2000
    if [[ $year -lt 2000 ]]; then
        # Print to stderr
        >&2 echo "$line"
    else
        # Print to stdout
        echo "$line"
    fi
done < "$1"

#!/bin/sh
# POSIX compliant shell script to generate 100 fake IDs

generate_fake_id() {
    counter=$1
    echo "ID-$(echo "fixed_seed_$counter" | cksum | cut -d ' ' -f 1)"
}

main() {
    for i in $(seq 1 12); do
        generate_fake_id $i
    done
}

main

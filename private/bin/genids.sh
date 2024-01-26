#!/bin/sh

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

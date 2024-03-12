#!/usr/bin/env bash

generate_fake_id() {
    counter=$1
    echo -n "ID-"
    echo "fixed_seed_$counter" \
        | cksum \
        | cut -d ' ' -f 1
}

for i in $(seq 1 12); do
    generate_fake_id $i
done

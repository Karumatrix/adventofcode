#!/bin/bash

if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <year> <file_extension>"
    exit 1
fi

YEAR=$1
EXTENSION=$2

mkdir -p "$YEAR"

for i in $(seq -w 1 25); do
    DAY_FOLDER="$YEAR/Day$i"
    mkdir -p "$DAY_FOLDER"
    touch "$DAY_FOLDER/part1$EXTENSION"
    touch "$DAY_FOLDER/part2$EXTENSION"
done

echo "Advent of Code structure for $YEAR created with $EXTENSION files"
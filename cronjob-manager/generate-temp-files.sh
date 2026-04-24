#!/bin/bash

TEMP_DIR="temp"
EXTENSIONS=("log" "txt" "tmp" "pdf" "csv" "json" "xml" "bak" "old" "data")
NUM_FILES=10

mkdir -p $TEMP_DIR

echo "--- Generating $NUM_FILES test files at $TEMP_DIR ---"

for ((i=0; i<NUM_FILES; i++)); do
    FILENAME="test_file_$((i+1)).${EXTENSIONS[$i]}"
    FILEPATH="$TEMP_DIR/$FILENAME"
    
    echo "ID: $RANDOM" > "$FILEPATH"
    echo "Timestamp: $(date --rfc-3339=seconds)" >> "$FILEPATH"
    head -c 50 /dev/urandom | base64 >> "$FILEPATH"

    AGE=$(( ( RANDOM % 55 ) + 5 ))

    touch -d "$AGE seconds ago" "$FILEPATH"

    echo "CREATED: $FILENAME ($AGE time ago)"
done

echo "----------------------------------------------------"
echo "Setup ready. Run worker to test."
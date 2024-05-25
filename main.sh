#!/bin/sh

allocate_memory() {
    # Get total and used memory in bytes
    TOTAL_MEM=$(free -b | awk '/Mem:/ {print $2}')
    USED_MEM=$(free -b | awk '/Mem:/ {print $3}')

    # Calculate memory to allocate
    TAKE=${TAKE:-15}
    MEMORY_TO_ALLOCATE=$(awk "BEGIN {print int($TOTAL_MEM * 0.$TAKE - $USED_MEM)}")

    if [ "$MEMORY_TO_ALLOCATE" -gt 0 ]; then
        # Allocate memory using dd
        dd if=/dev/zero of=/tmp/allocated_mem bs=1 count="$MEMORY_TO_ALLOCATE" 2>/dev/null
        echo "Allocated $MEMORY_TO_ALLOCATE bytes to reach target memory usage."
    else
        echo "No need to allocate memory."
    fi
}

# Initial memory allocation
allocate_memory
echo "Done!"

# Infinite loop with sleep and optional CPU-intensive task
while true; do
    sleep 86400  # Sleep for one day

    if [ -z "$NOCPUB" ]; then
        RESULT=1
        for i in $(seq 1 1000000); do
            RESULT=$((RESULT * i))
        done
    fi

    # Restart the script
    exec "$0" "$@"
done

#!/bin/bash

# Get total memory in bytes
total=$(free | awk '/^Mem:/{print $2}')
# Calculate the amount of memory to use (4% of total)
use=$((total / 4))
# Allocate memory
memory=$(head -c "$use" /dev/urandom)

while 1; do
  # Sleep for 20 seconds
  sleep 20
done

# Free memory
unset memory

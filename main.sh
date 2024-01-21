#!/bin/bash
total=$(free | awk '/^Mem:/{print $2}')
use=$((total / 5))
memory=$(head -c "$use" /dev/urandom)

echo "Done!"
while true; do
  sleep 20
done
unset memory

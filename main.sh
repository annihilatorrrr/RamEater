#!/bin/bash
total=$(free | awk '/^Mem:/{print $2}')
use=$((total / 20))
memory=$(head -c "$use" /dev/urandom)

echo "Done!"
while 1; do
  sleep 20
done
unset memory

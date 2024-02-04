#!/usr/bin/env bash
while true; do
  LISTEN=127.0.0.1:8080 gin -i \
    --build . \
    run
  echo "Command failed with exit code $?. Respawning.." >&2
  sleep 1
done

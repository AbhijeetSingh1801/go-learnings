#!/usr/bin/env bash
set -euo pipefail

declare  UP_COUNT
declare  DOWN_COUNT

key="https://example.com"

# safe read with default 0
current=${UP_COUNT[$key]:-0}   # <-- make sure this is spelled 'current'

# increment
UP_COUNT[$key]=$(( current + 1 ))

echo "current = $current"
echo "new value = ${UP_COUNT[$key]}"

#!/bin/bash

# Terminal color codes
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

LOG_FILE="./logs.txt"

SITES=(
  "https://www.google.com"
  "https://www.github.com"
  "https://nonexistentwebsite123.com"
)

declare -A UP_COUNT
declare -A DOWN_COUNT

FUN_FACTS=(
  "Did you know? Honey never spoils!"
  "ASCII art is cool 😎"
  "Stay positive! Even if a site goes down."
  "Fun fact: Bananas are berries, but strawberries aren't."
)

function random_fact {
    echo "${FUN_FACTS[$(( RANDOM % ${#FUN_FACTS[@]} ))]}"
}

while true; do
    clear
    echo -e "${YELLOW}==========================" >> "$LOG_FILE"
    echo "    🚦 UPTIME DASHBOARD 🚦" >> "$LOG_FILE"
    echo -e "==========================${NC}" >> "$LOG_FILE"

    echo -e "\n--- $(date) ---" >> "$LOG_FILE"

    for SITE in "${SITES[@]}"; do
        
        # Assume declare -A UP_COUNT DOWN_COUNT was done earlier
        key="$SITE"                            # safe key without quoting inside $(( ))
        HTTP_STATUS=$(curl -L -s -o /dev/null -w "%{http_code}" "$SITE")

        if [ "$HTTP_STATUS" -eq 200 ]; then
            # Get current value (default 0)
            current=${UP_COUNT[$key]:-0}
            # Increment and store back
            UP_COUNT[$key]=$(( current + 1 ))
            echo -e "${GREEN}[UP]   $SITE (${HTTP_STATUS})${NC}"
            echo "$(date +"%H:%M:%S") [UP] $SITE ($HTTP_STATUS)" >> "$LOG_FILE"
        else
            current=${DOWN_COUNT[$key]:-0}
            DOWN_COUNT[$key]=$(( current + 1 ))
            FACT=$(random_fact)
            echo -e "${RED}[DOWN] $SITE (${HTTP_STATUS}) → $FACT${NC}"
            echo "$(date +"%H:%M:%S") [DOWN] $SITE ($HTTP_STATUS) → $FACT" >> "$LOG_FILE"
        fi

    done

    sleep 5
done

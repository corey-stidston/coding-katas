#!/bin/bash

KATA_FILTER="${1:-*}"

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Running Go tests in all subdirectories..."
echo "----------------------------------------"

for dir in $(find . -type f -name "$KATA_FILTER.go" -exec dirname {} \; | sort -u); do
    if [ -f "$dir/go.mod" ]; then
        echo -e "\nTesting in ${GREEN}$dir${NC}"
        (cd "$dir" && go test -v ./...)
        if [ $? -eq 0 ]; then
            echo -e "${GREEN}✓ Tests passed in $dir${NC}"
        else
            echo -e "${RED}✗ Tests failed in $dir${NC}"
        fi
    fi
done

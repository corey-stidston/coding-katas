#!/bin/bash

# Script to run go fmt on all directories containing go.mod files

echo "Formatting Go code in all modules..."

# Find all directories containing go.mod files
find . -name "go.mod" -type f -exec dirname {} \; | while read -r dir; do
    echo "Formatting Go code in $dir"
    # Change to the directory
    pushd "$dir" > /dev/null
    
    # Run go fmt on all packages in the module
    go fmt ./...
    
    # Return to the original directory
    popd > /dev/null
done

echo "Formatting complete!"
#!/bin/bash

if [ -z "$1" ]; then
    echo "Please provide a kata name"
    exit 1
fi

KATA_NAME="${1/-/_}"
KATA_FOLDER="${KATA_NAME/_/-}"
KATA_DIR="$PWD/$KATA_FOLDER"

# Create directory
mkdir -p "$KATA_DIR"

# Create go.mod file
cat > "$KATA_DIR/go.mod" << EOF
module github.com/corey-stidston/coding-katas/$KATA_NAME

go 1.21
EOF

# Create main.go file
cat > $KATA_DIR/$KATA_NAME".go" << EOF
package main

func getHelloMessage() string {
	return "Hello, World!"
}

func main() {
}
EOF

# Create main_test.go file
cat > $KATA_DIR/$KATA_NAME"_test.go" << EOF
package main

import "testing"

func TestGetHelloMessage(t *testing.T) {
	expected := "Hello, World!"
	result := getHelloMessage()

	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

EOF

echo "Created new kata project in $KATA_DIR"

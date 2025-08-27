#!/bin/bash

# Fix duplicate type definitions in generated Go bindings

echo "Fixing duplicate type definitions..."

# Remove PackedUserOperation type definition from all generated files except types.go
for file in *.go; do
    if [[ "$file" != "types.go" && "$file" != "go.mod" ]]; then
        echo "Processing $file..."
        
        # Create a temporary file
        temp_file=$(mktemp)
        
        # Remove the PackedUserOperation type definition and its closing brace
        awk '
        BEGIN { skip = 0 }
        /^type PackedUserOperation struct/ { skip = 1; next }
        skip == 1 && /^}$/ { skip = 0; next }
        skip == 0 { print }
        ' "$file" > "$temp_file"
        
        # Replace the original file
        mv "$temp_file" "$file"
    fi
done

echo "Duplicate type definitions fixed!"
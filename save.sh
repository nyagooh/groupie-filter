#!/bin/bash

# Recursively iterate through all files in the current directory
find . -type f | while read -r file; do
  # Add each file to git
  git add "$file"
  
  # Commit with a message mentioning the filename
  git commit -m "pulled repo from git: added $(basename "$file")"
done

echo "All files added and committed."
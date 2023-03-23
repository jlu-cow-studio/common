#!/bin/bash

commit_msg="$1"
if [ -z "$commit_msg" ]
then
  echo "Please provide a commit message"
  exit 1
fi

# Get the latest tag
last_tag=$(git describe --abbrev=0 --tags)

# Extract the version number
if [[ $last_tag =~ ^v([0-9]+)\.([0-9]+)\.([0-9]+)$ ]]
then
  major=${BASH_REMATCH[1]}
  minor=${BASH_REMATCH[2]}
  patch=${BASH_REMATCH[3]}
else
  echo "Invalid tag format: $last_tag"
  exit 1
fi

# Increment the patch number
patch=$((patch + 1))

# Create the new tag
new_tag="v$major.$minor.$patch"

echo "Creating new tag: $new_tag"

# Commit changes and tag
git add .
git commit -m "$commit_msg"
git tag "$new_tag"
git push origin master "$new_tag"

echo "Done"
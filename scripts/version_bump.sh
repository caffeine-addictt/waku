#!/bin/bash

VERSION=$1
VERSION_FILES="README.md pkg/version/version.go"
VERSION_DIRS="www/docs scripts"
CURRENT_VERSION=$(grep -oP 'const Version = "\K[\d\.]+' 'pkg/version/version.go' | head -n 1)
SEMVER_REGEX=$(grep -oP "MustCompile\(\`\K.+(?=\`\))" "pkg/version/version_test.go" | head -n 1)

if [[ -z "$VERSION" ]]; then
  echo "Usage: make bump version=x.x.x"
  exit 1
fi

if [[ -z "$CURRENT_VERSION" ]]; then
  echo "No version found in cmd/global/version.go"
  exit 1
fi

if [[ -z "$SEMVER_REGEX" ]]; then
  echo "No semver regex found in cmd/global/version_test.go"
  exit 1
fi

# strip v prefix
VERSION=${VERSION#v}

if [ "$CURRENT_VERSION" = "$VERSION" ]; then
  echo "$CURRENT_VERSION already matches $VERSION"
  exit 0
fi

# check semver
if ! echo "$VERSION" | grep -qP "$SEMVER_REGEX"; then
  echo "$VERSION does not follow semver."
  exit 1
fi

# files
for file in $VERSION_FILES; do
  sed -i "s/$CURRENT_VERSION/$VERSION/g" "$file"
done

# dir
for dir in $VERSION_DIRS; do
  find "$dir" -type f -print0 | xargs -0 sed -i "s/$CURRENT_VERSION/$VERSION/g"
done

echo "bumped $CURRENT_VERSION -> $VERSION"
echo "Be sure to manually verify all updated files are correct before committing."

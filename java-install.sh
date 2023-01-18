#!/bin/bash
RELEASE_VERSION=$1

if [[ -z "$RELEASE_VERSION" ]]; then
   RELEASE_VERSION="0.0.0"
fi

echo "generating version: $RELEASE_VERSION"

cd java && ./gradlew build publishToMavenLocal -Pversion=$RELEASE_VERSION

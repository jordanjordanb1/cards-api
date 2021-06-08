#!/usr/bin/env

set -e

# Build dir
BUILD_DIR=./build

# Clears build directory
rm -rf $BUILD_DIR

# Make directory
mkdir -p $BUILD_DIR

# Build
go build -o $BUILD_DIR

exit

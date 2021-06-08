#!/usr/bin/env

set -e

BUILD_DIR=./build

(export GIN_MODE=release; $BUILD_DIR/cards-api)

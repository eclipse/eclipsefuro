#!/usr/bin/env bash
# exit when any command fails
set -e

# check for breaking changes in proto against main branch
buf breaking --against '.git#branch=main'
#!/usr/bin/env bash
# exit when any command fails
set -e

buf generate --path $(find dist/protos/ -type d | grep dist/protos/[^$] | tr '\n' , | sed 's/.$//')
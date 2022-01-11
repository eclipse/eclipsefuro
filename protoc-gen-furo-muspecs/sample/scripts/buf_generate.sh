#!/usr/bin/env bash
# exit when any command fails
set -e

buf generate --path $(find sourceprotos/ -type d | grep sourceprotos/[^$] | tr '\n' , | sed 's/.$//')
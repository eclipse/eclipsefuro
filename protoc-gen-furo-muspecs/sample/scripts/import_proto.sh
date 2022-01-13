#!/usr/bin/env bash
# exit when any command fails
set -e

buf generate --template ./buf.protoimport.yaml --path $(find sourceprotos/ -type d | grep sourceprotos/[^$] | tr '\n' , | sed 's/.$//')
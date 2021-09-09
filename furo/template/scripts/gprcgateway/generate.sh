#!/bin/bash

TARGETDIR=pkg/grpc-gateway
mkdir -p $TARGETDIR

# the registration of the services..
furo exportAsYaml | simple-generator -t scripts/gprcgateway/autoregister.go.tpl > $TARGETDIR/autoregister.go

# beautify
go fmt $TARGETDIR/autoregister.go

echo "gateway sources generated"
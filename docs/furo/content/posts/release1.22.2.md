---
title: "Release1.22.2"
date: 2020-12-01T06:02:22+01:00
---

# Release 1.22.2
- commit: eaee2ad6

## features:
*muType*: multiline type field definitions supported, this makes it easier to write good descriptions.

*muService*: multiline method definitions supported, this makes it easier to write good descriptions.    

*muService*: multiline query param field definitions supported, this makes it easier to write good descriptions.    

*muSpec2Spec*: flag –overwrite-spec-options to generate the default options based on the .furo config again.

## fixes: 
*muSpec2Spec* --overwrite-spec-options null pointer in flow mode.

*genEsModule* added query fields from request type to service.

## furoBEC
docker pull thenorstroem/furo-bec:v1.22.2
docker pull thenorstroem/furo-bec:latest

## homebrew
Formula udated (eb160b01b623336a9dede0ed8880315fa51e54b6aa4fc231156a0c4a05f94285)
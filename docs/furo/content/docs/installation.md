---
title: "Installation"
weight: 5
---


# Installation
You can install furo on your local machine. For working with specs only, this is enough. 
We recommend [furoBEC](./tools/BEC/) if you have to generate more then "just" the specs. furoBEC is a
docker image which has nearly all dependencies already installed to generate all the additional things you may need to.

## Installation with brew
    brew tap theNorstroem/tap
    brew install furo

## Installation with go
    GO111MODULE=on go get github.com/theNorstroem/furo@v1.xx.xx

## Installation from sources
    git clone git@github.com:theNorstroem/furo.git
    go install

## Working without installation
Use the docker image [furoBEC](./tools/BEC/)

    docker run -it --rm -v `pwd`:/specs thenorstroem/furo-bec
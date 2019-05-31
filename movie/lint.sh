#!/bin/sh
golangci-lint run --deadline 20m --enable-all --disable=goimports --disable=lll --disable=dupl --tests=false
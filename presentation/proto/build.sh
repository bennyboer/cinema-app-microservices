#!/bin/sh
command -v protoc >/dev/null 2>&1 || { echo >&2 "You must install the Protocol Buffer compiler (protoc) first in order to build. Install it or check your PATH variable. Exiting."; exit 1; }
protoc -I=. --micro_out=. --go_out=. service.proto
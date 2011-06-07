#!/usr/bin/bash
pushd ..
make
popd
touch sample.go
make

#!/bin/bash

export CURDIR=$PWD
export GOPATH=$CURDIR/../../../../

cd ./tests/c/bin/
./ctest rand
./cbenchmark

cp -f $CURDIR/tests/rand.bin $CURDIR/benchmarks

cd $CURDIR/benchmarks

go test -v -test.bench=".*" -count=1



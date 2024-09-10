#!/bin/bash

args=`find dataset -type f | xargs`

echo -e "Serial"
time bash go/serial/run.sh $args
echo
echo -e "Concurrent"
time bash go/concurrent/run.sh $args

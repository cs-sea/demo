#!/bin/sh

for loop in {1..200000} 
do
    ./producer
    echo "The value is: $loop"
done

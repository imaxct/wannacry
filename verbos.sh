#!/bin/bash
for file in $( find . ); do
: `./Main -opt e -o $file`
done

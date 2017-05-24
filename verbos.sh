#!/bin/bash
opt=${1-'e'}

if [[ "$opt" != "e" && "$opt" != "d" ]]; then
	echo "wrong parameter"
	exit 1
fi
for file in $( find . -type f ); do
: `./Main -opt $opt -o "$file"`
done

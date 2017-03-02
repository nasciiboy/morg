#! /bin/bash

for file in *.out
do
 mv "$file" "${file%.html.out}.html"
done

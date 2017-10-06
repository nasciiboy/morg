#! /bin/bash

for file in *.out
do
 mv "$file" "${file%.morg.out}.morg"
done

for file in *.out
do
 mv "$file" "${file%.html.out}.html"
done

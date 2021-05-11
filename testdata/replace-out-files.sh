#! /bin/bash

for file in *\.morg.out
do
 mv "$file" "${file%.morg.out}.morg"
done

for file in *\.html.out
do
 mv "$file" "${file%.html.out}.html"
done

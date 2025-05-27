#!/bin/bash

# orders files in migrations directory alphabetically,
# then loops through them and appends /tmp/schema.sql 
# with the contents of each file 
#
# Sqlc uses alphabetical order when determining the sequence
# of migrations 
compile-sql () {

  files=$(find ./internal/db/migrations -maxdepth 1 -mindepth 1 -type f -printf '%f\n' | sort -d)

   echo "" > ./tmp/schema.sql

   for file in $files 
      do 
	cat "./internal/db/migrations/$file" >> ./tmp/schema.sql
    done
}


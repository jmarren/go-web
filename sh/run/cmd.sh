#!/bin/bash

run () {
  shift
       # build app.ts
       npm run build --prefix ./web/js &&
       # cd into js-build
       cd ./pkg/js-build
       # run js-build to bundle up app.js and extensions into <root>/web/public/index.js
       go run main.go && 
       # cd back
       cd ../../
       # generate templ files
       templ generate &&
       # generate sqlc
       sqlc generate -f ./internal/db/sqlc.yaml &&
       # run app with air
       air -- -env="$1" 

}

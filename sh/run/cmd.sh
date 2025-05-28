#!/bin/bash

run () {
  shift
       # build app.ts
       npm run build --prefix $MYAPP_DIR/web/js && 

       # cd into app root directory
       cd $MYAPP_DIR

       # run js build
       go run ./cmd/js-build

       # generate templ files
       templ generate -path $MYAPP_DIR &&
       # generate sqlc
       sqlc generate -f $MYAPP_DIR/internal/db/sqlc.yaml &&
       # run app with air
       air -- -env="$1" 

}

#!/bin/bash

# I prefer to use a bash script to initialize a cli
# in order to prevent polluting my PATH
myapp () {

# source deploy script
source ./deploy/cmd.sh

# handle argument 1
 case "$1" in 
    # run air with all args
    run)
       # build app.ts
       npm run build --prefix ./web/js &&
       # cd into js-build
       cd ./pkg/js-build
       # run js-build to bundle up app.js and extensions into <root>/web/public/index.js
       go run main.go && 
       # cd back
       cd ../../
       # generate templates
       templ generate &&
       # run app with air
       air $@
      ;;
    # connect to instance with name $2
    connect)
      ssh $2 -F ./deploy/ansible/ssh/$2.ssh
      ;;
	  #  connect-dev)
	  #     case "$2" in
	  # app)
	  #     ssh test@127.0.0.1 -p 
    # run deploy with all args
    deploy)
	deploy $@
	;;
      *)
      echo "command '$1' not found. Expected myapp [run | deploy]"
      ;;
  esac
}

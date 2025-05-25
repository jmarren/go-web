# #!/bin/bash

# I prefer to use a bash script to initialize a cli
# in order to prevent polluting my PATH
myapp () {

# source deploy script
source ./deploy/cmd.sh

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
    # connect to instance
    connect)
      ssh $2 -F ./deploy/sshconfig
      ;;
    # run deploy with all args
    deploy)
	deploy $@
	;;
      *)
      echo "command '$1' not found. Expected myapp [run | deploy]"
      ;;
  esac
}

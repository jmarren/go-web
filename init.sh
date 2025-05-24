# #!/bin/bash

# I prefer to use a bash script to initialize a cli
# in order to prevent polluting my PATH
myapp () {

# source deploy script
source ./deploy/cmd.sh

 case "$1" in 
    # run air with all args
    run)
       npm run build --prefix ./web/js &&
       cd ./pkg/js-build
       go run main.go && 
       cd ../../
       # go build -o ./pkg/js-build/build/main ./pkg/js-build &&
       # ./pkg/js-build/build &&
       templ generate &&
       air $@
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

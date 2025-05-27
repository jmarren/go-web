#!/bin/bash

# I prefer to use a bash script to initialize a cli
# in order to prevent polluting my PATH
myapp () {

# source main script
source ./sh/cmd.sh

cmd "$@"

}


loadbalancer() {

   case "$1" in 
      run)
	  go run ./cmd/load-balancer/main.go
      ;;
      p)
	 cat ./init.sh
	 ;;
      *)
	 echo "command '$1' not found"
   esac


}


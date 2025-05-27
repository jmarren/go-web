#!/bin/bash

cmd () {

# handle argument 1
 case "$1" in 
    run)
     source ./sh/run/cmd.sh
     run "$@"
     ;;
    connect)
      source ./sh/connect/cmd.sh
      connect "$@"
      ;;
    deploy)
      source ./sh/deploy/cmd.sh
      deploy "$@"
      ;;
    play)
      source ./sh/play/cmd.sh
      play "$@"
      ;;
    *)
      echo "command '$1' not found. Expected myapp [run | deploy]"
      ;;
  esac
}




compile-sql () {

   find ./internal/db/migrations -maxdepth 1 -mindepth 1 -type f -printf '%f\n'

	  #
	  #  for file in /internal/db/migrations/*/     # list directories in the form "/tmp/dirname/"
	  #    do
	  # dir=${dir%*/}      # remove the trailing "/"
	  # echo "${dir##*/}"    # print everything after the final "/"
	  #    done

}


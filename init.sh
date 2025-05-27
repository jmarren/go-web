#!/bin/bash

# I prefer to use a bash script to initialize a cli
# in order to prevent polluting my PATH
myapp () {

# source main script
source ./sh/cmd.sh

cmd "$@"

# # source deploy script
# source ./deploy/cmd.sh

}

# # handle argument 1
#  case "$1" in 
#     # run air with all args
#     run)
#        # build app.ts
#        npm run build --prefix ./web/js &&
#        # cd into js-build
#        cd ./pkg/js-build
#        # run js-build to bundle up app.js and extensions into <root>/web/public/index.js
#        go run main.go && 
#        # cd back
#        cd ../../
#        # generate templates
#        templ generate &&
#
#        sqlc generate -f ./internal/db/sqlc.yaml &&
#        # run app with air
#        air -- -env="$2" 
#       ;;
#     # connect to instance with name $2
#     connect)
#       ssh $2 -F ./deploy/ansible/ssh/$2.ssh
#       ;;
#     connectdev)
#        connectdev $@
#    ;;
#     # run deploy with all args
#     deploy)
# 	deploy $@
# 	;;
#       *)
#       echo "command '$1' not found. Expected myapp [run | deploy]"
#       ;;
#   esac
# }
#
#
# function connectdev () {
#    shift 
#       case "$1" in
# 	  app)
# 	      ssh test@127.0.0.1 -p 201
# 	       ;;
# 	  db)
# 	      ssh test@127.0.0.1 -p 200
# 	      ;;
# 	  dbpsql) 
# 	      psql -U postgres -h 127.0.0.1 -d db1
# 	      ;;
# 	  *)
# 	    echo "command '$2' not found. Expected [ app | db | dbpsql ]"
# 	    ;;
# 	   esac
#
# }
#
#
# compile-sql () {
#
#    find ./internal/db/migrations -maxdepth 1 -mindepth 1 -type f -printf '%f\n'
#
# 	  #
# 	  #  for file in /internal/db/migrations/*/     # list directories in the form "/tmp/dirname/"
# 	  #    do
# 	  # dir=${dir%*/}      # remove the trailing "/"
# 	  # echo "${dir##*/}"    # print everything after the final "/"
# 	  #    done
#
# }
#

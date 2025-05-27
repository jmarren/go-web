#!/bin/bash

function connect () {
   shift 
      case "$1" in
	  app)
	      ssh app -F ./deploy/ansible/ssh/app.ssh
	      ;;
	  db)
	      ssh db -F ./deploy/ansible/ssh/db.ssh
	      ;;
	  devapp)
	      ssh test@127.0.0.1 -p 201
	       ;;
	  devdb)
	      ssh test@127.0.0.1 -p 200
	      ;;
	  devpsql) 
	      psql -U postgres -h 127.0.0.1 -d db1
	      ;;
	  *)
	    echo "command '$2' not found. Expected [ app | db | devapp | devdb | devpsql ]"
	    ;;
	   esac

}

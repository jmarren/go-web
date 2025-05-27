#!/bin/bash

play () {
  shift 

    case "$1" in 
      compsql)
	source ./sh/play/compile-sql.sh
	compile-sql
	;;
      db) 
	source ./sh/play/compile-sql.sh
	compile-sql
	ansible-playbook ./deploy/ansible/db.yaml
	;;
      p)
	cat ./sh/play/cmd.sh
	;; 
      *)
	echo "command '$1' not found. Expected [ db ] " 
	;;

     esac

}




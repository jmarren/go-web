#!/bin/bash

play () {
  shift 

    case "$1" in 
      db) 
	ansible-playbook ./deploy/ansible/db.yaml
	;;
      *)
	echo "command '$1' not found. Expected [ db ] " 
	;;

     esac

}

#!/bin/bash

function ans () { 
 shift

 case "$1" in 
    ping) 
	ansible ec2 -m ping -i ./deploy/ansible/inventory.yaml
	;;
    playdb)
	ansible-playbook ./deploy/ansible/db.yaml
	;;
    *) 
	echo "command '$1' not found"
	echo "expected [ ping ]"
	;;

  esac
}

#!/bin/bash
deploy () {

 shift

 case "$1" in 
    fresh)
        source ./sh/deploy/fresh.sh
        fresh "$@"
	;;
    destroy)
	terraform -chdir=./deploy/terraform destroy --auto-approve
	;;
    ans) 
	source ./deploy/ansible/cmd.sh
	ans "$@"
	;;
    *)
	echo "command '$1' not found"
	echo "usage:  deploy [ apply | destroy ]"
	;;
esac
}







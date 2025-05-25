# #!/bin/bash
deploy () {
	echo "deploy hit"
	echo $@

 case "$2" in 
    # run air with all args
    apply)
	# apply terraform config
	terraform -chdir=./deploy/terraform apply -auto-approve &&
	
        # get ec2 IPs and set them in env variables
	export PLAYFUL_IP=$(terraform -chdir=./deploy/terraform output -json | jq '.playful_1_instance_ip.value' | tr -d '"')
	# write sshconfig file using IP address(es)
	envsubst < ./deploy/sshconfig.template > ./deploy/sshconfig
	
	# run ansible playbook
	ansible-playbook ./deploy/ansible/playbook.yaml
	;;
   

esac

	
}

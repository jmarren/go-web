#!/bin/bash

fresh () {
    # apply terraform config
    terraform -chdir=./deploy/terraform apply -auto-approve &&
    
    # get ec2 IPs and set them in env variables
    export APP_SERVER_IP=$(terraform -chdir=./deploy/terraform output -json | jq '.app_server_ip.value' | tr -d '"')
     
    export DB_SERVER_IP=$(terraform -chdir=./deploy/terraform output -json | jq '.db_server_ip.value' | tr -d '"')
    # write sshconfig file using IP address(es)
    envsubst < ./deploy/ansible/templates/app.ssh.template > ./deploy/ansible/ssh/app.ssh

    envsubst < ./deploy/ansible/templates/db.ssh.template > ./deploy/ansible/ssh/db.ssh
    
    # run ansible playbook
    ansible-playbook ./deploy/ansible/app.yaml &&
    ansible-playbook ./deploy/ansible/db.yaml 
    
}

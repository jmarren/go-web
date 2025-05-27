FROM ubuntu:latest

# NOTE: You must manually ssh into the container using:
#	 	`ssh test@127.0.0.1 -p <SSH PORT>`
# 	at least once in order to add the container to 
#       known hosts in order for ansible to be able to connect



# REFERENCE:
# https://dev.to/s1ntaxe770r/how-to-setup-ssh-within-a-docker-container-i5i

RUN apt update && apt install openssh-server -y && apt -y install sudo

RUN useradd -rm -d /home/ubuntu -s /bin/bash -g root -G sudo -u 1001 test

RUN echo 'test:test' | chpasswd

RUN service ssh start

CMD ["/usr/sbin/sshd","-D"]


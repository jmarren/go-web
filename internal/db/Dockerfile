FROM ubuntu:latest

# REFERENCE:
# https://dev.to/s1ntaxe770r/how-to-setup-ssh-within-a-docker-container-i5i

RUN apt update && apt install openssh-server -y && apt -y install sudo

RUN useradd -rm -d /home/ubuntu -s /bin/bash -g root -G sudo -u 1001 test

RUN echo 'test:test' | chpasswd

RUN service ssh start

CMD ["/usr/sbin/sshd","-D"]


# FROM ubuntu:22.04
#
# # set shell to bash and error on any failure, unset variables, and pipefail
# SHELL ["/bin/bash", "-euo", "pipefail", "-c"]
#
# RUN apt-get update && apt-get install --no-install-recommends openssh-server sudo python3 pip -y && \
#     rm -rf /var/lib/apt/lists/* && \
#     pip install --no-cache-dir supervisor
#
# COPY ./ubuntu.ssh.supervisor.conf /etc/supervisor.conf
#
# RUN echo "root:root" | chpasswd && \
#     # Enable password authentication over SSH for root
#     sed 's/\(#\{0,1\}PermitRootLogin .*\)/PermitRootLogin yes/' /etc/ssh/sshd_config > /tmp/sshd_config && mv /tmp/sshd_config /etc/ssh/sshd_config && \
#     sed 's/\(#\{0,1\}PasswordAuthentication .*\)/PasswordAuthentication yes/' /etc/ssh/sshd_config > /tmp/sshd_config && mv /tmp/sshd_config /etc/ssh/sshd_config && \
#     # Start the SSH service to create the necessary run symlinks
#     service ssh start
#
# # Expose docker port 22
# EXPOSE 22
#
# CMD ["supervisord", "-c", "/etc/supervisor.conf"]
#


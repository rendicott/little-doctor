#!/bin/bash
# generate missing keys per this post http://serverfault.com/questions/158151/sshd-shuts-down-with-no-supported-key-exchange-algorithms-error

ssh-keygen -t rsa -f /etc/ssh/ssh_host_rsa_key
ssh-keygen -t ecdsa -f /etc/ssh/ssh_host_ecdsa_key
ssh-keygen -t ed25519 -f /etc/ssh/ssh_host_ed25519_key
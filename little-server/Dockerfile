FROM centos:7
ENV http_proxy http://10.114.16.12:80
ENV https_proxy http://10.114.16.12:80
ENV HTTP_PROXY http://10.114.16.12:80
ENV HTTPS_PROXY http://10.114.16.12:80
ENV no_proxy "localhost,127.0.0.1,ge.com,3.0.0.0/8,10.0.0.0/8"
RUN yum install epel-release -y
RUN yum install ansible -y
RUN yum install net-tools -y
RUN yum install python-pip -y
RUN pip install redis
RUN mkdir /usr/ansible
RUN yum install openssh -y
RUN ssh-keygen -b 2048 -t rsa -f /root/.ssh/id_rsa -q -N ""
ADD /keys/client-key.priv /root/.ssh/id_rsa
WORKDIR /usr/ansible
ADD /scripts/ansible-inventory.py /usr/ansible
ADD /scripts/ansible.cfg /etc/ansible/ansible.cfg
RUN yum install which -y
RUN yum install openssh-clients -y
ENTRYPOINT ["/bin/bash"]
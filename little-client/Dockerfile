FROM centos:7
ENV http_proxy http://10.114.16.12:80
ENV https_proxy http://10.114.16.12:80
ENV HTTP_PROXY http://10.114.16.12:80
ENV HTTPS_PROXY http://10.114.16.12:80
ENV no_proxy "localhost,127.0.0.1,ge.com,3.0.0.0/8,10.0.0.0/8"
RUN yum install epel-release -y
RUN yum install nginx -y
RUN yum install openssh -y
# set up a key so ansible can talk to this client
ADD ./keys/client-key.pub /root/client-key.pub
RUN ssh-keygen -b 2048 -t rsa -f /root/.ssh/id_rsa -q -N ""
RUN touch /root/.ssh/authorized_keys
RUN chmod 600 /root/.ssh/authorized_keys
RUN cat /root/client-key.pub /root/.ssh/authorized_keys

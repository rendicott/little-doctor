# modified from https://semaphoreci.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker
FROM golang:1.6
ENV http_proxy http://10.114.16.12:80
ENV https_proxy http://10.114.16.12:80
ENV HTTP_PROXY http://10.114.16.12:80
ENV HTTPS_PROXY http://10.114.16.12:80
ENV no_proxy "localhost,127.0.0.1,ge.com,3.0.0.0/8,10.0.0.0/8"
# Install beego and the bee dev tool
RUN go get github.com/astaxie/beego && go get github.com/beego/bee

# Expose the application on port 8080
EXPOSE 8080

# Set the entry point of the container to the bee command that runs the
# application and watches for changes

# set up a key so ansible can talk to this client
ADD ./keys/client-key.pub /root/client-key.pub
RUN ssh-keygen -b 2048 -t rsa -f /root/.ssh/id_rsa -q -N ""
RUN touch /root/.ssh/authorized_keys
RUN chmod 600 /root/.ssh/authorized_keys
RUN cat /root/client-key.pub >> /root/.ssh/authorized_keys

WORKDIR /tmp/
ADD scripts /tmp/
RUN /tmp/redis-prereqs.sh

RUN mkdir /app
WORKDIR /app
ADD /app/MathApp /app/MathApp
ADD /app/MathApp/views /app/views
ADD /app/MathApp/conf /app/conf
RUN go build /app/MathApp/main.go

RUN apt-get update
RUN apt-get install openssh-server -y
RUN apt-get install vim -y
#ENTRYPOINT "/tmp/post-ip-to-redis 172.17.0.1 6379 && /app/main"
ENTRYPOINT "/tmp/launcher.sh"
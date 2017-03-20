# little-doctor

An Ansible server with a client node that can be multiplied to test Ansible against many hosts.

## Usage

Build then run like so
```
docker-compose build
docker-compose up
```

To attach to the server container and start running ansible commands look up the id of the container and attach to it:
```
docker ps
docker exec -it <containerID> /bin/bash
```

Example:
```
1734021-C02RQ08ZG8WP:~ russellendicott$ docker ps
CONTAINER ID        IMAGE                               COMMAND                  CREATED             STATUS              PORTS                           NAMES
4bc3b0280282        littledoctor_little-doctor-server   "/bin/bash"              18 seconds ago      Up 16 seconds                                       littledoctor_little-doctor-server_1
a20d150b2148        littledoctor_little-doctor-client   "/tmp/launcher.sh"       18 seconds ago      Up 17 seconds       22/tcp, 8080/tcp                littledoctor_little-doctor-client_1
1398b3cdc313        littledoctor_little-doctor-redis    "/usr/bin/redis-se..."   19 seconds ago      Up 18 seconds       6379/tcp                        littledoctor_little-doctor-redis_1
1734021-C02RQ08ZG8WP:~ russellendicott$ docker exec -it 4bc3b0280282 /bin/bash
[root@4bc3b0280282 ansible]# ansible all -m raw -a 'who && hostname'
172.18.0.3 | SUCCESS | rc=0 >>
root     pts/0        Mar 20 20:31 (littledoctor_little-doctor-server_1.littledoctor_default)
a20d150b2148
Warning: Permanently added '172.18.0.3' (ECDSA) to the list of known hosts.
Shared connection to 172.18.0.3 closed.
```

If you want more clients you can spin them up like this in another terminal window from the source folder:
```
docker-compose scale little-doctor-client=8
```

Then jump back into your attached terminal from above and run the ansible command again. This time you'll have 8 hosts respond. 
```
[root@4bc3b0280282 ansible]# ansible all -m raw -a 'who && hostname'
172.18.0.3 | SUCCESS | rc=0 >>
root     pts/0        Mar 20 20:32 (littledoctor_little-doctor-server_1.littledoctor_default)
a20d150b2148
Shared connection to 172.18.0.3 closed.


172.18.0.6 | SUCCESS | rc=0 >>
root     pts/0        Mar 20 20:32 (littledoctor_little-doctor-server_1.littledoctor_default)
a59046a32c90
Warning: Permanently added '172.18.0.6' (ECDSA) to the list of known hosts.
Shared connection to 172.18.0.6 closed.


172.18.0.11 | SUCCESS | rc=0 >>
root     pts/0        Mar 20 20:32 (littledoctor_little-doctor-server_1.littledoctor_default)
03f9a8fb7693
Warning: Permanently added '172.18.0.11' (ECDSA) to the list of known hosts.
Shared connection to 172.18.0.11 closed.


172.18.0.5 | SUCCESS | rc=0 >>
root     pts/0        Mar 20 20:32 (littledoctor_little-doctor-server_1.littledoctor_default)
a1c68783fd43
Warning: Permanently added '172.18.0.5' (ECDSA) to the list of known hosts.
Shared connection to 172.18.0.5 closed.


172.18.0.9 | SUCCESS | rc=0 >>
root     pts/0        Mar 20 20:32 (littledoctor_little-doctor-server_1.littledoctor_default)
9a162dcabea9
Warning: Permanently added '172.18.0.9' (ECDSA) to the list of known hosts.
Shared connection to 172.18.0.9 closed.


172.18.0.7 | SUCCESS | rc=0 >>
root     pts/0        Mar 20 20:32 (littledoctor_little-doctor-server_1.littledoctor_default)
85468c6fa153
Warning: Permanently added '172.18.0.7' (ECDSA) to the list of known hosts.
Shared connection to 172.18.0.7 closed.


172.18.0.8 | SUCCESS | rc=0 >>
root     pts/0        Mar 20 20:32 (littledoctor_little-doctor-server_1.littledoctor_default)
396449f20d38
Warning: Permanently added '172.18.0.8' (ECDSA) to the list of known hosts.
Shared connection to 172.18.0.8 closed.


172.18.0.10 | SUCCESS | rc=0 >>
root     pts/0        Mar 20 20:32 (littledoctor_little-doctor-server_1.littledoctor_default)
c933e5df57bf
Warning: Permanently added '172.18.0.10' (ECDSA) to the list of known hosts.
Shared connection to 172.18.0.10 closed.


[root@4bc3b0280282 ansible]#
```

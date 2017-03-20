# little-doctor

An Ansible server with a client node that can be multiplied to test Ansible against many hosts.

![docker-compose scale ansible](./img/dc-scale.png?raw=true "Scaling docker clients for use with ansible")

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
1734021-C02RQ08ZG8WP:~ russellendicott$ docker exec -it littledoctor_little-doctor-server_1 /bin/bash
[root@990fa11f13ae ansible]# ansible all -m raw -a 'who && hostname'
172.18.0.3 | SUCCESS | rc=0 >>
root     pts/0        Mar 20 20:31 (littledoctor_little-doctor-server_1.littledoctor_default)
a20d150b2148
Warning: Permanently added '172.18.0.3' (ECDSA) to the list of known hosts.
Shared connection to 172.18.0.3 closed.
```

If you want more clients you can spin them up like this in another terminal window from the source folder:
```
docker-compose scale little-doctor-client=2 little-doctor-client-centos=2
```

Then jump back into your attached terminal from above and run the ansible command again. This time you'll have 4 hosts respond. 
```
[root@990fa11f13ae ansible]# ansible all -m raw -a 'python -c "import platform; print platform.platform()"'
172.18.0.6 | SUCCESS | rc=0 >>
Linux-4.9.12-moby-x86_64-with-debian-8.7
Shared connection to 172.18.0.6 closed.


172.18.0.5 | SUCCESS | rc=0 >>
Linux-4.9.12-moby-x86_64-with-debian-8.7
Shared connection to 172.18.0.5 closed.


172.18.0.3 | SUCCESS | rc=0 >>
Linux-4.9.12-moby-x86_64-with-centos-7.2.1511-Core
Shared connection to 172.18.0.3 closed.


172.18.0.7 | SUCCESS | rc=0 >>
Linux-4.9.12-moby-x86_64-with-centos-7.2.1511-Core
Shared connection to 172.18.0.7 closed.

[root@4bc3b0280282 ansible]#
```

There are some sample playbooks in the `/usr/ansible/playbooks` folder on the server. You can run them like so:
```
[root@c0797947481e ansible]# cd /usr/ansible/playbooks
[root@c0797947481e playbooks]# ansible-playbook -e targets=all sample-facts.yml

PLAY [Gathering facts] *********************************************************

TASK [setup] *******************************************************************
ok: [172.18.0.5]
ok: [172.18.0.7]
ok: [172.18.0.3]
ok: [172.18.0.6]

TASK [Show some os info from the facts module] *********************************
ok: [172.18.0.5] => {
    "msg": "I gathered facts and I can see that ansible_os_family = RedHat"
}
ok: [172.18.0.3] => {
    "msg": "I gathered facts and I can see that ansible_os_family = Debian"
}
ok: [172.18.0.6] => {
    "msg": "I gathered facts and I can see that ansible_os_family = Debian"
}
ok: [172.18.0.7] => {
    "msg": "I gathered facts and I can see that ansible_os_family = RedHat"
}

TASK [Show some virtualization info from the facts module] *********************
ok: [172.18.0.5] => {
    "msg": "I gathered facts and I can see that ansible_virtualization_type = docker"
}
ok: [172.18.0.6] => {
    "msg": "I gathered facts and I can see that ansible_virtualization_type = docker"
}
ok: [172.18.0.3] => {
    "msg": "I gathered facts and I can see that ansible_virtualization_type = docker"
}
ok: [172.18.0.7] => {
    "msg": "I gathered facts and I can see that ansible_virtualization_type = docker"
}

PLAY RECAP *********************************************************************
172.18.0.3                 : ok=3    changed=0    unreachable=0    failed=0
172.18.0.5                 : ok=3    changed=0    unreachable=0    failed=0
172.18.0.6                 : ok=3    changed=0    unreachable=0    failed=0
172.18.0.7                 : ok=3    changed=0    unreachable=0    failed=0

```

## Known Issues
* If you run a scale down command after scaling up your ansible inventory will still have the old machines so you'll get failures
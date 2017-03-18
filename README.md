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

If you want more clients you can spin them up like this:
```
docker-compose scale little-doctor-client=3
```
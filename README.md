# little-doctor

An Ansible server with a client node that can be multiplied to test Ansible against many hosts.

## Usage

Build then run like so
```
docker-compose build
docker-compose scale little-doctor-server=1 little-doctor-client=3
```

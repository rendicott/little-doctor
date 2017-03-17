# little-doctor

An Ansible server with a client node that can be multiplied to test Ansible against many hosts.

## Usage

Build then run like so
```
docker-compose build
docker-compose up
```

If you want more clients you can spin them up like this:
```
docker-compose scale little-doctor-client=3
```
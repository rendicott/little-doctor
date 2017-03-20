#!/bin/bash

/tmp/post-ip-to-redis redserver 6379
/usr/sbin/sshd
/app/main
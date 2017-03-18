#!/bin/bash

/tmp/post-ip-to-redis redserver 6379
/etc/init.d/ssh start
/app/main
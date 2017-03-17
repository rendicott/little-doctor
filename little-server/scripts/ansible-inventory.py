import redis
import json

# Sample JSON output for ansible inventory
# what we're building
"""
{
"10.1.1.2" : {"hosts": ["10.1.1.2"]},
"10.1.1.3" : {"hosts": ["10.1.1.3"]}
}
"""

dhosts = {}
r = redis.StrictRedis(host='redserver', port=6379, db=0)
ipaddrs = r.lrange('ipaddrs', 0, -1)
for ipa in ipaddrs:
    d = {"hosts": [ipa]}
    dhosts[ipa] = d

print json.dumps(dhosts)


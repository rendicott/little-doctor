/* post-ip-to-redis.go
Finds the local ip address for eth0 and posts it to a redis instance
given as param1 and port given as param2. It pushes it to the 
'ipaddrs' list in redis on DB0. 

Usage:
  go build post-ip-to-redis.go
  ./post-ip-to-redis 10.8.1.3 6379
*/

package main

import "gopkg.in/redis.v5"
import "net"
import "fmt"
import "os"


// getIp grabs the ipv4 address of the eth0 adapter and returns
// it as a string. Shamefully stolen from here:
// URL: https://github.com/mccoyst/myip/blob/master/myip.go
// URL: http://changsijay.com/2013/07/28/golang-get-ip-address/
func getIp() string {
	var istring string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				istring = ipnet.IP.String()
				os.Stdout.WriteString(istring + "\n")
			}
		}
	}
	return istring
}

// redisPost takes an ip address as a string and posts it to the
// redis server and port given as parameters two and three.
// Inspired by the 'gopkg.in/redis.v5' README
func redisPost(ip, rip, rport string) {
	// first setup the redis client info
	client := redis.NewClient(&redis.Options{
        Addr:     rip + ":" + rport,
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    pong, err1 := client.Ping().Result()
    fmt.Println(pong, err1)
    // Output: PONG <nil>
    fmt.Println("IP Addr" + ip)
	// push the ipaddress to the 'ipaddrs' list with Lpush
    err2 := client.LPush("ipaddrs", ip).Err()
    if err2 != nil {
        panic(err2)
    }
}

// main takes first two args from cmdline and spits some
// messages to stdout.
func main() {
	redisIp := os.Args[1]
	redisPort := os.Args[2]
	fmt.Println("Pushing to " + redisIp)
	ip := getIp()
	fmt.Println("IP ADDR IS : " + ip)
    redisPost(ip, redisIp, redisPort)
	}
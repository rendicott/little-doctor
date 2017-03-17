package main

import "gopkg.in/redis.v5"
import "net"
import "fmt"
import "os"

/*
URL: https://github.com/mccoyst/myip/blob/master/myip.go
URL: http://changsijay.com/2013/07/28/golang-get-ip-address/
*/

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

func redisPost(ip, rip, rport string) {
	client := redis.NewClient(&redis.Options{
        Addr:     rip + ":" rport,
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    pong, err1 := client.Ping().Result()
    fmt.Println(pong, err1)
    // Output: PONG <nil>
    fmt.Println("IP Addr" + ip)
    err2 := client.LPush("ipaddrs", ip).Err()
    if err2 != nil {
        panic(err2)
    }
}


func main() {
	redisIp := os.Args[1]
	redisPort := os.Args[2]
	ip := getIp()
	fmt.Println("IP ADDR IS : " + ip)
    redisPost(ip)
	}
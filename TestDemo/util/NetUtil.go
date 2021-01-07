package util

import (
	"net"
	"strconv"
	"strings"
)

func IPV4() (ip string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, value := range addrs {
		if ipnet, ok := value.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				return ip
			}
		}
	}
	return ""
}

func IPStr2Int(ipAddr string) (ipInt int64) {
	ip := strings.Replace(ipAddr, ".", "", -1)
	if ipD, err := strconv.ParseInt(ip, 10, 64); err != nil {
		panic(err)
	} else {
		ipInt = ipD
	}
	return ipInt
}

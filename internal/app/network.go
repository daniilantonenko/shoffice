package app

import (
	"fmt"
	"net"
)

func GetIp() []string {
	// get list of available addresses
	addr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var arr []string

	for _, addr := range addr {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			// check if IPv4 or IPv6 is not nil
			if ipnet.IP.To4() != nil || ipnet.IP.To16() == nil {
				// print available addresses
				arr = append(arr, ipnet.IP.String())
			}
		}
	}
	// TODO: return []string
	return arr
}

func IsIPv4(addr string) bool {
	trial := net.ParseIP(addr)
	return trial.To4() != nil
}

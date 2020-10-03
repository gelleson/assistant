package host

import (
	"net"
	"strings"
)

func GetIPAddress() string {
	return getLocalIP()
}

func getLocalIP() string {
	addressArr, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addressArr {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil && strings.Contains(ipnet.IP.String(), "192") {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

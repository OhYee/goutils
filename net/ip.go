package net

import (
	"errors"
	gnet "net"
	"strconv"
	"strings"
)

// ParseIPv4Port parse x.x.x.x:x to net.IP and uint16
func ParseIPv4Port(address string) (ip gnet.IP, port uint16, err error) {
	colonPos := strings.Index(address, ":")
	if colonPos == -1 {
		err = errors.New("Address must be formated as x.x.x.x:x")
		return
	}
	ip = gnet.ParseIP(address[:colonPos])
	if len(ip) != 16 {
		err = errors.New("Can not parse ip to net.IP")
		return
	}
	ip = ip[12:]
	_port, err := strconv.Atoi(address[colonPos+1:])
	if err != nil {
		err = errors.New("Can not parse port to int")
		return
	}
	if _port > 65535 || _port < 0 {
		err = errors.New("Port must between 0~65535")
	}
	port = uint16(_port)

	return
}

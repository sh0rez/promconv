package server

import (
	"net"
	"strings"
)

func AddrPort(netaddr net.Addr) (AttrAddress, AttrPort) {
	addr, port, ok := strings.Cut(netaddr.String(), ":")
	if !ok {
		return "", ""
	}
	return AttrAddress(addr), AttrPort(port)
}

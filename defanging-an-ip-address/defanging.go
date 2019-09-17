package defanging

import (
	"net"
	"strings"
)

func defangIPaddr(address string) string {
	if !isIPValid(address) {
		return ""
	}
	var d string
	for _, c := range address {
		if c == '.' {
			d += "[.]"
		} else {
			d += string(c)
		}
	}
	return d
}

func defangIPaddrStrings(address string) string {
	if !isIPValid(address) {
		return ""
	}
	return strings.ReplaceAll(address, ".", "[.]")
}

func isIPValid(address string) bool {
	if parsed := net.ParseIP(address); parsed != nil {
		return true
	}
	return false
}

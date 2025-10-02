package utils

import (
	"net"
	"regexp"
)

func IsIP(s string) bool {
	return net.ParseIP(s) != nil
}

func IsHostname(addr string) bool {
	hostnameRegex := regexp.MustCompile(`^(?i:[a-z0-9](?:[a-z0-9\-]{0,61}[a-z0-9])?)(?:\.(?i:[a-z0-9](?:[a-z0-9\-]{0,61}[a-z0-9])?))*$`)
	return len(addr) <= 253 && hostnameRegex.MatchString(addr)
}

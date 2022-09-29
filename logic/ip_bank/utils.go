package ip_bank

import (
	"fmt"
	"net"
)

func ToLocation(lat, lng float64) string {
	return fmt.Sprintf("[%f,%f]", lat, lng)
}

func IPVersion(s string) int {
	ip := net.ParseIP(s)
	if ip == nil {
		return 0
	}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '.':
			return 4
		case ':':
			return 6
		}
	}
	return 0
}

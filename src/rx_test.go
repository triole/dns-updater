package main

import "testing"

var (
	validIPv4Arr = []string{
		"127.0.0.1", "192.168.0.33",
	}
	validIPv6Arr = []string{
		"FE80:0000:0000:0000:0202:B3FF:FE1E:8329",
		"2345:425:2CA1:0000:0000:567:5673:23b5/64",
	}
)

func TestIsIPv4(t *testing.T) {
	for _, ip := range validIPv4Arr {
		validateIsIPv(4, ip, true, t)
	}
	for _, ip := range validIPv4Arr {
		validateIsIPv(6, ip, false, t)
	}
	for _, ip := range validIPv6Arr {
		validateIsIPv(6, ip, true, t)
	}
	for _, ip := range validIPv6Arr {
		validateIsIPv(4, ip, false, t)
	}
}

func validateIsIPv(ver int8, ip string, exp bool, t *testing.T) {
	var res bool
	if ver == 4 {
		res = isIPv4(ip)
	}
	if ver == 6 {
		res = isIPv6(ip)
	}
	if res != exp {
		t.Errorf("is valid ipv%v check failed: %s, expected %v", ver, ip, exp)
	}
}

func TestIsValidIPAddress(t *testing.T) {
	arr := append(validIPv4Arr, validIPv6Arr...)
	for _, ip := range arr {
		if !isValidIP(ip) {
			t.Errorf("is valid ip test failed, expected true for %s", ip)
		}
	}
}

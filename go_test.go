package main

import "testing"

func TestV6Make(t *testing.T) {
	res, err := randomIPv6AddressFromSubnet("2001:470:8a97::/48")
	if err != nil {
		t.Error("错误")
	}
	print(res)
}

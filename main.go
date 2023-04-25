package main

import (
	"flag"
	"fmt"
	reverse "kproxy/reverse"
	"math/rand"
	"net"
	"net/http"
	"net/url"
)

func randomIPv6AddressFromSubnet(subnet string) (net.IP, error) {
	_, ipNet, err := net.ParseCIDR(subnet)
	if err != nil {
		return nil, err
	}

	maskSize, _ := ipNet.Mask.Size()
	randBits := 128 - maskSize

	randomBytes := make([]byte, randBits/8)
	_, err = rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	ip := make(net.IP, len(ipNet.IP))
	copy(ip, ipNet.IP)

	for i := 0; i < len(randomBytes); i++ {
		byteIdx := (randBits - 1 - (i * 8)) / 8
		ip[15-byteIdx] |= randomBytes[i]
	}

	return ip, nil
}

func main() {
	bindAddr := flag.String("b", ":8080", "Bind IP:Port")
	ipv6Subnet := flag.String("i", "", "IPv6 subnet")
	flag.Parse()

	if *ipv6Subnet == "" {
		fmt.Println("IPv6 subnet must be provided.")
		return
	}

	http.ListenAndServe(*bindAddr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path, err := url.Parse("http://" + r.Host)
		if err != nil {
			panic(err)
			return
		}
		proxy := reverse.NewReverseProxy(path)
		ip, err := randomIPv6AddressFromSubnet(*ipv6Subnet)
		proxy.ServeHTTP(w, r, ip)
	}))

}

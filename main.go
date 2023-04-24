package main

import (
	"flag"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/url"
	"strings"
	"time"
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
func handleRequest(clientConn net.Conn, ipv6Subnet string) {
	defer clientConn.Close()

	var buffer [4096]byte
	n, err := clientConn.Read(buffer[:])
	if err != nil {
		fmt.Println("Error reading from client:", err)
		return
	}

	requestLine := strings.Split(string(buffer[:n]), "\r\n")[0]
	requestParts := strings.Split(requestLine, " ")
	if len(requestParts) != 3 {
		fmt.Println("Invalid request line:", requestLine)
		return
	}

	rawURL := requestParts[1]
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	host := parsedURL.Host
	if !strings.Contains(host, ":") {
		defaultPort := "80"
		if parsedURL.Scheme == "https" {
			defaultPort = "443"
		}
		host = host + ":" + defaultPort
	}

	randomIPv6, err := randomIPv6AddressFromSubnet(ipv6Subnet)
	if err != nil {
		fmt.Println("Error generating random IPv6 address:", err)
		return
	}

	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		LocalAddr: &net.TCPAddr{
			IP:   randomIPv6,
			Port: 0,
		},
	}

	targetConn, err := dialer.Dial("tcp", host)
	if err != nil {
		fmt.Println("Error connecting to target server:", err)
		return
	}
	defer targetConn.Close()

	targetConn.Write(buffer[:n])

	go func() {
		io.Copy(targetConn, clientConn)
	}()

	io.Copy(clientConn, targetConn)
}

func main() {
	bindAddr := flag.String("b", ":8080", "Bind IP:Port")
	ipv6Subnet := flag.String("i", "", "IPv6 subnet")
	flag.Parse()

	if *ipv6Subnet == "" {
		fmt.Println("IPv6 subnet must be provided.")
		return
	}

	listener, err := net.Listen("tcp", *bindAddr)
	if err != nil {
		fmt.Println("Error starting proxy server:", err)
		return
	}

	fmt.Printf("Proxy server started on %s\n", *bindAddr)
	for {
		clientConn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting client connection:", err)
			continue
		}

		go handleRequest(clientConn, *ipv6Subnet)
	}
}

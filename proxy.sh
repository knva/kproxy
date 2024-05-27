apt install ndppd
# 获取IPv6地址并提取前四个前缀
result=$(curl -s6 "http://6.ipw.cn/" 2>&1 | awk -F: '{print $1 ":" $2 ":" $3 ":" $4}')
ipv6_subnet="${result}::/64"
cat > /etc/ndppd.conf << EOF
route-ttl 30000
proxy enp1s0 {
  router yes
  timeout 500
  ttl 30000
  rule $ipv6_subnet {
    static
  }
}
EOF

sysctl net.ipv6.ip_nonlocal_bind=1

ip route add local  "$ipv6_subnet" dev enp1s0
ndppd &
#rm kproxy
#./kproxy -b 127.0.0.1:10808 -i "$ipv6_subnet" &

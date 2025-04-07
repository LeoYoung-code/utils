package addr

import (
	"fmt"
	"net"
)

// privateBlocks 存储私有网络IP范围
var privateBlocks []*net.IPNet

func init() {
	// 初始化私有网络IP范围
	cidrs := []string{
		"10.0.0.0/8",     // RFC1918
		"172.16.0.0/12",  // RFC1918
		"192.168.0.0/16", // RFC1918
		"100.64.0.0/10",  // RFC6598 CGNAT
		"fd00::/8",       // RFC4193 本地IPv6单播地址
	}

	for _, cidr := range cidrs {
		_, block, err := net.ParseCIDR(cidr)
		if err == nil {
			privateBlocks = append(privateBlocks, block)
		}
	}
}

// IsPrivateIP 检查IP是否为私有IP
func IsPrivateIP(ip net.IP) bool {
	if ip == nil {
		return false
	}

	for _, block := range privateBlocks {
		if block.Contains(ip) {
			return true
		}
	}
	return false
}

// IsPrivateIPString 检查IP字符串是否为私有IP
func IsPrivateIPString(ipAddr string) bool {
	ip := net.ParseIP(ipAddr)
	return IsPrivateIP(ip)
}

// ExtractTCPAddr 从TCP监听器中提取IP地址和端口
func ExtractTCPAddr(listener net.Listener) (string, int, error) {
	tcpAddr, ok := listener.Addr().(*net.TCPAddr)
	if !ok {
		return "", 0, fmt.Errorf("无法将地址转换为TCPAddr")
	}

	host := tcpAddr.IP.String()
	port := tcpAddr.Port

	ipAddr, err := Extract(host)
	if err != nil {
		return "", 0, err
	}

	return ipAddr, port, nil
}

// Extract 返回一个真实IP地址
// 如果提供了非0.0.0.0的地址，则返回该地址
// 否则尝试从网络接口获取合适的IP地址
func Extract(addr string) (string, error) {
	// 如果提供了有效的非通配地址，直接返回
	if len(addr) > 0 && addr != "0.0.0.0" && addr != "[::]" && addr != "::" {
		return addr, nil
	}

	// 获取所有网络接口
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", fmt.Errorf("获取网络接口失败: %v", err)
	}

	// 收集所有地址，将环回地址放在最后
	var addrs []net.Addr
	var loAddrs []net.Addr

	for _, iface := range ifaces {
		ifaceAddrs, err := iface.Addrs()
		if err != nil {
			// 忽略错误，接口可能已从系统中消失
			continue
		}

		if iface.Flags&net.FlagLoopback != 0 {
			loAddrs = append(loAddrs, ifaceAddrs...)
			continue
		}

		addrs = append(addrs, ifaceAddrs...)
	}

	// 将环回地址附加到最后
	addrs = append(addrs, loAddrs...)

	return extractIPFromAddrs(addrs)
}

// extractIPFromAddrs 从地址列表中提取合适的IP地址
func extractIPFromAddrs(addrs []net.Addr) (string, error) {
	var privateIP net.IP
	var publicIP net.IP

	for _, rawAddr := range addrs {
		var ip net.IP
		switch addr := rawAddr.(type) {
		case *net.IPAddr:
			ip = addr.IP
		case *net.IPNet:
			ip = addr.IP
		default:
			continue
		}

		if ip == nil {
			continue
		}

		if !IsPrivateIP(ip) {
			publicIP = ip
			continue
		}

		privateIP = ip
		break
	}

	// 优先返回私有IP
	if privateIP != nil {
		return privateIP.String(), nil
	}

	// 其次返回公网IP
	if publicIP != nil {
		return publicIP.String(), nil
	}

	return "", fmt.Errorf("未找到可用IP地址，且未提供明确的IP")
}

// IPs 返回所有已知的IP地址
func IPs() []string {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil
	}

	var ipAddrs []string

	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip == nil {
				continue
			}

			ipAddrs = append(ipAddrs, ip.String())
		}
	}

	return ipAddrs
}

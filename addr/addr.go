package addr

import (
	"fmt"
	"net"
	"strings"
)

// privateBlocks 存储私有网络IP范围
var privateBlocks []*net.IPNet

// 通配IP地址列表
var wildcardIPs = map[string]bool{
	"":        true,
	"0.0.0.0": true,
	"[::]":    true,
	"::":      true,
}

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
// 返回true表示是私有IP，false表示是公网IP或无效IP
func IsPrivateIP(ip net.IP) bool {
	if ip == nil {
		return false
	}

	if ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() {
		return true
	}

	for _, block := range privateBlocks {
		if block.Contains(ip) {
			return true
		}
	}
	return false
}

// IsPrivateIPString 检查IP字符串是否为私有IP
// 接受IP字符串作为输入，返回true表示是私有IP，false表示是公网IP或无效的IP字符串
func IsPrivateIPString(ipAddr string) bool {
	ip := net.ParseIP(ipAddr)
	return IsPrivateIP(ip)
}

// ExtractTCPAddr 从TCP监听器中提取IP地址和端口
// 返回提取的IP地址和端口号，如果提取失败则返回错误
func ExtractTCPAddr(listener net.Listener) (string, int, error) {
	if listener == nil {
		return "", 0, fmt.Errorf("监听器不能为空")
	}
	
	tcpAddr, ok := listener.Addr().(*net.TCPAddr)
	if !ok {
		return "", 0, fmt.Errorf("无法将地址转换为TCPAddr: %s", listener.Addr().String())
	}

	host := tcpAddr.IP.String()
	port := tcpAddr.Port

	ipAddr, err := Extract(host)
	if err != nil {
		return "", 0, fmt.Errorf("提取IP地址失败: %w", err)
	}

	return ipAddr, port, nil
}

// Extract 返回一个真实IP地址
// 如果提供了非通配地址，则返回该地址
// 否则尝试从网络接口获取合适的IP地址
func Extract(addr string) (string, error) {
	// 如果提供了有效的非通配地址，直接返回
	if !isWildcardIP(addr) {
		return addr, nil
	}

	// 获取所有网络接口
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", fmt.Errorf("获取网络接口失败: %w", err)
	}

	addrs, err := getInterfaceAddresses(ifaces)
	if err != nil {
		return "", err
	}

	return extractIPFromAddrs(addrs)
}

// isWildcardIP 检查地址是否为通配IP
func isWildcardIP(addr string) bool {
	return wildcardIPs[strings.TrimSpace(addr)]
}

// getInterfaceAddresses 获取所有网络接口的地址
// 返回两个切片：常规地址和环回地址
func getInterfaceAddresses(ifaces []net.Interface) ([]net.Addr, error) {
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
	return append(addrs, loAddrs...), nil
}

// extractIPFromAddrs 从地址列表中提取合适的IP地址
// 优先返回私有IP，其次是公网IP
func extractIPFromAddrs(addrs []net.Addr) (string, error) {
	var privateIP net.IP
	var publicIP net.IP

	for _, rawAddr := range addrs {
		ip := extractIPFromAddr(rawAddr)
		if ip == nil {
			continue
		}

		if IsPrivateIP(ip) {
			privateIP = ip
			break // 找到私有IP就停止
		} else {
			publicIP = ip
		}
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

// extractIPFromAddr 从单个地址对象中提取IP
func extractIPFromAddr(addr net.Addr) net.IP {
	switch v := addr.(type) {
	case *net.IPAddr:
		return v.IP
	case *net.IPNet:
		return v.IP
	default:
		return nil
	}
}

// IPs 返回所有已知的IP地址
// 包括所有网络接口的所有IP地址
func IPs() []string {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil
	}

	addrs, err := getInterfaceAddresses(ifaces)
	if err != nil {
		return nil
	}

	var ipAddrs []string
	for _, addr := range addrs {
		ip := extractIPFromAddr(addr)
		if ip != nil {
			ipAddrs = append(ipAddrs, ip.String())
		}
	}

	return ipAddrs
}

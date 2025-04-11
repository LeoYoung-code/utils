package http

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"utils/json"
)

// APICompareConfig 包含API比较的配置参数
type APICompareConfig struct {
	// 测试接口配置
	APIs map[string]string
	// 测试环境配置
	Environments map[string]Environment
	// 请求头配置
	Headers map[string]string
}

// Environment 环境配置
type Environment struct {
	EnvID    string // 环境标识
	CustomIP string // 自定义IP
}

// APIComparer API比较工具
type APIComparer struct {
	Config APICompareConfig
}

// NewAPIComparer 创建新的API比较器
func NewAPIComparer() *APIComparer {
	return &APIComparer{
		Config: DefaultConfig(),
	}
}

// DefaultConfig 返回默认配置
func DefaultConfig() APICompareConfig {
	return APICompareConfig{
		APIs: map[string]string{
			"XX接口": "/v1/index",
		},
		Environments: map[string]Environment{
			"sampleA": {EnvID: "1", CustomIP: "60.XXX.XXX"},
			"sampleB": {EnvID: "2", CustomIP: "60.XXX.XXX"},
		},
		Headers: map[string]string{
			"sign": "7937cad538b770da3c8532c570b387ca",
		},
	}
}

// TestAPI 测试所有API
func (a *APIComparer) TestAPI(domain string) {
	for title, path := range a.Config.APIs {
		a.printDivider(title, true)
		a.TestURL(path, domain)
		a.printDivider(title, false)
	}
}

// 打印分隔符
func (a *APIComparer) printDivider(title string, isBegin bool) {
	status := "Begin"
	if !isBegin {
		status = "End"
	}
	fmt.Printf("======================%s---%s=====================\n", title, status)
	if !isBegin {
		fmt.Println("\n")
	}
}

// TestURL 测试单个URL
func (a *APIComparer) TestURL(path, domain string) {
	fullURL := domain + path

	// 保存原始headers的副本
	headersCopy := make(map[string]string)
	for k, v := range a.Config.Headers {
		headersCopy[k] = v
	}

	// 测试环境A
	envA := a.Config.Environments["sampleA"]
	headersCopy["test-env"] = envA.EnvID
	responseA, err := a.makeRequest(fullURL, headersCopy, envA.CustomIP)
	if err != nil {
		fmt.Printf("测试环境%s请求错误: %v \n", envA.EnvID, err)
		return
	}
	fmt.Printf("测试环境%s响应: \n%v \n\n", envA.EnvID, responseA)

	// 测试环境B
	envB := a.Config.Environments["sampleB"]
	headersCopy["test-env"] = envB.EnvID
	responseB, err := a.makeRequest(fullURL, headersCopy, envB.CustomIP)
	if err != nil {
		fmt.Printf("测试环境%s请求错误: %v \n", envB.EnvID, err)
		return
	}
	fmt.Printf("测试环境%s响应: \n%v \n\n", envB.EnvID, responseB)

	// 比较结果
	a.compareResponses(responseA, responseB)
}

// 比较两个响应
func (a *APIComparer) compareResponses(responseA, responseB string) {
	differences := json.CompareJSON(responseA, responseB)
	if responseA == responseB || len(differences) == 0 {
		fmt.Println("两个环境的响应结果 【 相同 】")
	} else {
		fmt.Println("两个环境的响应结果【 不相同 】 ")
		fmt.Println("Differences:", differences)
	}
}

// 自定义 DialContext，用于指定 IP 地址访问域名
func customDialContext(ip string) func(ctx context.Context, network, addr string) (net.Conn, error) {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		// 提取主机名和端口
		_, port, err := net.SplitHostPort(addr)
		if err != nil {
			// 如果没有端口，假设是HTTPS使用443
			if strings.Contains(err.Error(), "missing port") {
				port = "443"
			} else {
				return nil, err
			}
		}

		// 使用指定的IP替换主机名
		return (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext(ctx, network, net.JoinHostPort(ip, port))
	}
}

// makeRequest 发送HTTP请求并返回响应
func (a *APIComparer) makeRequest(url string, headers map[string]string, ip string) (string, error) {
	// 创建一个自定义的 HTTP 客户端，禁用证书验证，并使用自定义的 DialContext
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		DialContext:     customDialContext(ip),
	}
	client := &http.Client{
		Timeout:   10 * time.Second,
		Transport: transport,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("创建请求失败: %v", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("请求失败: %v", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("关闭响应体失败: %v\n", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("非200状态码: %d %s", resp.StatusCode, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应失败: %v", err)
	}

	return string(body), nil
}

// RunTest 执行测试的入口函数
func RunTest(domain string) {
	comparer := NewAPIComparer()
	comparer.TestAPI(domain)
}

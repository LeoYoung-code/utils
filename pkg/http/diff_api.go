package http

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"utils/json"
)

// 测试参数构造
var urls = []map[string]string{
	{"XX接口": "/v1/index"},
}

var testEnv = map[string]string{
	"sampleA":   "1",                   // 测试环境 1
	"sampleB":   "2",                   // 测试环境 2
	"customIpA": "60.XXX.XXX",          // 自定义IP 1
	"customIpB": "60.XXX.XXX",          // 自定义IP 2
	"domain":    "https://XXX.XXX.XXX", // 自定义域名
}

var headers = map[string]string{
	"sign": "7937cad538b770da3c8532c570b387ca",
}

func testApi() {
	for _, urlM := range urls {
		for testTitle, url := range urlM {
			fmt.Printf("======================%s---Begin=====================\n", testTitle)
			testUrl(url)
			fmt.Printf("======================%s---End=====================\n\n\n", testTitle)
		}
	}
}

func testUrl(url string) {
	domain := testEnv["domain"]
	// 发起对测试环境1的请求
	headers["test-env"] = testEnv["sampleA"]
	responseTest, err := makeRequest(domain+url, headers, testEnv["customIpA"])
	if err != nil {
		fmt.Printf("测试环境%s请求错误: %v \n", headers["test-env"], err)
		return
	}
	fmt.Printf("测试环境%s响应: \n", headers["test-env"])
	fmt.Printf("%v \n\n", responseTest)

	// 发起对测试环境2的请求
	headers["test-env"] = testEnv["sampleB"]
	responseProd, err := makeRequest(domain+url, headers, testEnv["customIpB"])
	if err != nil {
		fmt.Printf("测试环境%s请求错误: %v \n", headers["test-env"], err)
		return
	}
	fmt.Printf("测试环境%s响应: \n", headers["test-env"])
	fmt.Printf("%v \n\n", responseProd)

	// 对比响应结果
	differences := json.CompareJSON(responseTest, responseProd)
	if responseTest == responseProd || len(differences) == 0 {
		fmt.Println("两个环境的响应结果 【 相同 】")
	} else {
		fmt.Println("两个环境的响应结果【 不相同 】 ")
		fmt.Println("Differences:", differences)
	}
}

// 自定义 DialContext，用于指定 IP 地址访问域名
func customDialContext(ip string) func(ctx context.Context, network, addr string) (net.Conn, error) {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		// 将域名替换为指定的 IP 地址
		return (&net.Dialer{}).DialContext(ctx, network, ip+":443")
	}
}

func makeRequest(url string, headers map[string]string, ip string) (string, error) {
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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("请求失败: ", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("非200状态码: %d %s", resp.StatusCode, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应失败: %v", err)
	}

	return string(body), nil
}

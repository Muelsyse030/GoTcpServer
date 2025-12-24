package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// 发起 GET 请求
	resp, err := http.Get("http://goproxy.cn") // 我们可以请求刚才配置的代理站看看
	if err != nil {
		fmt.Println(err)
		return
	}
	// ⚠️ 非常重要：必须关闭 Body，否则会内存/句柄泄露
	defer resp.Body.Close()

	// 读取响应内容
	body, _ := io.ReadAll(resp.Body)

	// 打印状态码和部分内容
	fmt.Println("Status:", resp.Status)
	fmt.Printf("Body length: %d bytes\n", len(body))
}

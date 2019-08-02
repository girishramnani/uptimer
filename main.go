package main

import (
	"fmt"

	"github.com/girishramnani/uptimer/pkg/req"
)

func main() {
	urlList := []string{
		"https://google.com",
		"https://youtube.com",
		"https://facebook.com",
		"https://baidu.com",
		"https://wikipedia.org",
		"https://taobao.com",
		"https://yahoo.com",
		"https://tmall.com",
		"https://amazon.com",
		"https://twitter.com",
		"https://live.com",
	}

	resps := req.GetAllUrls(urlList)

	for resp := range resps {
		fmt.Println(resp.RespCode, resp.URL)
	}

}

package services

import (
	"paper/quote-client/utils"
	"time"
)

// 按照输入数字，作为每分钟请求数，向服务器发请求
const (
	URL = "http://0.0.0.0:8000"
)

func SendRequestPerMin(reqs int, user string) {
	//一分钟内，向URL发reqs个请求; 每次发送reqs/60个请求，然后sleep 1s，直至这分钟发完reqs个请求
	//每次发送reqs/60个请求，每个请求间隔1s
	for i := 0; i < 60; i++ {
		if i == 59 {
			go SendRequestOneTime(URL, reqs-reqs/60*59, user)
			break
		}
		go SendRequestOneTime(URL, reqs/60, user)
		time.Sleep(1 * time.Second)
	}
}

func SendRequestOneTime(url string, reqs int, user string) {
	for i := 0; i < reqs/60; i++ {
		go SendRequest(url, user)
	}
}

func SendRequest(url, user string) {
	header := make(map[string]string)
	header[user] = user
	//发送Get请求
	utils.Get(url, header)
}

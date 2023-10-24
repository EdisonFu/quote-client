package services

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"paper/quote-client/utils"
	"strconv"
	"time"
)

// 按照输入数字，作为每分钟请求数，向服务器发请求
const (
	URL = "http://"
)

func SendRequestPerMin(reqs int) {
	//一分钟内，向URL发reqs个请求; 每次发送reqs/60个请求，然后sleep 1s，直至这分钟发完reqs个请求
	//每次发送reqs/60个请求，每个请求间隔1s
	for i := 0; i < 60; i++ {
		if i == 59 {
			go SendRequestOneTime(URL, reqs-reqs/60*59)
			break
		}
		go SendRequestOneTime(URL, reqs/60)
		time.Sleep(10 * time.Second)
	}
}

func SendRequestOneTime(url string, reqs int) {
	for i := 0; i < reqs; i++ {
		go SendRequest(url)
	}
}

func SendRequest(url string) {
	//发送Get请求
	utils.Get(url)
}

// 读取reqs.csv文件，获取每分钟请求数,并根据reqs列的数值，返回一个数组
func GetReqsPerMin() []int {
	// 打开 reqs.csv 文件
	file, err := os.Open("reqs.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 创建 CSV Reader
	reader := csv.NewReader(file)

	// 读取 CSV 数据行
	var reqs []int
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// 解析 reqs 列的数值
		req, err := strconv.Atoi(record[1])
		if err != nil {
			log.Fatal(err)
		}
		reqs = append(reqs, req)
	}
	return reqs
}

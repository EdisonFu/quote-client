package services

import (
	"encoding/csv"
	l4g "github.com/alecthomas/log4go"
	"io"
	"os"
	"strconv"
)

// 读取reqs.csv文件，获取每分钟请求数,并根据reqs列的数值，返回一个数组
func GetReqsPerMin() []int {
	// 打开 reqs.csv 文件
	file, err := os.Open("reqs.csv")
	if err != nil {
		l4g.Error("Error opening file: ", err)
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
			l4g.Error("Error reading record: ", err)
		}

		// 解析 reqs 列的数值
		req, err := strconv.Atoi(record[1])
		if err != nil {
			l4g.Error("Error parsing reqs: ", err)
			continue
		}
		reqs = append(reqs, req)
	}
	return reqs
}

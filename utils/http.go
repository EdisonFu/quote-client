package utils

import (
	"bytes"
	"encoding/json"
	l4g "github.com/alecthomas/log4go"
	"io/ioutil"
	"net/http"
	"time"
)

func Get(url string) []byte {
	client := &http.Client{}
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	//增加header选项
	reqest.Header.Add("Content-Type", "application/json")

	if err != nil {
		panic(err)
	}

	//处理返回结果
	var response *http.Response
	client.Timeout = time.Second * 10
	response, err = client.Do(reqest)
	if err != nil || response == nil {
		l4g.Error("Do put url:%s err:%s, Retry again", url, err.Error())
		time.Sleep(time.Second * 1)
		response, err = client.Do(reqest)
		if err != nil || response == nil {
			l4g.Error("Do put url:%s err:%s\n", url, err.Error())
			return nil
		}
	}
	defer response.Body.Close()

	buff, err := ioutil.ReadAll(response.Body)
	if err != nil {
		l4g.Error("Get read body err:%v", err)
	}
	return buff
}

func Post(url, token string, data interface{}) []byte {
	bin, err := json.Marshal(data)
	if err != nil {
		l4g.Error("post marshal req err:", err)
		return nil
	}

	client := &http.Client{}
	//提交请求
	reqest, err := http.NewRequest("POST", url, bytes.NewReader(bin))

	//增加header选项
	reqest.Header.Add("Content-Type", "application/json")
	reqest.Header.Add("token", token)

	if err != nil {
		panic(err)
	}
	//处理返回结果
	client.Timeout = time.Second * 10
	response, err := client.Do(reqest)
	if err != nil || response == nil {
		l4g.Error("Do POST url:%s err:%s\n", url, err.Error())
		time.Sleep(time.Second * 3)
		response, err = client.Do(reqest)
		if err != nil || response == nil {
			l4g.Error("Do POST url:%s err:%s\n", url, err.Error())
			return nil
		}
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		l4g.Error("post response err:", response.StatusCode)
	}

	buff, err := ioutil.ReadAll(response.Body)
	if err != nil {
		l4g.Error("Post read body err:%v", err)
	}
	return buff
}

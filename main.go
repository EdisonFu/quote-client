package main

import (
	l4g "github.com/alecthomas/log4go"
	"paper/quote-client/services"
)

func main() {
	l4g.LoadConfiguration("./log4go.xml")

	reqsList := services.GetReqsPerMin()

	for _, reqs := range reqsList {
		//每分钟发reqs个请求
		services.SendRequestPerMin(reqs)
	}
}

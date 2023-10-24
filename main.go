package main

import (
	"paper/quote-client/services"
)

func main() {
	reqsList := services.GetReqsPerMin()
	for _, reqs := range reqsList {
		//每分钟发reqs个请求
		services.SendRequestPerMin(reqs)
	}
}

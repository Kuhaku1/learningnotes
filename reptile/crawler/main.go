package main

import (
	"learningnotes/reptile/crawler/bilibili"
	"learningnotes/reptile/crawler/engine"
	"learningnotes/reptile/crawler/scheduler"
)

func main() {
	var seed []engine.Request

	seed = []engine.Request{
		{
			Url:   "http://www.bilibili.com",
			Parse: engine.NewFuncParser(bilibili.Getbangumi, "bilibili"),
		},
	}
	// e := engine.SimpleEngine{} //单机
	e := engine.ConcurrentEngine{
		MaxWorkerCount: 2,
		Scheduler:      &scheduler.SimpleScheduler{},
		RequestWorker:  engine.Worker,
	}
	e.Run(seed...)
}

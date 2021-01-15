package main

import (
	"concurrent-crawler/engine"
	"concurrent-crawler/scheduler"
	"concurrent-crawler/zhenai/parser"
)

func main() {
	concurrentEngine := engine.ConcurrentEngine{Scheduler: &scheduler.SimpleScheduler{}, WorkerCount: 10}
	concurrentEngine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}

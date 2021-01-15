package main

import (
	"concurrent-crawler/engine"
	"concurrent-crawler/scheduler"
	"concurrent-crawler/zhenai/parser"
)

func main() {
	concurrentEngine := engine.ConcurrentEngine{Scheduler: &scheduler.QueuedScheduler{}, WorkerCount: 100}
	concurrentEngine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}

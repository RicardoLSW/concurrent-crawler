package main

import (
	"concurrent-crawler/engine"
	"concurrent-crawler/persist"
	"concurrent-crawler/scheduler"
	"concurrent-crawler/zhenai/parser"
)

func main() {
	concurrentEngine := engine.ConcurrentEngine{Scheduler: &scheduler.QueuedScheduler{}, WorkerCount: 100, ItemChan: persist.ItemSaver()}
	//concurrentEngine.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})
	concurrentEngine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})
}

package main

import (
	"concurrent-crawler/engine"
	"concurrent-crawler/persist"
	"concurrent-crawler/scheduler"
	"concurrent-crawler/zhenai/parser"
)

func main() {
	itemSaver, err := persist.ItemSaver()
	if err != nil {
		panic(err)
	}
	concurrentEngine := engine.ConcurrentEngine{Scheduler: &scheduler.QueuedScheduler{}, WorkerCount: 10, ItemChan: itemSaver}
	//concurrentEngine.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})
	concurrentEngine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc: parser.ParseCity,
	})
}

package main

import (
	"runtime"
	"github.com/wzdxt/lianjia-spider/spider"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	spider.FollowChengjiao()
	time.Sleep(1 * time.Hour)
}



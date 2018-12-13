package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup
var urls = []string{
	"http://www.baidu.com",
	"http://www.alibaba.com",
	"https://www.tencent.com/zh-cn/index.html",
}

type SpiderJob struct {
	HasError   bool
	Url        string
	StatusCode string
}

func (s *SpiderJob) String() string {
	e := ""
	if s.HasError {
		e = "有错误"
	}
	return fmt.Sprintf("URL:%s \t 状态码：%s %s", s.Url, s.StatusCode, e)
}

func main() {
	responses := make(chan SpiderJob, len(urls))
	start := time.Now()
	log.Println("spider: START...")
	for _, url := range urls {
		wg.Add(1) // +1 信号量
		go func(url string) {
			defer wg.Done() // -1 信号量
			res, err := http.Get(url)
			if err != nil {
				responses <- SpiderJob{true, url, ""}
			} else {
				responses <- SpiderJob{false, url, res.Status}
			}
		}(url)
	}
	// closer
	go func() {
		wg.Wait() // 直到信号量 ==0
		close(responses)
	}()

	for range urls {
		response := <-responses
		log.Println(&response)
	}

	log.Printf("spider: DONE. took %v to run.\n", time.Now().Sub(start))
}

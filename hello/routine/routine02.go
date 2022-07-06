package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var urls = []string{
		"https://www.baidu.com/",
		"https://www.cnblogs.com/",
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetch(url, &wg)
	}

	wg.Wait()
	fmt.Println("数据已全部处理完成")
}

func fetch(url string, wg *sync.WaitGroup) (string, error) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(resp.Status)
	return resp.Status, nil
}

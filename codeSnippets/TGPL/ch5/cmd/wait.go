package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func WaitForServer(url string) error  {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries ++ {
		_, err := http.Get(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func main()  {
	fmt.Println(WaitForServer("http://example.com"))

	// 第三种策略：
	if err := WaitForServer("https://golang.org"); err != nil {
		log.Fatalf("site is down: %v\n", err)
	}

	fmt.Println("test")
}

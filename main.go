package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var sites = []string{
	"https://www.google.com/",
	"https://ya.ru/",
	"http://www.ru/",
	"https://example.ru/",
	"https://www.honeywell.com/",
	"https://github.com/",
	"https://pub.dev/",
	"https://stackoverflow.com/",
	"https://flutter.dev/",
	"https://gobyexample.com/",
}

func main() {
	for i := 0; ; i++ {
		select {
		case <-Attempt(i):
			break
		case <-time.After(6 * time.Second):
			println("Time to one try over!")
			break
		}
		println("waiting...")
		time.Sleep(10 * time.Second)
	}
}

func Attempt(i int) chan int {
	fmt.Printf("Attemp: %v\n", i)
	checked := make(chan int)
	go func() {
		defer operationTime(time.Now())
		checkAll()
		checked <- 0
	}()
	return checked
}

func checkAll() {
	statuses := make(chan HostStatus)

	go func() {
		defer close(statuses)
		var wg sync.WaitGroup
		for _, s := range sites {
			wg.Add(1)
			go func(s string) {
				defer wg.Done()
				isReachable(s, statuses)

			}(s)
		}
		wg.Wait()
	}()

	for s := range statuses {
		println(s.ToString())
		//println(s)
	}
}

func isReachable(s string, c chan HostStatus) {
	hs := HostStatus{host: s}
	hs.resp, hs.err = http.Get(s)
	c <- hs
}
func operationTime(start time.Time) {
	fmt.Printf("Operation succesed in %v\n", time.Now().Sub(start))
}

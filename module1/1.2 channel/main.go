package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)
	done := make(chan bool)

	go func() {
		var i = 0
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-done:
				return
			default:
				ch <- i
				i++
			}
		}
	}()

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-done:
				return
			case i := <-ch:
				fmt.Println(i)
			}
		}
	}()

	time.Sleep(100 * time.Second)
	done <- true
}

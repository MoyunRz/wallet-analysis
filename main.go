package main

import (
	"sync"
	"wallet-analysis/service"
)

func main() {
	var group = sync.WaitGroup{}
	group.Add(1)
	go func() {
		defer group.Done()
		service.ScanBlock()
	}()
	go func() {
		defer group.Done()
		service.StartSubscribe()
	}()
	group.Wait()

}

package main

import (
	"sync"
	"wallet-analysis/service"
)

func main() {
	var group = sync.WaitGroup{}
	group.Add(2)
	go func() {
		defer group.Done()
		service.ScanBlock()
	}()
	go func() {
		defer group.Done()
		service.StartSubscribe("0x6Cf015d91f18ec8E5bC5915366EA5e560Cbb6B31")
	}()
	group.Wait()
}
